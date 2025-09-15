package commands

import nadhi "nadhi.dev/binaries/gpm/helpers"


func Tidy() {
	nadhi.Logo()
	nadhi.Tidy()
	nadhi.Govar(true)
	nadhi.Success("Tidied go.mod and go.sum files.")
}