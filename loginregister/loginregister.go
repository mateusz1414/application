package loginregister

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
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

//IsLogined check if token is in session
func IsLogined(c *gin.Context) bool {
	store := ginsession.FromContext(c)
	_, ok := store.Get("jwt")
	if ok {
		return true
	}
	return false

}

//GetLanguage return actual language
func GetLanguage(c *gin.Context) string {
	language, ok := c.Get("language")
	if !ok {
		return "en"
	}
	return language.(string)
}

func Permission(c *gin.Context) string {
	store := ginsession.FromContext(c)
	permission, ok := store.Get("permission")
	if !ok {
		return ""
	}
	return permission.(string)

}
