package clients

import (
	model "activities-backend/models/activities"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetActivityById(activityId int) (model.Activity, error) {
	var activity model.Activity

	err := Db.Where("activity_id = ?", activityId).First(&activity).Error

	if err != nil {
		log.Println(err)
		return activity, err
	}

	log.Debug("Activity: ", activity)

	return activity, nil
}

func GetActivitiesByUserId(userId int) (model.Activities, error) {
	var activities model.Activities

	err := Db.Where("user_id = ?", userId).Find(&activities).Error

	if err != nil {
		log.Println(err)
		return activities, err
	}
	log.Debug("Activities: ", activities)

	return activities, nil
}

func InsertActivity(newActivity model.Activity) error {
	err := Db.Create(&newActivity).Error
	if err != nil {
		log.Println(err)
		return err
	}

	log.Debug("Activity inserted: ", newActivity)

	return nil
}

func DeleteActivityById(activityId int) error {
	var activity model.Activity
	err := Db.Where("activity_id = ?", activityId).Delete(&activity).Error
	if err != nil {
		log.Println(err)
		return err
	}

	log.Debug("Activity deleted with ID: ", activityId)

	return nil
}
