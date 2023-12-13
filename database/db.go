package database

import (
	activityClient "activities-backend/clients/activities"
	activityModel "activities-backend/models/activities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "activities_db" //variable de entorno para nombre de la base de datos
	DBUser := "activitiesdb"  //variable de entorno para el usuario de la base de datos
	//DBPass := ""
	DBPass := "password" //variable de entorno para la pass de la base de datos
	DBHost := "10.64.64.3"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	db.LogMode(true)

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	activityClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&activityModel.Activity{})

	log.Info("Finishing Migration Database Tables")
}