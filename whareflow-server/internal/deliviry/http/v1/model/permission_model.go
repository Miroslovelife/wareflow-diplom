package delivery

type Permission struct {
	Uuid        string `json:"uuid"`
	WareHouseId uint   `json:"warehouse_id"`
	Action      string `json:"action"`
}

type RoleResponse struct {
	Name        string
	Permissions []PermissionResponse
}

type PermissionResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type RoleReq struct {
	Name        string `json:"name"`
	UserName    string `json:"username"`
	Permissions []uint `json:"permissions"`
}
