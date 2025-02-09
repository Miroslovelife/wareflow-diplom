package domain

type User struct {
	Uuid        []byte `gorm:"table:users;column:uuid;primaryKey;default:gen_random_uuid()"`
	PhoneNumber string `gorm:"column:phone_number"`
	Username    string `gorm:"column:username"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	Surname     string `gorm:"column:surname"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	Role        string `gorm:"default:user"`
}
