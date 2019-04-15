package bdsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var TOKEN_SERVER string = "https://aip.baidubce.com/oauth/2.0/token"
var FACE_SERVER string = "https://aip.baidubce.com/rest/2.0/face/v3/detect"

var (
	TYPE_IMG_BASE64 = "BASE64"
	TYPE_IMG_URL    = "URL"
	TYPE_IMG_TOKEN  = "FACE_TOKEN"
)

type Token struct {
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	SessionKey   string `json:"session_key"`
	AccessToken  string `json:"access_token"`
	SessionScr   string `json:"session_secret"`
}

type Request struct {
	Img       string `json:"image"`
	ImgTyp    string `json:"image_type"`
	FaceField string `json:"face_field"`
}

type Response struct {
	ErrorCode int     `json:"error_code"`
	ErrorMsg  string  `json:"error_msg"`
	LogId     int64   `json:"log_id"`
	TimeStamp int64   `json:"timestamp"`
	Cached    int     `json:"cached"`
	Result    Results `json:"result"`
}
type Results struct {
	FaceNum  int64        `json:"face_num"`
	FaceList []FacePoints `json:"face_list"`
}

type Location struct {
	Left     float64 `json:"left"`
	Top      float64 `json:"top"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`
}

type Angle struct {
	Yaw   float64 `json:"yaw"`
	Pitch float64 `json:"pitch"`
	Roll  float64 `json:"roll"`
}

type Types struct {
	Type        string  `json:"type"`
	Probability float64 `json:"probability"`
}

type EyeStatus struct {
	Left  float64 `json:"left_eye"`
	Right float64 `json:"right_eye"`
}

type LandMark150 struct {
	CheekRight1    Coordinate `json:"cheek_right_1"`
	CheekRight3    Coordinate `json:"cheek_right_3"`
	CheekRight7    Coordinate `json:"cheek_right_7"`
	Chin1          Coordinate `json:"chin_1"`
	Chin2          Coordinate `json:"chin_2"`
	Chin3          Coordinate `json:"chin_3"`
	CheekLeft1     Coordinate `json:"cheek_left_1"`
	CheekLeft7     Coordinate `json:"cheek_left_7"`
	CheekLeft3     Coordinate `json:"cheek_left_3"`
	EyeRCornerR    Coordinate `json:"eye_right_corner_right"`
	EyeRELUp2      Coordinate `json:"eye_right_eyelid_upper_2"`
	EyeRELUp4      Coordinate `json:"eye_right_eyelid_upper_4"`
	EyeRELL4       Coordinate `json:"eye_right_eyelid_lower_4"`
	EyeRCornerL    Coordinate `json:"eye_right_corner_left"`
	EyeRBallCenter Coordinate `json:"eye_right_eyeball_center"`
	EyeBrowRCR     Coordinate `json:"eyebrow_right_corner_right"`
	EyeBrowRUp5    Coordinate `json:"eyebrow_right_upper_5"`
	EyeBrowRUp4    Coordinate `json:"eyebrow_right_upper_4"`
	EyeBrowRUp3    Coordinate `json:"eyebrow_right_upper_3"`
	EyeBrowRUp2    Coordinate `json:"eyebrow_right_upper_2"`
	EyeBrowRCL     Coordinate `json:"eyebrow_right_corner_left"`
	EyeBrowRL1     Coordinate `json:"eyebrow_right_lower_1"`
	EyeBrowRL2     Coordinate `json:"eyebrow_right_lower_2"`
	EyeBrowRL3     Coordinate `json:"eyebrow_right_lower_3"`
	EyeLCornerR    Coordinate `json:"eye_left_corner_right"`
	EyeLELUp2      Coordinate `json:"eye_left_eyelid_upper_2"`
	EyeLELUp4      Coordinate `json:"eye_left_eyelid_upper_4"`
	EyeLELL4       Coordinate `json:"eye_left_eyelid_lower_4"`
	EyeLCornerL    Coordinate `json:"eye_left_corner_left"`
	EyeLBallCenter Coordinate `json:"eye_left_eyeball_center"`
	EyeBrowLCR     Coordinate `json:"eyebrow_left_corner_right"`
	EyeBrowLUp2    Coordinate `json:"eyebrow_left_upper_2"`
	EyeBrowLUp3    Coordinate `json:"eyebrow_left_upper_3"`
	EyeBrowLUp4    Coordinate `json:"eyebrow_left_upper_4"`
	EyeBrowLUp5    Coordinate `json:"eyebrow_left_upper_5"`
	EyeBrowLCL     Coordinate `json:"eyebrow_left_corner_left"`
	EyeBrowLL1     Coordinate `json:"eyebrow_left_lower_1"`
	EyeBrowLL2     Coordinate `json:"eyebrow_left_lower_2"`
	EyeBrowLL3     Coordinate `json:"eyebrow_left_lower_3"`
	NoseRC1        Coordinate `json:"nose_right_contour_1"`
	NoseRC4        Coordinate `json:"nose_right_contour_4"`
	NoseRC5        Coordinate `json:"nose_right_contour_5"`
	NoseLC1        Coordinate `json:"nose_left_contour_1"`
	NoseLC4        Coordinate `json:"nose_left_contour_4"`
	NoseLC5        Coordinate `json:"nose_left_contour_5"`
	MouthCRO       Coordinate `json:"mouth_corner_right_outer"`
	MouthLUO6      Coordinate `json:"mouth_lip_upper_outer_6"`
	MouthCLO       Coordinate `json:"mouth_corner_left_outer"`
	MouthLLO6      Coordinate `json:"mouth_lip_lower_outer_6"`
	MouthLUI6      Coordinate `json:"mouth_lip_upper_inner_6"`
	MouthLLI6      Coordinate `json:"mouth_lip_lower_inner_6"`
	NoseBridge1    Coordinate `json:"nose_bridge_1"`
	NoseMidContour Coordinate `json:"nose_middle_contour"`
}

type FacePoints struct {
	FaceToken   string      `json:"face_token"`
	Locat       Location    `json:"location"`
	FaceProb    float64     `json:"face_probability"`
	Angle       Angle       `json:"angle"`
	Age         int32       `json:"age"`
	Beauty      float64     `json:"beauty"`
	Expression  Types       `json:"expression"`
	Faceshap    Types       `json:"face_shape"`
	Gender      Types       `json:"gender"`
	Glasses     Types       `json:"glasses"`
	Landmark150 LandMark150 `json:"landmark150"`
	Race        Types       `json:"race"`
	FaceType    Types       `json:"face_type"`
	Emotion     Types       `json:"emotion"`
	Eyestatus   EyeStatus   `json:"eye_status"`
}

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func GetToken(key, scr string) (string, error) {

	url := TOKEN_SERVER + "?grant_type=client_credentials&client_id=" + key + "&client_secret=" + scr + "&"

	rsp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	buf := new(bytes.Buffer)

	io.Copy(buf, rsp.Body)

	token := new(Token)

	if err = json.Unmarshal(buf.Bytes(), token); err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

//URL/FACE_TOKEN/BASE64
func GetFace(key, scr, img, typ, fields string) (*Response, error) {

	token, err := GetToken(key, scr)

	if err != nil {
		return nil, err
	}

	url := FACE_SERVER + "?access_token=" + token

	client := &http.Client{}

	reqParam := &Request{
		Img:       img,
		ImgTyp:    typ,
		FaceField: fields,
	}

	byt, err := json.Marshal(reqParam)
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(byt)

	req, err := http.NewRequest("POST", url, buff)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	rsp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()

	buf := new(bytes.Buffer)

	io.Copy(buf, rsp.Body)

	response := new(Response)

	if err = json.Unmarshal(buf.Bytes(), response); err != nil {
		return nil, err
	}

	if response.ErrorCode != 0 {
		return nil, errors.New(response.ErrorMsg)
	}

	return response, nil
}
