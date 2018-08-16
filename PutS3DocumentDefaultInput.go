package parser

//PutS3DocumentDefaultInput input object for putting object to s3
type PutS3DocumentDefaultInput struct {
	Region      string
	Bucket      string
	Key         string
	Content     []byte
	ContentType string
	ACL         string
}
