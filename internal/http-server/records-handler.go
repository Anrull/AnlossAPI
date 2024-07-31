package http_server

import (
	"AnlossAPI/internal/bot"
	"AnlossAPI/internal/http-server/structs"
	"AnlossAPI/internal/storage/sqlite"
	"AnlossAPI/pkg/env"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func addRecords(c echo.Context) error {
	var err error
	record := &structs.Record{}

	// Привязка параметров запроса к структуре Record
	if err = c.Bind(record); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if record.Stage == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "none stage")
	}
	if record.Olimp == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "none olimp")
	}
	if record.Sub == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "none sub")
	}
	if record.Teacher == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "none teacher")
	}

	fmt.Println(record)

	if record.Date == "nil" || record.Date == "null" {
		err = sqlite.AddRecord(
			record.Name, record.Class, record.Olimp, record.Sub, record.Teacher, record.Stage)
	} else {
		err = sqlite.AddRecord(
			record.Name, record.Class, record.Olimp, record.Sub, record.Teacher, record.Stage, record.Date)
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}

func getRecords(c echo.Context) error {
	record := &structs.GetRecord{}

	// Привязка параметров запроса к структуре Record
	if err := c.Bind(record); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if record.Stage == "" {
		record.Stage = "nil"
	}
	if record.Olimp == "" {
		record.Olimp = "nil"
	}
	if record.Sub == "" {
		record.Sub = "nil"
	}
	if record.Teacher == "" {
		record.Teacher = "nil"
	}

	records, err := sqlite.GetRecords(
		record.Name, record.Sub, record.Olimp, record.Stage, record.Teacher)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	jsonData, err := json.Marshal(records)
	if err != nil {
		Logger.Info("internal.http-server.getRecords" + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, string(jsonData))
}

func getRecordsCount(c echo.Context) error {
	record := &structs.GetRecord{}

	if err := c.Bind(record); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if record.Stage == "" {
		record.Stage = "nil"
	}
	if record.Olimp == "" {
		record.Olimp = "nil"
	}
	if record.Sub == "" {
		record.Sub = "nil"
	}
	if record.Teacher == "" {
		record.Teacher = "nil"
	}

	count, err := sqlite.GetRecordsCount(record.Name, record.Sub, record.Olimp, record.Stage, record.Teacher)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, fmt.Sprintf(`{"message": "ok", "count": %d}`, count))
}

func getAllRecords(c echo.Context) error {
	records, err := sqlite.GetAllRecords()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	jsonData, err := json.Marshal(records)
	if err != nil {
		Logger.Info("internal.http-server.getRecords" + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, string(jsonData))
}

func deleteAllRecords(c echo.Context) error {
	admin := structs.Admin{}

	if err := c.Bind(&admin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if env.GetValue("ADMIN_PASSWORD") != admin.Password {
		Logger.Info("unauthorized, password: " + admin.Password)
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	err := bot.SendFile(Config.RecordsPath, "records.db", "time")
	if err != nil {
		Logger.Info("StatusInternalServerError (bot send file db): " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = sqlite.DeleteAllRecords(Config.RecordsPath)

	if err != nil {
		Logger.Info("StatusInternalServerError (delete all records): " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, `{"message": "ok"}`)
}
