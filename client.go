package main

import (
	pb "bkgrpc/proto"
//	"bkgrpc/services"
	"context"
//	"flag"
	"fmt"
	"google.golang.org/grpc"
//	"time"
)

func main() {
	ctx :=context.Background()
	con,_:=grpc.Dial(":8085" , grpc.WithInsecure())
	fu :=pb.NewStringServicesClient(con)
	rep , _ :=fu.HealtStatus(ctx , &pb.HealthRequest{A: "a"})
	fmt.Println(rep)
}
