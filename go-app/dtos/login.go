package dtos

import "time"

type LoginData struct {
	LastLogin time.Time `json:"last_login"`
}
