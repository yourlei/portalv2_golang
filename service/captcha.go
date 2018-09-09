// package service
// captch module
// base64图片验证码
package service
import (
	"fmt"
	"time"
	"portal/database"
	"github.com/mojocn/base64Captcha"
)

var captchaList = make(map[string]string)
// 生成字符+数字的base64码
func GenerateCaptcha() (string, string){
	// config struct for Character
	// 字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height:             40,
		Width:              100,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumberAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         5,
	}
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	// 验证码id缓存到redis
	redisConn := database.RedisPool.Get()
	defer redisConn.Close()
	if _, err1 := redisConn.Do("PING"); err1 != nil {
		fmt.Println("error: ", err1)
	}
 
	if _, err := redisConn.Do("SET", idKeyC, time.Now().Unix(), "EX", 60); err != nil {
		fmt.Println("redis set failed:", err)
	}
	// 保存本次生成的验证码id
	go StoreCaptcha(idKeyC)

	return idKeyC, base64stringC
}


// VerifyCaptcha
// 检查输入的验证码是否正确
// @params: 
//   idKey: 验证码id
//   verifyValue: 验证码字符
// @return 
//   int 错误码 
//   string 错误信息
//
func VerifyCaptcha(idkey, verifyValue string) (int, string) {
	// 检查验证码是否失效
	if captchaList[idkey] == "" {
		return 10008, "该验证码已过期"
	}

	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
		// succuss
		return 0, "验证成功"
	} else {
		// fail
		return 10007, "验证码输入不正确或已失效"
	}
}
// StoreCaptcha
// @params uuid: 生成的验证码id
// 创建定时任务,60s后将该验证码从captchaList表移除
func StoreCaptcha(uuid string) {
	// 存储验证码id
	captchaList[uuid] = "1"

	timer := time.NewTimer(60 * time.Second)
	<-timer.C
	delete(captchaList, uuid)
}