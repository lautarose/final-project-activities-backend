package dtos

type InsertActivityDto struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

type InsertActivitiesDto []InsertActivityDto