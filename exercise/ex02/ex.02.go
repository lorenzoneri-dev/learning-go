/*
Piccolo progetto in Go per la gestione di un supermercato con semplicemente il nome del prodotto e il suo prezzo
Tutto viene salvato dentro struct prodotto ma una volta che il programma viene riavviato viene eliminato ogni singolo dato.
Questo progetto viene semplicamente gestito tramite un singolo file, niente import esterni di altri file. 

C'è una struct che serve per la struttura che rappresenta un prodotto con il suo nome e prezzo.
Func main():
- Nella prima parte andiamo a chiedere all'utente quanti sono i prodotti che deve inserire
- poi creiamo una slice che va da 0 fino al numero di prodotti (una slice in Go è una struttura dati dinamica che funge da "vista" su un array sottostante)
- Utilizzo di for per far inserire all'utente il nome e il prezzo del prodotto
- var p Prodotto serve per fare una nominazione della stucture Prodotto che si trova fuori della func main() per poterla utilizzare
- Poi iniziamo a chiedere all'utente i vari input
- Alla fine ci troviavo un append che serve per inserire i prodotti dentro la slice in questo modo da salvarli
- dopo essere usciti dal for stampiamo tutti i prodotti, per poter stampare tutti i prodotti senza che ci venga stampato solo l'utimo facciamo uso del ciclo for,
andiamo a ricavare usando la nominazione p per la struct Progetto quando sia lungo quindi il suo range in questo modo che il ciclo possa fermarsi ad un certo punto.
Nel print andiamo ad inserire %d serve per stampare un decimale che andra a fare da punto infatti vediamo 1) 2) ecc... %s serve per stampare il nome, %.2f serve per stampare il prezzo perché
f indica un numero float, .2 invece specifica di mostrare 2 cifre dopo la virgola. 
*/

/*
La struct sarebbe un insieme di elementi, come avevamo visto prima p è come se contenesse tutti i dati di Progetto essendo che progetto viene salvato come type
Quindi p non è un puntatore
*/


package main

import "fmt"

// Struttura che rappresenta un prodotto con nome e prezzo.
type Prodotto struct {
	Nome   string
	Prezzo float64
}

func main() {
	// Chiede quanti prodotti l'utente vuole inserire.
	var numeroProdotti int
	fmt.Println("Inserisci il numero di prodotti che devi inserire")
	fmt.Scan(&numeroProdotti)

	// Crea una slice vuota di prodotti con capacità iniziale pari al numero inserito.
	prodotti := make([]Prodotto, 0, numeroProdotti)

	// Ciclo che fa inserire il nome e il prezzo per ogni prodotto.
	for i := 0; i < numeroProdotti; i++ {
		var p Prodotto	//come noti Prodotto viene nomitato p questo viene utilizzato per inizializzarlo, p in questo momento contiene tutti i dati di prodotto e li puo chiamare per usarli

		fmt.Println("Inserisci il nome del prodotto: ")
		fmt.Scan(&p.Nome)//Come possiamo notare in questo punto viene chiamto p per inserire dati che si trovano all'interno della variabile che si trova dentro lo struct di Progetto

		fmt.Println("Inserisci il prezzo del prodotto: ")
		fmt.Scan(&p.Prezzo)

		// Aggiunge il prodotto appena inserito alla slice.
		prodotti = append(prodotti, p)
	}

	// Stampa tutti i prodotti salvati.
	fmt.Println("\nProdotti inseriti:")
	for i, p := range prodotti {
		fmt.Printf("%d) %s - %.2f €\n", i+1, p.Nome, p.Prezzo)
	}
}
