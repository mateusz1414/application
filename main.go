package main

import (
	"application/pages"
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
	server.GET("/showstudents", pages.ShowStudents)
	server.GET("/deletestudents", pages.DeleteStudents)
	server.GET("/addstudents", pages.AddStudents)
	server.GET("/editstudents", pages.EditStudents)
	server.GET("/register", pages.RegisterStudents)
	server.GET("/editstudentform/:studentID", pages.EditForm)
	//user := server.Group("user")
	//      /user/regiser
	/*	{
		user.GET("/register")
		user.GET("/login")
	}*/
	/*	server.GET("/addstudents", pages.AddStudents)*/

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
