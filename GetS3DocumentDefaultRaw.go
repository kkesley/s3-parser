package parser

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//GetS3DocumentDefaultRaw get s3 object and return output body
func GetS3DocumentDefaultRaw(region string, bucket string, key string) (*s3.GetObjectOutput, error) {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	svc := s3.New(sess)

	s3Output, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}
	return s3Output, nil
}
