package service

import (
	"fmt"
	"time"

	"github.com/mateo-tavera/unicorn-builder/entity"
	"github.com/mateo-tavera/unicorn-builder/stack"
	"github.com/mateo-tavera/unicorn-builder/util"
)

// Get all unicorns needed for the request regardless of whether they were created
// or collected from the store. Returns the list of unicorns
func GetUnicorns(amount int, names []string, adjectives []string) []entity.Unicorn {
	// Validate if there are available unicorns
	store := stack.GetUnicornStack()
	unicorns := stack.CheckUnicornStack(amount, store)
	// Create the reminder unicorns needed for the request
	if len(unicorns) < amount {
		remaining := amount - len(unicorns)
		liveUnicorns := createLiveUnicorns(names, adjectives, remaining)
		unicorns = append(unicorns, liveUnicorns...)
	}

	return unicorns
}

// Create the reminder unicorns needed for the request
func createLiveUnicorns(names []string, adjectives []string, amount int) []entity.Unicorn {
	unicorns := make([]entity.Unicorn, amount)

	for i := 0; i < amount; i++ {
		name := util.GenerateRandomName(adjectives, names)
		capabilities := util.GenerateRandomCapabilities(util.Capabilities, 3)

		unicorn := entity.Unicorn{
			Name:         name,
			Id:           stack.SetId(),
			Capabilities: capabilities,
		}
		fmt.Printf("- Creating unicorn with id %d...\n", unicorn.Id)

		unicorns[i] = unicorn
		time.Sleep(util.GetProductionTime()) // Simulates time consumed for creating a unicorn
		fmt.Printf("- Unicorn with id %d was successfully created\n", unicorn.Id)
	}

	return unicorns
}
