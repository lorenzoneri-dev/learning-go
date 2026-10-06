/*
Nel printf notiamo che ci sta %[2]s o %[1]s serve a dire quale variabile prendere seguendo l'ordine alla fine del print quindi in questo caso
"surname, name" infatti notiamo come %[2]s prende nome che si trova in posizioen due mentre %[1]s prende surname essendo che si trova in posizione uno
*/


package main 

import ("fmt"; "math"; "time")

func main() {
	var name, surname string = "Lorenzo", "Neri"
	var x = 4 * math.Atan(1)

	fmt.Printf("Today is %v. How are you doing? %s\n\n", time.Now().Format("02 Jan 2006"), name)
	fmt.Printf("Messing up with the arguments.\n")
	fmt.Printf(" Name :- %[2]s\tSurname :- %[1]s\n Name and Surname :- %[2]s %[1]s\n\n", surname, name)
	fmt.Printf("The value of π is %1.2f or %1.3[1]f or %1.4[1]f\n", x)
}