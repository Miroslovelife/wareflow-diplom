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
	RoleId       uint
	PermissionId uint
}

type WareHouseUserRole struct {
	WareHouseId uint
	UserUuid    string
	RoleId      uint
}
