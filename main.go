package main

import (
	"github.com/okeyaki/zipaddr/command"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	command.Execute()
}
