package domain

type User struct {
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	PLD       PLD    `json:"PLD,omitempty" bson:"PLD,omitempty"`
}
