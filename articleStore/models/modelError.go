package models

//Error ...
type Error struct {
	//format=int32,
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
