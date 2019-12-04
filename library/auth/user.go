package auth

const (
	//PermAccess access
	PermAccess = 0x1
	//PermUpdate update
	PermUpdate = 0x2
	//PermDelete delete
	PermDelete = 0x4
	//PermAppend delete
	PermAppend = 0x8
	//PermAll    all
	PermAll = 0xF
)

//UserPerm doc
//@Struct UserPerm @Summary User Permission table
//@Member (string) Authorized address
//@Member (int) Permission [1111]=>[append&delete&update&access]
type UserPerm struct {
	URI  string `xml:"uri" yaml:"uri" json:"uri"`
	Perm int    `xml:"permission" yaml:"permission" json:"permission"`
}

//isAccess doc
//@Method isAccess @Summary Whether to grant access
//@Return (bool)
func (slf *UserPerm) isAccess() bool {

	if (slf.Perm & PermAccess) > 0 {
		return true
	}
	return false
}

//isUpdate doc
//@Method isUpdate doc : Whether authorization can be updated
//@Return (bool)
func (slf *UserPerm) isUpdate() bool {
	if (slf.Perm & PermUpdate) > 0 {
		return true
	}
	return false
}

//isDelete doc
//@Method isDelete @Summary Whether authorization can be deleted
//@Return (bool)
func (slf *UserPerm) isDelete() bool {
	if (slf.Perm & PermDelete) > 0 {
		return true
	}
	return false
}

//isAppend doc
//@Method isAppend @Summary Whether authorization can be added
//@Return (bool)
func (slf *UserPerm) isAppend() bool {
	if (slf.Perm & PermAppend) > 0 {
		return true
	}
	return false
}

//User doc
//@Struct User Claims: User Online data
//@Member
type User struct {
	Key       string     `xml:"key" yaml:"key" json:"key"`
	Name      string     `xml:"name" yaml:"name" json:"name"`
	LoginTime int        `xml:"logintime" yaml:"logintime" json:"logintime"`
	Perm      []UserPerm `xml:"perms" yaml:"perms" json:"perms"`
}

//IsAccess doc
//@Method IsAccess @Summary Whether to grant access
//@Param  (string) Inspection path
//@Return (bool)
func (slf *User) IsAccess(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}

	return perm.isAccess()
}

//IsUpdate doc
//@Method IsUpdate @Summary Whether to grant update
//@Param  (string) Inspection path
//@Return (bool)
func (slf *User) IsUpdate(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}
	return perm.isUpdate()
}

//IsDelete doc
//@Method IsDelete @Summary Whether to grant delete
//@Param  (string) Inspection path
//@Return (bool)
func (slf *User) IsDelete(URI string) bool {
	perm := slf.getPerm(URI)
	if perm == nil {
		return true
	}
	return perm.isDelete()
}

//IsAppend doc
//@Method IsAppend @Summary Whether to grant append
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
