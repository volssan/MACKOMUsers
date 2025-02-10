package handler

import "time"

type SuccessResponse struct {
	Success bool `json:"success"`
}

type GetUsersByFilterRequest struct {
	FromDate time.Time `json:"from_date"`
	ToDate   time.Time `json:"to_date"`
	MinAge   int       `json:"min_age"`
	MaxAge   int       `json:"max_age"`
}
