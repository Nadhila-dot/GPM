package commands

import (
	"fmt"

	nadhi "nadhi.dev/binaries/gpm/helpers"
	"nadhi.dev/binaries/gpm/toml"
)


func List() string {
	fmt.Println(bold + blue + ` 
  ____ ____  __  __  
 / ___|  _ \|  \/  | 
| |  _| |_) | |\/| | ` + cyan + `a product of pkg.lat
` + blue + `| |_| |  __/| |  | | ` + green + `Go Package Manager 
` + blue + ` \____|_|   |_|  |_| ` + reset)
	fmt.Println(bold + yellow + "List of installed packages:" + reset)
	fmt.Println()
	pkgs := toml.GetPackages()
	for name, link := range pkgs {
		fmt.Printf("%s%s%s -> %s\n", bold, blue, name, link)
	}
	nadhi.VerifyPackagesAndRemoveOnError()
	return "Listed packages"
}