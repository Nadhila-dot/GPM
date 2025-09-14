package commands



import "fmt"

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	blue   = "\033[34m"
	cyan   = "\033[36m"
	green  = "\033[32m"
	yellow = "\033[33m"
	gray   = "\033[90m"
)

func Help() string {
	fmt.Println(bold + blue + ` 
  ____ ____  __  __  
 / ___|  _ \|  \/  | 
| |  _| |_) | |\/| | ` + cyan + `a product of pkg.lat
` + blue + `| |_| |  __/| |  | | ` + green + `Go Package Manager 
` + blue + ` \____|_|   |_|  |_| ` + reset)
	fmt.Println(gray + "A package Registar for Go Lang" + reset)
	fmt.Println()
	fmt.Println(bold + "Usage:" + reset + " " + cyan + "gpm <command> [parameters]" + reset + "\n")
	fmt.Println(bold + "Commands:" + reset)
	fmt.Println("  " + green + "refresh" + reset + "               " + gray + "Refresh the package list" + reset)
	fmt.Println("  " + green + "list" + reset + "                  " + gray + "List installed packages" + reset)
	fmt.Println("  " + green + "download" + reset + " <package>    " + gray + "Download a package" + reset)
	fmt.Println("  " + green + "install" + reset + " <package>     " + gray + "Install a package" + reset)
	fmt.Println("  " + green + "remove" + reset + " <package>      " + gray + "Remove a package" + reset)
	fmt.Println("  " + green + "help" + reset + "                  " + gray + "Show this help message" + reset)
	fmt.Println()
	fmt.Println("For more information, visit " + yellow + "https://go.pkg.lat" + reset)
	return ""
}