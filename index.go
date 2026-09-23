package main //indica che appartiene al file principale

import "fmt" //importa un pacchetto, il pacchetto fmt serve per stampare

func main(){ //da qui parte il programma

	var eta int
	fmt.Println("Inserisci la tua eta' in forma numerica: ")
	fmt.Scan(&eta)

	if eta >= 18{ //faccio una verifica
		fmt.Println("Sei maggiorenne")
	}else{
		fmt.Println("Sei minorenne")
	};
}


