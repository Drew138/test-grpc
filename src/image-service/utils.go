package imageService

import (
	"bytes"
	"errors"
	"image"
	"log"

	"github.com/drew138/go-graphics/filters"
	"github.com/drew138/go-graphics/filters/encoding"
	"github.com/drew138/go-graphics/filters/kernels"
	"github.com/drew138/test-grpc/src/messages"
)

func processImage(req *messages.ImageRequest, kernel kernels.Kernel, transformation int) (*messages.ImageResponse, error) {
	body := req.GetImage()
	img, format, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		log.Fatalf("Could not decode image: %v", err)
		return nil, err
	}
	var responseImg image.Image
	switch transformation {
	case filter:
		responseImg = filters.ApplyFilter(img, kernel)
	case negative:
		responseImg = filters.CreateNegativeImage(img)
	}

	buffer := new(bytes.Buffer)
	if err := encoding.EncodeImageToBytes(responseImg, format, buffer); err != nil {
		log.Fatalf("Could not encode resulting image %v: ", err)
		return nil, err
	}
	return &messages.ImageResponse{Image: buffer.Bytes()}, nil
}

func createKernel(req *messages.ImageRequest) (kernels.Kernel, error) {
	customKernel := req.GetKernel()
	if customKernel == nil {
		return nil, errors.New("kernel was not provided")
	}
	kernel := kernels.CreateKernelFromFloats(
		customKernel.GetRowOneColOne(),
		customKernel.GetRowOneColTwo(),
		customKernel.GetRowOneColThree(),
		customKernel.GetRowTwoColOne(),
		customKernel.GetRowTwoColTwo(),
		customKernel.GetRowTwoColThree(),
		customKernel.GetRowThreeColOne(),
		customKernel.GetRowThreeColTwo(),
		customKernel.GetRowThreeColThree(),
	)
	return kernel, nil
}
