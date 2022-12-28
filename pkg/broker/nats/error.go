package nats

import (
	"github.com/fenixvlabs/meshkit/pkg/errors"
)

const (
	ErrConnectCode        = "11000"
	ErrEncodedConnCode    = "11000"
	ErrPublishCode        = "11001"
	ErrPublishRequestCode = "11001"
	ErrQueueSubscribeCode = "11001"
)

func ErrConnect(err error) error {
	return errors.NewErrorDescription(ErrConnectCode, errors.Alert, []string{"Connection to broker failed"}, []string{err.Error()}, []string{"Endpoint might not be reachable"}, []string{"Make sure the NATS endpoint is reachable"})
}
func ErrEncodedConn(err error) error {
	return errors.NewErrorDescription(ErrEncodedConnCode, errors.Alert, []string{"Encoding connection failed with broker"}, []string{err.Error()}, []string{"Endpoint might not be reachable"}, []string{"Make sure the NATS endpoint is reachable"})
}
func ErrPublish(err error) error {
	return errors.NewErrorDescription(ErrPublishCode, errors.Alert, []string{"Publish failed"}, []string{err.Error()}, []string{"NATS is unhealthy"}, []string{"Make sure NATS is up and running"})
}
func ErrPublishRequest(err error) error {
	return errors.NewErrorDescription(ErrPublishRequestCode, errors.Alert, []string{"Publish request failed"}, []string{err.Error()}, []string{"NATS is unhealthy"}, []string{"Make sure NATS is up and running"})
}
func ErrQueueSubscribe(err error) error {
	return errors.NewErrorDescription(ErrQueueSubscribeCode, errors.Alert, []string{"Subscription failed"}, []string{err.Error()}, []string{"NATS is unhealthy"}, []string{"Make sure NATS is up and running"})
}
