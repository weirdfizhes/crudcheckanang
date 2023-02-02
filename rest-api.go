package main

import (
	"fmt"
	"net/http"
	
	"github.com/labstack/echo"
)

type Users struct {
	Email string `json:"email"`
	Nama string	`json:"nama"`
	Umur int	`json:"umur"`
	Alamat string	`json:"alamat"`
}

func main() {
	 route := echo.New()
	 
	 route.POST("/user/create_user", func(c echo.Context) error {
		var user = new(Users)
		c.Bind(user)
		fmt.Println(user)

		response := struct{
			Message string
			Data Users
		}{
			Message : "sukses user baru" ,
			Data : *user,
		}
		return c.JSON(http.StatusOK, response)
	 })

	 route.PUT("/user/update_user/:email", func(c echo.Context) error {
		user := new(Users)
		user.Email = c.Param("email")
		//update db
		response := struct{
			Message string
			Data Users
		}{
			Message : "sukses update",
			Data : *user,
		}
		return c.JSON(http.StatusOK, response)	
	})
	route.DELETE("/user/delete_user/:email", func(c echo.Context) error {
		user := new(Users)
		user.Email = c.Param("email")
		//del data
		response := struct {
			Message string
			ID string
		}{
			Message : "sukses hapus",
			ID : user.Email,
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("/user", func(c echo.Context) error {
		user := new(Users)
		user.Email = "maulanarahmadi888@gmail.com"
		user.Nama = "anang"
		user.Umur = 22
		user.Alamat = "jatayu"

		response := struct {
			Message string
			Data Users

		}{
			Message : "sukses liat",
			Data: *user ,
		}
		return c.JSON (http.StatusOK, response)

	})
	route.Start(":1080")
}


