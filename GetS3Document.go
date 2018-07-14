package parser

import (
	"github.com/aws/aws-lambda-go/events"
)

//GetS3Document get s3 object
func GetS3Document(region string, record events.S3EventRecord) ([]byte, error) {
	return GetS3DocumentDefault(region, record.S3.Bucket.Name, record.S3.Object.Key)
}
