package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func MolarHandlerFunc(c echo.Context) error {
	fmt.Println(c.Param("a"))
	return nil
}

func BalanceHandlerFunc(c echo.Context) error {
	return nil
}
