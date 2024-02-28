package Middlewares

//
//func IsLogin() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		res := e.Gin{C: c}
//		authorization := c.Request.Header.Get("Authorization")
//		if authorization == "" {
//			res.Fail(-1, "Please login first", "")
//			c.Abort()
//			return
//		}
//		//log.Println("authorization:", authorization)
//		kv := strings.Split(authorization, " ")
//		if len(kv) != 2 || kv[0] != "Bearer" {
//			res.Fail(-1, "authorization fail", "")
//			c.Abort()
//			return
//		}
//		token := kv[1]
//		vk, err := helper.JwtDncode(token, []byte(config.Configs.JwtSecret))
//		if err != nil {
//			res.Fail(-1, "authorization fail2", "")
//			c.Abort()
//			return
//		}
//		uidStr := helper.Interface2String(vk["uid"])
//		uid := int32(com.StrTo(uidStr).MustInt())
//		auto := Models.CheckAutoLoginData(uid)
//		nowTime := int32(time.Now().Unix())
//		if auto.Id <= 0 || auto.ExpireTime < nowTime {
//			res.Fail(-1, "Please login first", "")
//			c.Abort()
//			return
//		}
//
//		c.Set("uid", uid)
//		c.Next()
//	}
//}
