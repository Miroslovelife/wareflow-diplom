package domain

type WareHouse struct {
	Id       uint64 `gorm:"primaryKey;autoIncrement:true;column:id"`
	Address  string `gorm:"column:address"`
	Name     string `gorm:"column:name"`
	UuidUser string `gorm:"column:uuid_user;foreignKey:UuidUser;references:Uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
