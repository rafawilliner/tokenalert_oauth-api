package access_token

import (
	"net/http"
	"tokenalert_oauth-api/src/domain/access_token"
	"tokenalert_oauth-api/src/services"

	"github.com/gin-gonic/gin"
	"github.com/rafawilliner/tokenalert_utils-go/src/rest_errors"
)

func GetById(c *gin.Context) {
	
	accessToken, getErr := services.AccessTokenService.GetById(c.Param("access_token_id"))
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func Create(c *gin.Context) {
	
	var request access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	accessToken, err := services.AccessTokenService.Create(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, accessToken)
}