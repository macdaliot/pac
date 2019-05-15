package util

import (
	"strings"

	"github.com/PyramidSystemsInc/go/errors"
)

//ValidatePipelineType : Validates that the Pipeline arg is a value that PAC can manage.
func ValidatePipelineType(args []string) bool {
	// TODO as PAC grows we might want to consider extracting such properties from the source code.
	supportedPipelines := []string{"jenkins", "azure"}

	if len(args) == 1 {
		if !contains(supportedPipelines, args[0]) {
			errors.LogAndQuit("The type was set to an invalid value. The valid types are " + strings.Join(supportedPipelines[:], ", "))
		}
	} else if len(args) == 0 {
		errors.LogAndQuit("A type must be specifed after the 'automate' command. The valid types are " + strings.Join(supportedPipelines[:], ", "))
	} else if len(args) > 1 {
		errors.LogAndQuit("Only one type may be passed for each 'automate' command")
	}
	return true
}

// TODO the following functions are probably a better fit for a go/collctions helper lib, team to advise
func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func contains(vs []string, t string) bool {
	return index(vs, t) >= 0
}
