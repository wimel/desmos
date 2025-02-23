package types

// Subspaces module event types
const (
	EventTypeCreateSubspace              = "create_subspace"
	EventTypeEditSubspace                = "edit_subspace"
	EventTypeDeleteSubspace              = "delete_subspace"
	EventTypeCreateSection               = "create_section"
	EventTypeEditSection                 = "edit_section"
	EventTypeMoveSection                 = "move_section"
	EventTypeDeleteSection               = "delete_section"
	EventTypeCreateUserGroup             = "create_user_group"
	EventTypeEditUserGroup               = "edit_user_group"
	EvenTypeMoveUserGroup                = "move_user_group"
	EventTypeSetUserGroupPermissions     = "set_user_group_permissions"
	EventTypeDeleteUserGroup             = "delete_user_group"
	EventTypeAddUserToGroup              = "add_group_member"
	EventTypeRemoveUserFromGroup         = "remove_group_member"
	EventTypeSetUserPermissions          = "set_user_permissions"
	EventTypeGrantTreasuryAuthorization  = "grant_treasury_authorization"
	EventTypeRevokeTreasuryAuthorization = "revoke_treasury_authorization"
	EventTypeGrantAllowance              = "grant_allowance"
	EventTypeRevokeAllowance             = "revoke_allowance"

	AttributeValueCategory      = ModuleName
	AttributeKeySubspaceID      = "subspace_id"
	AttributeKeySubspaceName    = "subspace_name"
	AttributeKeySubspaceCreator = "subspace_creator"
	AttributeKeyCreationTime    = "creation_date"
	AttributeKeySectionID       = "section_id"
	AttributeKeyUserGroupID     = "user_group_id"
	AttributeKeyUser            = "user"
	AttributeKeyGranter         = "granter"
	AttributeKeyGrantee         = "grantee"
	AttributeKeyUserGrantee     = "user_grantee"
	AttributeKeyGroupGrantee    = "group_grantee"
)
