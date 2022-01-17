package controller

// import (
// 	"net/http"
// 	"net/url"
// 	"time"

// 	"github.com/tidwall/gjson"
// )

// // http底層介面
// type IhttpRequest interface {
// 	// 建立請求
// 	CurlWithRequest(req *http.Request, arg HttpArgsF) (res *http.Response, err error)
// }

// // 定義泛型回傳型別
// type ExternalAPIResult interface {
// 	~string | gjson.Result
// }

// // 呼叫外接API介面
// type IExternalCurl2[T ExternalAPIResult] interface {
// 	// 呼叫API
// 	curlAPI(rp RequestParam) (t T, err error)
// 	//寫操作記錄
// 	// writeAPIRecord(path, params, reqType, reqHeader, resHeader, resData, errorCode, errorText, logError string, status int64, durationTime time.Duration, trace *trace.CalculatedHttpTraceInfo) (err error)
// }

// // 路由常數的泛型
// type RestAPI interface {
// 	DurianExternalRESTAPI
// }

// // 請求需要用到的參數
// type RequestParam[T RestAPI] struct {
// 	Rest    T
// 	Query   *url.Values
// 	Body    []byte
// 	Payload map[string]interface{}
// }

// func RP(){
// 	rp := RequestParam{Rest:"",Query:"",Body:[]byte("sdsd")}
// }
// // http底層實作
// type httpRequestCurl struct {
// }

// func (h *httpRequestCurl) CurlWithRequest(req *http.Request, arg HttpArgsF) (res *http.Response, err error) {
// 	// To Do
// 	return
// }

// // http參數
// type HttpArgsF struct {
// 	// 超時
// 	CurlNoTimeout bool
// 	// 超時秒數
// 	CurlTimeout time.Duration
// 	// 緩存
// 	Transport *http.Transport
// }

// // 呼叫Durian的API實作
// type DurianExternalCurl[T ExternalAPIResult] struct {
// }

// // 呼叫Durian的API
// func (e *DurianExternalCurl[T]) curlAPI(rp RequestParam) (t T, err error) {

// 	// To Do

// 	res, err := http.NewRequest()
// 	if err != nil {
// 		return
// 	}

// 	// header, cookies

// 	hr := &httpRequestCurl{}

// 	start := time.Now()

// 	hr.CurlWithRequest(res)

// 	// 計算API耗時
// 	elapsed := time.Since(start)

// 	// 寫操作記錄
// 	// e.writeAPIRecord()
// 	// fmt.Print("a")
// 	return
// }

// // 寫操作記錄
// func (e *DurianExternalCurl) writeAPIRecord(path, params, reqType, reqHeader, resHeader, resData, errorCode, errorText, logError string, status int64, durationTime time.Duration, trace *trace.CalculatedHttpTraceInfo) (err error) {

// 	// To Do

// 	return
// }

// // 新增功能實作 CashOperation 現金操作
// func (e *DurianExternalCurl) CashOperation(rest DurianExternalRESTAPI) (err error) {

// 	// To Do

// 	rp := &RequestParam{}

// 	e.curlAPI(rp)

// 	return
// }

// type IExternalRESTAPI interface {
// 	Method() string
// 	URL() string
// }

// type DurianExternalRESTAPI string

// const (
// 	DurianAPICashOp DurianExternalRESTAPI = "cash_op"
// )

// // 請求方式
// func (d DurianExternalRESTAPI) Method() string {
// 	switch d {
// 	case DurianAPICashOp:
// 		return http.MethodPut
// 	default:
// 		return ""
// 	}
// }

// // API的URL
// func (d DurianExternalRESTAPI) URL() string {
// 	switch d {
// 	case DurianAPICashOp:
// 		return "/api/single_wallet/{user_id}/op"
// 	default:
// 		return ""
// 	}
// }
