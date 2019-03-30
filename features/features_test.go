package features

import (
	"testing"

	sdk "github.com/hongyuefan/facedetect/bdsdk"
)

var face *sdk.Response

var url string = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1553865142718&di=9ba81753dfa76365016b3a4f8b556870&imgtype=0&src=http%3A%2F%2Fgss0.baidu.com%2F94o3dSag_xI4khGko9WTAnF6hhy%2Fzhidao%2Fpic%2Fitem%2F8644ebf81a4c510fd9e29b926159252dd42aa517.jpg"

func init() {

	var err error

	face, err = sdk.GetFace("ARthitTObYeqbmIgbCnHlpDP", "CxEGjfrOqcZuF6GsqLTIUdlLXg8tO8OV", url, sdk.TYPE_IMG_URL, "age,beauty,expression,face_shape,gender,glasses,landmark150,race,quality,eye_status,emotion,face_type")

	if err != nil {
		panic(err)
	}
}

func TestEyeLength(t *testing.T) {

	t.Log(EyeLength(face.Result.FaceList[0].Landmark150))
}

func TestEyeDistance(t *testing.T) {

	t.Log(EyeDistance(face.Result.FaceList[0].Landmark150))
}

func TestNoseWidth(t *testing.T) {

	t.Log(NoseWidth(face.Result.FaceList[0].Landmark150))
}

func TestNoseLength(t *testing.T) {
	t.Log(NoseLength(face.Result.FaceList[0].Landmark150))
}

func TestRateDiff(t *testing.T) {

	t.Log(RateDiff(EyeLength(face.Result.FaceList[0].Landmark150), EyeDistance(face.Result.FaceList[0].Landmark150)))
}
