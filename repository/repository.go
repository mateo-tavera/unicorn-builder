package repository

import (
	"github.com/mateo-tavera/unicorn-builder/util"
)

// Store names in a slice of strings from the file
func GetNames() []string {
	return util.ReadFile("repository/petnames.txt")
}

// Store adjectives in a slice of strings from the file
func GetAdjectives() []string {
	return util.ReadFile("repository/adj.txt")
}

// Store capabilities in a slice of strings from the file
func GetCapabilities() []string {
	return util.ReadFile("repository/capabilities.txt")
}
