package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnswerRequest struct {
	Answers []string `json:"answers"`
}

func GetPpeScore(c *gin.Context) {
	// var answers []string

	// if err := c.Bind(&answers); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 	return
	// }

	// body, _ := io.ReadAll(c.Request.Body)
	// fmt.Println("Raw body:", string(body))
	var req AnswerRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}
	// // 🧮 Example: Count how many answers match "correct" ones
	correctAnswers := []string{
		"150ksc", "45ksc", "38ksc", "hardness", "0.4%", "increases", "Temperature", "Temperature increases",
	}

	score := 0
	for i, ans := range correctAnswers {
		if req.Answers[i] == ans {
			score++
		}
	}

	c.JSON(http.StatusOK, score)
}
