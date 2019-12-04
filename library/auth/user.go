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

//ConfigUserPerm doc
//@Summary user Permission config
//@Struct ConfigUserPerm
//@Member []UserPerm Arrays
type ConfigUserPerm struct {
	Data []UserPerm `xml:"config" yaml:"config" json:"config"`
}

//UserPerm doc
//@Summary User Permission table
//@Struct UserPerm
//@Member (string) Authorized address
//@Member (int) Permission [1111]=>[append&delete&update&access]
type UserPerm struct {
	URI  string `xml:"uri" yaml:"uri" json:"uri"`
	Perm int    `xml:"permission" yaml:"permission" json:"permission"`
}

//isAccess doc
//@Summary Whether to grant access
//@Method isAccess
//@Return (bool)
func (slf *UserPerm) isAccess() bool {

	if (slf.Perm & PermAccess) > 0 {
		return true
	}
	return false
}

//isUpdate doc
//@Whether authorization can be updated
//@Method isUpdate
//@Return (bool)
func (slf *UserPerm) isUpdate() bool {
	if (slf.Perm & PermUpdate) > 0 {
		return true
	}
	return false
}

//isDelete doc
//@Summary Whether authorization can be deleted
//@Method isDelete
//@Return (bool)
func (slf *UserPerm) isDelete() bool {
	if (slf.Perm & PermDelete) > 0 {
		return true
	}
	return false
}

//isAppend doc
//@Summary Whether authorization can be added
//@Method isAppend
//@Return (bool)
func (slf *UserPerm) isAppend() bool {
	if (slf.Perm & PermAppend) > 0 {
		return true
	}
	return false
}

//User doc
//@Summary User Online data
//@Struct User Claims
//@Member
type User struct {
	Key       string     `xml:"key" yaml:"key" json:"key"`
	Name      string     `xml:"name" yaml:"name" json:"name"`
	LoginTime int        `xml:"logintime" yaml:"logintime" json:"logintime"`
	Perm      []UserPerm `xml:"perms" yaml:"perms" json:"perms"`
}

//IsAccess doc
//@Summary Whether to grant access
//@Method IsAccess
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
//@Summary Whether to grant update
//@Method IsUpdate
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
//@Summary Whether to grant delete
//@Method IsDelete
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
//@Summary Whether to grant append
//@Method IsAppend
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
