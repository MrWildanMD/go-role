package options

import (
	"github.com/MrWildanMD/go-role/utils"
)

// RoleOption represents options when fetching roles.
type RoleOption struct {
	WithPermissions bool
	Pagination      *utils.Pagination
}
