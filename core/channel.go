package core

type Channel interface {
	GetName() string
	GetUsers() []string
	GetOps() []string
	GetVoiced() []string
	HasVoiceOrOp(name string) bool
	HasOp(name string) bool
	HasVoice(name string) bool
}

type channel struct {
	users  []string
	name   string
	ops    []string
	voiced []string
}

func (c *channel) GetName() string {
	return c.name
}

func (c *channel) GetUsers() []string {
	return c.users
}

func (c *channel) GetOps() []string {
	return c.ops
}

func (c *channel) GetVoiced() []string {
	return c.voiced
}

func (c *channel) HasVoiceOrOp(name string) bool {
	return c.HasOp(name) || c.HasVoice(name)
}

func (c *channel) HasOp(name string) bool {
	for _, v := range c.ops {
		if v == name {
			return true
		}
	}
	return false
}

func (c *channel) HasVoice(name string) bool {
	for _, v := range c.voiced {
		if v == name {
			return true
		}
	}
	return false
}