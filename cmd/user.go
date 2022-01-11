package main

type User struct {
	token string
}

func (pUser *User) SetToken(token string) {
	pUser.token = token
}
func (pUser User) GetToken() string {
	return pUser.token
}
