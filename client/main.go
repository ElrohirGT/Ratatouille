package main

import (
	"github.com/ElrohirGT/Ratatouille/internal/tui"
)

func main() {
	user, password := tui.StartAuthentication()
	
	println(user, password)
}
