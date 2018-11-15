package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go_web_demo/lib/common"
	"go_web_demo/lib/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//http响应通用数据
type CommonData struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
	Success    bool   `json:"success"`
}

//http响应message 返回json 格式数据
type SpecialCommonData struct {
	StatusCode int         `json:"statusCode"`
	Msg        interface{} `json:"msg"`
	Success    bool        `json:"success"`
}

// 有的接口statuscode和msg写反了， 所以用interface解析
type InterfaceCommonData struct {
	StatusCode interface{} `json:"statusCode"`
	Msg        interface{} `json:"msg"`
	Success    bool        `json:"success"`
}

type CommonDataString struct {
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
	Success    bool   `json:"success"`
}

// 并发请求
type ResultMapStruct struct {
	Url  string `json:"url"`
	Data []byte `json:"data"`
	Err  error  `json:"err"`
}

/**
 * 保存日志 请求参数和返回数据
 * @param  {[type]} c    *gin.Context              [description]
 * @param  {[type]} data map[string]interface{}) ([]byte,      error [description]
 * @return {[type]}      [description]
 */
func saveRespLog(c *gin.Context, data interface{}) error {
	requestUrl := c.Request.URL.String()

	reqStartTime, _ := c.Get("reqStartTime")
	costTime := common.End(reqStartTime.(time.Time))

	reqData, _ := c.Get("reqData")
	requestBody, err := json.Marshal(reqData)
	if err != nil {
		logger.Warnf("saveRespLog::errUrl$v reqData:%v err%v", reqData, err)
	}

	responseBody, err := json.Marshal(data)
	if err != nil {
		logger.Warnf("saveRespLog::errUrl$v responseBody:%v err%v", responseBody, err)
	}
	logger.Infof(` %s###%s###%s###%v`, requestUrl, string(requestBody), string(responseBody), costTime)
	return nil
}

// ResponseSuccess 成功返回AccountMessageContentDataStruct
func ResponseSuccess(c *gin.Context, data interface{}) {

	ret := map[string]interface{}{
		"statusCode": common.ERR_SUC.ErrorNo,
		"msg":        common.ERR_SUC.ErrorMsg,
		"success":    true,
		"data":       data,
	}

	RenderJson(c, ret)
	return
}

// 自定义返回格式，目的兼容老版本接口
func ResponseCustom(c *gin.Context, data interface{}) {

	// ret := map[string]interface{}{
	// 	"statusCode": common.ERR_SUC.ErrorNo,
	// 	"msg":        common.ERR_SUC.ErrorMsg,
	// 	"success":    true,
	// 	"data":       data,
	// }

	RenderJson(c, data)
	return
}

// ResponseError 失败返回
func ResponseError(c *gin.Context, err *common.Err) {

	ret := map[string]interface{}{
		"statusCode": err.ErrorNo,
		"msg":        err.ErrorMsg,
		"success":    false,
		"data":       []interface{}{},
	}

	RenderJson(c, ret)
	return
}

// ResponseCustomError 自定义返回
func ResponseCustomError(c *gin.Context, errMsg string) {

	ret := map[string]interface{}{
		"statusCode": common.ERR_CUSTOM.ErrorNo,
		"msg":        errMsg,
		"success":    false,
		"data":       []interface{}{},
	}

	RenderJson(c, ret)
	return
}

func ResponseStatusCodeStringError(c *gin.Context, info []byte) {

	var commonDataString CommonDataString
	var err error
	err = json.Unmarshal(info, &commonDataString)
	if err != nil {
		ResponseError(c, common.ERR_REMOTE_CURL)
		return
	}

	ResponseCustom(c, commonDataString)
}

func RenderJson(c *gin.Context, data interface{}) {
	saveRespLog(c, data) //保存日志
	c.Header("Content-Type", "application/json;charset=UTF-8")
	c.JSON(200, data)
	return
}
func GetFullUrl(urlStr string, params map[string]string) string {
	v, _ := url.Parse(urlStr)

	paramsData := url.Values{}
	if params != nil {
		for key, value := range params {
			paramsData.Set(key, value)
		}
	}

	v.RawQuery = paramsData.Encode()
	urlPath := v.String()

	return urlPath
}

// HttpGet get请求
func Get(urlStr string, params map[string]string) ([]byte, error) {
	urlPath := GetFullUrl(urlStr, params)

	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("response StatusCode is not StatusOK")
	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

// HttpPost post请求
func PostJson(url string, params []byte) ([]byte, error) {
	body := bytes.NewBuffer(params)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("response StatusCode is not StatusOK")
	}

	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respData, nil
}

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,HEAD,OPTIONS,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin,Content-Type,Accept,X-Requested-With,Authorization,If-None-Match,sid,source,token,app-type,set-token,Cookie")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			c.Header("Access-Control-Max-Age", "1728000")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

func PostJsonRequest(url string, requestData interface{}, responseData interface{}) (err error) {

	if !common.IsValidString(url) {
		return errors.New("PostJsonRequest url is invalid!")
	}

	var jsonData []byte
	jsonData, err = json.Marshal(requestData)
	if err != nil {
		logger.Warnf("Marshal json error:%s ", err.Error())
		return err
	}
	logger.Debugf("request:%+v, json:%+v\n", requestData, string(jsonData))
	var responseByte []byte
	responseByte, err = PostJson(url, jsonData)
	if err != nil {
		logger.Warnf("PostJson error:%s url:%s data:%s", err.Error(), url, string(responseByte))
		return err
	}
	responseByte = []byte(strings.Replace(string(responseByte), "\"data\":[]", "\"data\":{}", 1))
	err = json.Unmarshal(responseByte, responseData)
	logger.Debugf("url:%s", url)
	//logger.Debugf("request:%+v, json:%+v\n", requestData, string(responseByte))
	if err != nil {
		logger.Warnf("Unmarshal response error,data:%s err:%s ", string(responseByte), err.Error())
		return err
	}
	return nil
}

// 和PostJsonRequest 区别 是 data不做任何处理 原样返回
func PostJsonRequestOrigin(url string, requestData interface{}, responseData interface{}) (err error) {

	if !common.IsValidString(url) {
		return errors.New("PostJsonRequest url is invalid!")
	}

	var jsonData []byte
	jsonData, err = json.Marshal(requestData)
	if err != nil {
		logger.Warnf("Marshal json error:%s ", err.Error())
		return err
	}
	logger.Debugf("request:%+v, json:%+v\n", requestData, string(jsonData))
	var responseByte []byte
	responseByte, err = PostJson(url, jsonData)
	if err != nil {
		logger.Warnf("PostJson error:%s url:%s data:%s", err.Error(), url, string(responseByte))
		return err
	}
	err = json.Unmarshal(responseByte, responseData)
	logger.Debugf("url:%s", url)
	//logger.Debugf("request:%+v, json:%+v\n", requestData, string(responseByte))
	if err != nil {
		logger.Warnf("Unmarshal response error,data:%s err:%s ", string(responseByte), err.Error())
		return err
	}
	return nil
}

func PostJsonRequestReturnJsonByte(url string, requestData interface{}) ([]byte, error) {
	var jsonData []byte
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	logger.Debugf("request:%+v, json:%+v\n", requestData, string(jsonData))
	var responseByte []byte
	responseByte, err = PostJson(url, jsonData)
	logger.Debugf("url:%s", url)
	if err != nil {
		logger.Warnf("PostJson error url:%s data:%s error:%s", url, string(responseByte), err.Error())
		return nil, err
	}

	return responseByte, nil
}

//GetBodyParam 获取json格式的请求数据，keyStruct需要传入指针，该方法会对keyStruct赋值
func GetBodyParam(c *gin.Context, keyStruct interface{}) (err error) {
	err = c.ShouldBind(&keyStruct)
	c.Set("reqData", keyStruct)
	return
}

// 并发请求
// key url
// value map[参数: 参数值]
func PostMapRequestAsync(params map[string]interface{}) map[string]*ResultMapStruct {

	/*
		params := map[string]interface{}{
			"http://classtool.02.qadev.xuebadev.com/homework/getContent":   map[string]interface{}{"hw_id": 1},
			"http://rd01-course.qadev.xuebadev.com/course/getCourseRemark": map[string]interface{}{"course_id": 1, "type": 3},
		}
	*/

	finishNum := len(params)
	result := make(chan *ResultMapStruct, finishNum)
	resultMap := make(map[string]*ResultMapStruct, finishNum)
	for urlStr, param := range params {
		go func(urlStr string, param interface{}) {
			bytesArr, err := PostJsonRequestReturnJsonByte(urlStr, param)
			rms := ResultMapStruct{Data: bytesArr, Err: err, Url: urlStr}
			finishNum--
			result <- &rms
		}(urlStr, param)
	}

LOOP:
	for {
		select {
		case x := <-result:
			resultMap[x.Url] = x
			if finishNum <= 0 {
				break LOOP
			}
		}
	}

	return resultMap
}
