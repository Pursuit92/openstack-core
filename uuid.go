package core

import (
	"crypto/rand"
	"fmt"
)

type Uuid struct {
	b []byte
}

func Uuidgen() *Uuid {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return nil
	}

	b[8] = (b[8] | 0x80) & 0xBF
	b[6] = (b[6] | 0x40) & 0x4F
	return &Uuid{b}
}

func (u *Uuid) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		u.b[0:4], u.b[4:6], u.b[6:8], u.b[8:10], u.b[10:])
}

func (u *Uuid) Bytes() []byte {
	ret := make([]byte,16)
	for i,v := range u.b {
		ret[i] = v
	}
	return ret
}
