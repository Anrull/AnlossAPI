package http_server

import (
	"AnlossAPI/internal/http-server/structs"
	"AnlossAPI/pkg/timetable"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func getJson(c echo.Context) error {
	dataStruct := structs.GetJSON{}

	if err := c.Bind(&dataStruct); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println(dataStruct)

	data, err := os.ReadFile(dataStruct.Path)
	if err != nil {
		return c.String(http.StatusBadRequest, "Ошибка чтения файла")
	}

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return c.String(http.StatusBadRequest, "Ошибка парсинга JSON")
	}

	return c.JSON(http.StatusOK, jsonData)
}

func getTimetable(c echo.Context) error {
	data := structs.GetTimetable{}

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	day, week, err := timetable.FormatingDate(data.Day)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result, err := timetable.GetTimetableText(week, day, data.Class)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[interface{}]interface{}{"day": day, "week": week, "timetable": result})
}

func getTimetableTeacher(c echo.Context) error {
	data := structs.GetTimetableTeacher{}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	day, week, err := timetable.FormatingDate(data.Day)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result, err := timetable.GetTimetableTeachersText(data.Name, week, day)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[interface{}]interface{}{"day": day, "week": week, "timetable": result})
}

func getStages(c echo.Context) error {
	return c.JSON(http.StatusOK, map[interface{}]interface{}{"classes": timetable.Stages})
}

func getTrackerData(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"subjects": timetable.SubjectsForButton,
		"stages":   timetable.StagesTracker,
		"teacher":  timetable.TeacherTracker,
		"olimps":   timetable.TrackerOlimps,
	})
}

func getTeachers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"teachers": timetable.Teachers})
}
