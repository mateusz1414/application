package pages

import (
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func loadHtmlPage(c *gin.Context, contentPage string, rightPanelPage string) {
	c.HTML(200, "index.html", gin.H{
		"page-section": contentPage,
		"page-aside":   rightPanelPage,
	})
}

func isLogined(c *gin.Context) bool {
	store := ginsession.FromContext(c)
	_, ok := store.Get("apikey")
	if ok {
		return true
	}
	return false

}

//IndexHandler load template
func ShowStudents(c *gin.Context) {
	c.HTML(200, "contents/showtable", gin.H{
		"isLogined": isLogined(c),
	})
}

func AddStudents(c *gin.Context) {
	c.HTML(200, "contents/addtable", gin.H{
		"isLogined": isLogined(c),
	})
}

func DeleteStudents(c *gin.Context) {
	c.HTML(200, "contents/deletetable", gin.H{
		"isLogined": isLogined(c),
	})
}

func EditStudents(c *gin.Context) {
	c.HTML(200, "contents/edittable", gin.H{
		"isLogined": isLogined(c),
	})
}
