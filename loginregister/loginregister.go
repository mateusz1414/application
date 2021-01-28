package loginregister

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

type Credentials struct {
	Platforms map[string]Platform
	Google    Platform `json:"google"`
}

type Platform struct {
	Cid      string   `json:"cid"`
	Csecret  string   `json:"csecret"`
	Redirect string   `json:"redirect"`
	Scopes   []string `json:"scopes"`
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

//MakeConfig create platform config
func MakeConfig(platform Platform, endpoint oauth2.Endpoint) oauth2.Config {
	fmt.Println(platform)
	return oauth2.Config{
		ClientID:     platform.Cid,
		ClientSecret: platform.Csecret,
		RedirectURL:  platform.Redirect,
		Scopes:       platform.Scopes,
		Endpoint:     endpoint,
	}
}

//GenerateUrls return oauth urls
func generateSaveToken(c *gin.Context) (state string) {
	tok := make([]byte, 32)
	rand.Read(tok)
	state = base64.StdEncoding.EncodeToString(tok)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	return state
}

//OauthLogin c
func OauthLogin(c *gin.Context) {
	provider := c.Param("provider")
	fileinterface, _ := c.Get("credjson")
	file := fileinterface.([]byte)
	var cred Credentials
	json.Unmarshal(file, &cred)
	platform, endpoint := getData(provider, cred)
	config := oauth2.Config{
		ClientID:     platform.Cid,
		ClientSecret: platform.Csecret,
		RedirectURL:  platform.Redirect,
		Scopes:       platform.Scopes,
		Endpoint:     endpoint,
	}
	s := config.AuthCodeURL(generateSaveToken(c))
	fmt.Println(s)
	c.Redirect(302, s)
}

func getData(provider string, cred Credentials) (platform Platform, endpoint oauth2.Endpoint) {
	switch provider {
	case "google":
		platform = cred.Google
		endpoint = google.Endpoint
	}
	return platform, endpoint
}
