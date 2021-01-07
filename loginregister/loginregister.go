package loginregister

import (
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//User data about user
type User struct {
	Login           string `json:"login"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

//Result is return message from API
type Result struct {
	Status     int    `json:"Status"`
	Message    string `json:"Message"`
	ErrorCode  string `json:"ErrorCode"`
	UpdateRows string `json:"UpdateRows"`
	AuthToken  string `json:"AuthToken"`
	Error      string `json:"Error"`
}

type Session struct {
	IsLogined   bool
	UserID      int
	Email       string
	Permissions string
}

//IsLogined return data about user
func IsLogined(c *gin.Context) Session {
	s := Session{
		IsLogined: false,
	}
	session := sessions.Default(c)
	if session.Get("jwt") == nil {
		return s
	}
	if session.Get("permissions") == nil {
		return s
	}
	s.Permissions = session.Get("permissions").(string)
	if session.Get("userID") == nil {
		return s
	}
	s.UserID, _ = strconv.Atoi(session.Get("userID").(string))
	if session.Get("email") == nil {
		return s
	}
	s.Email = session.Get("email").(string)
	s.IsLogined = true
	return s

}

//GetLanguage return actual language
func GetLanguage(c *gin.Context) string {
	language, ok := c.Get("language")
	if !ok {
		return "en"
	}
	return language.(string)
}
