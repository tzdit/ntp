package models2

import (
	"github.com/labstack/echo/v4"
	"teleradiology/backend/package/log"
	"time"
)

func GetDateField(c echo.Context) (time.Time, error) {
	receivedDate := c.Param("date")
	if receivedDate == "" {
		log.Errorf("error invalid date %v", receivedDate)
	}

	filterDate, err := time.Parse("2006-01-02", receivedDate)
	if err != nil {
		log.Errorf("error parsing date %v", receivedDate)
	}

	return filterDate, nil
}

func GetIdField(c echo.Context) (string, error) {
	receivedUuid := c.Param("id")
	if receivedUuid == "" {
		log.Errorf("error invalid id %v", receivedUuid)
	}
	return receivedUuid, nil
}

func GetTableField(c echo.Context) (string, error) {
	receivedField := c.Param("field")
	if receivedField == "" {
		log.Errorf("error invalid field %v", receivedField)
	}
	return receivedField, nil
}
