package features

import (
	"math"

	sdk "github.com/hongyuefan/facedetect/bdsdk"
)

var (
	RATE_EYEBROW_WIDTH      float64 = 1
	RATE_EYEBROW_LENGTH     float64 = 1
	RATE_MOUTH_LENGTH       float64 = 1
	RATE_MOUTH_UPPER_WIDTH  float64 = 1
	RATE_MOUTH_LOWWER_WIDTH float64 = 1
	RATE_EYE_WIDTH          float64 = 1
	RATE_EYE_LENGTH         float64 = 1
	RATE_EYE_TO_BROW        float64 = 1
	RATE_NOSE_LENGTH        float64 = 1
	RATE_NOSE_WIDTH         float64 = 1
)

//数值越大眼越大
func RateEyeWidth(m sdk.LandMark150) float64 {
	return EyeWidth(m) / EyeLength(m) * 100
}

//数值越大眉眼距离越大
func RateEyeToBrow(m sdk.LandMark150) float64 {
	return EyeToBrow(m) / EyeWidth(m) * 100
}

//数值越大鹰钩鼻越明显
func RateNoseEagle(m sdk.LandMark150) float64 {
	return NoseEagle(m) / NoseWidth(m) * 100
}

//数值越大鼻翼越大
func RateNoseWidth(m sdk.LandMark150) float64 {
	return NoseWidth(m) / NoseBrigeWidth(m) * 100
}

//数值越大双眼距离显得越大
func RateEyeDistance(m sdk.LandMark150) float64 {
	return EyeDistance(m) / NoseWidth(m) * 100
}

//数值越大脸越宽
func RateFaceWidth(m sdk.LandMark150) float64 {
	return FaceWidth(m) / EyeWidth(m) * 100
}

//数值越大嘴越大
func RateMouseWidth(m sdk.LandMark150) float64 {
	return MouthLength(m) / NoseWidth(m) * 100
}

//数值越大嘴唇越厚
func RateMouthLipThickness(m sdk.LandMark150) float64 {
	return MouthLipThickness(m) / MouthLength(m) * 100
}

//眉尾挑高
func AngleEyeBrow(m sdk.LandMark150) float64 {
	return 180 / math.Pi * math.Atan(EyeBrowHeightDiff(m)/EyeBrowHighWidth(m))
}
