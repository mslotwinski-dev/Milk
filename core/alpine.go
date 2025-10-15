package core

import (
	"fmt"
	"strings"

	utility "milk/core/utility"
)

type AlpineJS struct {
	XOn EventHandlers

	XData State

	XText string
}

func (h *Html) IgnoreAlpineJS() *Html {
	h.IgnoreAlpine = true
	return h
}

func (c *Component) GetState(s State, key string) *Component {
	_, exist := s[key]

	if !exist {
		utility.Error("Key does not exist in state: %s", key)
		return c
	}

	c.Alpine.XText = key

	return c
}

func (a AlpineJS) Render() string {
	var b strings.Builder

	b.WriteString(a.XOn.Render())

	if len(a.XData) > 0 {
		b.WriteString(a.XData.Render())
	}

	if a.XText != "" {
		fmt.Fprintf(&b, ` x-text="%s"`, a.XText)
	}

	return b.String()

}
