package imageService

import (
	"context"
	"log"

	"github.com/drew138/go-graphics/filters/kernels"
	"github.com/drew138/test-grpc/src/messages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Type of process to be applied to image
const (
	filter = iota
	negative
)

type ImageServiceServer struct {
	messages.UnimplementedImageServiceServer
}

func (*ImageServiceServer) ApplyNegative(ctx context.Context, req *messages.ImageRequest) (*messages.ImageResponse, error) {
	res, err := processImage(req, nil, negative)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return res, nil
}

func (*ImageServiceServer) ApplyCustomFilter(ctx context.Context, req *messages.ImageRequest) (*messages.ImageResponse, error) {
	kernel, err := createKernel(req)
	if err != nil {
		log.Fatalf("Could not apply filter: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	res, err := processImage(req, kernel, filter)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error()) // Error already logged
	}
	return res, nil
}

func (*ImageServiceServer) ApplyGaussianBlurFilter(ctx context.Context, req *messages.ImageRequest) (*messages.ImageResponse, error) {
	res, err := processImage(req, kernels.GaussianBlur, filter)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error()) // Error already logged
	}
	return res, nil
}

func (*ImageServiceServer) ApplyBoxBlurFilter(ctx context.Context, req *messages.ImageRequest) (*messages.ImageResponse, error) {
	res, err := processImage(req, kernels.BoxBlur, filter)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error()) // Error already logged
	}
	return res, nil
}

func (*ImageServiceServer) ApplyEdgeDetectionFilter(ctx context.Context, req *messages.ImageRequest) (*messages.ImageResponse, error) {
	res, err := processImage(req, kernels.EdgeDetection, filter)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error()) // Error already logged
	}
	return res, nil
}

func (*ImageServiceServer) ApplySharpenFilter(ctx context.Context, req *messages.ImageRequest) (*messages.ImageResponse, error) {
	res, err := processImage(req, kernels.Sharpen, filter)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error()) // Error already logged
	}
	return res, nil
}
