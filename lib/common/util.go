package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func InArrayInt(element int, list []int) bool {
	for _, val := range list {
		if element == val {
			return true
		}
	}

	return false
}

func InArrayUint(element uint, list []uint) bool {
	for _, val := range list {
		if element == val {
			return true
		}
	}

	return false
}

func postJsonRequest(url string, requestData interface{}, responseData interface{}) (err error) {
	var jsonData []byte
	jsonData, err = json.Marshal(requestData)
	if err != nil {
		return err
	}

	var responseByte []byte

	body := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	responseByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseByte, responseData)
	if err != nil {
		return err
	}
	return nil
}

// GetJsonReturnStruct 获取接口返回结构体
// lib.GetJsonReturnStruct(config.Config.RemoteUrl.GetStudentEvaluate, req)
// return
func GetJsonReturnStruct(url string, req interface{}) {
	m := map[string]interface{}{}
	err := postJsonRequest(url, req, &m)
	me, _ := json.Marshal(m)
	fmt.Println(string(me))
	fmt.Println(err)
	fmt.Println("==============")
	fmt.Println()
	JsonToStruct(me)
}

func JsonToStruct(jsonByte []byte) {
	m := map[string]interface{}{}
	err := json.Unmarshal(jsonByte, &m)
	fmt.Println(err)
	fmt.Println("==============")
	fmt.Println()
	MapToStruct(m, 0)
}

func MapToStruct(m interface{}, count int) {
	data, ok := m.(map[string]interface{})
	count++
	if ok {
		fmt.Println(strings.Repeat("    ", count) + "------------------")
		for k, v := range data {
			fmt.Printf(strings.Repeat("    ", count)+"%s %v `json:%q`", StrFirstToUpper(k), reflect.TypeOf(v), k)
			fmt.Println()

			vData, ok := v.(map[string]interface{})
			if ok {
				fmt.Println(strings.Repeat("    ", count) + k)
				MapToStruct(vData, count)
			}

			list, ok := v.([]interface{})
			if ok && len(list) > 0 {
				listValue, ok := list[0].(map[string]interface{})
				if ok {
					fmt.Println(strings.Repeat("    ", count) + k)
					MapToStruct(listValue, count)
				}
			}
		}
		fmt.Println(strings.Repeat("    ", count) + "------------------")
	}
}

// strFirstToUpper 首字母及下划线大写
func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i]) // + string(vv[i+1])
			} else {
				upperStr += string(vv[i])
			}
		}

	}
	return upperStr
}

// struct转map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}

	return data
}

// 获取随机byte
func GetRandomByte(length int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	strLen := len(str)
	strBytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, strBytes[r.Intn(strLen)])
	}
	return result
}

// 获取随机byte
func GetRandomNumberyte(length int) []byte {
	str := "0123456789"
	strLen := len(str)
	strBytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, strBytes[r.Intn(strLen)])
	}
	return result
}

func InSliceString(sliceArr []string, val string) bool {
	var isIn bool
	for _, v := range sliceArr {
		log.Printf("===%+v==,%+v====", v, val)
		if v == val {
			isIn = true
			break
		}
	}
	return isIn
}

func IsChineseChar(nickName string) bool {
	matched, err := regexp.MatchString(`^[\p{Han}]+$`, nickName)

	if err != nil {
		return false
	}
	return matched
}

/**
 * 开始计时
 * return time
 */
func Start() time.Time {
	t := time.Now()
	return t
}

/**
 * 结束计时
 * return ms  转化成毫秒 保留三位小数
 */
func End(t time.Time) string {
	s := time.Since(t).Seconds() * 1000
	return fmt.Sprintf("%.3f", s)
}
