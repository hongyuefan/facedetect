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

	if err != nil {
		return sdk.FacePoints{}, err
	}
	return rsp.Result.FaceList[0], err
}
