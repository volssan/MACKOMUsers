package core

import "time"

type User struct {
	Firstname     string    `json:"first_name"`
	Lastname      string    `json:"last_name"`
	Age           int       `json:"age"`
	RecordingDate time.Time `json:"recording_date"`
}
