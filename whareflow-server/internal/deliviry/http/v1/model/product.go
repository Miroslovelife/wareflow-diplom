package delivery

type ProductModelRequest struct {
	Title       string `json:"title"`
	Count       uint64 `json:"count"`
	Description string `json:"description"`
	ZoneId uint64 `json:"zone_id"`
}

type ProductModelResponse struct {
	Uuid        string `json:"uuid"`
	Title       string `json:"title"`
	Count       uint64 `json:"count"`
	QrImage     string `json:"qr_path"`
	Description string `json:"description"`
	ZoneId      uint64 `json:"zone_id"`
}
