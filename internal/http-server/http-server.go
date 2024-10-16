package http_server

import (
	"AnlossAPI/internal/config"
	"github.com/labstack/echo/v4"
	"log/slog"
)

var Logger *slog.Logger
var Config *config.Config

func New(logger *slog.Logger, cfg *config.Config) {
	Logger = logger
	Config = cfg

	e := echo.New()

	e.GET("/addRecord", addRecords)
	e.GET("/getRecords", getRecords)
	e.GET("/getRecordsCount", getRecordsCount)
	e.GET("/getAllRecords", getAllRecords)
	e.GET("/deleteAllRecords", deleteAllRecords)

	e.GET("/checkSnils", checkSnils)
	e.GET("/addStudent", addStudent)

	e.GET("/getJson", getJson)

	e.GET("/getTrackerData", getTrackerData)
	e.GET("/getStages", getStages)
	e.GET("/getTimetableTeacher", getTimetableTeacher)
	e.GET("/getTimetable", getTimetable)
	e.GET("/getTeachers", getTeachers)

	e.Logger.Fatal(e.Start(cfg.Port))
}
