package models

type User struct {
	// Структура пользователя
	Id       int    `json:"-"`
	EMail    string `json:"e-mail" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Role     int    `json:"role" binding:"required"`
}
