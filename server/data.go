package main

import (
	"context"
	"grpc-server/pb"
)

var employee = &pb.EmployeeResponse{

	Id:            1,
	Name:          "张三",
	Height:        180,
	Weight:        180,
	Avatar:        []byte{},
	Email:         "123@qq.com",
	EmailVerified: true,
	PhoneNumbers:  []string{"151911231"},
	Gender:        1,
	// {
	// 	Id:            2,
	// 	Name:          "李四",
	// 	Height:        170,
	// 	Weight:        100,
	// 	Avatar:        []byte{},
	// 	Email:         "1234@qq.com",
	// 	EmailVerified: true,
	// 	PhoneNumbers:  []string{"1515911231"},
	// 	Gender:        1,
	// },
}

type employeeService struct {
}

func (e *employeeService) GetEmployee(ctx context.Context,
	in *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return employee, nil
}
