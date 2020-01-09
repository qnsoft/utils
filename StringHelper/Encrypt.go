package StringHelper

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"net/url"

	"github.com/qnsoft/utils/ErrorHelper"
	php2go "github.com/qnsoft/utils/Php2go"

	"github.com/qnsoft/mahonia"
)

//md5加密方法
func Md5(s string) string {
	//	h := md5.New()
	//h.Write([]byte(s))
	//return hex.EncodeToString(h.Sum(nil))
	return php2go.Md5(s)
}

/*
urlencode加密
*/
func UrlEncode(_key, _value string) string {
	v := url.Values{}
	//v.Add(_key, _value)
	v.Set(_key, _value)
	_str := v.Encode()
	return _str
}

/*
urldecode加密
*/
func UrlDecode(_body string) map[string][]string {
	_values, _ := url.ParseQuery(_body)
	return _values
}

/*
简单加密函数
*/
func Lock_url(txt, key string) string {
	_rt := ""
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-=+"
	nh := php2go.Rand(0, 64)
	ch := chars[nh]
	mdKey := php2go.Md5(key + string(ch))
	mdKey = php2go.Substr(mdKey, uint(nh%8), nh%8+7)
	txt = php2go.Base64Encode(txt)
	tmp := ""
	j := 0
	k := 0
	for i := 0; i < php2go.Strlen(txt); i++ {
		if k == php2go.Strlen(mdKey) {
			k = 0
		}
		j = (nh + php2go.Strpos(chars, string(txt[i]), 1) + php2go.Ord(string(mdKey[k+1]))) % 64
		tmp = tmp + string(chars[j])
	}
	_rt = php2go.URLEncode(string(ch) + tmp)
	return _rt
}

/*
简单解密函数
*/
func Unlock_url(txt, key string) string {
	_rt := ""
	txt, err := php2go.URLDecode(txt)
	ErrorHelper.CheckErr(err)
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-=+"
	ch := string(txt[0])
	nh := php2go.Strpos(chars, ch, 1)
	mdKey := php2go.Md5(key + string(ch))
	mdKey = php2go.Substr(mdKey, uint(nh%8), nh%8+7)
	txt = php2go.Substr(txt, 1, php2go.Strlen(txt))
	tmp := ""
	j := 0
	k := 0
	for i := 0; i < php2go.Strlen(txt); i++ {
		if k == php2go.Strlen(mdKey) {
			k = 0
		}
		j = php2go.Strpos(chars, string(txt[i]), 1) - nh - php2go.Ord(string(mdKey[k+1]))
		for j < 0 {
			j = j + 64
		}
		tmp = tmp + string(chars[j])
	}
	_rt, _ = php2go.Base64Decode(tmp)
	return _rt
}

/*
私钥签名
*/
func RsaSign(data, privateKey []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	//获取私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
}

/*
公钥验证
data 数据
signature 签名
publicKey 公钥字符串
*/
func RsaSignVer(data, signature, publicKey []byte) error {
	hashed := sha256.Sum256(data)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

/*
公钥加密
*/
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(origData)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

/*
私钥解密
*/
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(ciphertext)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//GBK和UTF8互转 UseNewEncoder("要转编码的字符串","gbk","utf-8")
func GBK_UTF8_Encoder(src string, oldEncoder string, newEncoder string) string {
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}
