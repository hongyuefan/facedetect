package features

import (
	sdk "github.com/hongyuefan/facedetect/bdsdk"
)

type Feature struct {
	Age     int32   `json:"age"`
	Gender  string  `json:"gender"`
	Beauty  float64 `json:"beauty"`
	Race    string  `json:"yellow"`
	Emotion string  `json:"emotion"`
	Ebrow   EyeBrow `json:"eyebrow"`
	Ey      Eye     `json:"eye"`
	Ns      Nose    `json:"nose"`
	Fc      Face    `json:"face"`
	Mth     Mouth   `json:"mouth"`
	Ch      Chin    `json:"chin"`
	RZ      RenZ    `json:"renzhong"`
}

type EyeBrow struct {
	Angle  float64 `json:"eyebrow_angle"`
	Width  float64 `json:"eyebrow_width"`
	Length float64 `json:"eyebrow_length"`
}

type Eye struct {
	Length float64 `json:"eye_length"`
	Width  float64 `json:"eye_width"`
	ToBrow float64 `json:"eye_to_brow"`
	ToEye  float64 `json:"eye_to_eye"`
}

type Nose struct {
	Length float64 `json:"nose_length"`
	Width  float64 `json:"nose_width"`
	Nasi   float64 `json:"nose_nasi"`
}

type Face struct {
	Shape  string  `json:"face_shape"`
	Length float64 `json:"face_length"`
	Width  float64 `json:"face_width"`
}

type Mouth struct {
	Length   float64 `json:"mouth_length"`
	Width    float64 `json:"mouth_width"`
	Lips     float64 `json:"mouth_lips"`
	ToNose   float64 `json:"mouth_to_nose"`
	LipsRate float64 `json:"mouth_lips_rate"`
}

type Chin struct {
	Length float64 `json:"chin_length"`
	Width  float64 `json:"chin_width"`
}

type RenZ struct {
	Length float64 `json:"renzhong_length"`
}

func GetFeatures(key, scr, url string) (sdk.FacePoints, error) {

	rsp, err := sdk.GetFace(key, scr, url, sdk.TYPE_IMG_URL, "age,beauty,expression,face_shape,gender,glasses,landmark150,race,quality,eye_status,emotion,face_type")

	return rsp.Result.FaceList[0], err
}

func GetFeature(key, scr, url string) (*Feature, error) {
	face, err := sdk.GetFace(key, scr, url, sdk.TYPE_IMG_URL, "age,beauty,expression,face_shape,gender,glasses,landmark150,race,quality,eye_status,emotion,face_type")
	if err != nil {
		return nil, err
	}

	ff := face.Result.FaceList[0]

	ft := new(Feature)

	ft.Age = ff.Age
	ft.Beauty = ff.Beauty
	ft.Gender = ff.Gender.Type
	ft.Emotion = ff.Emotion.Type
	ft.Race = ff.Race.Type

	ft.Ebrow.Angle = AngleEyeBrow(ff.Landmark150)
	ft.Ebrow.Length = RateEyeBrowEye(ff.Landmark150)
	ft.Ebrow.Width = RateEyeBrow(ff.Landmark150)

	ft.Ey.Width = RateEyeWidth(ff.Landmark150)
	ft.Ey.ToBrow = RateEyeToBrow(ff.Landmark150)
	ft.Ey.ToEye = RateEyeDistance(ff.Landmark150)

	ft.Ns.Width = RateNoseWidth(ff.Landmark150)
	ft.Ns.Nasi = RateNoseEagle(ff.Landmark150)

	ft.Mth.Length = RateMouseLength(ff.Landmark150)
	ft.Mth.Width = RateMouthLipThickness(ff.Landmark150)
	ft.Mth.Lips = RateMouthLip(ff.Landmark150)
	ft.Mth.ToNose = RateChinLength(ff.Landmark150)
	ft.Mth.LipsRate = RateMouthLip(ff.Landmark150)

	ft.Fc.Shape = ff.Faceshap.Type
	ft.Fc.Width = RateFaceWidth(ff.Landmark150)
	ft.Fc.Length = RateFaceLength(ff.Landmark150)

	ft.Ch.Length = RateChinLength(ff.Landmark150)
	ft.Ch.Width = RateChinWidth(ff.Landmark150)

	ft.RZ.Length = RateRenZLength(ff.Landmark150)

	return ft, nil
}
