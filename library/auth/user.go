package auth

const (
	PermAccess = 0x1
	PermUpdate = 0x2
	PermDelete = 0x4
	PermAppend = 0x8
	PermAll    = 0xF
)

//UserPerm desc
//@struct UserPerm desc: User Permission table
//@member (string) Authorized address
//@member (int) Permission [1111]=>[append&delete&update&access]
type UserPerm struct {
	URI  string `json:"uri"`
	Perm int    `json:"permission"`
}

//isAccess desc
//@method isAccess desc: Whether to grant access
//@return (bool)
func (slf *UserPerm) isAccess() bool {

	if (slf.Perm & PermAccess) > 0 {
		return true
	}
	return false
}

//isUpdate desc
//@method isUpdate desc : Whether authorization can be updated
//@return (bool)
func (slf *UserPerm) isUpdate() bool {
	if (slf.Perm & PermUpdate) > 0 {
		return true
	}
	return false
}

//isDelete desc
//@method isDelete desc: Whether authorization can be deleted
//@return (bool)
func (slf *UserPerm) isDelete() bool {
	if (slf.Perm & PermDelete) > 0 {
		return true
	}
	return false
}

//isAppend desc
//@method isAppend desc: Whether authorization can be added
//@return (bool)
func (slf *UserPerm) isAppend() bool {
	if (slf.Perm & PermAppend) > 0 {
		return true
	}
	return false
}

//User desc
//@struct User Claims: User Online data
//@member
type User struct {
	Key       string     `json:"key"`
	Name      string     `json:"name"`
	LoginTime int        `json:"logintime"`
	Perm      []UserPerm `json:"perms"`
}

//IsAccess desc
//@method IsAccess desc: Whether to grant access
//@param  (string) Inspection path
//@return (bool)
func (slf *User) IsAccess(URI string) bool {
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
func (slf *User) IsUpdate(URI string) bool {
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
func (slf *User) IsDelete(URI string) bool {
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
func (slf *User) IsAppend(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}
	return perm.isAppend()
}

func (slf *User) getPerm(URI string) *UserPerm {
	for _, v := range slf.Perm {
		if v.URI == URI {
			return &v
		}
	}
	return nil
}
