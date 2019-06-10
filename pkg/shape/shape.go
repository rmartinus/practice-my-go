package shape

// Shape defines shape methods
type Shape interface {
	Area() int
	ScaleBy(sc int)
}

// Square has equal lengths for each side
type Square struct {
	Length int
}

// Rectangle has width and height
type Rectangle struct {
	Height int
	Width  int
}

// Area calculates area of a square
func (s Square) Area() int {
	return s.Length * s.Length
}

// ScaleBy scales square by the given sc
func (s *Square) ScaleBy(sc int) {
	s.Length = s.Length * sc
}

// Area calculates area of a rectangle
func (r Rectangle) Area() int {
	return r.Height * r.Width
}

// ScaleBy scales rectangle by the given sc
func (r *Rectangle) ScaleBy(sc int) {
	r.Width = r.Width * sc
	r.Height = r.Height * sc
}
