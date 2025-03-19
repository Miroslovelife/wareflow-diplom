package domain

type Permission struct {
	Id   uint `gorm:"primary_key"`
	Name string
}

type Role struct {
	Id   uint `gorm:"primary_key"`
	Name string
}

type RolePermission struct {
	RoleId       uint `gorm:"column:role_id"`
	PermissionId uint `gorm:"column:permission_id"`
}

type WarehouseUserRole struct {
	WareHouseId uint   `gorm:"column:ware_house_id"`
	UserUuid    string `gorm:"column:user_id"`
	RoleId      uint   `gorm:"column:role_id"`
}
