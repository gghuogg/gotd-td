package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func mains() {
	str := "你好123abca"
	encryptstr, _ := RSAEncryptString(str, ".\\publickeyread\\rsa_public.pem")
	fmt.Println(encryptstr)

	decrypt, _ := RSADecryptString(encryptstr, ".\\publickeyread\\rsa_private_pkcs1.pem")
	fmt.Println(decrypt)

}

// RSA加密字节数组，返回字节数组
func RSAEncrypt(originalBytes []byte, filename string) ([]byte, error) {
	// 1、读取公钥文件，解析出公钥对象
	publicKey, err := ReadParsePublicKey(filename)
	fmt.Println(publicKey)
	if err != nil {
		return nil, err
	}
	// 2、RSA加密，参数是随机数、公钥对象、需要加密的字节
	// PKCS#1 v1.5 padding
	// Reader是一个全局共享的密码安全的强大的伪随机生成器
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, originalBytes)
}

// RSA解密字节数组，返回字节数组
func RSADecrypt(cipherBytes []byte, filename string) ([]byte, error) {
	// 1、读取私钥文件，解析出私钥对象
	privateKey, err := ReadParsePrivaterKey(filename)
	if err != nil {
		return nil, err
	}
	// 2、ras解密，参数是随机数、私钥对象、需要解密的字节
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherBytes)
}

// 读取公钥文件，解析出公钥对象
func ReadParsePublicKey(filename string) (*rsa.PublicKey, error) {
	// 1、读取公钥文件，获取公钥字节
	publicKeyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 2、解码公钥字节，生成加密块对象
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, errors.New("公钥信息错误！")
	}
	// 3、解析DER编码的公钥，生成公钥接口
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 4、公钥接口转型成公钥对象
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}

// 读取私钥文件，解析出私钥对象
func ReadParsePrivaterKey(filename string) (*rsa.PrivateKey, error) {
	// 1、读取私钥文件，获取私钥字节
	privateKeyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// 2、对私钥文件进行编码，生成加密块对象
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// RSA加密字符串，返回base64处理的字符串
func RSAEncryptString(originalText, filename string) (string, error) {
	cipherBytes, err := RSAEncrypt([]byte(originalText), filename)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// RSA 解密经过base64处理的加密字符串，返回加密前的明文
func RSADecryptString(cipherlText, filename string) (string, error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(cipherlText)
	originalBytes, err := RSADecrypt(cipherBytes, filename)
	if err != nil {
		return "", err
	}
	return string(originalBytes), nil
}





package main

import (
"crypto/rand"
"crypto/rsa"
"crypto/x509"
"encoding/base64"
"encoding/pem"
"fmt"
)

var publicKeyStr = `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAv4o7vRNFC/D5C+qGIU0afPsQexjpZa1e2cKBc6Wx4t6W3F0AIk7H
74mYNIy4c2lJXvjo6Bxp/L0/C7hH3yKnNf+jhbyPWbkfyiZf+ORjNc3yDU1+gio2
8BL4XaZNkgfLMlEwgH75pZ7M4Zi91tj89TrW4tyNmdEyNIFuJMu35SkO6MNNQzUj
+jYX+H+ossnOLsfc1SE4AosrIVwwL8m6eB72Au/yabGc9tabUiz6HMG01j8pJfEc
uZv7dWUxPGqG5k75I0ilvIUVMFIehSwQhSK9z7Yhh0iWt9NT5TpCv35TJ8jlVm91
jw8Cqpcsj6/mBy8kAXeklV3xXElDeewX3wIDAQAB
-----END RSA PUBLIC KEY-----
`
var privateKeyStr = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAv4o7vRNFC/D5C+qGIU0afPsQexjpZa1e2cKBc6Wx4t6W3F0A
Ik7H74mYNIy4c2lJXvjo6Bxp/L0/C7hH3yKnNf+jhbyPWbkfyiZf+ORjNc3yDU1+
gio28BL4XaZNkgfLMlEwgH75pZ7M4Zi91tj89TrW4tyNmdEyNIFuJMu35SkO6MNN
QzUj+jYX+H+ossnOLsfc1SE4AosrIVwwL8m6eB72Au/yabGc9tabUiz6HMG01j8p
JfEcuZv7dWUxPGqG5k75I0ilvIUVMFIehSwQhSK9z7Yhh0iWt9NT5TpCv35TJ8jl
Vm91jw8Cqpcsj6/mBy8kAXeklV3xXElDeewX3wIDAQABAoIBAQC3e3gWe6sc8U3h
eFvc8rt5FhiHkymh2R6Pg5/2ZLevGINzvTMvqkAk2q4PSO44wJckf1S9ZNqy3abT
V0iiG9QKCkVfC25XCNvHDlPJXR8cn5hL3fxePIo8GDALpg+Q4B5qqRPJQ9lqKmEh
iuvwMBTB5geLcrPCsX9ihMo6PVyAC9UbXFTwt/e/xArvzC8kHKTfkVlE7DLl1YCt
drlASoXlB3DlGU7OcDeSQXWHx9lj5lBGQPBHgmPhE32A+q4MeLF9AO5Sb7L6UxEn
BcihFdBQmAW7QcN9BzXJipxytugSOqGczGARo49znecAn1HaAYY2aEZXryHuTv2T
0UvL2EbBAoGBAMLMo4ucZqsuaRBCY7mP19Z4a+bZHgWc/5R0DXX2Y4lHnWzQlw/M
WxwM34avSrG4rdZtnakWJO0h8MWsmm9ntomP7pKqah3cj9bblfaVC4mR9UUL5nQX
kT7GnCMvh0WqqAkxLUnPkpWh9NTsttZJCCJHKhfBVyuLdLuFoToHFTaXAoGBAPu3
cnfKmvXFgXeF/tb6k5atgneyKrw7ZW1VP01zXoAzkSdpE+Pv04TFABn3ApLZ8QtT
oj1+t6L6ZsJI5bTusQeCLYd88FH5rAae44gxgfJUzBHYi2OFK5zWJ3VbZ442Hdym
WZf6p9O+yj6UogC8cGeLcIBNi3MYfTuA1ThBnNn5AoGAeM+7mklcb0vwQvI5sQrh
3JSArwH2mZqd867p7j/pyWkN3D+dgGHIG6RsOLacR4522abQAd7G9f53udDrUq5q
yfPnDTT9bRmuZ5116yDSr/ZCvBi+Sz+Wh+1bRoMijPVjE5hZMw62JXD1S9ynJEzU
65VhKKxy6IGr2uTscfWUlW8CgYBOhZ9ztwQ3/vCwNGCW93vq2R2F8VJbfbvaK2mf
44lrPKrfPGw0ArBdDodFkkR/QvVqryBBRrJKaQp9OPhPRNZ8nSNOzK67I8OCHTal
WaxLyCAQwRhaZ0R/nF6awXTAbClMl6gDPPH1n+K/OuZ+jEoUQu9JVudVdCI82aep
8O3ZIQKBgHFYTlevu6Td9Aq3ICwb3n4BCutSAyuaVtE4VFerM+ppDoFLCUJzz32y
SOjRjKtde9GR+NYjSFndJG6ijVu0WHFNxN7rOF20xEDQmyip72pu+TLTFztMb5FP
BCTxHTTjFfMBxI0onZQWaf6RvUrmQ1K/ugCjsvInutmTk8FqEUNL
-----END RSA PRIVATE KEY-----
`

func mainn() {
	//fmt.Println("privateKeyStr: ", privateKeyStr)
	// 解析 PEM 格式的私钥数据
	block, _ := pem.Decode([]byte(privateKeyStr))
	//fmt.Println("block: ", block)

	if block == nil {
		panic("failed to parse PEM block containing the private key")
	}
	// 解析 RSA 私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("jie xi si yao err: ", err)
		panic(err)
	}
	//fmt.Println(privateKey)

	// 解析pem格式的公钥数据
	blockPub, _ := pem.Decode([]byte(publicKeyStr))
	if blockPub == nil {
		panic("failed to parse PEM block containing the public key")
	}
	//fmt.Println("blockPub: ", blockPub)
	// 解析 DER 编码的公钥，得到一个 *rsa.PublicKey
	publicKey, err := x509.ParsePKCS1PublicKey(blockPub.Bytes)
	//publicKey, err := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if err != nil {
		panic(err)
	}
	//fmt.Println("publicKey: ", publicKey)

	//publicKeyLast := publicKey.(*rsa.PublicKey)

	//根据公钥加密
	myPasswd := []byte("yWdslfXfaFghVG6m")
	fmt.Println("myPasswd: ", myPasswd)
	fmt.Println("myPasswd: ", string(myPasswd))
	// 加密1
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, myPasswd)
	if err != nil {
		fmt.Println("err: ", err)
	}
	//fmt.Println("加密后：", ciphertext)
	//加密1

	JiaMiHouStr := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("加密后在base64 : ", JiaMiHouStr)

	// 对base64编码的字符串进行解码
	decoded, err := base64.StdEncoding.DecodeString(JiaMiHouStr)
	if err != nil {
		fmt.Println("Error decoding string:", err.Error())
		return
	}
	//fmt.Println("base64解码后： ", decoded)

	//根据私钥解密
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println("Origin passwd msg:", string(myPasswd))
	fmt.Println("decrypted message: ", string(decrypted))
}




















