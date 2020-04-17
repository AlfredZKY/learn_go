package main
import (
	"errors"
    "fmt"
    "crypto/cipher"
    "crypto/aes"
     "bytes"
    "encoding/base64"
)

//PKCS5Padding func
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

//PKCS5UnPadding func
func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

//AesEncrypt func
func AesEncrypt(origData, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()
    origData = PKCS5Padding(origData, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
    crypted := make([]byte, len(origData))
    blockMode.CryptBlocks(crypted, origData)
    return crypted, nil
}
//AesDecrypt func
func AesDecrypt(crypted, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    origData := make([]byte, len(crypted))
    blockMode.CryptBlocks(origData, crypted)
    origData = PKCS5UnPadding(origData)
    return origData, nil
}

// ReverseString func
func ReverseString(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
}

func main(){
	fmt.Println("startr encrypt!!!!")
    var asekey = []byte("hell0ohj ell0ohj")
    if 16 != len(asekey){
        fmt.Println(len(asekey))
        panic(errors.New("invalid key size must be 16"))
    }
	pass := []byte("12345678")
	xpass,err := AesEncrypt(pass,asekey)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(pass[:]))
	fmt.Println(string(xpass[:]))
	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%v\n",pass64)
	fmt.Println("startr decrypt!!!!")

	bytesPass,err := base64.StdEncoding.DecodeString(pass64)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bytesPass))
	tpass ,err := AesDecrypt(bytesPass,asekey)
	if err != nil{
		fmt.Println(err)
		return 
	}
    fmt.Printf("解密后：%s\n",tpass)
    str := "hellowoI"
    str1 := ReverseString(str)
    new := str + str1
    fmt.Println(new,len(new))
}