package auth

const (
	PermAccess = 0x1
	PermUpdate = 0x2
	PermDelete = 0x4
	PermAppend = 0x8
	PermAll    = 0xF
)

//UserPerm desc
//@Struct UserPerm desc: User Permission table
//@Member (string) Authorized address
//@Member (int) Permission [1111]=>[append&delete&update&access]
type UserPerm struct {
	URI  string `json:"uri"`
	Perm int    `json:"permission"`
}

//isAccess desc
//@Method isAccess desc: Whether to grant access
//@Return (bool)
func (slf *UserPerm) isAccess() bool {

	if (slf.Perm & PermAccess) > 0 {
		return true
	}
	return false
}

//isUpdate desc
//@Method isUpdate desc : Whether authorization can be updated
//@Return (bool)
func (slf *UserPerm) isUpdate() bool {
	if (slf.Perm & PermUpdate) > 0 {
		return true
	}
	return false
}

//isDelete desc
//@Method isDelete desc: Whether authorization can be deleted
//@Return (bool)
func (slf *UserPerm) isDelete() bool {
	if (slf.Perm & PermDelete) > 0 {
		return true
	}
	return false
}

//isAppend desc
//@Method isAppend desc: Whether authorization can be added
//@Return (bool)
func (slf *UserPerm) isAppend() bool {
	if (slf.Perm & PermAppend) > 0 {
		return true
	}
	return false
}

//User desc
//@Struct User Claims: User Online data
//@Member
type User struct {
	Key       string     `json:"key"`
	Name      string     `json:"name"`
	LoginTime int        `json:"logintime"`
	Perm      []UserPerm `json:"perms"`
}

//IsAccess desc
//@Method IsAccess desc: Whether to grant access
//@Param  (string) Inspection path
//@Return (bool)
func (slf *User) IsAccess(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}

	return perm.isAccess()
}

//IsUpdate desc
//@Method IsUpdate desc: Whether to grant update
//@Param  (string) Inspection path
//@Return (bool)
func (slf *User) IsUpdate(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}
	return perm.isUpdate()
}

//IsDelete desc
//@Method IsDelete desc: Whether to grant delete
//@Param  (string) Inspection path
//@Return (bool)
func (slf *User) IsDelete(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}
	return perm.isDelete()
}

//IsAppend desc
//@Method IsAppend desc: Whether to grant append
//@Param  (string) Inspection path
//@Return (bool)
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
