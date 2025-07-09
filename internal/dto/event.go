package dto

type InputMessage struct {
	BucketName string `json:"bucketName"`
	Key        string `json:"key"`
	Url        string `json:"url"`
}
