package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Rajesh-0907/go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/supabase-go"
	"golang.org/x/crypto/bcrypt"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func LogoutHandler(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
	})
	c.JSON(200, gin.H{"message": "Logged out"})
}

func GetUserInfo(c *gin.Context) {
	rollnoVal, _ := c.Get("rollno")
	rollno, _ := rollnoVal.(string)
	client, _ := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), &supabase.ClientOptions{})
	data, _, err := client.From("students").Select("*", "exact", false).
		Eq("rollno", rollno).
		Single().
		Execute()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var student models.SupabaseUser
	json.Unmarshal(data, &student)
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"name":        student.Name,
			"isloggedin":  student.Isloggedin,
			"issubmitted": student.Issubmitted,
			"score":       student.Score,
		},
	})
}

func RegisterUserHandler(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// ✅ Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// ✅ Insert into Supabase table (e.g., 'students')
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), &supabase.ClientOptions{})

	user := map[string]interface{}{
		"rollno":   req.RollNo,
		"name":     req.Name,
		"password": string(hashedPassword),
	}

	// ⚠️ Ensure 'rollno' is set as UNIQUE in Supabase schema
	client.From("students").Insert(user, false, "", "exact", "").Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}

func LoginUserHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	// Create Supabase client
	supaURL := os.Getenv("SUPABASE_URL")
	supaKey := os.Getenv("SUPABASE_KEY")
	client, err := supabase.NewClient(supaURL, supaKey, nil)
	if err != nil {
		log.Fatal("Supabase client creation failed:", err)
	}

	// Query student by roll number
	// var users []map[string]interface{}
	data, count, err := client.From("students").Select("*", "exact", false).
		Eq("rollno", req.RollNo).
		Single().
		Execute()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	var student models.SupabaseUser
	json.Unmarshal(data, &student)
	fmt.Println(count)

	// Check password
	hashedPassword := student.Password
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"rollno": student.RollNo,
		"name":   student.Name,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 24h expiry
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	// Set HttpOnly cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // true in production (HTTPS only)
		SameSite: http.SameSiteNoneMode,
		MaxAge:   86400,
	})
	// Successful login
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"name":        student.Name,
			"isloggedin":  student.Isloggedin,
			"issubmitted": student.Issubmitted,
		},
	})
}
