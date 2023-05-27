package sample

import (
	"math/rand"

	"github.com/bozkayasalihx/protobuf/pb"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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
  return rand.Intn(2) ==1;
}

func  randomStringFromFields(fields ...string) string {
  n := len(fields)
  return fields[rand.Intn(n)]
}

func randomString(size int) string {
  b := make([]byte, size) 
  for i := range b {
    if _, err := rand.Read(b[i:i+1]); err != nil {
      panic(err)
    }

    b[i] = charset[b[i]%byte(len(charset))]
  }

  return string(b)
}

func randomInt(min, max int) int {
  return  min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64{
  return min+rand.Float64()*(max-min)
}


