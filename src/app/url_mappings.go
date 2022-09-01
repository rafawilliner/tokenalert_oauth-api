package app

import (
	"tokenalert_oauth-api/src/controllers/access_token"
)

func mapUrls() {
	router.GET("/access_token/:access_token_id", access_token.GetById)
}