package parser

import (
	"bytes"
)

//GetS3DocumentDefault get s3 object
func GetS3DocumentDefault(region string, bucket string, key string) ([]byte, error) {
	output, err := GetS3DocumentDefaultRaw(region, bucket, key)

	if err != nil || output == nil {
		return make([]byte, 0), err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(output.Body)
	return buf.Bytes(), nil
}
