package core

import (
	"fmt"
	"strings"
)

type EventHandlers struct {
	Click string
}

func (c *Component) Click(handler string) *Component {
	c.Alpine.XOn.Click = handler
	return c
}

func (e EventHandlers) Render() string {
	var b strings.Builder

	if e.Click != "" {
		fmt.Fprintf(&b, ` @click="%s"`, e.Click)
	}

	return b.String()
}
