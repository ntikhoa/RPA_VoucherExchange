package services

import (
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type ImageService interface {
	UploadObjects(files []*multipart.FileHeader) ([]string, error)
}

type imageService struct {
	_BUCKET_NAME   string
	_BUCKET_REGION string
	s3session      *s3.S3
}

type UploadResult struct {
	FileName string
	Err      error
}

func NewImageService() ImageService {

	bucketName := os.Getenv("AWS_BUCKET_NAME")
	bucketRegion := os.Getenv("AWS_BUCKET_REGION")

	s3session := s3.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(bucketRegion),
		Credentials: credentials.NewEnvCredentials(),
	})))

	return &imageService{
		_BUCKET_NAME:   bucketName,
		_BUCKET_REGION: bucketRegion,
		s3session:      s3session,
	}
}

func (s *imageService) uploadObject(file *multipart.FileHeader, chanResFileName chan UploadResult) {
	id := uuid.New()
	file.Filename = id.String() + file.Filename
	f, err := file.Open()
	if err != nil {
		chanResFileName <- UploadResult{"", err}
		return
	}

	_, err = s.s3session.PutObject(&s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(s._BUCKET_NAME),
		Key:    aws.String(file.Filename),
	})

	if err != nil {
		chanResFileName <- UploadResult{"", err}
		return
	}
	chanResFileName <- UploadResult{file.Filename, nil}
}

func (s *imageService) UploadObjects(files []*multipart.FileHeader) ([]string, error) {

	uploadResultChan := make(chan UploadResult, len(files))
	for _, file := range files {
		go s.uploadObject(file, uploadResultChan)
	}
	var filesName []string
	for i := 0; i < len(files); i++ {
		uploadResult := <-uploadResultChan
		if uploadResult.Err != nil {
			return filesName, uploadResult.Err
		}
		filesName = append(filesName, uploadResult.FileName)
	}
	return filesName, nil
}
