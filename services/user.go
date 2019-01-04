package services

type user struct {
	nickname     string
	loginName    string
	host         string
	nickservName string
}

func (u *user) Nickname() string {
	return u.nickname
}

func (u *user) LoginName() string {
	return u.loginName
}

func (u *user) Hostname() string {
	return u.host
}

func (u *user) NickservAccount() string {
	return u.nickservName
}
