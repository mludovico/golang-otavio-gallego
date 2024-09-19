package vehicle

import "testing"

func TestDrive(t *testing.T) {
	t.Run("Car", func(t *testing.T) {
		car := Car{}
		speed := car.drive()
		if speed != 100 {
			t.Fatalf("When driving a car, expected 100, but got %f", speed)
		}
	})
	t.Run("Motorbike", func(t *testing.T) {
		motorbike := Motorbike{}
		speed := motorbike.drive()
		if speed != 80 {
			t.Fatalf("When driving a motorbike, expected 80, but got %f", speed)
		}
	})
}
