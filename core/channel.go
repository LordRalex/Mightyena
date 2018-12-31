package core

type Channel interface {
	GetName() string
	GetUsers() []User
	GetOps() []User
	GetVoiced() []User
	HasVoiceOrOp(name string) bool
	HasOp(name string) bool
	HasVoice(name string) bool
}

type channel struct {
	users  []User
	name   string
	ops    []User
	voiced []User
}

func (c *channel) GetName() string {
	return c.name
}

func (c *channel) GetUsers() []User {
	return c.users
}

func (c *channel) GetOps() []User {
	return c.ops
}

func (c *channel) GetVoiced() []User {
	return c.voiced
}

func (c *channel) HasVoiceOrOp(name string) bool {
	return c.HasOp(name) || c.HasVoice(name)
}

func (c *channel) HasOp(name string) bool {
	for _, v := range c.ops {
		if v.GetNickname() == name {
			return true
		}
	}
	return false
}

func (c *channel) HasVoice(name string) bool {
	for _, v := range c.voiced {
		if v.GetNickname() == name {
			return true
		}
	}
	return false
}