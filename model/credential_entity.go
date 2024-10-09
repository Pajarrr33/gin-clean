package model

import "time"

type Credential struct {
	Credential_id string 	`json:"id"`
	Email         string	`json:"email"`
	Password      string	`json:"password"`
	Created_at    time.Time	`json:"createdAt"`
	Updated_at 	  time.Time	`json:"updatedAt"`
}