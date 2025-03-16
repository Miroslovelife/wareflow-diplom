package domain

type Product struct {
	Uuid        []byte `gorm:"table:products;column:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string `gorm:"column:title"`
	Count       uint64 `gorm:"column:count"`
	QrPath      string `gorm:"column:qr"`
	Description string `gorm:"column:description"`
	ZoneId      uint64 `gorm:"column:zone_id"`
}
