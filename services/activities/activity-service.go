package services

import (
	activityClient "activities-backend/clients/activities"
	dto "activities-backend/dtos/activities"
	activityModel "activities-backend/models/activities"
	jwtUtils "activities-backend/utils/jwt"
	"strconv"
)

type activityService struct{}

type activityServiceInterface interface {
	GetActivityById(id int) (dto.GetActivityDto, error)
	GetActivitiesByUserId(auth string) (dto.GetActivitiesDto, error)
	InsertActivity(auth string, activity dto.InsertActivityDto) error
	DeleteActivityById(auth string, activityId int) error
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

func (s *activityService) GetActivitiesByUserId(auth string) (dto.GetActivitiesDto, error) {
	var activitiesDto dto.GetActivitiesDto
	
	// Verificar el token de autenticación
	claims, err := jwtUtils.VerifyToken(auth)
	if err != nil {
		return activitiesDto, err
	}

	// Obtener el ID del usuario del token
	id, err := strconv.Atoi(claims.Id)
	if err != nil {
		return activitiesDto, err
	}

	activities, err := activityClient.GetActivitiesByUserId(id)
	

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

func (s *activityService) InsertActivity(auth string, activityDto dto.InsertActivityDto) error {

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

func (s *activityService) DeleteActivityById(auth string, activityId int) error {
	err := activityClient.DeleteActivityById(activityId)
	if err != nil {
		return err
	}

	return nil
}