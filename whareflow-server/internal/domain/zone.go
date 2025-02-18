package domain

type Zone struct {
	Id          int `gorm:"primaryKey;autoIncrement:true;column:id"`
	Name        string
	Capacity    int
	WarehouseId int `gorm:"column:ware_house_id"`
}
