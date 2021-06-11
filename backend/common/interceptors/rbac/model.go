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

const (
	Basic 		  = "Basic"
	Admin 		  = "Admin"
	Verified 	  = "Verified"
	Agent 		  = "Agent"
	Nonregistered = "Nonregistered"
)

var (
	basic 		  = UserRole{ Id: "6bef2e66-0016-4a54-bfb6-6c916288bfb9", Name: Basic}
	admin 		  = UserRole{ Id: "0c4ebbce-3bde-431f-ad52-615143100caa", Name: Admin}
	verified 	  = UserRole{ Id: "d0ed9b15-ae57-4301-baad-95e0a47ec324", Name: Verified}
	agent 		  = UserRole{ Id: "85f45393-853e-4acb-b77e-f611056bd3df", Name: Agent}
	nonregistered = UserRole{ Id: "80de2f24-ce8f-4d83-81aa-48ced4421cfb", Name: Nonregistered}
)