package repository

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}


func GetUserByID(ID int64) *User {
	var user User
	db.Where("ID = ?", ID).Find(&user)
	return &user
}

func (user *User) CreateUser() *User {
	db.NewRecord(user)
	db.Create(&user)
	return user
}