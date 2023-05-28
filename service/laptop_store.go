package service

import (
	"errors"
	"sync"

	"github.com/bozkayasalihx/protobuf/pb"
	"github.com/bozkayasalihx/protobuf/sample/util"
)

var AllreadyExists = errors.New("this data allready exists in KV")
var NotFound = errors.New("data not found");


type LaptopStore interface {
  Save(l *pb.Laptop) error
  Find(id string)  (pb.Laptop, error)
}

type InMemoryLaptopStore struct {
	data map[string]*pb.Laptop
  mux  *sync.RWMutex
}


func NewInMemoryLaptopStore() *InMemoryLaptopStore {
  return &InMemoryLaptopStore{
    data: make(map[string]*pb.Laptop), 
    mux: &sync.RWMutex{},
  }
}


func (store *InMemoryLaptopStore) Save(l *pb.Laptop) error {
  store.mux.Lock()
  defer store.mux.Unlock()
  if _, ok := store.data[l.Id]; ok {
    return AllreadyExists
  }
  copy := util.DeepCopy(l)
  store.data[copy.Id] = copy
  return nil
}


func (store *InMemoryLaptopStore) Find(id string)  (*pb.Laptop,error){
  store.mux.RLock()
  defer store.mux.RUnlock()
  if _, ok := store.data[id]; !ok {
      return nil,  NotFound;
  }
  return store.data[id], nil
}




 
