package main

import (
	"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/kernelgarden/shallwe/config"
	"github.com/labstack/echo/middleware"
	"github.com/kernelgarden/shallwe/router"
)

func main() {

	db, err := initDB("mysql", fmt.Sprintf("root:tabstorage@/%s?charset=utf8", config.DBName))
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	router.InitRoutes(e)

	e.Static("/static", "static")
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	if err := e.Start(":3030"); err != nil {
		log.Println(err)
	}
}

func initDB(driver, connection string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(driver, connection)

	if err != nil {
		return nil, err
	}

	// TODO: sync db with model

	return db, nil
}