package main

import (
	"application/loginregister"
	"application/pages"
	"application/studentsactions"
	"application/translation"
	"os"
	"strings"

	"github.com/BurntSushi/toml"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var serverAddress = "http://localhost:8080"

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("translation/languages/active.pl.toml")
	server := gin.Default()
	server.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "templates",
		Extension:    ".tpl",
		Master:       "layouts/master",
		Partials:     []string{"page-aside/rightpanel"},
		DisableCache: false,
	})
	server.Static("/css", "./assets/css")
	server.Static("/js", "./assets/js")
	server.Use(ginsession.New())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{serverAddress}
	server.Use(cors.New(config))
	server.GET("session/:key", studentsactions.GetSession)
	server.Use(getLanguage(bundle))
	polish := server.Group("pl")
	{
		direct(polish)
	}
	english := server.Group("en")
	{
		direct(english)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}

func loginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		store := ginsession.FromContext(c)
		jwt, ok := store.Get("jwt")
		if !ok {
			return
		}
		c.Set("jwt", jwt.(string))
		c.Next()
	}
}

func direct(language *gin.RouterGroup) {
	language.GET("/", pages.ShowStudents)
	language.GET("/showstudents/", pages.ShowStudents)
	language.GET("/deletestudents/", pages.DeleteStudents)
	language.GET("/addstudents/", pages.AddStudents)
	language.GET("/editstudents/", pages.EditStudents)
	language.GET("/editstudentform/:studentID/", pages.EditForm)
	language.GET("/register/", pages.RegisterStudents)
	user := language.Group("user")
	{
		user.POST("/register/", loginregister.Register)
		user.POST("/login/", loginregister.Login)
	}
	/*studentaction := language.Group("action")
	{
		studentaction.Use(loginCheck())
		studentaction.POST("/add/", studentsactions.AddStudent)
		studentaction.GET("/del/:studentID", studentsactions.DelStudent)
		studentaction.POST("/edit/", studentsactions.EditStudent)
	}*/
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
