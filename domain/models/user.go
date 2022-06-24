package models

// User struct
type User struct {
	ID              string `json:"reference,omitempty" bson:"reference"`
	FirstName       string `json:"first_name,omitempty" bson:"first_name""`
	LastName        string `json:"last_name,omitempty" bson:"last_name""`
	Email           string `json:"email,omitempty" bson:"email"`
	Password        string `json:"-,omitempty"bson:"password"`
	HashedSecretKey string `json:"-,omitempty"bson:"hashed_secret_key"`
	DateOfBirth     string `json:"date_of_birth" bson:"date_of_birth"`
	CreatedAt       string `json:"created_at" bson:"created_at"`
}
