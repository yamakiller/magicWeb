package auth

import "strings"

//VerifyAdminProfile doc
//Summary Verify permissions
//param ([]UserPerm) permissions array
//param (string)     access uri
//return (bool)      yes/no
func VerifyAdminProfile(profiles []AdminUserProfile, uri string) bool {
	for _, v := range profiles {
		if strings.ToLower(v.URI) == "all" {
			return true
		}

		if uri == v.URI {
			return true
		}
	}

	return false
}
