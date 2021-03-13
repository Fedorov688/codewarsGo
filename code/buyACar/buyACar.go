package buyACar

import "math"

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {

	var cashInPants = 0.0
	var month = 0
	var PriceNew = float64(startPriceNew)
	var PriceOld = float64(startPriceOld)

	for cashInPants+PriceOld < PriceNew {
		PriceOld *= (100.0 - percentLossByMonth) / 100.0
		PriceNew *= (100.0 - percentLossByMonth) / 100.0
		cashInPants += float64(savingperMonth)
		if month%2 == 0 {
			percentLossByMonth += 0.5
		}
		month++
	}

	return [2]int{month, int(math.Round(cashInPants + PriceOld - PriceNew))}
}
