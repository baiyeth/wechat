package miniProgram

// Content 微信小程序相关API
type Content struct {
	*MiniProgram
}

func NewContent(mp *MiniProgram) *Content {
	return &Content{mp}
}

func (c *Content) CheckText(text string) (err error) {
	return c.GetMp().GetContentSecurity().CheckText(text)
}

func (c *Content) CheckImage(media string) (err error) {
	context := c.GetMp().GetContentSecurity()
	return context.CheckImage(media)
}
