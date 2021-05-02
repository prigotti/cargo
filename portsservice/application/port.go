package application

import (
	"context"

	"github.com/prigotti/cargo/common/pb"
	"github.com/prigotti/cargo/portsservice/domain"
	"google.golang.org/grpc"
)

// PortService is the application service mainly responsible for
// handling commands and queries (CQRS not being fully employed) to
// the Post domain model.
type PortService interface {
	CreateOrUpdateStream(stream pb.PortService_CreateOrUpdateStreamServer) error
	List(ctx context.Context, q *pb.ListQuery) (*pb.PortListData, error)
}

type portService struct {
	pb.UnimplementedPortServiceServer
	r domain.PortRepository
}

func NewPortService(server *grpc.Server, r domain.PortRepository) PortService {
	s := &portService{r: r}

	pb.RegisterPortServiceServer(server, s)

	return s
}

func CreateOrUpdateStream(stream pb.PortService_CreateOrUpdateStreamServer) error {
	return nil
}

func List(ctx context.Context, q *pb.ListQuery) (*pb.PortListData, error) {
	return nil, nil
}
