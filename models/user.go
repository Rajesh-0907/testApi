package models

type RegisterRequest struct {
	RollNo   string `json:"rollno"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	RollNo   string `json:"rollno"`
	Password string `json:"password"`
}

type SupabaseUser struct {
	RollNo      string `json:"rollno"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Score       int    `json:"score"`
	Isloggedin  bool   `json:"isloggedin"`
	Issubmitted bool   `json:"issubmitted"`
}
