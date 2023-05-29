package service

import (
	"context"
	"errors"
	"log"

	"github.com/bozkayasalihx/protobuf/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)


type LaptopServer struct {
  laptopStore LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
    laptopStore: store,
  }
}

func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse ,error) {
  laptop := req.GetLaptop()
  log.Printf("receivce create-laptop with id: %v", laptop.Id)

  if len(laptop.Id) > 0 {
    if _ ,err := uuid.Parse(laptop.Id); err != nil {
      return nil, status.Errorf(codes.InvalidArgument, "laptop id is not a valid uuid : %v", err);
    }else {
      id, _ := uuid.NewRandom()
      laptop.Id  = id.String()
    }
  }


  err := server.laptopStore.Save(laptop)
  if err != nil {
    code := codes.Internal
    if errors.Is(err, AllreadyExists) {
      code = codes.AlreadyExists
    }
    return nil, status.Error(code, "couldn't save data to db")
  }
  
  log.Printf("saved laptop with id: %s", laptop.Id)

  res := &pb.CreateLaptopResponse{
    Id: laptop.Id,
  }
  return res, nil
}
