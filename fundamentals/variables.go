package main

func VariableTut() {
	// ===== VARIABLES
	// syntax ~ var name type
	var myname string = "ayush"
	var v1, v2 string = "a", "b" // multiple variable declare
	var v3 = 22                  // let go decide variable type
	v4 := 44                     // shorthand syntax
	v4 = 76                      // reassign value
	pl("Unused variables:")
	pl("myname:", myname)
	pl("v1:", v1)
	pl("v2:", v2)
	pl("v3:", v3)
	pl("v4:", v4)
}
