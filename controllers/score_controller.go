package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Rajesh-0907/go-gin-boilerplate/models"

	"github.com/supabase-community/postgrest-go"
	"github.com/supabase-community/supabase-go"

	"github.com/gin-gonic/gin"
)

type AnswerRequest struct {
	Answers []string `json:"answers"`
}
type Student struct {
	Score int `json:"email"`
}

func GetTopScore(c *gin.Context) {
	client, _ := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), &supabase.ClientOptions{})
	resp, _, err := client.
		From("students").
		Select("rollno, name, score", "", false).
		Order("score", &postgrest.OrderOpts{
			Ascending: false,
		}). // order by score DESC
		Limit(10, "").
		Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	var students []models.TopScore
	if err := json.Unmarshal(resp, &students); err != nil {
		c.JSON(http.StatusBadRequest, err)

	}

	c.JSON(http.StatusOK, students)
}

func GetPpeScore(c *gin.Context) {
	rollnoVal, _ := c.Get("rollno")
	rollno, _ := rollnoVal.(string)
	var req AnswerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}
	correctAnswers := []string{
		"150ksc", "45ksc", "38ksc", "hardness", "0.4%", "increases", "Temperature", "Temperature increases",
	}

	score := 0
	for i, ans := range correctAnswers {
		if req.Answers[i] == ans {
			score++
		}
	}
	updatedData := map[string]interface{}{
		"score":       score,
		"issubmitted": true,
	}

	client, _ := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), &supabase.ClientOptions{})
	// client.From("students").Update()
	_, _, err := client.From("students").Update(updatedData, "exact", "").Eq("rollno", rollno).Execute()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, score)
}
