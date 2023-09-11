package util

import (
	"math/rand"
	"sync"
	"time"
)

// Global values
var idMutex sync.Mutex
var lastAssignedID int

// Creates a unicorn name using a random name and a random adjective
func GenerateRandomName(adjectives []string, names []string) string {
	return adjectives[rand.Intn(len(adjectives))] + "-" + names[rand.Intn(len(names))]
}

// Creates a list of 3 random unrepeated capabilities using a predefined list
func GenerateRandomCapabilities(capabilities []string, num int) []string {

	// Prevents to use more capabilities than possible
	if num > len(capabilities) {
		num = len(capabilities)
	}

	// Creates a random index list according to number of capabilities requested
	rand.Seed(time.Now().UnixNano())
	indices := rand.Perm(len(capabilities))[:num]

	// Select the unicorn capabilities according to the indexes randomly generated
	unicornCapabilities := make([]string, num)
	for i, n := range indices {
		unicornCapabilities[i] = capabilities[n]
	}

	return unicornCapabilities
}

// Simulates the creation of a unicorn
func GetProductionTime() time.Duration {
	return time.Duration(rand.Intn(3500)) * time.Millisecond // This value can be changed
}

// Assigns the request id in function of the last id created
func GenerateUniqueID() int {
	// Prevents race condition
	idMutex.Lock()
	defer idMutex.Unlock()

	lastAssignedID++
	return lastAssignedID
}
