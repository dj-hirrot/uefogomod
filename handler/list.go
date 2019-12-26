package handler

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "uefogomod/model"
)

func GetUser(c echo.Context) (err error) {
    id, _ := strconv.Atoi(c.Param("id"))
    return c.JSON(http.StatusOK, model.Users[id])
}

func ListUser(c echo.Context) (err error) {
    return c.JSON(http.StatusOK, model.Users)
}

