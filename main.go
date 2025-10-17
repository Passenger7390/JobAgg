package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Passenger7390/JobAgg/internal/jsearch"
	"github.com/Passenger7390/JobAgg/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()
	router.Use(cors.Default())
	jsclient := jsearch.NewClient(
		"https://jsearch.p.rapidapi.com/", 
		os.Getenv("JSEARCH_API"), 
		os.Getenv("JSEARCH_API_HOST"), 
		"job_title,employer_name,employer_website,job_publisher,job_employment_type,job_apply_link,apply_options,job_description,job_is_remote,job_posted_at_timestamp,job_city,job_state,job_country,job_min_salary,job_max_salary,job_salary_period", 
		nil)
		
	router.POST("/jobs", func(c *gin.Context) {
		var params models.SearchJobParams
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("SearchJob called with params:\n", params)
		_, res, err := jsclient.SearchJob(context.Background(), params)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, res)
	})

	router.Run()
}