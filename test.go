package main // Package declaration

import "fmt" // Import statement

func main() {
	empSal := map[string]int{
		"Neha": 20000,
		"Raj":  25000,
		"Atul": 15000}

	fmt.Println(empSal["Atul"]) // read
	empSal["Test"] = 1800025    //add
	empSal["Atul"] = 18000      //update
	delete(empSal, "Raj")       // delete
	fmt.Println(empSal)

	_, exist := empSal["test2"] // "_" mean i will define variable later; it is insted of  "test2, flag", if we would kept "test2, flag", we would have to use "test2".
	fmt.Println(exist)          // output: false  ; check specific key exst or not
}
