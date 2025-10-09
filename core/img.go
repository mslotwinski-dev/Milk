package core

import (
	utility "milk/core/utility"
)

func (c *Component) Src(src string) *Component {
	if c.Tag == "img" {
		c.ImageSrc = src

	} else {
		utility.Error("Src can only be set on img components")
	}
	return c
}

func (c *Component) Alt(alt string) *Component {
	if c.Tag == "img" {
		c.ImageAlt = alt

	} else {
		utility.Error("Alt can only be set on img components")
	}
	return c
}

func (c *Component) Width(width int) *Component {
	if c.Tag == "img" {
		c.ImageWidth = width

	} else {
		utility.Error("Width can only be set on img components")
	}
	return c
}

func (c *Component) Height(height int) *Component {
	if c.Tag == "img" {
		c.ImageHeight = height

	} else {
		utility.Error("Height can only be set on img components")
	}
	return c
}
