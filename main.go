package main

import (
	"fmt"
	"time"
)

const TrobarTotesLesSolucions = false
const TrobarUnCiclic = true
const IterationLog = 1000000

func main() {

	var pendents, bons, ciclics []Quadre
	var solucionat, sortirBucle bool
	var q, freq *Quadre
	var startTime time.Time
	var firstSoluctionTime time.Duration
	var MaxValue int

	startTime = time.Now()

	solucionat = false
	sortirBucle = false
	pendents = Firsts()
	freq = Init()
	MaxValue = 0

	for i := 0; len(pendents) != 0 && !sortirBucle; i++ {
		//agafem l'últim q així tindirem a avançar i tenir els grups de pendents més controlats
		var lpendents = len(pendents) - 1
		q = &pendents[lpendents]
		if q.Final() {
			if !solucionat {
				firstSoluctionTime = time.Since(startTime)
			}
			bons = append(bons, *q)
			solucionat = true
			if q.Ciclic() {
				ciclics = append(ciclics, *q)
				if !TrobarTotesLesSolucions {
					sortirBucle = true
				}
			}
			if !TrobarTotesLesSolucions && !TrobarUnCiclic {
				sortirBucle = true
			}
		}
		if MaxValue < q.lastNum {
			MaxValue = q.lastNum
		}
		pendents = pendents[:lpendents]
		pendents = append(pendents, q.NextStep()...)

		if i%IterationLog == 0 {
			fmt.Printf("MaxValue: %d Pendents: %d Iteracions: %dM Bons: %d Ciclics: %d Temps: %s\n",
				MaxValue, len(pendents), i/1000000, len(bons), len(ciclics), time.Since(startTime))
		}
	}

	if TrobarTotesLesSolucions {
		for _, q := range bons {
			hFreqBons, wFreqBons := q.WhereIsFirst()
			freq.cela[hFreqBons][wFreqBons]++
		}
		fmt.Println("Freqüencis de Bons")
		freq.Draw()
		freq = Init()
		for _, q := range ciclics {
			hFreqCiclics, wFreqCiclics := q.WhereIsFirst()
			freq.cela[hFreqCiclics][wFreqCiclics]++
		}
		fmt.Println("Freqüencia de ciclics")
		freq.Draw()

		fmt.Printf("Total Resultat bons: %d (Ciclics: %d), primera solució en %s, tot el process %s\n",
			len(bons), len(ciclics), firstSoluctionTime, time.Since(startTime))
	} else {
		q.Draw()
		fmt.Printf("Temps de calcul %s\n", time.Since(startTime))
	}

}
