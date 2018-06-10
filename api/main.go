package main

import (
	"fmt"
	"inandout/api/controller"
	"os"
	"runtime"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	db, err := dbConnect()
	if err != nil {
		fmt.Println("dbConnect error", err)
		os.Exit(1)
	}
	defer db.Close()
	fmt.Println("DB Connected")

	e := echo.New()

	e.HideBanner = true
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	controller.InitController(e, db)
	initServer(e)
}

func dbConnect() (db *gorm.DB, err error) {
	dbURL := fmt.Sprintf("root:%s@tcp(localhost:3306)/inandout?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_PASS"))
	db, err = gorm.Open("mysql", dbURL)

	if err != nil {
		return nil, err
	}

	return db, err
}
