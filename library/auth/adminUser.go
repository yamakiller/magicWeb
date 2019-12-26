package auth

const (
	//ProfileAccess access
	ProfileAccess = 0x1
	//ProfileUpdate update
	ProfileUpdate = 0x2
	//ProfileDelete delete
	ProfileDelete = 0x4
	//ProfileAppend delete
	ProfileAppend = 0x8
	//ProfileAll    all
	ProfileAll = 0xF
)

//ConfigAdminUserProfile doc
//@Summary user Permission config
//@Member []UserPerm Arrays
type ConfigAdminUserProfile struct {
	Items []AdminUserProfile `xml:"items" yaml:"items" json:"items"`
}

//AdminUserProfileItems doc
//@Summary user Permission config
//@Member []UserPerm Arrays
type AdminUserProfileItems struct {
	Items []AdminUserProfile `xml:"items" yaml:"items" json:"items"`
}

//AdminUserProfile doc
//@Summary User Permission table
//@Member (string) Authorized address
//@Member (int) Permission [1111]=>[append&delete&update&access]
type AdminUserProfile struct {
	URI  string `xml:"uri" yaml:"uri" json:"uri"`
	Auth int    `xml:"auth" yaml:"auth" json:"auth"`
}

//isAccess doc
//@Summary Whether to grant access
//@Method isAccess
//@Return (bool)
func (slf *AdminUserProfile) isAccess() bool {

	if (slf.Auth & ProfileAccess) > 0 {
		return true
	}
	return false
}

//isUpdate doc
//@Whether authorization can be updated
//@Method isUpdate
//@Return (bool)
func (slf *AdminUserProfile) isUpdate() bool {
	if (slf.Auth & ProfileUpdate) > 0 {
		return true
	}
	return false
}

//isDelete doc
//@Summary Whether authorization can be deleted
//@Method isDelete
//@Return (bool)
func (slf *AdminUserProfile) isDelete() bool {
	if (slf.Auth & ProfileDelete) > 0 {
		return true
	}
	return false
}

//isAppend doc
//@Summary Whether authorization can be added
//@Method isAppend
//@Return (bool)
func (slf *AdminUserProfile) isAppend() bool {
	if (slf.Auth & ProfileAppend) > 0 {
		return true
	}
	return false
}

//User doc
//@Summary User Online data
//@Struct User Claims
//@Member
type User struct {
	Key       string             `xml:"key" yaml:"key" json:"key"`
	Name      string             `xml:"name" yaml:"name" json:"name"`
	LoginTime int                `xml:"logintime" yaml:"logintime" json:"logintime"`
	Profile   []AdminUserProfile `xml:"profiles" yaml:"profiles" json:"profiles"`
}

//IsAccess doc
//@Summary Whether to grant access
//@Method IsAccess
//@Param  (string) Inspection path
//@Return (bool)
func (slf *User) IsAccess(URI string) bool {
	perm := slf.getProfile(URI)
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
	perm := slf.getProfile(URI)
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
	perm := slf.getProfile(URI)
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
	perm := slf.getProfile(URI)
	if perm == nil {
		return true
	}
	return perm.isAppend()
}

func (slf *User) getProfile(URI string) *AdminUserProfile {
	for _, v := range slf.Profile {
		if v.URI == URI {
			return &v
		}
	}
	return nil
}
