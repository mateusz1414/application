package pages

import (
	"application/loginregister"

	"github.com/gin-gonic/gin"
)

//ShowStudents page with table of students
func ShowStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/showtable", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
	})
}

//AddStudents page with add student form
func AddStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/addtable", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
	})
}

//DeleteStudents page with table to delete students
func DeleteStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/deletetable", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
	})
}

//EditStudents page with table to edit students
func EditStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/edittable", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
	})
}

//RegisterStudents load page to register or main page
func RegisterStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	language := loginregister.GetLanguage(c)
	if loginregister.IsLogined(c) {
		c.Redirect(302, "/"+language+"/")
		return
	}
	c.HTML(200, "contents/register", gin.H{
		"isLogined":   false,
		"language":    language,
		"translation": translation.(map[string]string),
	})
}

//EditForm page with form to change data of student
func EditForm(c *gin.Context) {
	translation, _ := c.Get("translation")
	language := loginregister.GetLanguage(c)
	if !loginregister.IsLogined(c) {
		c.Redirect(302, "/"+language+"/register")
		return
	}
	c.HTML(200, "contents/editform", gin.H{
		"isLogined":   true,
		"language":    language,
		"translation": translation.(map[string]string),
	})
}
