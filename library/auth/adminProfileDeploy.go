package auth

var (
	defaultAdminProfile = &ConfigAdminUserProfileItems{}
)

//WithAdminProfileDeploy Set admin user profile deploy data
func WithAdminProfileDeploy(config *ConfigAdminUserProfileItems) {
	defaultAdminProfile = config
}

//GetAdminProfileDeploy Returns admin user profile deploy data
func GetAdminProfileDeploy() *ConfigAdminUserProfileItems {
	return defaultAdminProfile
}
