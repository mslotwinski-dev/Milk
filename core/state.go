package core

import (
	"encoding/json"
	"fmt"
	"milk/core/utility"

	"regexp"
)

type State map[string]any

func (c *Component) State(data State) *Component {
	c.Alpine.XData = data
	return c
}

func (s State) Mutate(key string, value any) string {

	re := regexp.MustCompile(`[+\-*/&^!@#$%_=<>?~]`)
	search := re.ReplaceAllString(key, "")

	_, exists := s[search]
	if !exists {
		utility.Error("Key does not exist in state: %s", search)
		return ""
	}

	return fmt.Sprintf("%s=%v", key, value)
}

func (s State) Render() string {
	data, err := json.Marshal(s)
	if err != nil {
		return ` x-data="{}"`
	}
	return fmt.Sprintf(" x-data='%s'", data)
}
