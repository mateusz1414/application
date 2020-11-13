package studentsactions

import (
	"encoding/json"
	"fmt"
	"net/http"
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
