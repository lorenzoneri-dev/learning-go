package main //indica che appartiene al file principale

import "fmt" //importa un pacchetto, il pacchetto fmt serve per stampare

func main(){ //da qui parte il programma

	var eta int = 12
	if eta >= 18{ //faccio una verifica
		fmt.Println("Sei maggiorenne")
	}else{
		fmt.Println("Sei minorenne")
	}

	var x int //dichiaro la variabile x

	for x < 10 { //ciclo for fino a 10 che permette di stampare sul termile numeri da 1 a 10
		x++
		fmt.Println(x)
	}
}


