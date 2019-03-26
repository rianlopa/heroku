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
	
	db, err := sql.Open("postgres", os.Getenv("postgresql-infinite-10767"))
  		if err != nil {
    		log.Fatal(err)
	}
	
	defer db.Close()
	
	qr, err := db.Query(`INSERT INTO public."Name"(nombres) VALUES ('Jean');`)
	
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
