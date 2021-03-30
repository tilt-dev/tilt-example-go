// +build !skipintegration

package integration

import (
	"fmt"
	"testing"
	"time"
)

func TestSomethingSlow(t *testing.T) {
	fmt.Println("Transponder log partitioned")
	time.Sleep(time.Second)
	fmt.Println("Cascading generator transistorized")
	time.Sleep(time.Second)
	fmt.Println("Checking the harmonix proxy processor")
	time.Sleep(time.Second)
	fmt.Println("Is network capacity inverted?")
	time.Sleep(time.Second)
	fmt.Println("\t- transmission technician acquired")
	time.Sleep(time.Second)
	fmt.Println("\t- transmission technician acquired")
	time.Sleep(time.Second)
	fmt.Println("All is well!")
}

func TestOtherSlowThing(t *testing.T) {
	fmt.Println("Verifying scan bus frequency")
	time.Sleep(time.Second)
	fmt.Println("Logisticating data")
	time.Sleep(time.Second)
	fmt.Println("Broadband plasma may be dangerous, performing safety check")
	time.Sleep(time.Second)
	fmt.Println("Watch out for geese")
	time.Sleep(time.Second)
	fmt.Println("All is well!")
}
