package main

import "strategy/pkg"

var (
	start      = 10
	end        = 200
	strategies = []pkg.Strategy{
		&pkg.PublicStrategy{},
		&pkg.RouteStrategy{},
		&pkg.WalkStrategy{},
	}
)

func main() {
	navigator := &pkg.Navigator{}
	for _, s := range strategies {
		navigator.SetStrategy(s)
		navigator.Route(start, end)
	}
}
