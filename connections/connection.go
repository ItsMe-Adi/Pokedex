package connections

import (
	"fmt"
	"io"
	"os"

	"pokedex/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GinSetup() {

	setupLogOutput()
	server := gin.Default()
	fmt.Println("GIN connected Successfully")
	server.Static("templates/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	routes.InitializeApiRoutes(server)
	routes.InitializeViewRoutes(server)
	server.Run() // listen and serve on 0.0.0.0:8080
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func DbConnection() *gorm.DB {
	dsn := "root:`DB-PASSWORD`@tcp(127.0.0.1:3306)/anime?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 100,
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("DB connected Successfully")
	return db
}
