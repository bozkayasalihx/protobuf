package sample

import (

	"github.com/bozkayasalihx/protobuf/pb"
	"github.com/google/uuid"
)

 
func NewKeyboard() *pb.Keyboard {
  return &pb.Keyboard{
    Layout: randomKeyboardLayout(),
    Backlit: randomBool(),
  }
}


func NewCPU() *pb.CPU {
  return  &pb.CPU{
    Brand: randomStringFromFields("Intel", "AMD"),
    Name: randomString(6),
    NumberCores: uint32(randomInt(4,8)),
    NumberThreads: uint32(randomInt(5,12)),
    MinGhz: randomFloat64(2.0, 5.0),  
    MaxGhz: randomFloat64(2.0, 5.0),
  }
}
 
func NewGPU()  *pb.GPU {
  return &pb.GPU{
    Brand: randomStringFromFields("Intel", "AMD"),
    Name: randomString(6),
    MinGhz: randomFloat64(2.0, 5.0),  
    MaxGhz: randomFloat64(2.0, 5.0),
    Memory: &pb.Memory{
      Value: uint64(randomInt(8,32)),
      Unit: pb.Memory_GIGABYTE,
    },
  }
}

func NewRAM() *pb.Memory {
  return &pb.Memory{
    Value: uint64(randomInt(8,32)),
    Unit: pb.Memory_GIGABYTE,
  }
}


func NewSSD() *pb.Storage {
  return &pb.Storage{
    Memory: &pb.Memory{
      Value: uint64(randomInt(8,32)),
      Unit: pb.Memory_GIGABYTE,
    },
    Driver: pb.Storage_SSD,
  }
}

func NewHDD() *pb.Storage {
  return &pb.Storage{
    Memory: &pb.Memory{
      Value: uint64(randomInt(8,32)),
      Unit: pb.Memory_GIGABYTE,
    },
    Driver: pb.Storage_HDD,
  }
}


func NewLaptop() *pb.Laptop {
  id ,_ := uuid.NewRandom()
  
  return &pb.Laptop{
      Id: id.String(), 
      Brand: randomStringFromFields("Intel", "AMD"),
      Name: randomString(6),
      Cpu: NewCPU(), 
      Ram:  NewRAM(), 
      Gpus: []*pb.GPU{NewGPU(), NewGPU()}, 
      Storages: []*pb.Storage{NewSSD(), NewHDD()},
      Screen: &pb.Screen{
        SizeInch: float32(10) , 
        Resolution: &pb.Screen_Resolution{
          Width: 100,
          Height: 100, 
        },
        Multitouch: true,
      }, 
      Keyboard: NewKeyboard(), 
      PriceUsd: float64(randomInt(1500, 3000)), 
      ReleaseYear: uint32(10),
  }
}


