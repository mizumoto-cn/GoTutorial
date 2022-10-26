// All identifiers defined within a package are visible throughout that package.
// When importing a package you can access only its exported identifiers.
// An identifier is exported if it begins with a capital letter.
package main

import "github/mizumoto-cn/gotut/4.ObjectOriented/4.Public_Private_GoModules/timer"

func main() {
	timestamp := timer.Timestamp{}
	timestamp.Set()

	// You cant do that
	// if timestamp
}
