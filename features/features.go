package features

import (
	"math"

	sdk "github.com/hongyuefan/facedetect/bdsdk"
)

func distance(p1, p2 sdk.Coordinate) float64 {

	x := math.Abs(p1.X - p2.X)
	y := math.Abs(p1.Y - p2.Y)

	return math.Sqrt(x*x + y*y)
}

func RateDiff(origin, num float64) float64 {
	return math.Abs(num-origin) / origin * 100
}

func EyeLength(m sdk.LandMark150) float64 {

	ll := distance(m.EyeLCornerL, m.EyeLCornerR)

	rl := distance(m.EyeRCornerL, m.EyeRCornerR)

	return (ll + rl) / 2
}

func EyeDistance(m sdk.LandMark150) float64 {
	return distance(m.EyeLCornerR, m.EyeRCornerL)
}

func NoseLength(m sdk.LandMark150) float64 {
	return distance(m.NoseBridge1, m.NoseMidContour)
}

func NoseWidth(m sdk.LandMark150) float64 {
	return distance(m.NoseLC4, m.NoseRC4)
}

func MouthLength(m sdk.LandMark150) float64 {
	return distance(m.MouthCRO, m.MouthCLO)
}

func MouthUpLipWith(m sdk.LandMark150) float64 {
	return distance(m.MouthLUI6, m.MouthLUO6)
}

func MouthLowLipWith(m sdk.LandMark150) float64 {
	return distance(m.MouthLLI6, m.MouthLLO6)
}

func EyeToBrow(m sdk.LandMark150) float64 {
	return (distance(m.EyeBrowLL2, m.EyeLELUp2) + distance(m.EyeBrowRL2, m.EyeRELUp2)) / 2
}
