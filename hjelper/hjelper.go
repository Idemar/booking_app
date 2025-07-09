package hjelper

import "strings"

func ErGyldig(brukerForNavn string, brukerEtterNavn string, brukerEmail string, brukerBillett uint, antallBilletterTilgjengelig uint) (bool, bool, bool) {
	erGyldigNavn := len(brukerForNavn) >= 2 && len(brukerEtterNavn) >= 2
	erGyldigEmail := strings.Contains(brukerEmail, "@") && strings.Contains(brukerEmail, ".")
	erGyldigBillett := brukerBillett > 0 && brukerBillett <= antallBilletterTilgjengelig
	return erGyldigNavn, erGyldigEmail, erGyldigBillett
}
