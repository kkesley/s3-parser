package parser

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//GetS3DocumentDefault get s3 object
func GetS3DocumentDefault(region string, bucket string, key string) ([]byte, error) {
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
		return make([]byte, 0), err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(s3Output.Body)
	return buf.Bytes(), nil
}
