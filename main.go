package main

import "fmt"

func main() {

	var pendents, bons []Quadre
	var solucionat bool

	solucionat = false
	q := Init()
	q.First()

	pendents = append(pendents, *q)

	for i := 0; len(pendents) != 0; i++ {
		//agafem el primer
		q = &pendents[0]
		if q.Final() {
			bons = append(bons, *q)
			solucionat = true
		}
		pendents = pendents[1:]
		pendents = append(pendents, q.NextStep()...)
		fmt.Printf("Pendents: %d - Iteracions: %d - Solucionat: %t\n", len(pendents), i, solucionat)
	}
	for _, q := range bons {
		q.Draw()
	}
	fmt.Printf("Total Resultat bons: %d\n", len(bons))
}
