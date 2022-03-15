package rsa_tools

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m\n"
	NoticeColor  = "\033[1;36m%s\033[0m\n"
	WarningColor = "\033[1;33m%s\033[0m\n"
	ErrorColor   = "\033[1;31m%s\033[0m\n"
	DebugColor   = "\033[0;36m%s\033[0m\n"
	SuccessColor = "\033[1;32m%s\033[0m\n"
)

func GetPublicKey() (*rsa.PublicKey, error) {
	// todo get pubPem from license
	pubPem := TestPublicKeyPem()
	return ParseRsaPublicKeyFromPemStr(pubPem)
}

// TestPrivateKey just for test
func TestPrivateKey() (*rsa.PrivateKey, error) {
	priPem := TestPrivateKeyPem()
	return ParseRsaPrivateKeyFromPemStr(priPem)
}

func SignPKCS1v15(data []byte, priKey *rsa.PrivateKey) []byte {
	//对原文进行hash散列
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	//实现签名
	sign, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("generate sign error: %v \n", err)
	}
	return sign
}

func SignPSS(data []byte, priKey *rsa.PrivateKey) []byte {
	// todo 用一行代码完成：sha256.Sum256(data)  对原文进行hash散列
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	opts := rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.SHA256}

	// 实现签名
	sign, err := rsa.SignPSS(rand.Reader, priKey, crypto.SHA256, hashed, &opts)
	if err != nil {
		fmt.Printf("generate sign error: %v \n", err)
	}
	return sign
}

func VerifyPKCS1v15(plain []byte, signature string, pubKey *rsa.PublicKey) (bool, error) {
	result, err := ParseSignature(signature)
	if err != nil {
		fmt.Printf("decode signature error: %v \n", err)
	}

	signVer := []byte(result["sign_ver"])
	licSN := []byte(result["lic_sn"])
	nonce := []byte(result["nonce"])
	timestamp := []byte(result["timestamp"])

	signStr := append(plain, signVer...)
	signStr = append(signStr, licSN...)
	signStr = append(signStr, nonce...)
	signStr = append(signStr, timestamp...)

	h := sha256.New()
	h.Write(signStr)
	hashed := h.Sum(nil)

	sign, err := base64.StdEncoding.DecodeString(result["sign"])
	if err != nil {
		fmt.Printf("decode sign error: %v \n", err)
	}

	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed, sign)
	if err != nil {
		return false, err
	}
	return true, nil
}

func VerifyPSS(plain []byte, signature string) (bool, error) {
	result, err := ParseSignature(signature)
	if err != nil {
		fmt.Printf("decode signature error: %v \n", err)
	}

	signVer := []byte(result["sign_ver"])
	licSN := []byte(result["lic_sn"])
	nonce := []byte(result["nonce"])
	timestamp := []byte(result["timestamp"])

	signStr := append(plain, signVer...)
	signStr = append(signStr, licSN...)
	signStr = append(signStr, nonce...)
	signStr = append(signStr, timestamp...)

	h := sha256.New()
	h.Write(signStr)
	hashed := h.Sum(nil)

	opts := rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.SHA256}

	sign, err := base64.StdEncoding.DecodeString(result["sign"])
	if err != nil {
		fmt.Printf("decode sign error: %v \n", err)
	}

	pub, err := GetPublicKey()
	if err != nil {
		fmt.Printf("get public key error: %v \n", err)
		return false, err
	}

	err = rsa.VerifyPSS(pub, crypto.SHA256, hashed, sign, &opts)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ParseSignature(signBase64 string) (map[string]string, error) {
	signBytes, err := base64.StdEncoding.DecodeString(signBase64)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	pairs := strings.Split(string(signBytes), ",")
	for _, pair := range pairs {
		// find first '='
		// note: sign value (base64) contains '='
		index := strings.Index(pair, "=")
		// not found =, first is '='
		if index <= 0 {
			return nil, err
		}
		k := pair[0:index]
		k = strings.TrimSpace(k)

		v := pair[index+1 : len(pair)]
		v = strings.TrimSpace(v)
		if k == "" || v == "" {
			return nil, err
		}

		result[k] = v
	}
	return result, nil
}

func GenerateSignInfo() {
	licSN := []byte("cc649a78-5d49-4213-aa0b-fb068a29f096")
	timestamp := []byte("1623329779886")
	imgPath := checkFileExist("./test.png")
	priKeyPath := checkFileExist("./pri.key")
	nonce := []byte("tr4NKMBhBNmKKXCpBRBw4bdaPhYyrhiC25B6dm723xtWbjY7RFfHW3Csc2pNsGsM")
	signVer := []byte("1")

	args := os.Args
	if len(args) > 1 {
		licSN = []byte(os.Args[1])
		if os.Args[1] == "-h" {
			fmt.Println("Usage:  ./generate [lic_sn] [imagePath] [priKeyPath]")
			fmt.Println()
			fmt.Println("lic_sn")
			fmt.Println("\t    公钥对应的：License SN（若lic_sn不存在，将使用内置sn：cc649a78-5d49-4213-aa0b-fb068a29f096）")
			fmt.Println()
			fmt.Println("imagePath")
			fmt.Println("\t    工具会自动读取当前目录下的：test.png（若test.png不存在，将使用\"Hello world\"作为消息明文）")
			fmt.Println("\t    可以手动指定要load的图片路径，示例：./generate cc649a78-5d49-4213-aa0b-fb068a29f096 ../xxx/test2.png")
			fmt.Println()
			fmt.Println("priKeyPath")
			fmt.Println("\t    工具会自动读取当前目录下的：pri.key（若当前目录pri.key不存在，将使用内置私钥）")
			fmt.Println("\t    可以手动指定要load的私钥路径，示例：./generate cc649a78-5d49-4213-aa0b-fb068a29f096 ../xxx/test2.png  ../yyy/pri2.key")
			flag.PrintDefaults()
			os.Exit(0)
		}
	}

	if len(args) > 2 {
		imgPath = os.Args[2]
		if _, err := os.Stat(imgPath); os.IsNotExist(err) {
			fmt.Printf("%v文件不存在 \033[1;33m\033[0m\n", imgPath)
			os.Exit(0)
		}
	}

	if len(args) > 3 {
		priKeyPath = os.Args[3]
		if _, err := os.Stat(priKeyPath); os.IsNotExist(err) {
			fmt.Printf("%v文件不存在 \033[1;33m\033[0m\n", priKeyPath)
			os.Exit(0)
		}
	}

	blob := getBlob(imgPath)

	hasher := sha256.New()
	hasher.Write(blob)
	blobHash := hasher.Sum(nil)

	signStr := append(blobHash, signVer...)
	signStr = append(signStr, licSN...)
	signStr = append(signStr, nonce...)
	signStr = append(signStr, timestamp...)

	// 256 bytes
	priKey := getPriKey(priKeyPath)
	sign := SignPKCS1v15(signStr, priKey)
	base64Sign := base64.StdEncoding.EncodeToString(sign)
	signature := fmt.Sprintf("timestamp=%v,sign_ver=%v,lic_sn=%v,nonce=%v,sign=%v", string(timestamp), string(signVer), string(licSN), string(nonce), base64Sign)
	base64Signature := base64.StdEncoding.EncodeToString([]byte(signature))

	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Printf("blob_hash: \033[1;34m%v\033[0m\n", blobHash)
	fmt.Printf("origin_signature: \033[1;32m%v\033[0m\n", signature)

	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Printf("blob_digest: \033[1;34m%v\033[0m\n", base64.StdEncoding.EncodeToString(blobHash))
	fmt.Printf("signature: \033[1;33m%v\033[0m\n", base64Signature)

	fmt.Println(VerifyPKCS1v15(blobHash, base64Signature, &priKey.PublicKey))
}

func checkFileExist(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ""
	} else {
		return path
	}
}

func getBlob(imgPath string) []byte {
	if imgPath != "" {
		fmt.Printf("...签名原文: \033[1;33m%v\033[0m\n", imgPath)
		blob, err := ioutil.ReadFile(imgPath)
		if err != nil {
			fmt.Printf(ErrorColor, "图片读取失败")
			os.Exit(0)
		}
		return blob
	} else {
		plain := "Hello World"
		fmt.Printf("...签名原文: \033[1;33m%v\033[0m\n", plain)
		return []byte(plain)
	}
}

func getPriKey(priKeyPath string) *rsa.PrivateKey {
	if priKeyPath != "" {
		fmt.Printf("...签名私钥: \u001B[1;33m%v\033[0m\n", priKeyPath)
		blob, err := ioutil.ReadFile(priKeyPath)
		if err != nil {
			fmt.Printf(ErrorColor, "私钥文件读取失败")
			os.Exit(0)
		}

		pri, err := ParseRsaPrivateKeyFromPemStr(string(blob))
		if err != nil {
			fmt.Printf(ErrorColor, "私钥文件转换失败，请使用标准pem格式（包含begin和end）")
			os.Exit(0)
		}
		return pri
	} else {
		fmt.Printf("...签名私钥: \u001B[1;33m%v\033[0m\n", "内置测试私钥")
		pri, err := TestPrivateKey()
		if err != nil {
			fmt.Printf(ErrorColor, "测试私钥转换失败")
			os.Exit(0)
		}
		return pri
	}
}

func GenerateRsa() {
	//生成私钥
	pri, _ := rsa.GenerateKey(rand.Reader, 2048)

	priStr := ExportRsaPrivateKeyAsPemStr(pri)
	priStrFormat1 := strings.Replace(priStr, "\n", "", -1)
	priStrFormat2 := strings.Replace(priStrFormat1, "-----", "\n\n", -1)

	fmt.Println(priStrFormat2)

	fmt.Println()
	//生成公钥
	pub := &pri.PublicKey
	pubStr, err := ExportRsaPublicKeyAsPemStr(pub)
	pubStrFormat1 := strings.Replace(pubStr, "\n", "", -1)
	pubStrFormat2 := strings.Replace(pubStrFormat1, "-----", "\n\n", -1)
	if err != nil {
		fmt.Println(ErrorColor, err)
	}
	fmt.Println(pubStrFormat2)
}



func GenerateSignature() {
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano()/1000000)
	nonce := "cc649a785d494213aa0bfb068a29f096"
	key := "1ef1b4c60e0f458d9772babcddbdc99d"
	secret := "cd2e12efffc849929c0f4e0f468ed0b2"

	arr := []string{timestamp, nonce, key}
	sort.Strings(arr)

	data := strings.Join(arr[:], "")

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	signature := hex.EncodeToString(h.Sum(nil))

	fmt.Printf("key=%v,nonce=%v,timestamp=%v,signature=%v\n", key, nonce, timestamp, signature)
}
