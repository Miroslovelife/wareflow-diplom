package delivery

type WarehouseModelRequest struct {
	Name    string `json:"id"`
	Address string `json:"address"`
}

type WarehouseModelResponse struct {
	Id      uint64 `json:"id"`
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Employer struct {
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
}
