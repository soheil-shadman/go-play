package controllers

type signupParameters struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type loginParameters struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type checktokenParameters struct {
	Token string `json:"token"`
}
