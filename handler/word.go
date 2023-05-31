package handler

import (
	"net/http"
	"quote/pkg"
	"quote/utils"

	"github.com/gin-gonic/gin"
)

func HandlerGetRandomWord(c *gin.Context) {
	result, err := pkg.RandomWordsGet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildResponse("Error", nil))
		return
	}
	c.JSON(http.StatusOK, utils.BuildResponse("Success", result))
}
