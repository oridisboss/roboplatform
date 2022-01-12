package main

var MyUser User

type User struct {
	token string
}

func (pUser *User) SetToken(token string) {
	pUser.token = token
	SaveSetting("token", token)
	GRPCConnect()
}

func (pUser User) GetToken() string {
	if pUser.token == "" {
		pUser.token = GetSetting("token")
	}
	return pUser.token
}
