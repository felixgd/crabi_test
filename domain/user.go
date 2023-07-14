package domain

type User struct {
	Email     string `json:"email" bson:"email" validate:"required,email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty" validate:"required"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty" validate:"required"`
	PLD       PLD    `json:"PLD,omitempty" bson:"PLD,omitempty"`
}
