package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

const (
	ASE_256_CBC_KEY_PREFIX string = "base64:"
)

var ErrHmac256Check = errors.New("hmac256 校验失败")

type jsonData struct {
	Iv    string `json:"iv"`
	Value string `json:"value"`
	Mac   string `json:"mac"`
}

// laravel aes256cbc加密
func Encrypt(key string, data string) (string, error) {
	// 获取key
	keyByte, err := getKey(key)
	if err != nil {
		return "", err
	}
	// 选取加密算法
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	// php 的字符串序列化
	data, err = SerializeStr(data)
	if err != nil {
		return "", err
	}

	// pkcs7补位
	text := PKCS7Padding([]byte(data), block.BlockSize())
	// 随机iv
	iv := GetRandomByte(block.BlockSize())

	cfb := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(text))
	cfb.CryptBlocks(ciphertext, text)

	//hmac 256
	ciphertextStr := EncodeBase64(ciphertext)
	ivStr := EncodeBase64(iv)
	mac, err := Hmac256String(ivStr+ciphertextStr, keyByte)
	if err != nil {
		return "", err
	}

	jsonStr, _ := json.Marshal(jsonData{ivStr, ciphertextStr, mac})
	return EncodeBase64(jsonStr), nil
}

func Decrypt(key string, b64 string) (string, error) {
	// 获取key
	keyByte, err := getKey(key)
	if err != nil {
		return "", err
	}

	// hmac256 hash 校验
	payload, err := getJsonPayload(keyByte, b64)
	if err != nil {
		return "", err
	}

	ciphertext, err := DecodeBase64(payload.Value)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	iv, err := DecodeBase64(payload.Iv)
	if err != nil {
		return "", err
	}
	text := make([]byte, len(ciphertext))

	cfb := cipher.NewCBCDecrypter(block, iv)
	cfb.CryptBlocks(text, ciphertext)

	text = PKCS7UnPadding(text)

	textStr, err := UnSerializeStr(string(text))
	if err != nil {
		return "", err
	}
	return textStr, nil
}

// key base64 解密
func getKey(key string) ([]byte, error) {
	if strings.HasPrefix(key, ASE_256_CBC_KEY_PREFIX) {
		key = strings.Replace(key, ASE_256_CBC_KEY_PREFIX, "", 1)
		keyByte, err := DecodeBase64(key)
		if err != nil {
			return nil, err
		}
		return keyByte, nil
	}

	return []byte(key), nil
}

// hmac256 str
func Hmac256String(s string, key []byte) (string, error) {
	h := hmac.New(sha256.New, key)
	_, err := io.WriteString(h, s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// 获取随机字符byte
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

// hmac256校验
func getJsonPayload(key []byte, b64 string) (*jsonData, error) {
	payload := &jsonData{}
	b64Byte, err := DecodeBase64(b64)
	if err != nil {
		return nil, err
	}

	fmt.Println()
	fmt.Println(string(b64Byte))
	fmt.Println()

	err = json.Unmarshal(b64Byte, &payload)
	if err != nil {
		return nil, err
	}
	if payload.Iv == "" || payload.Value == "" || payload.Mac == "" {
		return nil, ErrHmac256Check
	}

	calcMac, err := Hmac256String(payload.Iv+payload.Value, key)

	if err != nil {
		return nil, err
	}
	if calcMac != payload.Mac {
		return nil, ErrHmac256Check
	}

	return payload, nil
}

// 公共的一些
// 补0
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	padtext = append(text, padtext...)
	return padtext
}

func PKCS7UnPadding(text []byte) []byte {
	length := len(text)
	unpadding := int(text[length-1])
	return text[:(length - unpadding)]
}

func SerializeStr(data string) (string, error) {
	encoder := php_serialize.NewSerializer()
	val, err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return val, nil
}

func UnSerializeStr(data string) (string, error) {
	decoder := php_serialize.NewUnSerializer(data)
	val, err := decoder.Decode()
	if err != nil {
		return "", err
	}
	if strVal, ok := val.(string); ok {
		return strVal, nil
	}
	return "", err
}

func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func DecodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}
