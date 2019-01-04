package core

type User interface {
	//Gets the current nickname for a user.
	//In IRC, this follows the format of "nickname!loginname@hostname".
	//This name is able to be changed while the user is connected to the IRC server.
	Nickname() string

	//Gets the login name for a user.
	//In IRC, this follows the format of "nickname!loginname@hostname".
	//This does not change once a user is connected.
	LoginName() string

	//Gets the hostname for a user.
	//In IRC, this follows the format of "nickname!loginname@hostname".
	//This does not change once a user is connected.
	Hostname() string

	//Gets the nickserv account name for this user.
	//This is only supported if the IRC server has such a system.
	//An empty string indicates the user is not logged into an account.
	NickservAccount() string
}

