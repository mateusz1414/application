package studentsactions

import (
	"application/loginregister"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

//Results is struct data from API
type Results struct {
	Status       int       `json:"Status"`
	TotalResults int       `json:"TotalResults"`
	Students     []Student `json:"Student"`
	ErrorCode    string    `json:"ErrorCode"`
}

//StudentsList return list of students in API
func StudentsList() ([]Student, error) {
	result := Results{}
	endpoint := "https://studenci.herokuapp.com/student"
	response, err := http.Get(endpoint)
	if err != nil {
		return result.Students, fmt.Errorf("Api connection error")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		err = json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			return result.Students, err
		}
		return result.Students, fmt.Errorf(result.ErrorCode)
	}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return result.Students, err
	}
	return result.Students, nil
}

//GetStudent getting student with id
func (s *Student) GetStudent() error {
	//should be correct in API
	students, err := StudentsList()
	if err != nil {
		return err
	}
	for _, student := range students {
		if student.StudentID == s.StudentID {
			*s = student
			return nil
		}
	}
	return fmt.Errorf("Student not found")
}

func (s *Student) getDataWithForm(c *gin.Context) {
	s.StudentFirstName = c.PostForm("studentFirstName")
	s.StudentLastName = c.PostForm("studentLastName")
	s.StudentFaciulty = c.PostForm("studentFaciulty")
	s.DateOfBrith = c.PostForm("studentDateOfBrith")
	s.StudentGender = c.PostForm("studentGender")
}

func sendHTTPRequest(student Student, endpoint string, method string, jwt string) (loginregister.Result, error) {
	result := loginregister.Result{}
	jsonRequest, err := json.Marshal(student)
	if err != nil {
		return result, fmt.Errorf("Nie można wysłać zapytania")
	}
	request, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return result, fmt.Errorf("Nie można wysłać zapytania")
	}
	request.Header.Set("Authorization", "Bearer "+jwt)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return result, err
	}
	if response.StatusCode != 200 {
		return result, fmt.Errorf(result.ErrorCode)
	}
	return result, nil

}

func clearJWT(c *gin.Context) {
	store := ginsession.FromContext(c)
	store.Delete("jwt")
}

//AddStudent add student to database
func AddStudent(c *gin.Context) {
	language := loginregister.GetLanguage(c)
	jwt, ok := c.Get("jwt")
	if !ok {
		loginregister.SetError(c, "loginErrorFirst", "MustLoginInError")
		c.Redirect(302, "/"+language+"/register/")
		return
	}
	student := Student{}
	student.getDataWithForm(c)
	endpoint := "https://studenci.herokuapp.com/student/"
	result, err := sendHTTPRequest(student, endpoint, http.MethodPost, jwt.(string))
	if result.Error != "" {
		clearJWT(c)
		loginregister.SetError(c, "loginErrorFirst", "MustLoginInError")
		c.Redirect(302, "/"+language+"/register/")
		return
	}
	if err != nil {
		loginregister.SetError(c, "addError", "AddError")
		c.Redirect(302, "/"+language+"/addstudents/")
		return
	}
	c.Redirect(302, "/"+language+"/")

}

//DelStudent delete student from database
func DelStudent(c *gin.Context) {
	language := loginregister.GetLanguage(c)
	jwt, ok := c.Get("jwt")
	if !ok {
		loginregister.SetError(c, "loginErrorFirst", "MustLoginInError")
		c.Redirect(302, "/"+language+"/register/")
		return
	}
	student := Student{}
	studentIDString, ok := c.Params.Get("studentID")
	if !ok {
		c.Redirect(302, "/"+language+"/deletestudents/")
		return
	}
	var err error
	student.StudentID, err = strconv.Atoi(studentIDString)
	if err != nil {
		c.Redirect(302, "/"+language+"/deletestudents/")
		return
	}
	endpoint := "https://studenci.herokuapp.com/student/" + studentIDString
	result, err := sendHTTPRequest(student, endpoint, http.MethodDelete, jwt.(string))
	if result.Error != "" {
		clearJWT(c)
		loginregister.SetError(c, "loginErrorFirst", "MustLoginInError")
		c.Redirect(302, "/"+language+"/register/")
		return
	}
	if err != nil {
		fmt.Println(result)
		c.Redirect(302, "/"+language+"/deletestudents/")
		return
	}
	c.Redirect(302, "/"+language+"/")
}

//EditStudent change data of student in database
func EditStudent(c *gin.Context) {
	language := loginregister.GetLanguage(c)
	jwt, ok := c.Get("jwt")
	if !ok {
		loginregister.SetError(c, "loginErrorFirst", "MustLoginInError")
		c.Redirect(302, "/"+language+"/register/")
		return
	}
	student := Student{}
	student.getDataWithForm(c)
	studentIDString := strconv.Itoa(student.StudentID)
	endpoint := "https://studenci.herokuapp.com/student/" + studentIDString
	result, err := sendHTTPRequest(student, endpoint, http.MethodPut, jwt.(string))
	if result.Error != "" {
		clearJWT(c)
		loginregister.SetError(c, "loginErrorFirst", "MustLoginInError")
		c.Redirect(302, "/"+language+"/register/")
		return
	}
	if err != nil {
		loginregister.SetError(c, "editError", "EditError")
		c.Redirect(302, "/"+language+"/editstudentform/"+studentIDString+"/")
		return
	}

	c.Redirect(302, "/"+language+"/")

}
