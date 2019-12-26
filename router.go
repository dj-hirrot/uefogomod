package main

import(
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "uefogomod/handler"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/users", handler.ListUser)
    e.GET("/users/:id", handler.GetUser)
    e.POST("/users", handler.CreateUser)
    e.PUT("/users/:id", handler.UpdateUser)
    e.DELETE("/users/:id", handler.DeleteUser)

    e.Logger.Fatal(e.Start(":1324"))
}

