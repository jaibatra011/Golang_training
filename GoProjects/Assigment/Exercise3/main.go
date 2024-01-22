package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle   vehicle
	fourWheel bool
}

type sedan struct {
	vehicle vehicle
	luxury  bool
}

func main() {
	vehicle1 := vehicle{2, "white"}
	truck1 := truck{vehicle1, true}

	vehicle2 := vehicle{4, "black"}
	sedan1 := sedan{vehicle2, true}

	fmt.Printf("%+v\n", truck1)
	fmt.Printf("%+v\n", sedan1)

	fmt.Println("Color of truck is", truck1.vehicle.color)
	fmt.Println("Color of sedan is", sedan1.vehicle.color)
}
