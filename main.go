package main

import (
	"fmt"
	"strings"
)

func main() {
	kursNavn := "Go konferanse 2025"
	const kursBillettAntall = 50
	var antallBilletterTilgjengelig uint = 50
	var bookings []string

	velkommen(kursNavn, kursBillettAntall)

	for {
		brukerForNavn, brukerEtterNavn, brukerEmail, brukerBillett := hentBrukerInformasjon()

		erGyldigNavn, erGyldigEmail, erGyldigBillett := erGyldig(brukerForNavn, brukerEtterNavn, brukerEmail, brukerBillett, antallBilletterTilgjengelig)

		if erGyldigNavn && erGyldigEmail && erGyldigBillett {
			antallBilletterTilgjengelig = antallBilletterTilgjengelig - brukerBillett
			bookings = append(bookings, brukerForNavn+" "+brukerEtterNavn)

			fmt.Printf("Takk for din bestilling %v %v du bestilte %v billett(er), du vil motta en bekreftelse på e-post: %v\n", brukerForNavn, brukerEtterNavn, brukerBillett, brukerEmail)

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

func erGyldig(brukerForNavn string, brukerEtterNavn string, brukerEmail string, brukerBillett uint, antallBilletterTilgjengelig uint) (bool, bool, bool) {
	erGyldigNavn := len(brukerForNavn) >= 2 && len(brukerEtterNavn) >= 2
	erGyldigEmail := strings.Contains(brukerEmail, "@") && strings.Contains(brukerEmail, ".")
	erGyldigBillett := brukerBillett > 0 && brukerBillett <= antallBilletterTilgjengelig
	return erGyldigNavn, erGyldigEmail, erGyldigBillett
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
