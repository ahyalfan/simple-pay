package util

import "math"

//bikin algoritma perpindahan dari yg lama ke yg baru

func GetDistanceTwo(lat1, lon1, lat2, lon2 float64) float64 {
	const (
		earthRadius = 6371e3 // Radius of the Earth in meters
	)

	dlat := lat2 - lat1
	dlon := lon2 - lon1

	a := (dlat/2)*(dlat/2) +
		(dlon/2)*(dlon/2)*math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := earthRadius * c

	return distance

}

// cara kedua
func GetDistance(lat1, lon1, lat2, lon2 float64) float64 {
	R := 6371e3 // Radius of the Earth in meters

	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	dlat := lat2Rad - lat1Rad
	lon1Rad := lon1 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180
	dlon := lon2Rad - lon1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c
	return distance
}

// dan masih banyak cara lagi
