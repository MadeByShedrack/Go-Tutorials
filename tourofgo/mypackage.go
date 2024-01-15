package tourofgo

import (
	"fmt"
	"math"
	"math/rand"
)

func MyPackage() {

	favoriteNumber := rand.Intn(12)
	fmt.Printf("Random Number -> %v\n", favoriteNumber)

	mySquare := math.Sqrt(2)
	fmt.Printf("My square root -> %v\n", mySquare)

	mathPi := math.Pi
	fmt.Printf("Math Pi -> %v\n", mathPi)
}
