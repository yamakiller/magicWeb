package auth

import "strings"

//VerifyPerm doc
//method VerifyPerm @Summary Verify permissions
//param ([]UserPerm) permissions array
//param (string)     access uri
//param (int)        need permissions
//return (bool)      yes/no
func VerifyPerm(perms []UserPerm, uri string, need int) bool {
	for _, v := range perms {
		if strings.ToLower(v.URI) == "all" && (v.Perm&need) != 0 {
			return true
		}

		if strings.Index(uri, v.URI) >= 0 && (v.Perm&need) != 0 {
			return true
		}
	}

	return false
}
