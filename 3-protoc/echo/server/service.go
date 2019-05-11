package main

import (
  "context"
  "log"
  "../proto"
  //pb "github.com/vvatanabe/go-grpc-basics/echo/proto"
)

type echoService struct{}

func (s *echoService) Echo(ctx context.Context,
  req *echo.EchoRequest) (*echo.EchoResponse, error){
  log.Println(req)
  return &echo.EchoResponse{
    Message: req.GetMessage()}, nil
}
