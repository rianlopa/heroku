package main

import (
	"log"
	"fmt"
	"net/http"
	"os"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

func main() {
	
	db, err := sql.Open("postgres", os.Getenv("postgres://intvoldyplkcsz:3ceecdb0d7313dc8a2649b0ad2bc9f4cea69987b0f6a90cd81f902999723e217@ec2-107-20-177-161.compute-1.amazonaws.com:5432/d3ff2t14i54049"))
  		if err != nil {
    		log.Fatal(err)
	}
	
	defer db.Close()
	
	var name = "Jean Wong"

	qr, err := db.Query(`INSERT INTO Public."Name"(nombres)
		VALUES ($1);`, name)
	
	fmt.Printf("Query Result: %v\n", qr)
	
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
