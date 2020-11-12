package pages

import (
	"application/studentsactions"
	"fmt"

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
