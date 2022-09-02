package app

import (
	"tokenalert_oauth-api/src/datasources/mongodb/access_token_db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	access_token_db.InitDataBase()
	router.Run(":8080")
}