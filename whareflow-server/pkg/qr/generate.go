package qr

import "github.com/skip2/go-qrcode"

type Generator struct {
	gen qrcode.QRCode
}

// @Summary Reg
// @Description Reg new user
// @Tags auth
// @Router /auth/reg [post]

func (g *Generator) Generate() ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode("fff", qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return png, nil
}
