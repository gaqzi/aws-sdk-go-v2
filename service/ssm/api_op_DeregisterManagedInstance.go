// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterManagedInstanceInput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to the managed instance when you registered it using the
	// activation process.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterManagedInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterManagedInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterManagedInstanceInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterManagedInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterManagedInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterManagedInstance = "DeregisterManagedInstance"

// DeregisterManagedInstanceRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Removes the server or virtual machine from the list of registered servers.
// You can reregister the instance again at any time. If you don't plan to use
// Run Command on the server, we suggest uninstalling SSM Agent first.
//
//    // Example sending a request using DeregisterManagedInstanceRequest.
//    req := client.DeregisterManagedInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeregisterManagedInstance
func (c *Client) DeregisterManagedInstanceRequest(input *DeregisterManagedInstanceInput) DeregisterManagedInstanceRequest {
	op := &aws.Operation{
		Name:       opDeregisterManagedInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterManagedInstanceInput{}
	}

	req := c.newRequest(op, input, &DeregisterManagedInstanceOutput{})
	return DeregisterManagedInstanceRequest{Request: req, Input: input, Copy: c.DeregisterManagedInstanceRequest}
}

// DeregisterManagedInstanceRequest is the request type for the
// DeregisterManagedInstance API operation.
type DeregisterManagedInstanceRequest struct {
	*aws.Request
	Input *DeregisterManagedInstanceInput
	Copy  func(*DeregisterManagedInstanceInput) DeregisterManagedInstanceRequest
}

// Send marshals and sends the DeregisterManagedInstance API request.
func (r DeregisterManagedInstanceRequest) Send(ctx context.Context) (*DeregisterManagedInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterManagedInstanceResponse{
		DeregisterManagedInstanceOutput: r.Request.Data.(*DeregisterManagedInstanceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterManagedInstanceResponse is the response type for the
// DeregisterManagedInstance API operation.
type DeregisterManagedInstanceResponse struct {
	*DeregisterManagedInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterManagedInstance request.
func (r *DeregisterManagedInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
