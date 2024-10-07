package controllers

import (
	"context"
	"fmt"
	"os"
	"src/index/src/common"
	"src/index/src/db"
	pb "src/index/src/gRPC/user"
	"src/index/src/models"
	"src/index/src/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	common.Server
}

func (s *Server) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	utils.Logger.Printf("[REGISTER]Received: %v", in.GetUsername())

	if in.GetUsername() == "" || in.GetPassword() == "" {
		return nil, fmt.Errorf("username and password cannot be empty")
	}

	hashedPassword, err := utils.HashPassword(in.GetPassword())
	if err != nil {
		utils.Logger.Printf("[REGISTER]Error hashing password: %v", err)
		return nil, fmt.Errorf("error hashing password")
	}

	hotel_id, _ := strconv.ParseUint(in.GetHotelId(), 10, 64)
	user := models.User{
		Username:   in.GetUsername(),
		Password:   hashedPassword,
		Contact_no: in.GetContactNo(),
		Email:      in.GetEmail(),
		Name:       in.GetName(),
		Role:       in.GetRole(),
		Hotel_ID:   uint(hotel_id),
	}
	result := db.DB.Create(&user)
	if result.Error != nil {
		utils.Logger.Printf("[REGISTER]Error saving user: %v", result.Error)
		return nil, fmt.Errorf("error saving user")
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": in.GetUsername(),
		"role":     in.GetRole(),
		"exp":      time.Now().Add(time.Hour * 24 * 10).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("jwt_secret")))
	if err != nil {
		utils.Logger.Printf("[REGISTER]Error generating token: %v", err)
		return nil, fmt.Errorf("error generating token")
	}

	redisClient, err := utils.GetRedisClientFromContext(ctx)
	if err != nil {
		utils.Logger.Errorf("[REGISTER]Error retrieving Redis client from context: %v", err)
		return nil, fmt.Errorf("error retrieving Redis client")
	}

	redisKey := fmt.Sprintf("session:%s:%s", in.GetUsername(), uuid.NewString())

	err = redisClient.Set(ctx, redisKey, tokenString, time.Hour*24*10).Err()
	if err != nil {
		utils.Logger.Printf("[REGISTER]Error saving token in redis: %v", err)
		return nil, fmt.Errorf("error saving token in redis")
	}

	md := metadata.Pairs("session-token", redisKey)
	metadata.NewOutgoingContext(ctx, md)

	utils.Logger.Printf("[REGISTER]User %v registered successfully", in.GetUsername())
	return &pb.RegisterUserResponse{Message: "Hello " + in.GetUsername()}, nil
}
