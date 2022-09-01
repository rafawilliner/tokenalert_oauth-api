package access_token

import (
	"net/http"
	"tokenalert_oauth-api/src/services"
	"github.com/gin-gonic/gin"
)

func GetById(c *gin.Context) {
	
	accessToken, getErr := services.AccessTokenService.GetById(c.Param("access_token_id"))
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
