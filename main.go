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
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

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
	server.Static("/assets", "./assets/css")
	server.Use(ginsession.New())
	//	server.Use(translations())
	server.Use(getLanguage(bundle))
	/*	defaultLanguage := server.Group("/")
		{
			defaultLanguage.Use(setLanguage("", bundle))
			direct(defaultLanguage)
			/*		defaultLanguage.GET("/", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"hello": "world",
					})
				})
		}/*/
	polish := server.Group("pl")
	{
		//polish.Use(setLanguage("pl", bundle))
		direct(polish)
	}
	english := server.Group("en")
	{
		//english.Use(setLanguage("en", bundle))
		direct(english)
	}
	/*	server.GET("/", pages.ShowStudents)
		server.GET("/showstudents/", pages.ShowStudents)
		server.GET("/deletestudents/", pages.DeleteStudents)
		server.GET("/addstudents/", pages.AddStudents)
		server.GET("/editstudents/", pages.EditStudents)
		server.GET("/editstudentform/:studentID/", pages.EditForm)
		server.GET("/register/", pages.RegisterStudents)
		user := server.Group("user")
		{
			user.POST("/register/", loginregister.Register)
			user.POST("/login/", loginregister.Login)
		}
		studentaction := server.Group("action")
		{
			studentaction.Use(loginCheck())
			studentaction.POST("/add/", studentsactions.AddStudent)
			studentaction.GET("/del/:studentID", studentsactions.DelStudent)
			studentaction.POST("/edit/", studentsactions.EditStudent)
		}*/

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
func cos(*gin.Context) {

}

func translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		//	bundle.MustLoadMessageFile("")

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
	studentaction := language.Group("action")
	{
		studentaction.Use(loginCheck())
		studentaction.POST("/add/", studentsactions.AddStudent)
		studentaction.GET("/del/:studentID", studentsactions.DelStudent)
		studentaction.POST("/edit/", studentsactions.EditStudent)
	}
}

func setLanguage(language string, bundle *i18n.Bundle) gin.HandlerFunc {
	return func(c *gin.Context) {
		if language == "" || !isAccepted(language) {
			language = ""
			languages := c.GetHeader("Accept-Language")
			languagesArray := strings.Split(languages, ",")
			for i := 1; i < len(languagesArray); i++ {
				lang := strings.Split(languagesArray[i], ";")
				if isAccepted(lang[0]) {
					language = lang[0]
					break
				}
			}
		}
		if language == "" {
			language = "en"
		}
		localizer := i18n.NewLocalizer(bundle, language)
		//fmt.Println(*localizer)
		c.Set("translation", translation.LoadTranslation(localizer))
		c.Set("language", language)
		c.Next()
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
		//fmt.Println(*localizer)
		c.Set("translation", translation.LoadTranslation(localizer))
		c.Set("language", language)
		c.Next()
		//languages := c.GetHeader("Accept-Language")
		//localizer := i18n.NewLocalizer(bundle, "", languages)
		//fmt.Println(language.ParseAcceptLanguage(languages))
		//c.Set("language", language)
		//c.Next()
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
