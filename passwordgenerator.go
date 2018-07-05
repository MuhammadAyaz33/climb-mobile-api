package main

import (
	"favorites"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.POST("/addcomments", favorites.AddComments)

	e.Logger.Fatal(e.Start(":8082"))
	fmt.Println("start...")
	// p := "12345678"
	// hash := hashAndSalt([]byte(p))
	// fmt.Println(hash, "/n")
	//notification.AddMentorCreatContributionHistory("5b20dc52379bad4f4023dcaee")
}

//encrypt password
func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
