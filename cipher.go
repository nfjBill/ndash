package ndash

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm3"
	"strings"
)

func Sha256(buf []byte) string {
	h := sha256.New()
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sm4GetKey(str string) string {
	str = strings.ToLower(str)
	hMd5 := md5.New()
	hMd5.Write([]byte(str))
	md5Str := hex.EncodeToString(hMd5.Sum(nil))
	hSm3 := sm3.New()
	hSm3.Write([]byte(md5Str))
	sum := hSm3.Sum(nil)

	sm3Str := strings.ToUpper(hex.EncodeToString(sum))

	return sm3Str[6:38]
}

func SMAbsAf(str string) string {
	return Sm4GetKey(StrReverse(Md5(Md5(str)[2:12] + Sm4GetKey(str)[3:13]))[6:16])
}

func SHA256HashCode(info []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(info)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

	//方法二：
	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
	//return hashcode2
}
