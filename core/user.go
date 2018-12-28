package core

type User interface {
	//Gets the current nickname for a user.
	//In IRC, this follows the format of "nickname!loginname@hostname".
	//This name is able to be changed while the user is connected to the IRC server.
	GetNickname() string

	//Gets the login name for a user.
	//In IRC, this follows the format of "nickname!loginname@hostname".
	//This does not change once a user is connected.
	GetLoginName() string

	//Gets the hostname for a user.
	//In IRC, this follows the format of "nickname!loginname@hostname".
	//This does not change once a user is connected.
	GetHostname() string

	//Gets the nickserv account name for this user.
	//This is only supported if the IRC server has such a system.
	//An empty string indicates the user is not logged into an account.
	GetNickservAccount() string

	//Gets the channels this user belongs to
	GetChannels() []Channel
}

type user struct {
	nickname     string
	loginName    string
	host         string
	nickservName string
	channels     []Channel
}

func (u *user) GetNickname() string {
	return u.nickname
}

func (u *user) GetLoginName() string {
	return u.loginName
}

func (u *user) GetHostname() string {
	return u.host
}

func (u *user) GetNickservAccount() string {
	return u.nickservName
}

func (u *user) GetChannels() []Channel {
	return u.channels
}