package sample

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/rand"

	"github.com/bozkayasalihx/protobuf/pb"
)

const Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_QWERTZ
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomStringFromFields(fields ...string) string {
	n := len(fields)
	return fields[rand.Intn(n)]
}

func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("couldn't read the bytes : %v\n", err)
	}
	return b
}

func randomString(size int) string {
	letterRunes := []rune(Base62Chars)
	var bb bytes.Buffer
	bb.Grow(size)
	l := uint32(len(letterRunes))
	for i := 0; i < size; i++ {
		bb.WriteRune(letterRunes[binary.LittleEndian.Uint32(Bytes(4))%l])
	}
	return bb.String()
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
