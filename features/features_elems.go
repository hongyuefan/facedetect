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

func EyeWidth(m sdk.LandMark150) float64 {
	return (distance(m.EyeRELUp4, m.EyeRELL4) + distance(m.EyeLELUp4, m.EyeLELL4)) / 2
}

func EyeDistance(m sdk.LandMark150) float64 {
	return distance(m.EyeLCornerR, m.EyeRCornerL)
}

func NoseBrigeWidth(m sdk.LandMark150) float64 {
	return distance(m.NoseLC1, m.NoseRC1)
}

func NoseLength(m sdk.LandMark150) float64 {
	return distance(m.NoseBridge1, m.NoseMidContour)
}

func NoseWidth(m sdk.LandMark150) float64 {
	return distance(m.NoseLC4, m.NoseRC4)
}

func NoseEagle(m sdk.LandMark150) float64 {
	return m.NoseMidContour.Y - (m.NoseLC5.Y+m.NoseRC5.Y)/2
}

func EyeHeightDiff(m sdk.LandMark150) float64 {
	return ((m.EyeRCornerR.Y - m.EyeRCornerL.Y) + (m.EyeLCornerL.Y - m.EyeLCornerR.Y)) / 2
}

func EyeLengthDiff(m sdk.LandMark150) float64 {
	return math.Abs((m.EyeRCornerR.X-m.EyeRCornerL.X)+(m.EyeLCornerR.X-m.EyeLCornerL.X)) / 2
}

func EyeBrowDistence(m sdk.LandMark150) float64 {
	return distance(m.EyeBrowRCL, m.EyeBrowLCR)
}

//眉尾与眉头差值
func EyeBrowHeightDiff(m sdk.LandMark150) float64 {
	return ((m.EyeBrowLCR.Y - m.EyeBrowLCL.Y) + (m.EyeBrowRCL.Y - m.EyeBrowRCR.Y)) / 2
}

func EyeBrowLengthDiff(m sdk.LandMark150) float64 {
	return math.Abs((m.EyeBrowLCL.X-m.EyeBrowLCR.X)+(m.EyeBrowRCL.X-m.EyeBrowRCR.X)) / 2
}

func EyeBrowHighWidth(m sdk.LandMark150) float64 {
	return (distance(m.EyeBrowLUp4, m.EyeBrowLL3) + distance(m.EyeBrowRUp4, m.EyeBrowRL3)) / 2
}

//眉毛最大高度
func EyeBrowNormalWidth(m sdk.LandMark150) float64 {
	return (distance(m.EyeBrowLUp2, m.EyeBrowLL1) + distance(m.EyeBrowRUp2, m.EyeBrowRL1)) / 2
}

func EyeBrowLength(m sdk.LandMark150) float64 {
	return (distance(m.EyeBrowLCL, m.EyeBrowLCR) + distance(m.EyeBrowRCL, m.EyeBrowRCR)) / 2
}

func MouthLength(m sdk.LandMark150) float64 {
	return distance(m.MouthCRO, m.MouthCLO)
}

func MouthLipThickness(m sdk.LandMark150) float64 {
	return distance(m.MouthLUO6, m.MouthLLO6)
}

func MouthHeightDiff(m sdk.LandMark150) float64 {
	return math.Abs((m.MouthCLO.Y-m.MouthLUI6.Y)+(m.MouthCRO.Y-m.MouthLUI6.Y)) / 2
}

func MouthLengthDiff(m sdk.LandMark150) float64 {
	return math.Abs(m.MouthCLO.X - m.MouthLUI6.X)
}

func MouthUpLipWidth(m sdk.LandMark150) float64 {
	return distance(m.MouthLUI6, m.MouthLUO6)
}

func MouthLowLipWidth(m sdk.LandMark150) float64 {
	return distance(m.MouthLLI6, m.MouthLLO6)
}

func MouthToNose(m sdk.LandMark150) float64 {
	return distance(m.NoseMidContour, m.MouthLUO6)
}

func EyeToBrow(m sdk.LandMark150) float64 {
	return (distance(m.EyeBrowLL2, m.EyeLELUp2) + distance(m.EyeBrowRL2, m.EyeRELUp2)) / 2
}

func FaceHeightFrame(m sdk.LandMark150) float64 {
	return (m.EyeBrowLUp3.Y - m.NoseMidContour.Y) / (m.NoseMidContour.Y - m.Chin2.Y)
}

func FaceWidth(m sdk.LandMark150) float64 {
	return (distance(m.EyeLCornerL, m.CheekLeft1) + distance(m.EyeRCornerR, m.CheekRight1)) / 2
}

func FaceMidLength(m sdk.LandMark150) float64 {
	return (math.Abs(m.EyeBrowLUp5.Y-m.NoseMidContour.Y) + math.Abs(m.EyeBrowRUp5.Y-m.NoseMidContour.Y)) / 2
}

func FaceLowLength(m sdk.LandMark150) float64 {
	return distance(m.NoseMidContour, m.Chin2)
}

func ChinLength(m sdk.LandMark150) float64 {
	return distance(m.MouthLLO6, m.Chin2)
}

func ChinWidth(m sdk.LandMark150) float64 {
	return distance(m.Chin1, m.Chin3)
}

func CheekWidth(m sdk.LandMark150) float64 {
	return distance(m.CheekLeft1, m.CheekRight1)
}
