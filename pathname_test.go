package gopathname

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetFileNames(t *testing.T) {
	r := gin.Default()

	r.SetHTMLTemplate(GetTemplate("test"))

	r.GET("/", func(c *gin.Context) {

		var query struct {
			File string `form:"file" binding:"required"`
		}

		if err := c.Bind(&query); err != nil {
			t.Error(err)
		}

		c.HTML(200, query.File, nil)
	})

	r.Run(":80")
}
