package loginregister

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

type User struct {
	Login           string `json:"login"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

type Result struct {
	Status     int    `json:"Status"`
	Message    string `json:"Message"`
	ErrorCode  string `json:"ErrorCode"`
	UpdateRows string `json:"UpdateRows"`
	AuthToken  string `json:"AuthToken"`
	Error      string `json:"Error"`
}

func IsLogined(c *gin.Context) bool {
	store := ginsession.FromContext(c)
	_, ok := store.Get("jwt")
	if ok {
		return true
	}
	return false

}

func GetLanguage(c *gin.Context) string {
	language, ok := c.Get("language")
	if !ok {
		return "en"
	}
	return language.(string)
}

func (u *User) GetPostInformation(c *gin.Context) int {
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

func SendHttpRequest(user User, endpoint string) (Result, error) {

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

func SetError(c *gin.Context, key string, res Result, redirect string) {
	store := ginsession.FromContext(c)
	message := fmt.Sprintf("%v \n %v", res.Message, res.ErrorCode)
	store.Set(key, message)
	store.Save()
	c.Redirect(302, redirect)
}

func Login(c *gin.Context) {
	language := GetLanguage(c)
	key := "loginError"
	res := Result{
		Message:   "Należy się zalogować",
		ErrorCode: "",
	}
	if IsLogined(c) {
		c.Redirect(302, "/"+language+"/")
		return
	}
	user := User{}
	count := user.GetPostInformation(c)
	if count != 1 {
		SetError(c, key, res, "/"+language+"/register")
		return
	}
	endpoint := "http://studenci.herokuapp.com/user/login"
	res, err := SendHttpRequest(user, endpoint)
	if err != nil {
		SetError(c, key, res, "/"+language+"/register")
		return
	}
	store := ginsession.FromContext(c)
	store.Set("jwt", res.AuthToken)
	err = store.Save()
	if err != nil {
		SetError(c, key, Result{
			Message:   "Wystąpił błąd podczas logowania",
			ErrorCode: "",
		}, "/"+language+"/register")
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

func Register(c *gin.Context) {
	language := GetLanguage(c)
	key := "registerError"
	res := Result{
		Message:   "Wystąpił błąd podczas rejestracji",
		ErrorCode: "",
	}
	if IsLogined(c) {
		c.Redirect(302, "/"+language+"/")
		return
	}
	user := User{}
	count := user.GetPostInformation(c)
	if count != 0 {
		SetError(c, key, res, "/"+language+"/register")
		return
	}
	endpoint := "http://studenci.herokuapp.com/user/register"
	res, err := SendHttpRequest(user, endpoint)
	if err != nil {
		SetError(c, key, res, "/"+language+"/register")
		return
	}

	c.Redirect(302, "/"+language+"/register")
}
