package commands

import (
	"fmt"
)

var VERSION = "1.0.6"
var GitRev = ""

type VersionCommand struct {
}

//Execute - returns the version
func (c *VersionCommand) Execute([]string) error {
	suffix := ""
	if GitRev != "" {
		suffix = fmt.Sprintf("-rev-%s\n", GitRev)
	}
	fmt.Println(fmt.Sprintf("%s%s",VERSION, suffix))
	return nil
}
