package middlewares

import "github.com/gin-gonic/gin"

func LocalizationMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		lang := c.GetHeader("Accept-Language")

		if lang != "ar" {
			lang = "en"
		}

		c.Set("lang", lang)

		c.Next()
	}
}