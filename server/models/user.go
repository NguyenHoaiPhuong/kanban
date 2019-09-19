package models

// User account
type User struct {
	Model     `json:",inline" bson:",inline"`
	FirstName string `json:"FirstName" bson:"FirstName"`
	LastName  string `json:"LastName" bson:"LastName"`
	UserName  string `json:"UserName" bson:"UserName" required:"true"`
	Email     string `json:"Email" bson:"Email"`
	Password  string `json:"Password" bson:"Password" required:"true"`
}
