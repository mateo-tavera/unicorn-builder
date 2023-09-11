package stack

import (
	"fmt"
	"sync"
	"time"

	"github.com/mateo-tavera/unicorn-builder/entity"
	"github.com/mateo-tavera/unicorn-builder/repository"
	"github.com/mateo-tavera/unicorn-builder/util"
)

// Global variable to be configurated in the service layer
var unicornStack *UnicornStack

type UnicornStack struct {
	mu       sync.Mutex
	Unicorns []entity.Unicorn
}

func NewUnicornStack() *UnicornStack {
	return &UnicornStack{}
}

func GetUnicornStack() *UnicornStack {
	return unicornStack
}

func SetUnicornStack(stack *UnicornStack) {
	unicornStack = stack
}

// Validates if there are enough unicorns in the store to satisfy the requirements
// and returns the ones that satisfy
func CheckUnicornStack(amount int, stack *UnicornStack) []entity.Unicorn {
	// Prevents race condition
	stack.mu.Lock()
	defer stack.mu.Unlock()
	// Loop through the store collecting the useful unicorns
	unicorns := []entity.Unicorn{}
	for len(stack.Unicorns) > 0 && len(unicorns) < amount {
		unicorn := pop(stack)
		unicorns = append(unicorns, unicorn)
	}

	return unicorns
}

// Takes the last added value from the stack
func pop(stack *UnicornStack) entity.Unicorn {
	unicorn := stack.Unicorns[len(stack.Unicorns)-1]
	fmt.Printf("- Unicorn with id %d was taken from the store\n", unicorn.Id)
	stack.Unicorns = stack.Unicorns[:len(stack.Unicorns)-1]

	return unicorn
}

// Continously add a new unicorn to the store
func CreateBacklogUnicorns(stack *UnicornStack) {
	for {
		// Input
		names := repository.GetNames()
		adjectives := repository.GetAdjectives()
		restockTime := 15 * time.Second // This value can be changed
		// Simulates time consumed in restocking
		time.Sleep(restockTime)
		// Generate unicorn atributes
		name := util.GenerateRandomName(adjectives, names)
		listOfCapabilities := repository.GetCapabilities()
		capabilities := util.GenerateRandomCapabilities(listOfCapabilities, 3)
		// Create unicorn
		unicorn := entity.Unicorn{
			Name:         name,
			Id:           SetId(),
			Capabilities: capabilities,
		}
		// Push the unicorn into the store
		stack.mu.Lock()
		stack.Unicorns = append(stack.Unicorns, unicorn)
		fmt.Printf("- Unicorn with id %d was added to the store\n", unicorn.Id)
		stack.mu.Unlock()

	}
}

// Returns the unique identifier generated
func SetId() int {
	return util.GenerateUniqueID()
}
