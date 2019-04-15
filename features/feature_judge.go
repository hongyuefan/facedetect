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

//数值越大眼越大，一般为 20
func RateEyeSize(m sdk.LandMark150) float64 {
	return EyeLength(m) / CheekWidth(m) * 100
}

//数值越大眼越圆，一般为 1/3
func RateEyeWidth(m sdk.LandMark150) float64 {
	return EyeWidth(m) / EyeLength(m) * 100
}

//眼尾挑高角度，角度越大越挑
func AngleEye(m sdk.LandMark150) float64 {
	return 180 / math.Pi * math.Atan(EyeHeightDiff(m)/EyeLengthDiff(m))
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
	return NoseWidth(m) / EyeLength(m) * 100
}

//数值越大双眼距离显得越大
func RateEyeDistance(m sdk.LandMark150) float64 {
	return EyeDistance(m) / EyeLength(m) * 100
}

//数值越大脸越宽
func RateFaceWidth(m sdk.LandMark150) float64 {
	return FaceWidth(m) / EyeLength(m) * 100
}

//数值越大嘴越大
func RateMouseLength(m sdk.LandMark150) float64 {
	return MouthLength(m) / NoseWidth(m) * 100
}

//数值越大嘴唇越厚
func RateMouthLipThickness(m sdk.LandMark150) float64 {
	return MouthLipThickness(m) / MouthLength(m) * 100
}

//眉尾挑高角度，角度越大越挑
func AngleEyeBrow(m sdk.LandMark150) float64 {
	return 180 / math.Pi * math.Atan(EyeBrowHeightDiff(m)/EyeBrowLengthDiff(m))
}

//眉毛与眼睛长度之比，越大眉毛越长
func RateEyeBrowEye(m sdk.LandMark150) float64 {
	return EyeBrowLength(m) / EyeLength(m) * 100
}

//眉毛与眉毛之间的距离与眼长之比
func RateEyeBrowToEyeBrow(m sdk.LandMark150) float64 {
	return EyeBrowDistence(m) / EyeLength(m) * 100
}

//眉毛宽度最大比值
func RateEyeBrow(m sdk.LandMark150) float64 {
	return EyeBrowHighWidth(m) / EyeBrowNormalWidth(m) * 100
}

//下嘴唇与上嘴唇之比
func RateMouthLip(m sdk.LandMark150) float64 {
	return MouthLowLipWidth(m) / MouthUpLipWidth(m) * 100
}

//嘴角角度
func AngleMouth(m sdk.LandMark150) float64 {
	return 180 / math.Pi * math.Atan(MouthHeightDiff(m)/MouthLengthDiff(m))
}

//嘴唇到鼻子距离与嘴厚度比，数值越大距离越远
func RateMouseToNose(m sdk.LandMark150) float64 {
	return MouthToNose(m) / MouthLipThickness(m) * 100
}

//脸中庭与下庭之比，越大脸越长
func RateFaceLength(m sdk.LandMark150) float64 {
	return FaceMidLength(m) / FaceLowLength(m) * 100
}

//下巴宽度与鼻子宽度之比，数值越大下巴越宽
func RateChinWidth(m sdk.LandMark150) float64 {
	return ChinWidth(m) / NoseWidth(m) * 100
}

//嘴到下巴与鼻尖到嘴之比
func RateChinLength(m sdk.LandMark150) float64 {
	return distance(m.MouthLLI6, m.Chin2) / distance(m.NoseMidContour, m.MouthLUI6) * 100
}

//人中长度
func RateRenZLength(m sdk.LandMark150) float64 {
	return distance(m.NoseMidContour, m.MouthLUI6) / distance(m.MouthLLI6, m.Chin2) * 100
}
