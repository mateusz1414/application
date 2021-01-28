package main

import (
	"application/loginregister"
	"application/pages"
	"application/studentsactions"
	"application/translation"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var serverAddress = "https://studentscode.online/"

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("translation/languages/active.pl.toml")
	server := gin.Default()

	//temlates
	server.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "templates",
		Extension: ".tpl",
		Master:    "layouts/master",
		//Partials:     []string{"page-aside/rightpanel"},
		DisableCache: false,
	})

	//static files
	server.Static("/css", "./assets/css")
	server.Static("/js", "./assets/js")

	//session
	store := cookie.NewStore([]byte("thisIsGoLanguage"))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 3600,
	})
	server.Use(sessions.Sessions("go_session_id", store))

	//cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	server.Use(cors.New(config))

	//session directores
	server.GET("session/:key", studentsactions.GetSession)
	server.DELETE("session/:key", studentsactions.ClearKey)
	server.POST("session/", studentsactions.SetSession)

	//OAuth directory
	file, err := ioutil.ReadFile("./config/credentials.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	server.GET("/auth/:provider", sendConfig("credjson", file), loginregister.OauthLogin)

	/*c	onf := &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://localhost:8080/auth/google/",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	fmt.Println(conf.AuthCodeURL("run"))
	server.GET("/auth/:provider", func(c *gin.Context) {
		tok, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
		if err != nil {
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!")
			return
		}
		fmt.Println(tok)
		client := conf.Client(oauth2.NoContext, tok)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!")
			return
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		log.Println("Resp body: ", string(data))
	})*/

	//languages
	server.Use(getLanguage(bundle))
	polish := server.Group("pl")
	{
		direct(polish)
	}
	english := server.Group("en")
	{
		direct(english)
	}

	//server start
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}

func sendConfig(name string, file []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(name, file)
		c.Next()
	}
}

func sess() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		//session.Options(sessions.Options{MaxAge: 3600})
		session.Save()
	}
}

func direct(language *gin.RouterGroup) {
	language.GET("/", pages.ShowStudents)
	language.GET("/students/", pages.ShowStudents)
	language.GET("/teachers/", pages.ShowTeachers)
	language.GET("/deletestudents/", pages.DeleteStudents)
	language.GET("/addstudents/", pages.AddStudents)
	language.GET("/editstudents/", pages.EditStudents)
	language.GET("/editstudentform/:studentID/", pages.EditForm)
	language.GET("/register/", authMiddleWeare(nil), pages.Register)
	language.GET("/login/", authMiddleWeare(nil), pages.Login)
	language.GET("/login/:status", authMiddleWeare(nil), pages.Login)
	language.GET("/getgrades/", authMiddleWeare([]string{"student"}), pages.GetGrades)
	language.GET("/addgrades/", authMiddleWeare([]string{"teacher"}), pages.AddGrades)
	language.GET("/modify/", authMiddleWeare([]string{"dean"}), pages.Modify)
	language.GET("/user/", authMiddleWeare([]string{"teacher", "student", "dean", "user"}), pages.UserPanel)
	language.GET("/logout/", studentsactions.Logout)
}

func sendCredentials(cred loginregister.Credentials) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func authMiddleWeare(permisssion []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := loginregister.IsLogined(c)
		language, _ := c.Get("language")
		if !user.IsLogined && permisssion != nil {
			//cooke należy się zalogować
			c.Redirect(302, "/"+language.(string)+"/login/")
			return
		} else if user.IsLogined && permisssion == nil {
			c.Redirect(302, "/"+language.(string)+"/")
			return
		} else if user.IsLogined && !having(user.Permissions, permisssion) {
			fmt.Println("aa")
			//cooke nie masz uprawnien
			c.Redirect(302, "/"+language.(string)+"/")
			return
		}
	}
}

func getLanguage(bundle *i18n.Bundle) gin.HandlerFunc {
	return func(c *gin.Context) {
		directores := strings.Split(c.Request.URL.Path, "/")
		language := directores[1]
		if !isAccepted(language) {
			language = ""
			accept := c.GetHeader("Accept-Language")
			acceptArray := strings.Split(accept, ",")
			for i := 1; i < len(acceptArray); i++ {
				lang := strings.Split(acceptArray[i], ";")
				if isAccepted(lang[0]) {
					language = lang[0]
					break
				}
			}
			if language == "" {
				language = "en"
			}
			c.Redirect(302, "/"+language+c.Request.URL.Path)
			return
		}
		localizer := i18n.NewLocalizer(bundle, language)
		c.Set("translation", translation.LoadTranslation(localizer))
		c.Set("language", language)
		c.Next()
	}
}

func isAccepted(language string) bool {
	switch language {
	case "pl":
		return true
	case "en":
		return true
	}
	return false
}

func having(have string, requied []string) bool {
	for _, value := range requied {
		if value == have {
			return true
		}
	}
	return false
}
