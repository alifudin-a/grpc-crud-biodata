package grpc

import (
	"context"
	pb "grpc-crud-biodata/proto/biodata"
)

// Server : initialize server
type Server struct {
}

// CreateBiodata : server for create
func (s *Server) CreateBiodata(ctx context.Context, req *pb.CreateBiodataRequest) (*pb.CreateBiodataResponse, error) {

	return &pb.CreateBiodataResponse{}, nil
}

// ReadBiodata : server for read
func (s *Server) ReadBiodata(ctx context.Context, req *pb.ReadBiodataRequest) (*pb.ReadBiodataResponse, error) {
	return &pb.ReadBiodataResponse{}, nil
}

// UpdateBiodata : server for update
func (s *Server) UpdateBiodata(ctx context.Context, req *pb.UpdateBiodataRequest) (*pb.UpdateBiodataResponse, error) {
	return &pb.UpdateBiodataResponse{}, nil
}

// DeleteBiodata : server for delete
func (s *Server) DeleteBiodata(ctx context.Context, req *pb.DeleteBiodataRequest) (*pb.DeleteBiodataResponse, error) {
	return &pb.DeleteBiodataResponse{}, nil
}

// ListBiodata : server for list
func (s *Server) ListBiodata(ctx context.Context, req *pb.ListBiodataRequest) (*pb.ListBiodataResponse, error) {
	return &pb.ListBiodataResponse{}, nil
}

// NewServer for initialize server
func NewServer() *Server {
	return &Server{}
}
