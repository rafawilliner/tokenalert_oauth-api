package users

type User struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	TelegramUser string `json:"telegram_user"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
