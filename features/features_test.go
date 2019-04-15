package features

import (
	"encoding/json"
	"testing"

	sdk "github.com/hongyuefan/facedetect/bdsdk"
)

var face *sdk.Response

//var url string = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1553940337163&di=8cff3e7ddce8aa20e928f08e57878c73&imgtype=0&src=http%3A%2F%2Fp4.yokacdn.com%2Fpic%2Fpeople%2Fspotlight%2F2013%2FU424P41T8D202939F430DT20130609150919_maxw808.jpg"

//var url string = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1553972164983&di=14264c7f7fc2bece2f3d990bef1e645b&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F01dcd05a169503a801213490d187d7.jpg%402o.jpg"

var url string = "https://timgsa.baidu.com/timg?image&quality=80&size=b10000_10000&sec=1554716664&di=6755891a1f1713b5eee03e307289914d&src=http://gss0.baidu.com/-vo3dSag_xI4khGko9WTAnF6hhy/zhidao/pic/item/e824b899a9014c082301d7680d7b02087af4f4c1.jpg"

//var url string = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1554018064712&di=468136627bd6e8d4953efc97d118435b&imgtype=0&src=http%3A%2F%2Fg.hiphotos.baidu.com%2Fzhidao%2Fpic%2Fitem%2Fc8177f3e6709c93d8ec881ed993df8dcd10054f6.jpg"

//var url string = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1554018350528&di=36bab3d6018aeeff80b976ae586bb268&imgtype=0&src=http%3A%2F%2Fi0.sinaimg.cn%2Fty%2Fj%2Fp%2F2010-04-01%2FU2035P6T12D4917194F44DT20100401180023.jpg"

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

func TestMouthLength(t *testing.T) {
	t.Log(MouthLength(face.Result.FaceList[0].Landmark150))
}

func TestMouthUpLip(t *testing.T) {
	t.Log(MouthUpLipWidth(face.Result.FaceList[0].Landmark150))
}

func TestRateDiff(t *testing.T) {

	t.Log(RateDiff(EyeLength(face.Result.FaceList[0].Landmark150), EyeDistance(face.Result.FaceList[0].Landmark150)))
}

func TestNoseNasi(t *testing.T) {
	t.Log(NoseEagle(face.Result.FaceList[0].Landmark150))
}

func TestEyeBrowHeightDiff(t *testing.T) {
	t.Log(EyeBrowHeightDiff(face.Result.FaceList[0].Landmark150))
}

func TestEyeBrowLengthDiff(t *testing.T) {
	t.Log(EyeBrowLengthDiff(face.Result.FaceList[0].Landmark150))
}

func TestRateEyeDistance(t *testing.T) {
	t.Log("数值越大双眼距离显得越大:", RateEyeDistance(face.Result.FaceList[0].Landmark150))
}

func TestRateEyeToBrow(t *testing.T) {
	t.Log("数值越大眉眼距离越大:", RateEyeToBrow(face.Result.FaceList[0].Landmark150))
}

func TestRateEyeWidth(t *testing.T) {
	t.Log("数值越大眼越大:", RateEyeWidth(face.Result.FaceList[0].Landmark150))
}

func TestRateFaceWidth(t *testing.T) {
	t.Log("数值越大脸越宽:", RateFaceWidth(face.Result.FaceList[0].Landmark150))
}

func TestRateMouseWidth(t *testing.T) {
	t.Log("数值越大嘴越大:", RateMouseLength(face.Result.FaceList[0].Landmark150))
}

func TestRateMouthLipThickness(t *testing.T) {
	t.Log("数值越大越厚:", RateMouthLipThickness(face.Result.FaceList[0].Landmark150))
}

func TestRateNoseEagle(t *testing.T) {
	t.Log("数值越大鹰钩鼻越明显:", RateNoseEagle(face.Result.FaceList[0].Landmark150))
}

func TestRateNoseWidth(t *testing.T) {
	t.Log("数值越大鼻翼越大:", RateNoseWidth(face.Result.FaceList[0].Landmark150))
}

func TestAngleEyeBrow(t *testing.T) {
	t.Log("眉毛角度：", AngleEyeBrow(face.Result.FaceList[0].Landmark150))
}

func TestGetFeature(t *testing.T) {

	ft, err := GetFeatures("ARthitTObYeqbmIgbCnHlpDP", "CxEGjfrOqcZuF6GsqLTIUdlLXg8tO8OV", url)
	if err != nil {
		panic(err)
	}

	byt, _ := json.Marshal(ft)

	t.Log(string(byt))
}
