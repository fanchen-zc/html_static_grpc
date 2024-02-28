package service

import (
	"github.com/mojocn/base64Captcha"
)

// configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// base64Captcha create http handler
// https://captcha.mojotv.cn/
func GenerateCaptchaHandler(width, height int) map[string]interface{} {
	param := configJsonBody{
		Id:          "",
		CaptchaType: "math",
		VerifyValue: "",
		DriverMath: &base64Captcha.DriverMath{
			Height:          height,
			Width:           width,
			NoiseCount:      0,
			ShowLineOptions: 2,                            //线条干扰
			Fonts:           []string{"wqy-microhei.ttc"}, //其他字体比较难认
			//Fonts:      []string{"ApothecaryFont.ttf", "DeborahFancyDress.ttf", "Flim-Flam.ttf", "wqy-microhei.ttc"},
		},
	}
	//param := configJsonBody{
	//	Id:          "",
	//	CaptchaType: "digit",
	//	VerifyValue: "",
	//	DriverDigit: &base64Captcha.DriverDigit{
	//		Height:   80,
	//		Width:    240,
	//		Length:   4,
	//		MaxSkew:  0.7,
	//		DotCount: 80,
	//	},
	//}
	var driver base64Captcha.Driver
	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	Captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := Captcha.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	return body
}

// base64Captcha verify http handler
func CaptchaVerifyHandle(id, verify string) bool {
	//verify the captcha
	return store.Verify(id, verify, true)
}
