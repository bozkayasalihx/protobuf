package service

import (
	"context"
	"log"

	"github.com/bozkayasalihx/protobuf/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)



type LaptopServer struct {
  store LaptopStore
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
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



  return nil ,nil

}
