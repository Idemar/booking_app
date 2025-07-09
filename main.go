package main

import (
	"booking_app/hjelper"
	"fmt"
	"strconv"
)

const kursBillettAntall = 50

var kursNavn = "Go konferanse 2025"
var antallBilletterTilgjengelig uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	velkommen(kursNavn, kursBillettAntall)

	for {
		brukerForNavn, brukerEtterNavn, brukerEmail, brukerBillett := hentBrukerInformasjon()

		erGyldigNavn, erGyldigEmail, erGyldigBillett := hjelper.ErGyldig(brukerForNavn, brukerEtterNavn, brukerEmail, brukerBillett, antallBilletterTilgjengelig)

		if erGyldigNavn && erGyldigEmail && erGyldigBillett {
			bookBillett(brukerForNavn, brukerEtterNavn, brukerEmail, brukerBillett)

			fmt.Printf("Booking liste: %v\n", bookings)
			fmt.Println("Antall billetter igjen:", antallBilletterTilgjengelig)
			fmt.Println(" ")
		} else {
			if !erGyldigNavn {
				fmt.Println("Fornavn og etternavn må være minst 2 tegn langt.")
			}
			if !erGyldigEmail {
				fmt.Println("E-post må inneholde '@' og '.'")
			}
			if !erGyldigBillett {
				fmt.Println("Antall billetter må være mellom 1 og", antallBilletterTilgjengelig)
			}
		}

	}
}

func velkommen(kursNavn string, antallBilletter uint) {
	fmt.Printf("Velkommen til %v booking programmet.\n", kursNavn)
	fmt.Printf("Kurset har totalt %v billetter.\n", antallBilletter)
	fmt.Printf("Meld deg på kurs her. Det er %v billetter igjen for booking.\n", antallBilletter)
}

func hentBrukerInformasjon() (string, string, string, uint) {
	var brukerForNavn string
	var brukerEtterNavn string
	var brukerEmail string
	var brukerBillett uint

	fmt.Println()
	fmt.Println("----------Ny booking----------")

	fmt.Println("Skriv inn fornavn:")
	fmt.Scan(&brukerForNavn)

	fmt.Println("Skriv inn etternavn:")
	fmt.Scan(&brukerEtterNavn)

	fmt.Println("Skriv inn e-post:")
	fmt.Scan(&brukerEmail)

	fmt.Println("Hvor mange billetter ønsker du å bestille?")
	fmt.Scan(&brukerBillett)

	return brukerForNavn, brukerEtterNavn, brukerEmail, brukerBillett
}

func bookBillett(brukerForNavn string, brukerEtterNavn string, brukerEmail string, brukerBillett uint) {
	antallBilletterTilgjengelig = antallBilletterTilgjengelig - brukerBillett

	var brukerData = make(map[string]string)
	brukerData["fornavn"] = brukerForNavn
	brukerData["etternavn"] = brukerEtterNavn
	brukerData["epost"] = brukerEmail
	brukerData["antallBilletter"] = strconv.FormatUint(uint64(brukerBillett), 10)

	bookings = append(bookings, brukerData)

	fmt.Printf("Takk for din bestilling %v %v du bestilte %v billett(er), du vil motta en bekreftelse på e-post: %v\n", brukerForNavn, brukerEtterNavn, brukerBillett, brukerEmail)

}
