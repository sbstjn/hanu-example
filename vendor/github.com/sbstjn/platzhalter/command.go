package platzhalter

import (
	"errors"
	"regexp"
	"strings"
)

const (
	commandParameterValue   = "a-zA-Z0-9\\-"
	commandParameterKeyName = "a-zA-Z\\-"
)

// Command has a Text
type Command struct {
	Text string
}

// Keys parses defined key names
func (c *Command) Keys() []string {
	rp := regexp.MustCompile("<([" + commandParameterKeyName + "]+)>")
	keys := rp.FindAllString(c.Text, -1)

	for i := 0; i < len(keys); i++ {
		keys[i] = strings.Replace(keys[i], "<", "", -1)
		keys[i] = strings.Replace(keys[i], ">", "", -1)
	}

	return keys
}

// Values returns a list of values in the request matching the command
func (c *Command) Values(request string) ([]string, error) {
	if !c.Matches(request) {
		return nil, errors.New("Command does not match Request")
	}

	rp := c.Regexp()
	keys := rp.FindAllStringSubmatch(request, -1)

	return keys[0][1:], nil
}

// Matches checks if the command matches for given request
func (c *Command) Matches(request string) bool {
	return c.Regexp().MatchString(request)
}

// Regexp returns a regular expression matching the command
func (c *Command) Regexp() *regexp.Regexp {
	template := c.Text
	pattern := "<([" + commandParameterKeyName + "]+)\\?*>"
	rp := regexp.MustCompile(pattern)
	keys := rp.FindAllString(template, -1)

	for i := 0; i < len(keys); i++ {
		template = strings.Replace(template, keys[i], "(["+commandParameterValue+"]+)", -1)
	}

	re, _ := regexp.Compile("^" + template + "$")
	return re
}

// Param returns the value of a parameter in a request matching the command
func (c *Command) Param(request string, name string) string {
	keys := c.Keys()
	values, _ := c.Values(request)

	if len(keys) != len(values) {
		return ""
	}

	for i := 0; i < len(keys); i++ {
		if keys[i] == name {
			return values[i]
		}
	}

	return ""
}
