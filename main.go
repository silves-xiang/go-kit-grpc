package main
import(
	"bkgrpc/router"
	"bkgrpc/services"
	"bkgrpc/endpoint"
	"google.golang.org/grpc"
	"bkgrpc/proto"
	"log"
	"net"
)
func main () {
	svc := services.ServicesA{}
	endpoints := endpoint.EndpointA{
		ConcatEndpoint: endpoint.MakeConcatEndpoint(svc),
		DiffEndpoint: endpoint.MakeDiffEndpoint(svc),
		HealthEndpoint: endpoint.MakeHealthEndpoint(svc),
	}
	r := router.NewRouter(endpoints)
	lis , err := net.Listen("tcp" , ":8085")
	if err != nil {
		log.Println(err)
		return
	}
	grpcserver := grpc.NewServer()
	gproto.RegisterStringServicesServer(grpcserver , r)
	grpcserver.Serve(lis)
}
