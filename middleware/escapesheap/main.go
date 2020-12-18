package main

// "fmt"

// User sds
type User struct {
	ID     int64
	Name   string
	Avatar string
}

// GetUserInfo dsd
func GetUserInfo() *User {
	return &User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
}

// GetUserInfos dsa
func GetUserInfos(u *User) *User {
	return u
}

func main() {
	// _ = GetUserInfo()
	_ = GetUserInfos(&User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"})
	// str := new(string)
	// *str = "EDDYCJY"
	// fmt.Println(str)
}
