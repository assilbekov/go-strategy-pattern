package pkg

import "fmt"

type RouteStrategy struct {
	//
}

func (r *RouteStrategy) Route(startPoint, endPoint int) {
	avSpeed := 50
	trafficJam := 10
	totalDistance := endPoint - startPoint
	totalTime := totalDistance * trafficJam / avSpeed

	fmt.Printf("Road A: [%d] to Road B: [%d], Traffic Jam: [%d], Average Speed: [%d]\n", startPoint, endPoint, trafficJam, avSpeed)
	fmt.Println("Total time: ", totalTime)
}
