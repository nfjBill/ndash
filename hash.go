package ndash

import (
	"encoding/hex"
	"github.com/speps/go-hashids/v2"
)

func HashID(buf []byte) (string, error) {
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	hashCode := hex.EncodeToString(buf)
	e, err := h.EncodeHex(hashCode)
	//fmt.Println(e)
	//d, _ := h.DecodeWithError(e)
	//fmt.Println(d)
	return e, err
}
