package models

type NewConnectionReq struct {
	Name     string `json:"name"`
	Urls     string `json:"urls"`
	Username string `json:"username"`
	Password string `json:"password"`
}
