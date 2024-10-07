package controllers

import (
	"context"
	"fmt"
	"src/index/src/db"
	"src/index/src/models"
	"src/index/src/utils"
	"time"

	pb "src/index/src/gRPC/hotel"
)

type HotelServer struct {
	pb.UnimplementedHotelServiceServer
}

func (s *HotelServer) AddHotel(ctx context.Context, in *pb.AddHotelRequest) (*pb.AddHotelResponse, error) {
	utils.Logger.Printf("[ADD HOTEL] Received: %v", in.GetName())

	if in.GetName() == "" || in.GetAddress() == "" {
		return nil, fmt.Errorf("name and address cannot be empty")
	}

	hotel := models.Hotel{
		Name:        in.GetName(),
		Address:     in.GetAddress(),
		Phone:       in.GetContactNo(),
		Email:       in.GetEmail(),
		Website:     in.GetWebsite(),
		Description: in.GetDescription(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := db.DB.Create(&hotel)
	if result.Error != nil {
		utils.Logger.Printf("[ADD HOTEL] Error saving hotel: %v", result.Error)
		return nil, fmt.Errorf("error saving hotel")
	}

	return &pb.AddHotelResponse{Message: "Hotel Added Successfuly"}, nil
}

func (s *HotelServer) GetHotel(ctx context.Context, in *pb.GetHotelRequest) (*pb.GetHotelResponse, error) {
	utils.Logger.Printf("[GET HOTEL] Received: %v", in.GetId())

	var hotel models.Hotel
	result := db.DB.First(&hotel, in.GetId())
	if result.Error != nil {
		utils.Logger.Printf("[GET HOTEL] Error getting hotel: %v", result.Error)
		return nil, fmt.Errorf("error getting hotel")
	}

	return &pb.GetHotelResponse{
		Id:          uint64(hotel.ID),
		Name:        hotel.Name,
		Address:     hotel.Address,
		ContactNo:   hotel.Phone,
		Email:       hotel.Email,
		Website:     hotel.Website,
		Description: hotel.Description,
	}, nil
}

func (s *HotelServer) UpdateHotel(ctx context.Context, in *pb.UpdateHotelRequest) (*pb.UpdateHotelResponse, error) {
	utils.Logger.Printf("[UPDATE HOTEL] Received: %v", in.GetId())

	var hotel models.Hotel
	result := db.DB.First(&hotel, in.GetId())
	if result.Error != nil {
		utils.Logger.Printf("[UPDATE HOTEL] Error getting hotel: %v", result.Error)
		return nil, fmt.Errorf("error getting hotel")
	}

	hotel.Name = in.GetName()
	hotel.Address = in.GetAddress()
	hotel.Phone = in.GetContactNo()
	hotel.Email = in.GetEmail()
	hotel.Website = in.GetWebsite()
	hotel.Description = in.GetDescription()
	hotel.UpdatedAt = time.Now()

	result = db.DB.Save(&hotel)
	if result.Error != nil {
		utils.Logger.Printf("[UPDATE HOTEL] Error updating hotel: %v", result.Error)
		return nil, fmt.Errorf("error updating hotel")
	}

	return &pb.UpdateHotelResponse{Message: "Hotel Updated Successfuly"}, nil
}

func (s *HotelServer) DeleteHotel(ctx context.Context, in *pb.DeleteHotelRequest) (*pb.DeleteHotelResponse, error) {

	utils.Logger.Printf("[DELETE HOTEL] Received: %v", in.GetId())

	var hotel models.Hotel
	result := db.DB.First(&hotel, in.GetId())
	if result.Error != nil {
		utils.Logger.Printf("[DELETE HOTEL] Error getting hotel: %v", result.Error)
		return nil, fmt.Errorf("error getting hotel")
	}

	result = db.DB.Delete(&hotel)
	if result.Error != nil {
		utils.Logger.Printf("[DELETE HOTEL] Error deleting hotel: %v", result.Error)
		return nil, fmt.Errorf("error deleting hotel")
	}

	return &pb.DeleteHotelResponse{Message: "Hotel Deleted Successfuly"}, nil
}

func (s *HotelServer) ListHotels(ctx context.Context, in *pb.ListHotelsRequest) (*pb.ListHotelsResponse, error) {
	utils.Logger.Printf("[LIST HOTELS] Received")

	var hotels []models.Hotel
	result := db.DB.Find(&hotels)
	if result.Error != nil {
		utils.Logger.Printf("[LIST HOTELS] Error getting hotels: %v", result.Error)
		return nil, fmt.Errorf("error getting hotels")
	}

	hotelsResponse := []*pb.Hotel{}
	for _, hotel := range hotels {
		hotelsResponse = append(hotelsResponse, &pb.Hotel{
			Id:          uint64(hotel.ID),
			Name:        hotel.Name,
			Address:     hotel.Address,
			ContactNo:   hotel.Phone,
			Email:       hotel.Email,
			Website:     hotel.Website,
			Description: hotel.Description,
		})
	}

	return &pb.ListHotelsResponse{Hotels: hotelsResponse}, nil
}