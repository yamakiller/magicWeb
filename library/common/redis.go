package common

//GetRdsOnlineKey doc
//Summary Returns Online User key
//Method GetOnlineKey
//Param (string) user id
//Return (string) user key
func GetAdminRdsOnlineKey(userid string) string {
	return "adminuser:" + userid
}
