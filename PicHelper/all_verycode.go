package PicHelper

import (
	"fmt"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

/*
线上案例地址
在线Demo [Playground Powered by Vuejs+elementUI+Axios](http://captcha.mojotv.cn)
*/
//configJsonBody json request body.
type ConfigJsonBody struct {
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
func GenerateCaptchaHandler(param ConfigJsonBody) (string, string, error) {
	var driver base64Captcha.Driver
	//choose driver
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
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	return id, b64s, err
}

// base64Captcha verify http handler
func Verify(param ConfigJsonBody) bool {
	var _rt = false
	if store.Verify(param.Id, param.VerifyValue, true) {
		_rt = true
	}
	return _rt
}

/*
数字验证码
*/
func Pic_verifycode_digit(w http.ResponseWriter, r *http.Request) (string, string) {
	idKeyD, base64stringD, _ := GenerateCaptchaHandler(ConfigJsonBody{
		Id:          "222222",
		CaptchaType: "string",
		//VerifyValue   string
		//DriverAudio * base64Captcha.DriverAudio,
		//DriverString  *base64Captcha.DriverString
		//DriverChinese *base64Captcha.DriverChinese
		//DriverMath    *base64Captcha.DriverMath
		DriverDigit: &base64Captcha.DriverDigit{
			Height:   80,
			Width:    240,
			MaxSkew:  0.7,
			DotCount: 80,
			Length:   4,
		},
	})
	//base64stringD := "11111" //base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return idKeyD, base64stringD
}

/*
数字+字母验证码
*/
func Pic_verifycode_character(w http.ResponseWriter, r *http.Request) (string, string) {
	idKeyD, base64stringD, _err := GenerateCaptchaHandler(ConfigJsonBody{
		Id:          "222222",
		CaptchaType: "string",
		//VerifyValue   string
		//DriverAudio * base64Captcha.DriverAudio,
		//DriverString  *base64Captcha.DriverString
		//DriverChinese *base64Captcha.DriverChinese
		//DriverMath    *base64Captcha.DriverMath
		DriverDigit: &base64Captcha.DriverDigit{
			Height:   80,
			Width:    240,
			MaxSkew:  0.7,
			DotCount: 80,
			Length:   4,
		},
	})
	fmt.Println(_err)
	//base64stringD := "11111" //base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return idKeyD, base64stringD
}

/*
数字+字母验证码
*/
func Pic_verifycode_audio(w http.ResponseWriter, r *http.Request) (string, string) {
	idKeyD, base64stringD, _ := GenerateCaptchaHandler(ConfigJsonBody{
		Id:          "222222",
		CaptchaType: "string",
		//VerifyValue   string
		//DriverAudio * base64Captcha.DriverAudio,
		//DriverString  *base64Captcha.DriverString
		//DriverChinese *base64Captcha.DriverChinese
		//DriverMath    *base64Captcha.DriverMath
		DriverDigit: &base64Captcha.DriverDigit{
			Height:   80,
			Width:    240,
			MaxSkew:  0.7,
			DotCount: 80,
			Length:   4,
		},
	})
	//base64stringD := "11111" //base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return idKeyD, base64stringD
}
