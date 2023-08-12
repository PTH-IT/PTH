package model

type RegisterUserResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Token struct {
	Authorization string `json:"Authorization"`
	Type          string `json:"type"`
}
