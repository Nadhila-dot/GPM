package nadhi

import fmt "fmt"

func Logo() string {
	fmt.Println(bold + blue + ` 
  ____ ____  __  __  
 / ___|  _ \|  \/  | 
| |  _| |_) | |\/| | ` + cyan + `Simplify your Go modules
` + blue + `| |_| |  __/| |  | | ` + green + `Go Package Manager
` + blue + ` \____|_|   |_|  |_| ` + reset)
	fmt.Println(gray + "A package manager for Go Lang" + reset)
	fmt.Println("Made for " + WhatOs() + " with " + Govar(false))
	fmt.Println()


	return ""
}