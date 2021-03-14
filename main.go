package main

import "fmt"

const TrobarTotesLesSolucions = true

func main() {

	var pendents, bons []Quadre
	var solucionat, sortirBucle bool

	solucionat = false
	sortirBucle = false
	q := Init()
	q.First()

	pendents = append(pendents, *q)

	for i := 0; len(pendents) != 0 && !sortirBucle; i++ {
		//agafem l'últim q així tindirem a avançar i tenir els grups de pendents més controlats
		var lpendents = len(pendents) - 1
		q = &pendents[lpendents]
		if q.Final() {
			bons = append(bons, *q)
			solucionat = true
			if !TrobarTotesLesSolucions {
				sortirBucle = true
			}
		}
		pendents = pendents[:lpendents]
		pendents = append(pendents, q.NextStep()...)
		fmt.Printf("MVActual: %d Pendents: %d - Iteracions: %d - Solucionat: %t\n", q.lastNum, len(pendents), i, solucionat)
	}
	for _, q := range bons {
		q.Draw()
	}
	fmt.Printf("Total Resultat bons: %d\n", len(bons))
}
