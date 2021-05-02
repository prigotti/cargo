package application

import (
	"context"
	"io"

	"github.com/prigotti/cargo/common/pb"
	"github.com/prigotti/cargo/portsservice/domain"
	"go.mongodb.org/mongo-driver/mongo"
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
	r          domain.PortRepository // Used by commands only
	collection *mongo.Collection     // Used by queries only
}

// NewPortService is the factory for PortService.
func NewPortService(server *grpc.Server, r domain.PortRepository, database *mongo.Database, portCollection string) PortService {
	collection := database.Collection(portCollection)

	s := &portService{r: r, collection: collection}

	pb.RegisterPortServiceServer(server, s)

	return s
}

// CreateOrUpdateStream will handle streamed create or update commands.
func (s *portService) CreateOrUpdateStream(stream pb.PortService_CreateOrUpdateStreamServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
			p, err := stream.Recv()
			if err == io.EOF {
				return stream.SendAndClose(
					&pb.SendBatchResponse{Message: "port data synchronization finished"},
				)
			}
			if err != nil {
				return err
			}

			dp, err := s.r.GetWithID(stream.Context(), p.Id)
			if err != domain.ErrAggregateNotFound && err != nil {
				return err
			} else if err == domain.ErrAggregateNotFound {
				dp, err := domain.NewPort(
					p.Id, p.Name, p.City, p.Country, p.Alias, p.Regions,
					p.Coordinates, p.Province, p.Timezone, p.Unlocs, p.Code,
				)
				if err == domain.ErrInvalidAggregateData {
					continue
				} else if err != nil {
					return err
				}

				if err := s.r.Save(stream.Context(), dp); err != nil {
					return err
				}
			} else {
				if err := dp.Update(
					p.Name, p.City, p.Country, p.Alias, p.Regions,
					p.Coordinates, p.Province, p.Timezone, p.Unlocs, p.Code,
				); err != nil {
					return err
				}

				if err := s.r.Save(stream.Context(), dp); err != nil {
					return err
				}
			}
		}
	}
}

// ListPorts is a query method and does not act on the domain model,
// so we'll just let it use a database handler for now.
// This is done so that the PortRepository interface doesn't get
// overwhelmed by pagination, filtering and other querying concerns.
// Ideally we'd have a port/interface for generic object queries.
func (s *portService) List(ctx context.Context, q *pb.ListQuery) (*pb.PortListData, error) {
	return nil, nil
}
