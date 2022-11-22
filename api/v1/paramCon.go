package v1

import (
	"github.com/gin-gonic/gin"
	"paramboard_go/service"
	util "paramboard_go/utilss"
)

func AddParam(c *gin.Context) {
	var paramSer service.ParamSer
	if err := c.ShouldBind(&paramSer); err == nil {
		res := paramSer.AddParam(c.Request.Context())
		c.JSON(0, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func GetParam(c *gin.Context) {
	var paramSer service.ParamSer
	if err := c.ShouldBind(&paramSer); err == nil {
		res := paramSer.FindByAuthorToken(c.Request.Context())
		c.JSON(0, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DelParam(c *gin.Context) {
	var paramSer service.ParamSer
	if err := c.ShouldBind(&paramSer); err == nil {
		res := paramSer.DelParam(c.Request.Context())
		c.JSON(0, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
