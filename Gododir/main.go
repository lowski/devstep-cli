package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("build", nil, func(c *do.Context) {
		c.Run("make coverage")
		c.Run("make vet")
		c.Run("make build-ci")
	}).Src("**/*.go")
}

func main() {
	do.Godo(tasks)
}
