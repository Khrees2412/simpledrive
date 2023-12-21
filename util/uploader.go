package util

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	AwsAccessKey = os.Getenv("AWS_ACCESS_KEY")
	AwsSecretKey = os.Getenv("AWS_SECRET_KEY")
	AwsBucket    = os.Getenv("AWS_BUCKET")
)

func connectAWS() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String("eu-north-1"),
			Credentials: credentials.NewStaticCredentials(AwsAccessKey, AwsSecretKey, ""),
		})
	if err != nil {
		panic(err)
	}

	return sess
}

func UploadFile(file *multipart.FileHeader) (*s3manager.UploadOutput, error) {
	buffer, err := file.Open()

	if err != nil {
		log.Error(fmt.Sprintf("Error opening file: %s", err))
	}
	defer buffer.Close()

	uploader := s3manager.NewUploader(connectAWS())

	data, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AwsBucket),     // Bucket to be used
		Key:    aws.String(file.Filename), // Name of the file to be saved
		Body:   buffer,                    // File
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}
