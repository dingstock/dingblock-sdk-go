package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/dingstock/dingblock-sdk-go/modal"
)

var ()

type Sign interface {
	Sign(data string) (string, error)
	Verify(originalData, signData string) error
}

type RsaSign struct {
	DingBlockPubKey string
	UserPrivateKey  string
}
type AesSign struct {
	AppId     string
	AppSecret string
}

func (AesSign) GetSignStr(bizData string) {
	return
}

func (a *AesSign) Encrypt(plainText string) ([]byte, error) {
	return modal.AesEcbEncrypt([]byte(plainText), []byte(a.AppSecret))
}

func (a *AesSign) Decrypt(signData []byte) (originalData []byte, err error) {
	return modal.AesEcbDecrypt(signData, []byte(a.AppSecret))
}

func (r *RsaSign) baseSign(s string) (string, error) {
	fmt.Printf("【盯链】签名字符串:  %s \n", s)
	decodeString, _ := base64.StdEncoding.DecodeString(r.UserPrivateKey)
	privateKey, err := x509.ParsePKCS8PrivateKey(decodeString)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		return "", err
	}
	h := sha256.New()
	h.Write([]byte(s))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hash)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	out := base64.StdEncoding.EncodeToString(signature)
	//out := hex.EncodeToString(signature)
	return out, nil
}

func (r *RsaSign) ResponseSign(data modal.PublicResponse) (string, error) {
	s := fmt.Sprintf("bizData=%s&code=%d&msg=%s&nonce=%s&timestamp=%d", data.BizData, data.Code, data.Msg, data.Nonce, data.Timestamp)
	fmt.Printf("【盯链】签名字符串:  %s \n", s)
	return r.baseSign(s)
}

func (r *RsaSign) RequestSign(data modal.PublicRequest) (string, error) {
	s := fmt.Sprintf("appId=%s&bizData=%s&method=%s&nonce=%s&timestamp=%d", data.AppId, data.BizData, data.Method, data.Nonce, data.Timestamp)
	fmt.Printf("【盯链】签名字符串:  %s \n", s)
	return r.baseSign(s)
}

func (r *RsaSign) Verify(originalData, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		println(err.Error())
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(r.DingBlockPubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		println(err.Error())
		return err
	}
	hash := sha256.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash.Sum(nil), sign)
}
