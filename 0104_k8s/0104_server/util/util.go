package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type API_Error struct {
	Code int
	Msg  string
	Data interface{}
}

func Msg(c *gin.Context, code int, msg string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, API_Error{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

const Code_ok = 1
const Code_Param_Invalid = 2
const Code_DB_Conn = 3
const Code_Session_Invalid = 4
const Code_Server_Invalid = 5

// Token----------------------------------------
// func GetCookie(c *gin.Context) {
// 	token, err := c.Cookie("token")
// 	if err != nil {
// 		fmt.Println(err)
// 		Msg(c, Code_Param_Invalid, "請重新登錄", nil)
// 		return
// 	}
// 	url := "http://localhost:8082/middlewareAuth"
// 	method := "GET"

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, nil)
// 	if err != nil {
// 		log.Warn().Caller().Err(err).Msg("NewRequest Error")
// 		Msg(c, 5, "請求無效", nil)
// 		return
// 	}
// 	req.Header.Add("Content-Type", "application/json")
// 	req.AddCookie(&http.Cookie{Name: "token", Value: token})
// }

// // 生成token ---------------------------------------------
// func GenToken(SessionID, username string) (string, error) {
// 	file, err := os.Open("./config/TokenData.json")
// 	if err != nil {
// 		return "", err
// 	}
// 	data := json.NewDecoder(file)
// 	err = data.Decode(&Tk)
// 	if err != nil {
// 		return "", err
// 	}

// 	t := MyClaims{
// 		SessionID,
// 		username, // 自訂Header
// 		jwt.StandardClaims{ // 設定payload
// 			ExpiresAt: time.Now().Add(time.Duration(Tk.TokenExpireDuration) * time.Second).Unix(),
// 			Issuer:    "Larry",
// 		},
// 	}
// 	// 選擇編碼模式
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
// 	// 用指定的SecretKey加密獲得Token字串
// 	return token.SignedString([]byte(Tk.MySecret))
// }

// // 解析Token ---------------------------------------------
// func ParseToken(tokenString string) (*MyClaims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(Tk.MySecret), nil
// 	})
// 	if err != nil {
// 		expired := strings.Contains(err.Error(), "token is expired")
// 		if expired {
// 			return token.Claims.(*MyClaims), err
// 		}
// 		return nil, err
// 	}
// 	// 驗證claims正確就回傳
// 	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
// 		return claims, nil
// 	}
// 	return nil, errors.New("Invalid Token")
// }

// // 生成加密亂碼 --------------------------------------------
// func BcryptPassword(data string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	bcryptString := string(hash)
// 	return bcryptString, nil
// }
