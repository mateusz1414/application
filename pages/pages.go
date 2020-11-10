package pages

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//IndexHandler load template
func ShowStudents(c *gin.Context) {
	fmt.Println(c.Param("p"))
	c.HTML(200, "index.html", gin.H{})
}
