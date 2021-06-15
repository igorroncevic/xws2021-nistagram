package rbac

import "gorm.io/gorm"

func SetupUsersRBAC(db *gorm.DB) error {
	dropUsersTables(db)
	err := db.AutoMigrate(&UserRole{}, Permission{}, RolePermission{})
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		userRoles := []UserRole{basic, admin, verified, agent, nonregistered}
		result := db.Create(&userRoles)
		if result.Error != nil {
			return result.Error
		}

		permissions := []Permission{
			getUserById, getUsernameById, getAllUsers,
			updateUserPassword, updateUserProfile, searchUser,
			updatePrivacy, blockUser, unBlockUser, checkIfBlocked, checkUserProfilePublic, getAllPublicUsers, approveAccount, submitVerificationRequest,
		}
		result = db.Create(&permissions)
		if result.Error != nil {
			return result.Error
		}

		rolePermissions := []RolePermission{
			basicGetUserById, agentGetUserById, adminGetUserById, verifiedGetUserById, nonregisteredGetUserById,
			basicGetUsernameById, adminGetUsernameById, agentGetUsernameById, verifiedGetUsernameById, nonregisteredGetUsernameById,
			basicGetAllUsers, agentGetAllUsers, adminGetAllUsers, verifiedGetAllUsers, nonregisteredGetAllUsers,
			basicUpdateUserPassword, agentUpdateUserPassword, verifiedUpdateUserPassword,
			basicUpdateUserProfile, agentUpdateUserProfile, verifiedUpdateUserProfile,
			basicSearchUser, agentSearchUser, adminSearchUser, verifiedSearchUser, nonregisteredSearchUser,
			basicUpdatePrivacy, agentUpdatePrivacy, verifiedUpdatePrivacy,
			basicBlockUser, agentBlockUser, verifiedBlockUser,
			basicUnBlockUser, agentUnBlockUser, verifiedUnBlockUser,
			basicCheckIfBlocked, agentCheckIfBlocked, adminCheckIfBlocked, verifiedCheckIfBlocked, nonregisteredCheckIfBlocked,
			basicCheckUserProfilePublic, agentCheckUserProfilePublic, adminCheckUserProfilePublic, verifiedCheckUserProfilePublic, nonregisteredCheckUserProfilePublic,
			basicGetAllPublicUsers, agentGetAllPublicUsers, adminGetAllPublicUsers, verifiedGetAllPublicUsers, nonregisteredGetAllPublicUsers,
			basicApproveAccount, adminApproveAccount, agentApproveAccount,
			basicSubmitVerificationRequest, agentSubmitVerificationRequest,
		}
		result = db.Create(&rolePermissions)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return err
}

func dropUsersTables(db *gorm.DB) {
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
	getUserById        = Permission{Id: "992d5bf5-3e7f-4c8e-a76a-ad8444c0b32e", Name: "GetUserById"}
	getUsernameById    = Permission{Id: "c9295278-8fe8-40e6-9c9d-75543d48fa94", Name: "GetUsernameById"}
	getAllUsers        = Permission{Id: "26e41719-e309-4591-bb7e-3291b59c6dd4", Name: "GetAllUsers"}
	updateUserProfile  = Permission{Id: "48719e11-38de-4935-a93a-a61886c9303e", Name: "UpdateUserProfile"}
	updateUserPassword = Permission{Id: "50db6a87-483e-4d97-b320-9ff68235280a", Name: "UpdateUserPassword"}
	searchUser         = Permission{Id: "afbbf68f-ad1d-45db-8345-37ab619eea33", Name: "SearchUser"}
	approveAccount     = Permission{Id: "15e9a950-8581-4aa6-81c1-ae722c527051", Name: "ApproveAccount"}

	updatePrivacy          = Permission{Id: "3ce13d71-30e2-4cca-8a48-8a5ee1b6a42e", Name: "UpdatePrivacy"}
	blockUser              = Permission{Id: "9ec3fb89-28d6-4789-82b8-f247706cb6a0", Name: "BlockUser"}
	unBlockUser            = Permission{Id: "bf4632b1-e3ae-41d5-a04a-4bac73b7a2ef", Name: "UnBlockUser"}
	checkIfBlocked         = Permission{Id: "ce7b4f42-02ce-4c92-bcc4-529972173a4b", Name: "CheckIfBlocked"}
	checkUserProfilePublic = Permission{Id: "f2d282a9-c171-47f4-935d-32875fa61c8a", Name: "CheckUserProfilePublic"}
	getAllPublicUsers      = Permission{Id: "a30c3350-4b8f-4773-bd90-32c17e88d221", Name: "GetAllPublicUsers"}

	submitVerificationRequest = Permission{Id: "1d867c15-595b-4a69-b8ad-7135457bc402", Name: "SubmitVerificationRequest"}
)

var (
	basicGetUserById         = RolePermission{RoleId: basic.Id, PermissionId: getUserById.Id}
	agentGetUserById         = RolePermission{RoleId: agent.Id, PermissionId: getUserById.Id}
	adminGetUserById         = RolePermission{RoleId: admin.Id, PermissionId: getUserById.Id}
	verifiedGetUserById      = RolePermission{RoleId: verified.Id, PermissionId: getUserById.Id}
	nonregisteredGetUserById = RolePermission{RoleId: nonregistered.Id, PermissionId: getUserById.Id}

	basicGetUsernameById         = RolePermission{RoleId: basic.Id, PermissionId: getUsernameById.Id}
	agentGetUsernameById         = RolePermission{RoleId: agent.Id, PermissionId: getUsernameById.Id}
	adminGetUsernameById         = RolePermission{RoleId: admin.Id, PermissionId: getUsernameById.Id}
	verifiedGetUsernameById      = RolePermission{RoleId: verified.Id, PermissionId: getUsernameById.Id}
	nonregisteredGetUsernameById = RolePermission{RoleId: nonregistered.Id, PermissionId: getUsernameById.Id}

	basicApproveAccount = RolePermission{RoleId: basic.Id, PermissionId: approveAccount.Id}
	agentApproveAccount = RolePermission{RoleId: agent.Id, PermissionId: approveAccount.Id}
	adminApproveAccount = RolePermission{RoleId: admin.Id, PermissionId: approveAccount.Id}

	basicUpdateUserProfile    = RolePermission{RoleId: basic.Id, PermissionId: updateUserProfile.Id}
	agentUpdateUserProfile    = RolePermission{RoleId: agent.Id, PermissionId: updateUserProfile.Id}
	verifiedUpdateUserProfile = RolePermission{RoleId: verified.Id, PermissionId: updateUserProfile.Id}

	basicUpdateUserPassword    = RolePermission{RoleId: basic.Id, PermissionId: updateUserPassword.Id}
	agentUpdateUserPassword    = RolePermission{RoleId: agent.Id, PermissionId: updateUserPassword.Id}
	verifiedUpdateUserPassword = RolePermission{RoleId: verified.Id, PermissionId: updateUserPassword.Id}

	basicGetAllUsers         = RolePermission{RoleId: basic.Id, PermissionId: getAllUsers.Id}
	agentGetAllUsers         = RolePermission{RoleId: agent.Id, PermissionId: getAllUsers.Id}
	adminGetAllUsers         = RolePermission{RoleId: admin.Id, PermissionId: getAllUsers.Id}
	verifiedGetAllUsers      = RolePermission{RoleId: verified.Id, PermissionId: getAllUsers.Id}
	nonregisteredGetAllUsers = RolePermission{RoleId: nonregistered.Id, PermissionId: getAllUsers.Id}

	basicSearchUser         = RolePermission{RoleId: basic.Id, PermissionId: searchUser.Id}
	agentSearchUser         = RolePermission{RoleId: agent.Id, PermissionId: searchUser.Id}
	adminSearchUser         = RolePermission{RoleId: admin.Id, PermissionId: searchUser.Id}
	verifiedSearchUser      = RolePermission{RoleId: verified.Id, PermissionId: searchUser.Id}
	nonregisteredSearchUser = RolePermission{RoleId: nonregistered.Id, PermissionId: searchUser.Id}

	basicUpdatePrivacy    = RolePermission{RoleId: basic.Id, PermissionId: updatePrivacy.Id}
	agentUpdatePrivacy    = RolePermission{RoleId: agent.Id, PermissionId: updatePrivacy.Id}
	verifiedUpdatePrivacy = RolePermission{RoleId: verified.Id, PermissionId: updatePrivacy.Id}

	basicBlockUser    = RolePermission{RoleId: basic.Id, PermissionId: blockUser.Id}
	agentBlockUser    = RolePermission{RoleId: agent.Id, PermissionId: blockUser.Id}
	verifiedBlockUser = RolePermission{RoleId: verified.Id, PermissionId: blockUser.Id}

	basicUnBlockUser    = RolePermission{RoleId: basic.Id, PermissionId: unBlockUser.Id}
	agentUnBlockUser    = RolePermission{RoleId: agent.Id, PermissionId: unBlockUser.Id}
	verifiedUnBlockUser = RolePermission{RoleId: verified.Id, PermissionId: unBlockUser.Id}

	basicCheckIfBlocked         = RolePermission{RoleId: basic.Id, PermissionId: checkIfBlocked.Id}
	agentCheckIfBlocked         = RolePermission{RoleId: agent.Id, PermissionId: checkIfBlocked.Id}
	adminCheckIfBlocked         = RolePermission{RoleId: admin.Id, PermissionId: checkIfBlocked.Id}
	verifiedCheckIfBlocked      = RolePermission{RoleId: verified.Id, PermissionId: checkIfBlocked.Id}
	nonregisteredCheckIfBlocked = RolePermission{RoleId: nonregistered.Id, PermissionId: checkIfBlocked.Id}

	basicCheckUserProfilePublic         = RolePermission{RoleId: basic.Id, PermissionId: checkUserProfilePublic.Id}
	agentCheckUserProfilePublic         = RolePermission{RoleId: agent.Id, PermissionId: checkUserProfilePublic.Id}
	adminCheckUserProfilePublic         = RolePermission{RoleId: admin.Id, PermissionId: checkUserProfilePublic.Id}
	verifiedCheckUserProfilePublic      = RolePermission{RoleId: verified.Id, PermissionId: checkUserProfilePublic.Id}
	nonregisteredCheckUserProfilePublic = RolePermission{RoleId: nonregistered.Id, PermissionId: checkUserProfilePublic.Id}

	basicGetAllPublicUsers         = RolePermission{RoleId: basic.Id, PermissionId: getAllPublicUsers.Id}
	agentGetAllPublicUsers         = RolePermission{RoleId: agent.Id, PermissionId: getAllPublicUsers.Id}
	adminGetAllPublicUsers         = RolePermission{RoleId: admin.Id, PermissionId: getAllPublicUsers.Id}
	verifiedGetAllPublicUsers      = RolePermission{RoleId: verified.Id, PermissionId: getAllPublicUsers.Id}
	nonregisteredGetAllPublicUsers = RolePermission{RoleId: nonregistered.Id, PermissionId: getAllPublicUsers.Id}

	basicSubmitVerificationRequest = RolePermission{RoleId: basic.Id, PermissionId: submitVerificationRequest.Id}
	agentSubmitVerificationRequest = RolePermission{RoleId: agent.Id, PermissionId: submitVerificationRequest.Id}
)
