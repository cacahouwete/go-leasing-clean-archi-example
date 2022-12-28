package dto

type Schedule struct {
	BeginAt    string `json:"beginAt" binding:"required,datetime=2006-01-02T15:04:05Z07:00" pattern:"2006-01-02T15:04:05Z07:00"`
	EndAt      string `json:"endAt" binding:"required,datetime=2006-01-02T15:04:05Z07:00" pattern:"2006-01-02T15:04:05Z07:00"`
	CustomerID string `json:"customerId" binding:"required"`
	CarID      string `json:"carId" binding:"required"`
}

type ScheduleUpdate struct {
	BeginAt string `json:"beginAt" binding:"required,datetime=2006-01-02T15:04:05Z07:00" pattern:"2006-01-02T15:04:05Z07:00"`
	EndAt   string `json:"endAt" binding:"required,datetime=2006-01-02T15:04:05Z07:00" pattern:"2006-01-02T15:04:05Z07:00"`
}
