package delivery

type JWTCustomClaims struct {
	Uuid []byte `json:"uuid"`
	Name string `json:"name"`
}

type JWTCustomRefreshClaims struct {
	Uuid []byte `json:"uuid"`
}
