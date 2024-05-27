package services

type AWSService interface {
	ListResources() error
}
