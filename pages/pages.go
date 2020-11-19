package pages

import (
	"application/loginregister"
	"application/studentsactions"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func getLoginError(c *gin.Context) (returnMessage string) {
	store := ginsession.FromContext(c)
	message, ok := store.Get("loginError")
	if !ok {
		return ""
	}
	returnMessage = message.(string)
	store.Delete("loginError")
	return returnMessage
}

func getRegisterError(c *gin.Context) (returnMessage string) {
	store := ginsession.FromContext(c)
	message, ok := store.Get("registerError")
	if !ok {
		return ""
	}
	returnMessage = message.(string)
	store.Delete("registerError")
	return returnMessage
}

//ShowStudents page with table of students
func ShowStudents(c *gin.Context) {
	students := []studentsactions.Student{}
	students, err := studentsactions.StudentsList()
	translation, _ := c.Get("translation")
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(200, "contents/showtable", gin.H{
		"isLogined":    loginregister.IsLogined(c),
		"studentsList": students,
		"language":     loginregister.GetLanguage(c),
		"translation":  translation.(map[string]string),
	})
}

//AddStudents page with add student form
func AddStudents(c *gin.Context) {
	c.HTML(200, "contents/addtable", gin.H{
		"isLogined": loginregister.IsLogined(c),
		"language":  loginregister.GetLanguage(c),
	})
}

//DeleteStudents page with table to delete students
func DeleteStudents(c *gin.Context) {
	students := []studentsactions.Student{}
	students, err := studentsactions.StudentsList()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(200, "contents/deletetable", gin.H{
		"isLogined":    loginregister.IsLogined(c),
		"studentsList": students,
		"language":     loginregister.GetLanguage(c),
	})
}

//EditStudents page with table to edit students
func EditStudents(c *gin.Context) {
	students := []studentsactions.Student{}
	students, err := studentsactions.StudentsList()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(200, "contents/edittable", gin.H{
		"isLogined":    loginregister.IsLogined(c),
		"studentsList": students,
		"language":     loginregister.GetLanguage(c),
	})
}

//RegisterStudents load page to register or main page
func RegisterStudents(c *gin.Context) {
	language := loginregister.GetLanguage(c)
	if loginregister.IsLogined(c) {
		c.Redirect(302, "/"+language+"/")
		return
	}
	c.HTML(200, "contents/register", gin.H{
		"isLogined":     false,
		"loginError":    strings.Split(getLoginError(c), "\n"),
		"registerError": strings.Split(getRegisterError(c), "\n"),
		"language":      language,
	})
}

//EditForm page with form to change data of student
func EditForm(c *gin.Context) {
	language := loginregister.GetLanguage(c)
	if !loginregister.IsLogined(c) {
		c.Redirect(302, "/"+language+"/register")
		return
	}
	studentIDString, ok := c.Params.Get("studentID")
	if !ok {
		c.Redirect(302, "/"+language+"/editstudents")
		return
	}
	studentIDInt, err := strconv.Atoi(studentIDString)
	if err != nil {
		fmt.Println("Server convert error:", err.Error())
		c.Redirect(500, "/"+language+"/")
		return
	}
	student := studentsactions.Student{}
	student.StudentID = studentIDInt
	err = student.GetStudent()
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(500, "/"+language+"/")
		return
	}
	c.HTML(200, "contents/editform", gin.H{
		"isLogined": true,
		"student":   student,
		"language":  language,
	})
}
