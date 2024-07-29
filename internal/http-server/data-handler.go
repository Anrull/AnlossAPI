package http_server

import (
	"AnlossAPI/internal/http-server/structs"
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
