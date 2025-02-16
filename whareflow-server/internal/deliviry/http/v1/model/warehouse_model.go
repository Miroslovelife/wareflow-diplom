package delivery

type WarehouseModelRequest struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type WarehouseModelResponse struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}
