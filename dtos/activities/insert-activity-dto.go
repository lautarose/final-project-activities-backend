package dtos

type InsertActivityDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type InsertActivitiesDto []InsertActivityDto