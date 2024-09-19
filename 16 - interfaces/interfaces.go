package main

type Vehicle interface {
	// starts the vehicle and returns true if it started successfully
	start() bool

	// stsrts driving the vehicle and returns the speed
	drive() float32
}

type Car struct {
}

type Motorbike struct {
}

func (c Car) start() bool {
	return true
}

func (c Car) drive() float32 {
	return 100
}

func (m Motorbike) start() bool {
	return true
}

func (m Motorbike) drive() float32 {
	return 80
}

func generic(interf interface{}) {
	switch interf.(type) {
	case Car:
		println("Car")
	case Motorbike:
		println("Motorbike")
	default:
		println("Unknown")
	}
}

func main() {
	var v Vehicle
	v = Car{}
	v.start()
	v.drive()

	v = Motorbike{}
	v.start()
	v.drive()

	generic(v)

}
