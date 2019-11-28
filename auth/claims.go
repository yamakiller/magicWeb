package auth

import "github.com/dgrijalva/jwt-go"

const (
	permAccess = 0x1
	permUpdate = 0x2
	permDelete = 0x4
	permAppend = 0x8
)

//ClaimPerm desc
//@struct ClaimPerm desc: User Permission table
//@member (string) Authorized address
//@member (int) Permission [1111]=>[append&delete&update&access]
type ClaimPerm struct {
	URI  string `json:"uri"`
	Perm int    `json:"permission"`
}

//isAccess desc
//@method isAccess desc: Whether to grant access
//@return (bool)
func (slf *ClaimPerm) isAccess() bool {

	if (slf.Perm & permAccess) > 0 {
		return true
	}
	return false
}

//isUpdate desc
//@method isUpdate desc : Whether authorization can be updated
//@return (bool)
func (slf *ClaimPerm) isUpdate() bool {
	if (slf.Perm & permUpdate) > 0 {
		return true
	}
	return false
}

//isDelete desc
//@method isDelete desc: Whether authorization can be deleted
//@return (bool)
func (slf *ClaimPerm) isDelete() bool {
	if (slf.Perm & permDelete) > 0 {
		return true
	}
	return false
}

//isAppend desc
//@method isAppend desc: Whether authorization can be added
//@return (bool)
func (slf *ClaimPerm) isAppend() bool {
	if (slf.Perm & permAppend) > 0 {
		return true
	}
	return false
}

//Claims desc
//@struct User Claims: User Online data
//@member
type Claims struct {
	Key       string      `json:"key"`
	Name      string      `json:"name"`
	LoginTime int         `json:"logintime"`
	Perm      []ClaimPerm `json:"auth"`
	jwt.StandardClaims
}

//IsAccess desc
//@method IsAccess desc: Whether to grant access
//@param  (string) Inspection path
//@return (bool)
func (slf *Claims) IsAccess(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}

	return perm.isAccess()
}

//IsUpdate desc
//@method IsUpdate desc: Whether to grant update
//@param  (string) Inspection path
//@return (bool)
func (slf *Claims) IsUpdate(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}
	return perm.isUpdate()
}

//IsDelete desc
//@method IsDelete desc: Whether to grant delete
//@param  (string) Inspection path
//@return (bool)
func (slf *Claims) IsDelete(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}
	return perm.isDelete()
}

//IsAppend desc
//@method IsAppend desc: Whether to grant append
//@param  (string) Inspection path
//@return (bool)
func (slf *Claims) IsAppend(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}
	return perm.isAppend()
}

func (slf *Claims) getPerm(URI string) *ClaimPerm {
	for _, v := range slf.Perm {
		if v.URI == URI {
			return &v
		}
	}
	return nil
}
