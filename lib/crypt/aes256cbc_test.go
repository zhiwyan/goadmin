package crypt

import (
	"fmt"
	"goadmin/lib/config"
	"log"
	"strconv"
	"testing"
)

func init() {
	err := config.InitConfig("../../config/classroom.toml")
	if err != nil {
		panic(err)
	}
}

func TestEncrypt(t *testing.T) {
	var hwId int64 = 300
	token, err := Encrypt(config.Config.AppKey, strconv.FormatInt(hwId, 10)+"&"+"100")
	if err != nil {
		log.Println("token---", token)
		log.Println("err---", err)
	}
	log.Println("token---", token)

	hwIdDecrypt, err := Decrypt(config.Config.AppKey, token)
	if err != nil {
		log.Println("hwIdDecrypt---", hwIdDecrypt)
		log.Println("err---", err)
	}
	log.Println("hwIdDecrypt :", hwIdDecrypt)
}

func TestDecrypt(t *testing.T) {
	//token := "eyJpdiI6Ik92dHFLc2QweWllYmhNeWF4VHFJQVE9PSIsInZhbHVlIjoiOFptQk1GOUppa3o2VTNcL1J0UzU1dlE9PSIsIm1hYyI6IjQyZGVlMzUwNmYwZTc0NmMyNWVlY2JhNDMwMGUyMDA4OTQ0YzRjMzEyMjM3NmYzZWUwZWRkMmVhZDAzYzM1YmUifQ=="
	token := "eyJpdiI6Ik1HbG9PR3BrYUdwbGRHWnZkbXAwZEE9PSIsInZhbHVlIjoiY2tsdXV4VjJmSjkzQ3JvMTRNYk9JQT09IiwibWFjIjoiZTEwNjk4MjI2ZTc5YmEzMzFkZTA2YmI3MmJjNjZmYTkwMWZmMzEzZDY4YmYwMzRiZmJlMWUxMGUwNTkwYmRhYiJ9"
	hwIdDecrypt, err := Decrypt(config.Config.AppKey, token)
	if err != nil {
		log.Println("hwIdDecrypt---", hwIdDecrypt)
		log.Println("err---", err)
	}
	log.Println("hwIdDecrypt :", hwIdDecrypt)
}

func TestHmac256String(t *testing.T) {
	//jsonStr := `{"iv":"OvtqKsd0yiebhMyaxTqIAQ==","value":"8ZmBMF9Jikz6U3\/RtS55vQ==","mac":"42dee3506f0e746c25eecba4300e2008944c4c3122376f3ee0edd2ead03c35be"}`

	b64 := "42dee3506f0e746c25eecba4300e2008944c4c3122376f3ee0edd2ead03c35be"
	b64Byte, err := DecodeBase64(b64)

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println(string(b64Byte))
	fmt.Println()
}

func TestSerializeByte(t *testing.T) {
	str := "30"
	serStr, err := SerializeStr(str)
	if err != nil {
		log.Println("serStr---", serStr)
		log.Println("err---", err)
	}
	log.Println("serStr---", serStr)
}
