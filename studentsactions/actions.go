package studentsactions

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Student is struct of data about student
type Student struct {
	StudentID        int    `json:"idstudenta"`
	StudentFirstName string `json:"imiestudenta"`
	StudentLastName  string `json:"nazwiskostudenta"`
	DateOfBrith      string `json:"datastudenta"`
	StudentFaciulty  string `json:"wydzialstudenta"`
	StudentGender    string `json:"plecstudenta"`
}

//Session struct values from json
type Session struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Results is struct data from API
type Results struct {
	Status       int       `json:"Status"`
	TotalResults int       `json:"TotalResults"`
	Students     []Student `json:"Student"`
	ErrorCode    string    `json:"ErrorCode"`
}

//GetSession return session data
func GetSession(c *gin.Context) {
	key := c.Param("key")
	session := sessions.Default(c)
	value := session.Get(key)
	c.JSON(200, gin.H{
		key: value,
	})
}

//SetSession set values in session
func SetSession(c *gin.Context) {
	change := []Session{}
	err := c.ShouldBindJSON(&change)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, change)
		return
	}
	err = SaveSession(change, c)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "filed",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func SaveSession(change []Session, c *gin.Context) error {
	session := sessions.Default(c)
	for _, element := range change {
		session.Set(element.Key, element.Value)
	}
	session.Save()
	return nil
}

//ClearKey remove value in session
func ClearKey(c *gin.Context) {
	key := c.Param("key")
	session := sessions.Default(c)
	session.Delete(key)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("permissions")
	session.Delete("jwt")
	session.Delete("userID")
	session.Delete("email")
	session.Save()
	language := getLanguage(c)
	c.Redirect(302, "/"+language+"/")
}

func getLanguage(c *gin.Context) string {
	language, ok := c.Get("language")
	if !ok {
		return "en"
	}
	return language.(string)
}
