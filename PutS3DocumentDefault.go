package parser

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//PutS3DocumentDefault put byte array to s3 bucket
func PutS3DocumentDefault(region string, bucket string, key string, content []byte) error {
	sess := session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(region),
		}),
	)
	svc := s3.New(sess)
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(content),
	})
	return err
}
