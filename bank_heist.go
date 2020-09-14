package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// random number with seed
	rand.Seed(time.Now().UnixNano())
	// keeps track of whether how successful our heist is
	isHeistOn := true
	// 0-99 and we'll use it to simulate if we made it past the guards or not
	eludedGuards := rand.Intn(100)

	/*
	First Conditional (Sneak past guards)
	Let's say we have a 50% chance of making it past the guards
	 */
	if (eludedGuards >= 50) {
		fmt.Println("Looks like you've managed top make it past the guards. " +
			"Good job, but remember, this is the first step.")
	} else {
		// cancel the heist! The guards made us
		isHeistOn = false
		fmt.Println("Well that sucked. You didn't make it far at all! :(")
	}

	/*
	Second Conditional (Open the Vault)
	We'll use this value to determine if we are able to open the vault
	Opening the vault is harder than sneaking past the guards and we only have a 30%
	  chance of doing it.
	 */
	openedVault := rand.Intn(100)
	if (isHeistOn && (openedVault >= 70)) {
		fmt.Println("Grab and GO! GO! GO!")
	} else if (isHeistOn) {
		isHeistOn = false
		fmt.Println("Cracking the vault was unsuccessful.")
	}

	/*
	Third Conditional (Leaving)
	We still need to make it out with the money without getting caught by any of the bank's
	  security.  Sensory alarms (lasers), security cameras, extra guards, cops, et cetera
	 */

	// keeps track of whether or not the heist is on
	fmt.Println("The heist is currently a Go: ", isHeistOn)

}
