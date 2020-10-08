package main

import (
	"wait4it/cmd"
	"wait4it/input/env"
	"wait4it/input/flag"
	"wait4it/model"
)

func main() {
	c := &model.CheckContext{}
	c = env.Parse(c)
	c = flag.Parse(c)
	cmd.RunCheck(*c)
}
