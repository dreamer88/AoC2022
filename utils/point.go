package utils

import "math"

type Point2 struct {
	X, Y int
}

type Point3 struct {
	X, Y, Z int
}

type Rect struct {
	Min, Max Point2
}

type Cube struct {
	Min, Max Point3
}

type Point2Set = *HashSet[Point2]
type Point3Set = *HashSet[Point3]

func NewPoint2Set() Point2Set {
	return NewHashSet[Point2]()
}

func NewPoint3Set() Point3Set {
	return NewHashSet[Point3]()
}

func NewPoint2(x int, y int) Point2 {
	return Point2{x, y}
}

func (p Point2) Add(other Point2) Point2 {
	return Point2{p.X + other.X, p.Y + other.Y}
}

func (p Point2) Sub(other Point2) Point2 {
	return Point2{p.X - other.X, p.Y - other.Y}
}

func (p Point2) Up() Point2 {
	return Point2{p.X, p.Y - 1}
}

func (p Point2) Down() Point2 {
	return Point2{p.X, p.Y + 1}
}

func (p Point2) Left() Point2 {
	return Point2{p.X - 1, p.Y}
}

func (p Point2) Right() Point2 {
	return Point2{p.X + 1, p.Y}
}

func ContainingRect(a Point2, b Point2) Rect {
	x_0, x_1 := MinMax(a.X, b.X)
	y_0, y_1 := MinMax(a.Y, b.Y)

	return Rect{Point2{x_0, y_0}, Point2{x_1, y_1}}
}

func EmptyBoundingRect() Rect {
	return Rect{
		Point2{math.MaxInt, math.MaxInt},
		Point2{math.MinInt, math.MinInt},
	}
}

func (r *Rect) UpdateWithPoint(p Point2) {
	r.Min.X = Min(r.Min.X, p.X)
	r.Min.Y = Min(r.Min.Y, p.Y)
	r.Max.X = Max(r.Max.X, p.X)
	r.Max.Y = Max(r.Max.Y, p.Y)
}

func (r Rect) ContainsPoint(p Point2) bool {
	return r.Min.X <= p.X && p.X <= r.Max.X &&
		r.Min.Y <= p.Y && p.Y <= r.Max.Y
}

func (r Rect) WrapPoint(p Point2) Point2 {
	w, h := r.Max.X-r.Min.X+1, r.Max.Y-r.Min.Y+1
	return Point2{SaneMod(p.X-r.Min.X, w) + r.Min.X, SaneMod(p.Y-r.Min.Y, h) + r.Min.Y}
}

func (r Rect) IteratePoints(each_point func(Point2), end_row func(int)) {
	for y := r.Min.Y; y <= r.Max.Y; y++ {
		for x := r.Min.X; x <= r.Max.X; x++ {
			each_point(Point2{x, y})
		}

		if end_row != nil {
			end_row(y)
		}
	}
}

func NewPoint3(x int, y int, z int) Point3 {
	return Point3{x, y, z}
}

func (p Point3) Add(other Point3) Point3 {
	return Point3{p.X + other.X, p.Y + other.Y, p.Z + other.Z}
}

func (p Point3) Sub(other Point3) Point3 {
	return Point3{p.X - other.X, p.Y - other.Y, p.Z - other.Z}
}

func (p Point3) Up() Point3 {
	return Point3{p.X, p.Y - 1, p.Z}
}

func (p Point3) Down() Point3 {
	return Point3{p.X, p.Y + 1, p.Z}
}

func (p Point3) Left() Point3 {
	return Point3{p.X - 1, p.Y, p.Z}
}

func (p Point3) Right() Point3 {
	return Point3{p.X + 1, p.Y, p.Z}
}

func (p Point3) Forward() Point3 {
	return Point3{p.X, p.Y, p.Z + 1}
}

func (p Point3) Back() Point3 {
	return Point3{p.X, p.Y, p.Z - 1}
}

func EmptyBoundingCube() Cube {
	return Cube{
		Point3{math.MaxInt, math.MaxInt, math.MaxInt},
		Point3{math.MinInt, math.MinInt, math.MinInt},
	}
}

func (c *Cube) UpdateWithPoint(p Point3) {
	c.Min.X = Min(c.Min.X, p.X)
	c.Min.Y = Min(c.Min.Y, p.Y)
	c.Min.Z = Min(c.Min.Z, p.Z)
	c.Max.X = Max(c.Max.X, p.X)
	c.Max.Y = Max(c.Max.Y, p.Y)
	c.Max.Z = Max(c.Max.Z, p.Z)
}

func (c Cube) ContainsPoint(p Point3) bool {
	return c.Min.X <= p.X && p.X <= c.Max.X &&
		c.Min.Y <= p.Y && p.Y <= c.Max.Y &&
		c.Min.Z <= p.Z && p.Z <= c.Max.Z
}
