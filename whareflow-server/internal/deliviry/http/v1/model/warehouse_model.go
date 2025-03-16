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
