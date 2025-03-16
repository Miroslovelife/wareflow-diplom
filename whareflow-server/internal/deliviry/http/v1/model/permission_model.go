package delivery

type Permission struct {
	Uuid        string `json:"uuid"`
	WareHouseId uint   `json:"warehouse_id"`
	Action      string `json:"action"`
}
