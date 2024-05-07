package pkg

import "fmt"

type RouteStrategy struct {
	//
}

func (r *RouteStrategy) Route(startPoint, endPoint int) {
	fmt.Println("Route strategy: %", startPoint, endPoint)
	avSpeed := 50
	trafficJam := 10
	totalDistance := endPoint - startPoint
	totalTime := totalDistance * trafficJam / avSpeed

	fmt.Println("Total distance: ", totalDistance)
	fmt.Println("Total time: ", totalTime)
}
