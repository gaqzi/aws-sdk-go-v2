// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Client provides the API operation methods for making requests to
// Kinesis. See this package's package overview docs
// for details on the service.
//
// The client's methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Client struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*Client)

// Used for custom request initialization logic
var initRequest func(*Client, *aws.Request)

const (
	ServiceName = "Kinesis" // Service's name
	ServiceID   = "Kinesis" // Service's identifier
	EndpointsID = "kinesis" // Service's Endpoint identifier
)

// New creates a new instance of the client from the provided Config.
//
// Example:
//     // Create a client from just a config.
//     svc := kinesis.New(myConfig)
func New(config aws.Config) *Client {
	svc := &Client{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				EndpointsID:   EndpointsID,
				SigningName:   "kinesis",
				SigningRegion: config.Region,
				APIVersion:    "2013-12-02",
				JSONVersion:   "1.1",
				TargetPrefix:  "Kinesis_20131202",
			},
		),
	}

	if config.Retryer == nil {
		svc.Retryer = retry.NewStandard()
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a client operation and runs any
// custom request initialization.
func (c *Client) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
