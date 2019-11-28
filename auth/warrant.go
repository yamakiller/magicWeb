package auth

type Warrant struct {
}

//@method Enter desc: sign in
//@param (string)  claim key/id
//@param (string)  claim name
//@param (int)     enter time
//@param ([]ClaimPerm) Permission array
//@param (int)  expire time util/Minute
func (slf *Warrant) Enter(key, name string, time, perm []ClaimPerm, expire int) (string, error) {

	//expireTime := nowTime.Add(time)

	return "", nil
}
