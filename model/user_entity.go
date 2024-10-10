package model

import (
	"time"
)

type User struct {
	User_id       string 			`json:"userId"`
	Credential    Credential 		`json:"credential"`
	Name          string 			`json:"name"`
	Age           int    			`json:"age"`
	Gender        string 			`json:"gender"`
	Created_at    time.Time			`json:"createdAt"`
	Updated_at 	  time.Time			`json:"updatedAt"`
}