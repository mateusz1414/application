package pages

import (
	"application/loginregister"

	"github.com/gin-gonic/gin"
)

//ShowStudents page with table of students
func ShowStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/showstudents", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
	})
}

//ShowTeachers page with table of teachers
func ShowTeachers(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/showteachers", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
	})
}

//AddStudents page with add student form
func AddStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/addtable", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
	})
}

//DeleteStudents page with table to delete students
func DeleteStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/deletetable", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
	})
}

//EditStudents page with table to edit students
func EditStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/edittable", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
	})
}

//Register load page to register or main page
func Register(c *gin.Context) {
	translation, _ := c.Get("translation")
	language := loginregister.GetLanguage(c)
	c.HTML(200, "contents/register", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    language,
		"translation": translation.(map[string]string),
	})
}

func Login(c *gin.Context) {
	translation, _ := c.Get("translation")
	language := loginregister.GetLanguage(c)
	c.HTML(200, "contents/login", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    language,
		"translation": translation.(map[string]string),
	})
}

func GetGrades(c *gin.Context) {
	translation, _ := c.Get("translation")
	language := loginregister.GetLanguage(c)
	c.HTML(200, "contents/getgrades", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    language,
		"translation": translation.(map[string]string),
	})
}

func AddGrades(c *gin.Context) {
	translation, _ := c.Get("translation")
	language := loginregister.GetLanguage(c)
	c.HTML(200, "contents/addgrades", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    language,
		"translation": translation.(map[string]string),
	})
}

func UserPanel(c *gin.Context) {
	translation, _ := c.Get("translation")
	language := loginregister.GetLanguage(c)
	c.HTML(200, "contents/userpanel", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    language,
		"translation": translation.(map[string]string),
	})
}

func Modify(c *gin.Context) {
	translation, _ := c.Get("translation")
	language := loginregister.GetLanguage(c)
	c.HTML(200, "contents/modify", gin.H{
		"user":        loginregister.IsLogined(c),
		"language":    language,
		"translation": translation.(map[string]string),
	})
}

//EditForm page with form to change data of student
func EditForm(c *gin.Context) {
	/*	translation, _ := c.Get("translation")
		language := loginregister.GetLanguage(c)
			if !loginregister.IsLogined(c) {
				c.Redirect(302, "/"+language+"/register")
				return
			}
			c.HTML(200, "contents/editform", gin.H{
				"isLogined":   true,
				"language":    language,
				"translation": translation.(map[string]string),
			})*/
}
