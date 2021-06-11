package rbac

type UserRole struct {
	Id   string `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type Permission struct {
	Id   string `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type RolePermission struct {
	RoleId       string `gorm:"primaryKey"`
	PermissionId string `gorm:"primaryKey"`
}