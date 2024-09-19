package vehicle

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
