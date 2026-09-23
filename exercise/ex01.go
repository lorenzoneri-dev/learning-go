package main //se non è main il file non si esegue

import "fmt"

func main() {
	var nome string
	fmt.Println("Inserisci il tuo nome: ")
	fmt.Scan(&nome) //Per l'input
	fmt.Println("Ciao, ", nome)
}
