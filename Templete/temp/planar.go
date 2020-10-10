// 二维的，planar

// A Point is a Lon/Lat 2d point.
type Point [2]float64
// A MultiPoint represents a set of points in the 2D Eucledian or Cartesian plane.
type MultiPoint []Point

// LineString represents a set of points to be thought of as a polyline.
type LineString []Point
// MultiLineString is a set of polylines.
type MultiLineString []LineString

// Ring represents a set of ring on the earth.
type Ring LineString

// Polygon is a closed area. The first LineString is the outer ring.
// The others are the holes. Each LineString is expected to be closed
// ie. the first point matches the last.
type Polygon []Ring
// MultiPolygon is a set of polygons.
type MultiPolygon []Polygon

// A Bound represents a closed box or rectangle.
// To create a bound with two points you can do something like:
//	orb.MultiPoint{p1, p2}.Bound()
type Bound struct {
	Min, Max Point
}

// Clone will make a deep copy of the geometry.
func Clone(g Geometry) Geometry {
	...
}

// 二维关系

// Contains 包含
// Intersects 相交
// Disjoint 相离
// Touches 邻接
// Overlaps 重合
// Union 并集
// Intersection 交集
// Difference 求差


// 数据定义

// JoinStyle is the style of the joint of two line segments.
type JoinStyle int

// CapStyle is the style of the cap at the end of a line segment.
type CapStyle int