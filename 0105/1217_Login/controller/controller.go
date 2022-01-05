package controller

import (
	"time"

	"token/db"
	"token/redis"
	"token/util"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type RegisterData struct {
	Username string `json:"username" form:"username" binding:"required,min=6,max=12"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=12"`
}

type LoginData struct {
	Username string `json:"username" form:"username" binding:"required,min=6,max=12"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=12"`
}

//註冊
func Register(c *gin.Context) {
	var user db.User
	var reg RegisterData
	// 使用者傳帳密進來，解析到struct
	err := c.ShouldBindJSON(&reg)
	if err != nil {
		util.Msg(c, util.Code_Param_Invalid, "帳號密碼格式錯誤", nil)
		return
	}
	// 確認帳號是否存在，不存在就Create
	checkOK, err := util.CheckUsername(reg.Username)
	if err != nil {
		util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#1", nil)
		return
	}
	if !checkOK {
		// 密碼加密
		regPassword, err := util.BcryptPassword(reg.Password)
		if err != nil {
			log.Warn().Caller().Err(err).Msg("BcryptPassword")
			util.Msg(c, util.Code_Param_Invalid, "資料錯誤", nil)
			return
		}
		// 創建帳密key、value
		user = db.User{Username: reg.Username, Password: regPassword}
		err = db.DB.Create(&user).Error
		if err != nil {
			log.Warn().Caller().Err(err).Msg("DB Create")
			util.Msg(c, util.Code_DB_Conn, "註冊失敗，請重新嘗試#1", nil)
			return
		}
		util.Msg(c, util.Code_ok, "註冊成功", nil)
	} else {
		// 帳號已存在
		util.Msg(c, util.Code_Param_Invalid, "帳號已存在", nil)
		return
	}
}

// 登入
func Login(c *gin.Context) {
	var login LoginData
	// 使用者傳帳密進來，解析到struct
	err := c.ShouldBindJSON(&login)
	if err != nil {
		util.Msg(c, util.Code_Param_Invalid, "帳號密碼格式錯誤", nil)
		return
	}
	// 確認帳號是否存在
	checkOK, err := util.CheckUsername(login.Username)
	if err != nil {
		log.Warn().Caller().Err(err).Msg("DB Where.Find")
		util.Msg(c, util.Code_Param_Invalid, "資料存取錯誤#1", nil)
		return
	}
	// check = 0，資料不存在，check = 1，資料已存在
	if checkOK {
		data, err := util.GetUserData(login.Username)
		if err != nil {
			log.Warn().Caller().Err(err).Msg("DB Where.Find")
			util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#1", nil)
			return
		}
		user := data.(db.User)
		// 資料庫取出的密碼 與 使用者輸入的密碼 驗證
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if err != nil {
			log.Warn().Caller().Err(err).Msg("CompareHashAndPassword")
			util.Msg(c, util.Code_Param_Invalid, "密碼錯誤", nil)
			return
		}
		// 密碼對就生成Session、UserName加進Token回傳
		sessionID, err := util.BcryptPassword(user.Username)
		if err != nil {
			log.Warn().Caller().Err(err).Msg("BcryptPassword")
			util.Msg(c, util.Code_Param_Invalid, "參數無效", nil)
			return
		}
		token, err := util.GenToken(sessionID, user.Username)
		if err != nil {
			log.Warn().Caller().Err(err).Msg("GenToken")
			util.Msg(c, util.Code_Param_Invalid, "Token創建失敗", nil)
			return
		}
		// 把Token傳到前端cookie
		// cookie名字、cookie內容、存活時間(直接打數字)，設定-1 = 刪除、path使用路經、作用host、是否只能https協議發送到服務端、HttpOnly=true 就不能被js獲取到、
		c.SetCookie("token", token, util.Tk.CookieTime, "/", "localhost", false, true)
		// 創好的Token當Key，session ID當Value存進Redis
		err = redis.Client.Set(sessionID, user.Username, time.Duration(util.Tk.TokenExpireDuration)*time.Second).Err()
		if err != nil {
			log.Warn().Caller().Err(err).Msg("Redis Set")
			util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#2", nil)
			return
		}
		util.Msg(c, util.Code_ok, user.Username, token)
	} else {
		util.Msg(c, util.Code_Param_Invalid, "查無此帳號", nil)
		return
	}
}

// 登出
func Logout(c *gin.Context) {
	// token, ok := c.Get("token")
	// if !ok {
	// 	util.Msg(c, util.Code_Param_Invalid, "請先登錄帳號", nil)
	// 	return
	// }
	// username, ok := c.Get("username")
	// if !ok {
	// 	util.Msg(c, util.Code_Param_Invalid, "請先登錄帳號", nil)
	// 	return
	// }
	var sess util.SessionData
	err := c.ShouldBindJSON(&sess)
	if err != nil {
		util.Msg(c, 2, "Auth Error", nil)
		return
	}
	// c.SetCookie("token", sess.Token, -1, "/", "localhost", false, true)
	// sessionid, ok := c.Get("sessionid")
	// if !ok {
	// 	util.Msg(c, util.Code_Param_Invalid, "驗證失敗", nil)
	// 	return
	// }
	err = redis.Client.Del(sess.SessionID, sess.Username).Err()
	if err != nil {
		log.Warn().Caller().Err(err).Msg("Redis Del")
		util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#2", nil)
		return
	}
	util.Msg(c, util.Code_ok, "登出成功", nil)
}

// 確認當前使用者名稱
func LoginStruct(c *gin.Context) {
	var ok bool
	u, ok := c.Get("username")
	if !ok {
		util.Msg(c, util.Code_Param_Invalid, "", nil)
		return
	}
	username, ok := u.(string)
	if !ok {
		util.Msg(c, util.Code_Param_Invalid, "", nil)
		return
	}
	util.Msg(c, util.Code_ok, username, nil)
}
