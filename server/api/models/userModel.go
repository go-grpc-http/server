package models

type User struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	UserName  string `json:"userName" bson:"userName"`
	Password  string `json:"password" bson:"password"`
	Email     string `json:"email" bson:"email"`
	Address   string `json:"address" bson:"address"`
	Mobile    int    `json:"mobile" bson:"mobile"`
	Id        string `json:"id" bson:"id"`
}

var UserList []User

type Login struct {
	UserName string `json:"userName" bson:"userName"`
	Password string `json:"password" bson:"password"`
}
