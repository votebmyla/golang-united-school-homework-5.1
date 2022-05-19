package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() {
	// implement me

	// return s.start.y

}

func (s *Square) Area() uint {
	// implement me
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	// implement me
	return s.a * 4
}
