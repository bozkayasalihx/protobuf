package service_test

import (
	"context"
	"testing"

	"github.com/bozkayasalihx/protobuf/pb"
	"github.com/bozkayasalihx/protobuf/sample"
	"github.com/bozkayasalihx/protobuf/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestCreateLaptopServer(t *testing.T) {
	t.Parallel()

	withoutIdLaptop := sample.NewLaptop()
	withoutIdLaptop.Id = ""

	failureLaptop := sample.NewLaptop()
	failureLaptop.Id = "invalid_id"

	cases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},

		{
			name:   "success_without_id",
			laptop: withoutIdLaptop,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},

		{
			name:   "success_without_id",
			laptop: withoutIdLaptop,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},

		{
			name:   "failure_invalid_id",
			laptop: failureLaptop,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},

		{
			name:   "duplicate_invalid_id",
			laptop: failureLaptop,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.AlreadyExists,
		},
	}

	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := service.NewLaptopServer(tc.store)
			resp, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.NotEmpty(t, resp.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, resp.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, resp)
			}
		})
	}

}
