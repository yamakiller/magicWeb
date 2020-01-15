package auth

var (
	defaultAdminUserRoleTable = &AdminUserRoleTable{}
)

//AdminUserRole admin user role config/data
type AdminUserRole struct {
	Name         string `xml:"name" yaml:"name" json:"name"`
	Introduction string `xml:"introduction" yaml:"introduction" json:"introduction"`
}

type AdminUserRoleTable struct {
	Roles []AdminUserRole `xml:"roles" yaml:"roles" json:"roles"`
}

func WithAdminUserRoleTable(table *AdminUserRoleTable) {
	defaultAdminUserRoleTable = table
}

func GetAdminUserRoleTable() *AdminUserRoleTable {
	return defaultAdminUserRoleTable
}
