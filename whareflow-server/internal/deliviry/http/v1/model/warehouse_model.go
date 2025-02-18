package delivery

type WarehouseModelRequest struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type WarehouseModelResponse struct {
	Id      uint64 `json:"id"`
	Address string `json:"address"`
	Name    string `json:"name"`
}
