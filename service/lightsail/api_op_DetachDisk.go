// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DetachDiskInput struct {
	_ struct{} `type:"structure"`

	// The unique name of the disk you want to detach from your instance (e.g.,
	// my-disk).
	//
	// DiskName is a required field
	DiskName *string `locationName:"diskName" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachDiskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachDiskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachDiskInput"}

	if s.DiskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiskName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachDiskOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the time stamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DetachDiskOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachDisk = "DetachDisk"

// DetachDiskRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Detaches a stopped block storage disk from a Lightsail instance. Make sure
// to unmount any file systems on the device within your operating system before
// stopping the instance and detaching the disk.
//
// The detach disk operation supports tag-based access control via resource
// tags applied to the resource identified by disk name. For more information,
// see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DetachDiskRequest.
//    req := client.DetachDiskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DetachDisk
func (c *Client) DetachDiskRequest(input *DetachDiskInput) DetachDiskRequest {
	op := &aws.Operation{
		Name:       opDetachDisk,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachDiskInput{}
	}

	req := c.newRequest(op, input, &DetachDiskOutput{})
	return DetachDiskRequest{Request: req, Input: input, Copy: c.DetachDiskRequest}
}

// DetachDiskRequest is the request type for the
// DetachDisk API operation.
type DetachDiskRequest struct {
	*aws.Request
	Input *DetachDiskInput
	Copy  func(*DetachDiskInput) DetachDiskRequest
}

// Send marshals and sends the DetachDisk API request.
func (r DetachDiskRequest) Send(ctx context.Context) (*DetachDiskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachDiskResponse{
		DetachDiskOutput: r.Request.Data.(*DetachDiskOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachDiskResponse is the response type for the
// DetachDisk API operation.
type DetachDiskResponse struct {
	*DetachDiskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachDisk request.
func (r *DetachDiskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
