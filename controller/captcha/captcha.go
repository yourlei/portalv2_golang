package captcha

import (
	"portal/service"

	"github.com/gin-gonic/gin"
)
// 生成图片验证码
func CreatePngData(c *gin.Context)  {
	uuid, base64String := service.GenerateCaptcha()
	
	c.JSON(200, gin.H{
		"code": 0,
		"error": gin.H{
			"msg": "",
		},
		"data": gin.H{
			"uuid": uuid, 
			"image": base64String,
		},
	})
}