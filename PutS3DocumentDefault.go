package parser

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//PutS3DocumentDefault put byte array to s3 bucket
func PutS3DocumentDefault(input PutS3DocumentDefaultInput) error {
	sess := session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(input.Region),
		}),
	)
	svc := s3.New(sess)
	inputs3 := &s3.PutObjectInput{
		Bucket: aws.String(input.Bucket),
		Key:    aws.String(input.Key),
		Body:   bytes.NewReader(input.Content),
	}
	if len(input.ACL) > 0 {
		inputs3.ACL = aws.String(input.ACL)
	}
	if len(input.ContentType) > 0 {
		inputs3.ContentType = aws.String(input.ContentType)
	}
	_, err := svc.PutObject(inputs3)
	return err
}
