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
		"permission":  loginregister.Permission(c),
	})
}

//AddStudents page with add student form
func AddStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/addtable", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
		"permission":  loginregister.Permission(c),
	})
}

//DeleteStudents page with table to delete students
func DeleteStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/deletetable", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
		"permission":  loginregister.Permission(c),
	})
}

//EditStudents page with table to edit students
func EditStudents(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/edittable", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
		"permission":  loginregister.Permission(c),
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
		"permission":  loginregister.Permission(c),
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
		"permission":  loginregister.Permission(c),
	})
}

func ShowDepartaments(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/departaments", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
		"permission":  loginregister.Permission(c),
	})
}

func ShowTeachers(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/teachers", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
		"permission":  loginregister.Permission(c),
	})
}

func ShowGrades(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/grades", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
		"permission":  loginregister.Permission(c),
	})
}

func ShowGradesForTeacher(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/gradest", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
		"permission":  loginregister.Permission(c),
	})
}

func ShipmentPermission(c *gin.Context) {
	translation, _ := c.Get("translation")
	c.HTML(200, "contents/permission", gin.H{
		"isLogined":   loginregister.IsLogined(c),
		"language":    loginregister.GetLanguage(c),
		"translation": translation.(map[string]string),
		"permission":  loginregister.Permission(c),
	})
}
