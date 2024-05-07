package pkg

import "fmt"

type PublicStrategy struct {
	//
}

func (r *PublicStrategy) Route(startPoint, endPoint int) {
	avSpeed := 50
	trafficJam := 1
	totalDistance := endPoint - startPoint
	totalTime := totalDistance * trafficJam / avSpeed

	fmt.Printf("Road A: [%d] to Road B: [%d], Traffic Jam: [%d], Average Speed: [%d]\n", startPoint, endPoint, trafficJam, avSpeed)
	fmt.Println("Total time: ", totalTime)
}
