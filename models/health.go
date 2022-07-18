package models

type HealthResponse struct {
	Status int `json:"status"`
	Message string `json:"msg"`
}