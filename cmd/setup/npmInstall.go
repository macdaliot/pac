package setup

import "github.com/PyramidSystemsInc/go/commands"

func NpmInstall() {
	commands.Run("npm i", "core")
}
