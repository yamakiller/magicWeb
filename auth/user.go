package auth

const (
	permAccess = 0x1
	permUpdate = 0x2
	permDelete = 0x4
	permAppend = 0x8
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
	if (slf.Perm & permAccess) > 0 {
		return true
	}
	return false
}

//isUpdate desc
//@method isUpdate desc : Whether authorization can be updated
//@return (bool)
func (slf *UserPerm) isUpdate() bool {
	if (slf.Perm & permUpdate) > 0 {
		return true
	}
	return false
}

//isDelete desc
//@method isDelete desc: Whether authorization can be deleted
//@return (bool)
func (slf *UserPerm) isDelete() bool {
	if (slf.Perm & permDelete) > 0 {
		return true
	}
	return false
}

//isAppend desc
//@method isAppend desc: Whether authorization can be added
//@return (bool)
func (slf *UserPerm) isAppend() bool {
	if (slf.Perm & permAppend) > 0 {
		return true
	}
	return false
}

//User desc
//@struct User desc: User Online data
//@member
type User struct {
	Key       string     `json:"key"`
	Name      string     `json:"name"`
	LoginTime int        `json:"logintime"`
	LastTime  int        `json:"lasttimie"`
	Perm      []UserPerm `json:"auth"`
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
