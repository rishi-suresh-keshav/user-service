package model

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

//var db *gorm.DB
//
//func init() {
//	config.Connect()
//	db = config.GetDB()
//	//db.AutoMigrate(&Book{})
//}

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