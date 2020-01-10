package auth

//ConfigAdminUserProfile doc
//@Summary user Permission config
//@Member []UserPerm Arrays
type ConfigAdminUserProfileItems struct {
	Items []ConfigAdminUserProfile `xml:"items" yaml:"items" json:"items"`
}

type ConfigAdminUserProfile struct {
	Name  string `xml:"name" yaml:"name" json:"name"`
	URI   string `xml:"uri" yaml:"uri" json:"uri"`
	Local string `xml:"local" yaml:"local" json:"local"`
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
type AdminUserProfile struct {
	URI string `xml:"uri" yaml:"uri" json:"uri"`
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
		return false
	}

	return true
}

func (slf *User) getProfile(URI string) *AdminUserProfile {
	for _, v := range slf.Profile {
		if v.URI == URI {
			return &v
		}
	}
	return nil
}
