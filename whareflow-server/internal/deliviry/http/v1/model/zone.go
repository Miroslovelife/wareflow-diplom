package delivery

type ZoneModelRequest struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type ZoneModelResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}
