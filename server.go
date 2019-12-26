package main

import(
    "net/http"
    "github.com/labstack/echo"
)

type User struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Email string `json:email`
}

func main() {
    e := echo.New()
    e.GET("/", hello)
    e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello")
}
