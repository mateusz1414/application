package loginregister

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

func (u *User) getPostInformation(c *gin.Context) int {
	var ok bool
	count := 0
	u.Login, ok = c.GetPostForm("user")
	if !ok {
		count++
	}
	u.Password, ok = c.GetPostForm("password")
	if !ok {
		count++
	}
	u.ConfirmPassword, ok = c.GetPostForm("confirmpassword")
	if !ok {
		count++
	}
	return count
}

func sendHTTPRequest(user User, endpoint string) (Result, error) {

	res := Result{}
	jsonReq, err := json.Marshal(user)
	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonReq))
	if err != nil {
		return res, fmt.Errorf("Cannot send request")
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return res, fmt.Errorf("Cannot send request")
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, fmt.Errorf("Server running error")
	}
	if resp.StatusCode != 200 {
		return res, fmt.Errorf(res.Message + res.ErrorCode)
	}

	return res, nil
}

//SetError save error in session
func SetError(c *gin.Context, key string, value string) {
	store := ginsession.FromContext(c)
	store.Set(key, value)
	store.Save()
}

func setErrorMessage(errorCode string) string {
	switch errorCode {
	case "Taki urzytkownik już istnieje":
		return "ThisUserAlreadyExists"
	case "Hasła nie są jednakowe":
		return "PasswordsDoNotMatch"
	case "Hasło jest zbyt krótkie":
		return "PasswordIsTooShort"
	case "Login jest zbyt krótki":
		return "LoginIsTooShort"
	}
	return ""
}

//Login login user in API
func Login(c *gin.Context) {
	language := GetLanguage(c)
	redirectURL := "/" + language + "/register/"
	res := Result{}
	if IsLogined(c) {
		c.Redirect(302, "/"+language+"/")
		return
	}
	user := User{}
	count := user.getPostInformation(c)
	if count != 1 {
		c.Redirect(302, redirectURL)
		return
	}
	endpoint := "http://studenci.herokuapp.com/user/login"
	res, err := sendHTTPRequest(user, endpoint)
	if err != nil {
		SetError(c, "LoginErrorFirst", "LoginFailed")
		SetError(c, "LoginErrorSecond", "IncorrectLoginOrPassword")
		c.Redirect(302, redirectURL)
		return
	}
	store := ginsession.FromContext(c)
	store.Set("jwt", res.AuthToken)
	err = store.Save()
	if err != nil {
		SetError(c, "LoginErrorFirst", "LoginFailed")
		c.Redirect(302, redirectURL)
		return
	}
	c.Redirect(302, "/"+language+"/")
	/*bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)*/
	//c.Redirect(200, "/")

}

//Register register user in API
func Register(c *gin.Context) {
	language := GetLanguage(c)
	redirectURL := "/" + language + "/register/"
	res := Result{}
	if IsLogined(c) {
		c.Redirect(302, "/"+language+"/")
		return
	}
	user := User{}
	count := user.getPostInformation(c)
	if count != 0 {
		c.Redirect(302, "/"+language+"/")
		return
	}
	endpoint := "http://studenci.herokuapp.com/user/register/"
	res, err := sendHTTPRequest(user, endpoint)
	if err != nil {
		SetError(c, "RegisterErrorFirst", "RegistrationFailed")
		SetError(c, "RegisterErrorSecond", setErrorMessage(res.ErrorCode))
		c.Redirect(302, redirectURL)
		return
	}

	c.Redirect(302, redirectURL)
}
