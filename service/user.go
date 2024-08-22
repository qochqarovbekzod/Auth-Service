package service

import (
	"context"
	pb "users/generated/users"
	"users/storage/postgres"
)

type UserService struct {
	pb.UnimplementedAuthServiceServer
	UserRepo *postgres.UserRepo
}

func NewUserServer(repo *postgres.UserRepo) *UserService {
	return &UserService{UserRepo: repo}
}

func (u *UserService) ViewProfile(ctx context.Context, in *pb.ViewProfileRequest) (*pb.ViewProfileResponse, error) {
	return u.UserRepo.ViewProfile(in.Id)
}

func (u *UserService) EditProfile(ctx context.Context, in *pb.EditProfileRequeste) (*pb.EditProfileResponse, error) {
	return u.UserRepo.EditProfile(in)
}

func (u *UserService) ChangeUserType(ctx context.Context, in *pb.ChangeUserTypeRequeste) (*pb.ChangeUserTypeResponse, error) {
	return u.UserRepo.ChangeUserType(in)
}

func (u *UserService) GetAllUsers(ctx context.Context, in *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	return u.UserRepo.GetAllUsers(in)
}

func (u *UserService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return u.UserRepo.DeleteUser(in.Id)
}

func (u *UserService) PasswordReset(ctx context.Context, in *pb.PasswordResetRequest) (*pb.PasswordResetResponse, error) {
	return &pb.PasswordResetResponse{}, nil
}

func (u *UserService) TokenGeneration(ctx context.Context, in *pb.TokenGenerationRequest) (*pb.TokenGenerationResponse, error) {
	return &pb.TokenGenerationResponse{}, nil
}

func (u *UserService) TokenCancellation(ctx context.Context, in *pb.TokenCancellationRequest) (*pb.TokenCancellationResponse, error) {
	return &pb.TokenCancellationResponse{}, nil
}
