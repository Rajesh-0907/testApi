package controllers

import (
	"net/http"

	"github.com/Rajesh-0907/go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"
)

func GetPpeQuestions(c *gin.Context) {
	questions := []models.Question{
		{
			ID:       1,
			Question: "what is the pressure at msv?",
			Options:  []string{"150ksc", "160ksc", "170ksc", "180ksc"},
			InputID:  []string{"option1a", "option1b", "option1c", "option1d"},
		},
		{
			ID:       2,
			Question: "what is the pressure at CEP?",
			Options:  []string{"50ksc", "45ksc", "60ksc", "80ksc"},
			InputID:  []string{"option2a", "option2b", "option2c", "option2d"},
		},
		{
			ID:       3,
			Question: "what is the pressure at HRH?",
			Options:  []string{"50ksc", "45ksc", "36ksc", "38ksc"},
			InputID:  []string{"option3a", "option3b", "option3c", "option3d"},
		},
		{
			ID:       4,
			Question: "Steam should not penetrate Turbine blade. Which property of turbine material does this define?",
			Options:  []string{"malleability", "ductility", "brittle", "hardness"},
			InputID:  []string{"option4a", "option4b", "option4c", "option4d"},
		},
		{
			ID:       5,
			Question: "Carbon content present in low carbon steel?",
			Options:  []string{"0.6%", "0.4%", "0.2%", "0.1%"},
			InputID:  []string{"option5a", "option5b", "option5c", "option5d"},
		},
		{
			ID:       6,
			Question: "Ductility decreases as carbon content_______",
			Options:  []string{"increases", "decreases", "Both of the above", "None of the above"},
			InputID:  []string{"option6a", "option6b", "option6c", "option6d"},
		},
		{
			ID:       7,
			Question: "Viscosity is indirectly proportional to ________",
			Options:  []string{"Temperature", "Pressure", "Volume", "Flow"},
			InputID:  []string{"option7a", "option7b", "option7c", "option7d"},
		},
		{
			ID:       8,
			Question: "Vapour Pressure increases when _______",
			Options:  []string{"Pressure increases", "Pressure decreases", "Temperature increases", "Temperature decreases"},
			InputID:  []string{"option8a", "option8b", "option8c", "option8d"},
		},
	}

	c.JSON(http.StatusOK, questions)
}

func GetAptitudeQuestions(c *gin.Context) {
	questions := []models.Question{
		{
			ID:       1,
			Question: "P and Q are two positive integers such that PQ = 64. Which of the following cannot be the value of P + Q?",
			Options:  []string{"16", "20", "35", "65"},
			InputID:  []string{"option1a", "option1b", "option1c", "option1d"},
		},
		{
			ID:       2,
			Question: "What is the sum of squares of squares of the digits from 1 to 9",
			Options:  []string{"105", "260", "285", "385"},
			InputID:  []string{"option2a", "option2b", "option2c", "option2d"},
		},
		{
			ID:       3,
			Question: "what is the pressure at HRH?",
			Options:  []string{"50ksc", "45ksc", "36ksc", "38ksc"},
			InputID:  []string{"option3a", "option3b", "option3c", "option3d"},
		},
		{
			ID:       4,
			Question: "Steam should not penetrate Turbine blade. Which property of turbine material does this define?",
			Options:  []string{"malleability", "ductility", "brittle", "hardness"},
			InputID:  []string{"option4a", "option4b", "option4c", "option4d"},
		},
		{
			ID:       5,
			Question: "Carbon content present in low carbon steel?",
			Options:  []string{"0.6%", "0.4%", "0.2%", "0.1%"},
			InputID:  []string{"option5a", "option5b", "option5c", "option5d"},
		},
		{
			ID:       6,
			Question: "Ductility decreases as carbon content_______",
			Options:  []string{"increases", "decreases", "Both of the above", "None of the above"},
			InputID:  []string{"option6a", "option6b", "option6c", "option6d"},
		},
		{
			ID:       7,
			Question: "Viscosity is indirectly proportional to ________",
			Options:  []string{"Temperature", "Pressure", "Volume", "Flow"},
			InputID:  []string{"option7a", "option7b", "option7c", "option7d"},
		},
		{
			ID:       8,
			Question: "Vapour Pressure increases when _______",
			Options:  []string{"Pressure increases", "Pressure decreases", "Temperature increases", "Temperature decreases"},
			InputID:  []string{"option8a", "option8b", "option8c", "option8d"},
		},
	}

	c.JSON(http.StatusOK, questions)
}
