package pkg

import "fmt"

type WalkStrategy struct {
	//
}

func (w *WalkStrategy) Route(startPoint, endPoint int) {
	avSpeed := 5
	trafficJam := 1
	totalDistance := endPoint - startPoint
	totalTime := totalDistance * trafficJam / avSpeed

	fmt.Printf("Road A: [%d] to Road B: [%d], Traffic Jam: [%d], Average Speed: [%d]\n", startPoint, endPoint, trafficJam, avSpeed)
	fmt.Println("Total time: ", totalTime)
}
