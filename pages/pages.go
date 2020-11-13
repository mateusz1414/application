package pages

import (
	"application/studentsactions"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func isLogined(c *gin.Context) bool {
	store := ginsession.FromContext(c)
	_, ok := store.Get("apikey")
	if ok {
		return true
	}
	return false

}

//ShowStudents page with table of students
func ShowStudents(c *gin.Context) {
	students := []studentsactions.Student{}
	students, err := studentsactions.StudentsList()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(200, "contents/showtable", gin.H{
		"isLogined":    isLogined(c),
		"studentsList": students,
	})
}

//AddStudents page with add student form
func AddStudents(c *gin.Context) {
	c.HTML(200, "contents/addtable", gin.H{
		"isLogined": isLogined(c),
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
		"isLogined":    isLogined(c),
		"studentsList": students,
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
		"isLogined":    isLogined(c),
		"studentsList": students,
	})
}

//RegisterStudents load page to register or main page
func RegisterStudents(c *gin.Context) {
	if isLogined(c) {
		c.Redirect(301, "/")
		return
	}
	c.HTML(200, "contents/register", gin.H{
		"isLogined": false,
	})
}

//EditForm page with form to change data of student
func EditForm(c *gin.Context) {
	/*	if !isLogined(c) {
		c.Redirect(301, "/register")
		return
	}*/
	fmt.Println("wwwwwwwwwwwwwwwwwwwwww")
	studentIDString, ok := c.Params.Get("studentID")
	if !ok {
		c.Redirect(301, "/editstudents")
		return
	}
	studentIDInt, err := strconv.Atoi(studentIDString)
	if err != nil {
		fmt.Println("Server convert error:", err.Error())
		c.Redirect(500, "/")
		return
	}
	student := studentsactions.Student{}
	student.StudentID = studentIDInt
	err = student.GetStudent()
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(500, "/")
		return
	}
	c.HTML(200, "contents/editform", gin.H{
		"isLogined": true,
		"student":   student,
	})
}
