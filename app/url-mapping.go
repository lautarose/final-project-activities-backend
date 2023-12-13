package app

import (
	activityController "activities-backend/controllers/activities"
)

// MapUrls maps the urls
func MapUrls() {

	// Users Mapping
	router.GET("/activity/:id", activityController.GetActivityById)
	router.GET("/activities/user/:id", activityController.GetActivitiesByUserId)
	router.POST("/activity", activityController.InsertActivity)

}