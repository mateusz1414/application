package studentsactions

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
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
	store := ginsession.FromContext(c)
	value, _ := store.Get(key)
	//	a, _ := store.Get("jwt")
	c.JSON(200, gin.H{
		key: value,
	})
}

//SetSession set values in session
func SetSession(c *gin.Context) {
	session := []Session{}
	err := c.ShouldBindJSON(&session)
	if err != nil {
		c.JSON(500, gin.H{})
		return
	}
	store := ginsession.FromContext(c)
	for _, value := range session {
		store.Set(value.Key, value.Value)
	}
	err = store.Save()
	if err != nil {
		c.JSON(500, gin.H{})
		return
	}
}

//ClearKey remove value in session
func ClearKey(c *gin.Context) {
	key := c.Param("key")
	store := ginsession.FromContext(c)
	store.Delete(key)
}
