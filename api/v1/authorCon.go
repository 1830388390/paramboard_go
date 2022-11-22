package v1

import (
	"github.com/gin-gonic/gin"
	"paramboard_go/service"
	util "paramboard_go/utilss"
)

func AddAuthor(c *gin.Context) {
	var authorSer service.AuthorSer
	if err := c.ShouldBind(&authorSer); err == nil {
		res := authorSer.AddAuthor(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
