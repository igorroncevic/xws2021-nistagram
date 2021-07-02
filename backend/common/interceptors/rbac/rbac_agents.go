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
			createProduct, getUserByUsernameAgent, getAllProductsByAgentId, getAllProducts,
		}
		result = db.Create(&permissions)
		if result.Error != nil {
			return result.Error
		}

		rolePermissions := []RolePermission{
			agentCreateProduct,
			agentGetUserByUsernameAgent, basicGetUserByUsernameAgent,
			basicGetAllProductsByAgentId, agentGetAllProductsByAgentId,
			basicGetAllProducts, agentGetAllProducts,
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
	createProduct           = Permission{Id: "dc1e031f-c276-42de-87c6-786e358fc51e", Name: "CreateProduct"}
	getUserByUsernameAgent  = Permission{Id: "98ec230f-7995-4970-87be-a71e2efbcbd2", Name: "GetUserByUsername"}
	getAllProductsByAgentId = Permission{Id: "9f0095b3-0c0a-42c8-908a-dc05757a1453", Name: "GetAllProductsByAgentId"}
	getAllProducts          = Permission{Id: "d149a012-4b67-4525-82cb-66e9b2fccd54", Name: "GetAllProducts"}
)

var (
	agentCreateProduct = RolePermission{RoleId: agent.Id, PermissionId: createProduct.Id}

	agentGetUserByUsernameAgent = RolePermission{RoleId: agent.Id, PermissionId: getUserByUsernameAgent.Id}
	basicGetUserByUsernameAgent = RolePermission{RoleId: basic.Id, PermissionId: getUserByUsernameAgent.Id}

	agentGetAllProductsByAgentId = RolePermission{RoleId: agent.Id, PermissionId: getAllProductsByAgentId.Id}
	basicGetAllProductsByAgentId = RolePermission{RoleId: basic.Id, PermissionId: getAllProductsByAgentId.Id}

	basicGetAllProducts = RolePermission{RoleId: basic.Id, PermissionId: getAllProducts.Id}
	agentGetAllProducts = RolePermission{RoleId: agent.Id, PermissionId: getAllProducts.Id}
)
