package services

import (
	activityClient "activities-backend/clients/activities"
	dto "activities-backend/dtos/activities"
	activityModel "activities-backend/models/activities"
)

type activityService struct{}

type activityServiceInterface interface {
	GetActivityById(id int) (dto.GetActivityDto, error)
	GetActivitiesByUserId(userId int) (dto.GetActivitiesDto, error)
	InsertActivity(dto.InsertActivityDto) error
}

var (
	ActivityService activityServiceInterface
)

func init() {
	ActivityService = &activityService{}
}

func (s *activityService) GetActivityById(id int) (dto.GetActivityDto, error) {

	activity, err := activityClient.GetActivityById(id)
	var activityDto dto.GetActivityDto

	if err != nil {
		return activityDto, err
	}
	if activity.ActivityID == 0 {
		return activityDto, nil
	}
	activityDto.ActivityID = activity.ActivityID
	activityDto.UserID = activity.UserID
	activityDto.Title = activity.Title
	activityDto.Description = activity.Description
	activityDto.IsDone = activity.IsDone

	return activityDto, nil
}

func (s *activityService) GetActivitiesByUserId(userId int) (dto.GetActivitiesDto, error) {
	activities, err := activityClient.GetActivitiesByUserId(userId)
	var activitiesDto dto.GetActivitiesDto

	if err != nil {
		return activitiesDto, err
	}

	for _, activity := range activities {
		var activityDto dto.GetActivityDto
		activityDto.ActivityID = activity.ActivityID
		activityDto.UserID = activity.UserID
		activityDto.Title = activity.Title
		activityDto.Description = activity.Description
		activityDto.IsDone = activity.IsDone
		activitiesDto = append(activitiesDto, activityDto)
	}

	return activitiesDto, nil
}

func (s *activityService) InsertActivity(activityDto dto.InsertActivityDto) error {

	newActivity := activityModel.Activity{
		UserID:      activityDto.UserID,
		Title:       activityDto.Title,
		Description: activityDto.Description,
		IsDone:      activityDto.IsDone,
	}

	err := activityClient.InsertActivity(newActivity)
	if err != nil {
		return err
	}

	return nil
}