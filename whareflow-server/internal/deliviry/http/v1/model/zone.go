package delivery

type ZoneModelRequest struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	WarehouseId int    `json:"warehouse_id"`
}
