package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Psedudo Random Number Generator
	// SEED : starting point of generating a sequence of random numbers <- reproducibility by setting a seed for a sequence
	// True random numbers are generated from physical processes

	//CONSIDERATIONS : Thread safety / Cryptographic security
	fmt.Println(rand.Intn(101)) //autoseed from value "0"

	// with a seeded value
	// seed should be changing to generate different random numbers else random value will be fixed
	val := rand.New(rand.NewSource(50))
	fmt.Println(val.Intn(101)) // the random value remain fixed
	//changing seed value
	val2 := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(val2.Intn(80))
	fmt.Println(rand.Float64()) // between 0.0 - 1.0

}
