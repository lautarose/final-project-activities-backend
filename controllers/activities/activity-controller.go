package controllers

import (
	"net/http"
	"strconv"

	service "activities-backend/services/activities"

	dto "activities-backend/dtos/activities"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetActivityById(c *gin.Context) {
	log.Debug("Activity id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	activityDto, err := service.ActivityService.GetActivityById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, activityDto)
		return
	}

	c.JSON(http.StatusOK, activityDto)
}

func GetActivitiesByUserId(c *gin.Context) {

	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	activitiesDto, err := service.ActivityService.GetActivitiesByUserId(auth)

	if err != nil {
		c.JSON(http.StatusBadRequest, activitiesDto)
		return
	}

	c.JSON(http.StatusOK, activitiesDto)
}

func InsertActivity(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	var activtyDto dto.InsertActivityDto

	if err := c.ShouldBindJSON(&activtyDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ActivityService.InsertActivity(auth, activtyDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar actividad"})
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteActivityById(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = service.ActivityService.DeleteActivityById(auth, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la actividad"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Actividad eliminada correctamente"})
}
