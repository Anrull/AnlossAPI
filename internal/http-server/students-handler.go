package http_server

import (
	"AnlossAPI/internal/http-server/structs"
	"AnlossAPI/internal/storage/sqlite"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func checkSnils(c echo.Context) error {
	student := structs.CheckSnils{}

	if err := c.Bind(&student); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	flag, name, stage := sqlite.CheckSnils(student.Snils, false)

	if flag {
		return c.String(http.StatusOK, fmt.Sprintf(`{"name": "%s", "stage": "%s"}`, name, stage))
	}

	return c.JSON(http.StatusNoContent, fmt.Sprintf(`{"message": "no student with this (%s) hashing snils"}`, student.Snils))
}

func addStudent(c echo.Context) error {
	student := structs.AddStudent{}

	if err := c.Bind(&student); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println(student)

	if student.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "student name is required")
	}
	if student.Class == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "student class is required")
	}
	if student.Snils == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "student snils is required")
	}

	if err := sqlite.AddStudent(student.Name, student.Class, student.Snils); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, `{"message": "student added"}`)
}
