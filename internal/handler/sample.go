package handler

//
//import (
//	"fmt"
//	"github.com/bufbuild/protovalidate-go"
//	pb "github.com/jnst/supercell-id/pkg/gen/authentication/v1"
//)
//
//func main() {
//	req := &pb.RegisterRequest{
//		Email: "aaa@example.jp",
//	}
//
//	v, err := protovalidate.New()
//	if err != nil {
//		fmt.Println("failed to initialize validator:", err)
//	}
//
//	if err = v.Validate(req); err != nil {
//		fmt.Println("validation failed:", err)
//	} else {
//		fmt.Println("validation succeeded")
//	}
//}
