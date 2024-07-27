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

	// Метод POST add record to db
	// http://localhost:8080/addRecord?name=Иван&class=10A&olimp=Математика&sub=Алгебра&teacher=Петров&stage=Школьный
	e.POST("/addRecord", addRecords)

	// Метод GET records
	// "http://localhost:8080/addRecord?name=Иван&class=10A&olimp=Математика&sub=Алгебра&teacher=Петров&stage=Школьный"
	e.GET("/getRecords", getRecords)

	//http://localhost:8080/getRecordsCount?name=Иван
	e.GET("/getRecordsCount", getRecordsCount)

	e.GET("/getAllRecords", getAllRecords)

	e.GET("/deleteAllRecords", deleteAllRecords)

	e.GET("/checkSnils", checkSnils)

	e.GET("/addStudent", addStudent)

	// Запуск сервера на порту 8080
	e.Logger.Fatal(e.Start("localhost:8080"))
}
