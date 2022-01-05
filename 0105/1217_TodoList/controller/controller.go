package controller

import (
	"net/http"
	"strconv"
	"todoList/db"
	"todoList/util"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PageData struct {
	Totle     int64 `json:"totle"`
	TotlePage int   `json:"totlepage"`
	// PageSize  int   `json:"pagesize"`
}

// type PageAll struct {
// 	Code int         `json:"code"`
// 	Msg  string      `json:"msg"`
// 	Data interface{} `json:"data"`
// 	Page interface{} `json:"page"`
// }

// func PageMsg(c *gin.Context, code int, msg string, data interface{}, page interface{}) {
// 	c.AbortWithStatusJSON(http.StatusOK, PageAll{
// 		Code: code,
// 		Msg:  msg,
// 		Data: data,
// 		Page: page,
// 	})
// }

// 新增一筆資料
func Insert(c *gin.Context) {
	var todo db.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		util.Msg(c, util.Code_Param_Invalid, "帳號密碼格式錯誤", nil)
		return
	}
	// 與UserName一起存入DB
	username := util.GetUserName(c)
	todo = db.Todo{Title: todo.Title, UserId: username}
	err = db.DB.Create(&todo).Error
	if err != nil {
		log.Warn().Caller().Err(err).Msg("DB Create")
		util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#1", nil)
		return
	}
	todos := util.FindAll(c, username)
	util.Msg(c, util.Code_ok, "新增成功", todos)
}

// 取得全部資料
func Get(c *gin.Context) {
	var todos []db.Todo
	username := util.GetUserName(c)
	err := db.DB.Where("user_id", username).Find(&todos).Error
	if err != nil {
		log.Warn().Caller().Err(err).Msg("DB Where.Find")
		util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#1", nil)
		return
	} else {
		util.Msg(c, util.Code_ok, "取得資料", todos)
	}
}

// 取得該使用者的資料總筆數、頁數、資料
func GetPage(c *gin.Context) {
	DDB := db.DB
	var pages []db.Todo
	var p PageData
	var totle int64

	// 預設4筆資料
	pageSize := 4
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		util.ErrMsg(c, util.Code_Param_Invalid, "參數無效", nil, err)
		return
	}
	username := util.GetUserName(c)
	// 資料總筆數
	err = DDB.Where("user_id = ?", username).Find(&pages).Count(&totle).Error
	if err != nil {
		log.Warn().Caller().Err(err).Msg("DB Create.Find.Count")
		util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#1", nil)
		return
	}
	strInt64 := strconv.FormatInt(totle, 10)
	t, err := strconv.Atoi(strInt64)
	if err != nil {
		return
	}
	// 總頁數
	totlePage := (t / pageSize)
	if t%pageSize == 0 {
		p = PageData{Totle: totle, TotlePage: totlePage}
	} else {
		a := totlePage + 1
		p = PageData{Totle: totle, TotlePage: a}
	}
	// 查詢資料庫
	if page > 0 && pageSize > 0 {
		DDB = db.DB.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	err = DDB.Where("user_id = ?", username).Find(&pages).Error
	if err != nil {
		log.Warn().Caller().Err(err).Msg("DB Where")
		util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#1", nil)
		return
	}
	// Msg 使用者名稱，Data ，Page 此帳號的資料總筆數、需要生成的頁面數
	// PageMsg(c, 1, username, pages, p)
	c.JSON(http.StatusOK, gin.H{
		"Code": util.Code_ok,
		"Msg":  username,
		"Data": pages,
		"Page": p,
	})
}

// 更新資料
func Put(c *gin.Context) {
	var todo db.Todo
	id, ok := c.Params.Get("id")
	if !ok {
		util.Msg(c, util.Code_Param_Invalid, "id無效", nil)
		return
	}
	err := db.DB.Where("id =?", id).First(&todo).Error
	if err != nil {
		log.Warn().Caller().Err(err).Msg("DB Where.First")
		util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#1", nil)
		return
	}
	err = c.BindJSON(&todo)
	if err != nil {
		util.ErrMsg(c, util.Code_Param_Invalid, "格式錯誤", nil, err)
		return
	}
	err = db.DB.Model(&todo).Where("id = ?", todo.Id).Update("title", todo.Title).Error
	if err != nil {
		log.Warn().Caller().Err(err).Msg("DB Save")
		util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#1", nil)
		return
	}
	username := util.GetUserName(c)
	todos := util.FindAll(c, username)
	util.Msg(c, util.Code_ok, "更新成功", todos)
}

// 刪除資料
func Del(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		util.Msg(c, util.Code_Param_Invalid, "參數無效", nil)
		return
	}
	err := db.DB.Where("id = ?", id).Delete(db.Todo{}).Error
	if err != nil {
		log.Warn().Caller().Err(err).Msg("DB Where.Delete")
		util.Msg(c, util.Code_DB_Conn, "資料存取錯誤#1", nil)
		return
	}
	username := util.GetUserName(c)
	todos := util.FindAll(c, username)
	util.Msg(c, util.Code_ok, "刪除成功", todos)
}
