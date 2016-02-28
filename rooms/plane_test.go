package rooms

import (
	"math/rand"
	"testing"
)

func TestPlanePrinting(t *testing.T) {
	var testPlane Plane
	obfunk := []func(int) Object{
		HorizontalDoubleWall,
		VerticalDoubleWall,
	}
	for i := 0; i < 20; i++ {
		testPlane.New(rand.Intn(20)+1, rand.Intn(20)+1)
		for o := 0; o < i; o++ {
			x := rand.Intn(testPlane.Width())
			y := rand.Intn(testPlane.Height())
			obj := obfunk[rand.Intn(len(obfunk))](rand.Intn(10))
			testPlane.AddObject(x, y, obj)
		}
		testPlane.Print()
	}
}
