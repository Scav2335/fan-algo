package sorts

type Point2D struct {
	X, Y float64
}

func ccw(a, b, c Point2D) int {
	area2 := (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
	if area2 < 0 {
		return -1
	} else if area2 > 0 {
		return 1
	} else {
		return 0
	}
}
