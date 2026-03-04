package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	doctorv1 "github.com/eswar/grpc-doctor-booking/gen/go/doctor/v1"
	commonv1 "github.com/eswar/grpc-doctor-booking/gen/go/common/v1"
)

type DoctorServer struct {
	doctorv1.UnimplementedDoctorServiceServer
}

func NewDoctorServer() *DoctorServer {
	return &DoctorServer{}
}

func (s *DoctorServer) ListDoctors(ctx context.Context, req *doctorv1.ListDoctorsRequest) (*doctorv1.ListDoctorsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &doctorv1.ListDoctorsResponse{
		Doctors: []*doctorv1.Doctor{
			{Id: "doc_1", Name: "Dr. Asha Rao", Specialization: req.Specialization, YearsExperience: 8},
			{Id: "doc_2", Name: "Dr. Kiran Mehta", Specialization: req.Specialization, YearsExperience: 12},
		},
		Page: &commonv1.PageResponse{
			NextPageToken: "",
		},
	}

	return resp, nil
}
