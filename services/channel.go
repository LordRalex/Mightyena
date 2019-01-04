package services

type channel struct {
	users  []string
	name   string
	ops    []string
	voiced []string
}


func (c *channel) Name() string {
	return c.name
}

func (c *channel) Users() []string {
	return c.users
}

func (c *channel) Ops() []string {
	return c.ops
}

func (c *channel) Voiced() []string {
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

func (c *channel) removeUser(name string) {
	newUsers := make([]string, 0)

	for _, v := range c.users {
		if v != name {
			newUsers = append(newUsers, v)
		}
	}
	c.users = newUsers

	newVoiced := make([]string, 0)

	for _, v := range c.voiced {
		if v != name {
			newVoiced = append(newVoiced, v)
		}
	}
	c.voiced = newVoiced

	newOps := make([]string, 0)

	for _, v := range c.ops {
		if v != name {
			newOps = append(newOps, v)
		}
	}
	c.ops = newOps
}
