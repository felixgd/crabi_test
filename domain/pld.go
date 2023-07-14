package domain

type PLD struct {
	IsBlacklisted bool `json:"is_in_blacklist,omitempty" bson:"is_in_blacklist,omitempty"`
}

type PLDPayload struct {
	Email     string `json:"email" bson:"email"`
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
}
