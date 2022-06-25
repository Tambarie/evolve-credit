package models

type UpdateUserInfo struct {
	ID          string `json:"id,omitempty" bson:"id"`
	FirstName   string `json:"first_name,omitempty" bson:"first_name""`
	LastName    string `json:"last_name,omitempty" bson:"last_name""`
	Email       string `json:"email,omitempty" bson:"email"`
	Password    string `json:"-,omitempty"bson:"password"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
	UpdateAt    string `json:"update_at"bson:"update_at"`
}
