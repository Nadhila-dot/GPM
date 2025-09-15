package nadhi

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

func Error(message string) {
	// ANSI escape code for red: \033[31m, bold: \033[1m, reset: \033[0m
	fmt.Printf("%s\033[31m[Error] %s%s\n", bold, message, reset)
}

func RedLoading(message string) {
	// ANSI escape code for red: \033[31m, bold: \033[1m, reset: \033[0m
	fmt.Printf("%s\033[31m[⊗] %s%s\n", bold, message, reset)
}

func Success(message string) {
	// ANSI escape code for green: \033[32m, bold: \033[1m, reset: \033[0m
	fmt.Printf("%s\033[32m[Success] %s%s\n", bold, message, reset)
}


func SuccessCheck(message string) {
	// ANSI escape code for dark green: \033[32;2m, bold: \033[1m, reset: \033[0m
	fmt.Printf("%s\033[32;2m[✓] %s%s\n", bold, message, reset)
}

func LoadingCheck(message string) {
	// ANSI escape code for dark yellow: \033[33;2m, bold: \033[1m, reset: \033[0m
	fmt.Printf("%s\033[33;2m[↻] %s%s\n", bold, message, reset)
}

func Loading(message string) {
	// ANSI escape code for yellow: \033[33m, bold: \033[1m, reset: \033[0m
	fmt.Printf("%s\033[33m[Loading] %s%s\n", bold, message, reset)
}

func Hint(message string) {
	// ANSI escape code for light gray: \033[90m, bold: \033[1m, reset: \033[0m
	fmt.Printf("%s\033[90m[Hint] %s%s\n", bold, message, reset)
}



