package main

import (
	"commander/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	TOKEN = "e1816a06169975c188181ddb8a07e1f54c0232c54709c3a590a6f6a266fac125"
)

type Req struct {
	Title   string `json:"title"  binding:"required"`
	Content string `json:"content" binding:"required"`
}

type Rsp struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func main() {
	r := gin.Default()
	r.POST("/api/v1/post_msg", func(c *gin.Context) {
		var req Req
		var rsp Rsp

		if err := c.ShouldBindJSON(&req); err != nil {
			rsp.ErrCode = 400
			rsp.ErrMsg = err.Error()
			c.JSON(http.StatusBadRequest, rsp)
			return
		}

		if req.Title == "" || req.Content == "" {
			rsp.ErrCode = 400
			rsp.ErrMsg = "invalid title or content"
			c.JSON(http.StatusBadRequest, rsp)
			return
		}

		_ = utils.DingDingMsg(TOKEN, req.Title, req.Content)

		c.JSON(http.StatusOK, rsp)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
