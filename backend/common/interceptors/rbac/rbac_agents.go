package rbac

import "gorm.io/gorm"

func SetupAgentsRBAC(db *gorm.DB) error {
	dropAgentTables(db)
	err := db.AutoMigrate(&UserRole{}, Permission{}, RolePermission{})
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		userRoles := []UserRole{basic, agent, nonregistered}
		result := db.Create(&userRoles)
		if result.Error != nil {
			return result.Error
		}

		permissions := []Permission{
			createProduct,
		}
		result = db.Create(&permissions)
		if result.Error != nil {
			return result.Error
		}

		rolePermissions := []RolePermission{
			agentCreateProduct,
		}
		result = db.Create(&rolePermissions)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return err
}

func dropAgentTables(db *gorm.DB) {
	if db.Migrator().HasTable(&UserRole{}) {
		db.Migrator().DropTable(&UserRole{})
	}
	if db.Migrator().HasTable(&Permission{}) {
		db.Migrator().DropTable(&Permission{})
	}
	if db.Migrator().HasTable(&RolePermission{}) {
		db.Migrator().DropTable(&RolePermission{})
	}
}

var (
	createProduct = Permission{Id: "dc1e031f-c276-42de-87c6-786e358fc51e", Name: "CreateProduct"}
)

var (
	agentCreateProduct = RolePermission{RoleId: agent.Id, PermissionId: createProduct.Id}
)
