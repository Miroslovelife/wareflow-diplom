package delivery

// UserReg example
type UserReg struct {
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

type UserLoginByEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginByPhoneNumber struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
