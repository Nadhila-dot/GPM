package commands

import nadhi "nadhi.dev/binaries/gpm/helpers"


func Tidy() {
	nadhi.Logo()
	nadhi.Tidy()
	nadhi.Govar(true) // <--- 
	// why am I passing true here? Well when u pass true it prints some text like checking and etc..
	// Im a lazy person
	nadhi.Success("Tidied go.mod and go.sum files.")
}