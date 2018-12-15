package haversine

import (
	"math"
)

// 地球的平均半径
const (
	SphereRadiusInMetres float64 = 6371000 // 以米为单位的地球平均半径
	SphereRadiusKM       float64 = 6371    // 以千米为单位的地球平均半径
	SphereRadiusMile     float64 = 3958    // 以英里为单位的地球平均半径
)

// DegreesToRadians 度数转弧度
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Point 球体上以经纬度表示的点
type Point struct {
	Lat float64 // Latitude 维度
	Lon float64 // Longitude 经度
}

// toRadians 点的方法,把度数转换为弧度
func (p Point) toRadians() Point {
	return Point{
		Lat: DegreesToRadians(p.Lat),
		Lon: DegreesToRadians(p.Lon),
	}
}

// Delta 点的方法,计算两点的弧度差异
func (p Point) Delta(point Point) Point {

	return Point{
		Lat: p.Lat - point.Lat,
		Lon: p.Lon - point.Lon,
	}
}

// Distance 点的方法,计算与另外一点的距离
func (p Point) Distance(remote Point) (float64, float64, float64) {
	return Distance(p, remote)
}

// Distance 计算两点之间的距离,分别以米,千米,英里为单位返回
func Distance(origin, position Point) (m, km, mi float64) {
	// 两个点的经纬度都转换为弧度
	o := origin.toRadians()
	p := position.toRadians()
	// delta 计算出两个点的弧度差异,作为一个虚拟点
	d := o.Delta(p)
	a := math.Pow(math.Sin(d.Lat/2), 2) + math.Pow(math.Sin(d.Lon/2), 2)*math.Cos(o.Lat)*math.Cos(p.Lat)
	// c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	c := 2 * math.Asin(math.Sqrt(a))

	m = float64(SphereRadiusInMetres * c)
	km = float64(SphereRadiusKM * c)
	mi = float64(SphereRadiusMile * c)
	return
}
