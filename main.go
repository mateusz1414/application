package main

import (
	"application/loginregister"
	"application/pages"
	"application/studentsactions"
	"os"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func main() {
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
	server.GET("/", pages.ShowStudents)
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
