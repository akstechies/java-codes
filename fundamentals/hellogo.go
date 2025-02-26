// Package is a collection of code
package main // main means this code to be taken up by terminal

// import multiple packages ~ single packages can be done like - import "fmt"
import (
	"fmt"
	// "runtime"
)

// here above I can define a global variable that will use fmt.Println so that I dont't have to type it everytime
var pl = fmt.Println

// when go program runs it looks for main fucntion to be executed

func main() {

	fmt.Println("Hello GO!")
	// now using println variable
	// pl("Hello Ayush")

	// 1st - Input Tut
	// GetInput()

	// 2nd - Variables Tut
	// VariableTut()

	// 3rd - Data Types
	// DataTypes()

	// 4th - casting
	// CastingInGo()

	// fmt.Println(runtime.NumCPU()) //  use for getting the number of processors utilized in the program

	// 5th - Conditional
	// Conditional()

	// 6th - Strings
	// StringsTut()

	// 7th - Runes
	// RunesTut()

	// 8th - Time
	// TimeInGo()

	// 9th - Math
	// MathInGo()

	// 10th - ForLoop
	// ForLoop()

	// 11th - Arrays
	// ArraysInGO()

	// 12th - Slices
	// SlicesInGo()

	// 13th - Functions
	// FunctionsInGo()

	// 14th - Pointers
	// PointersinGo()

	// 15th - Files
	// FilesIO()

	// 16th - Packages and Modules
	// PackagesModules()

	// 17th - Maps
	// MapsInGo()

	// 18th - Generics and Constraints
	// GenericsConstraints()

	// 19th - Structs
	// StructsTut()

	// 20th - Getter Setter (Encapsulation)
	// GetterSetterTut()

	// 21st - Interface
	// InterfaceTut()

	// 22nd - Concurrency
	// GoroutinesTut()

	// 23rd - Channel
	ChannelsTut()

	// 23rd - Mutex
	// MutexInGo()
	// AccountMutex()

	// 24th - ClosuresInGo
	// ClosuresInGo()

	// 25th - Recursion
	// RecursionInGo()

	// 26th - Unit Test
	// UnitTestTut()
}
