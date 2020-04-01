// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type WithdrawByoipCidrInput struct {
	_ struct{} `type:"structure"`

	// The address range, in CIDR notation.
	//
	// Cidr is a required field
	Cidr *string `type:"string" required:"true"`
}

// String returns the string representation
func (s WithdrawByoipCidrInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *WithdrawByoipCidrInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "WithdrawByoipCidrInput"}

	if s.Cidr == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cidr"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type WithdrawByoipCidrOutput struct {
	_ struct{} `type:"structure"`

	// Information about the address pool.
	ByoipCidr *ByoipCidr `type:"structure"`
}

// String returns the string representation
func (s WithdrawByoipCidrOutput) String() string {
	return awsutil.Prettify(s)
}

const opWithdrawByoipCidr = "WithdrawByoipCidr"

// WithdrawByoipCidrRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Stops advertising an address range that is provisioned as an address pool.
// You can perform this operation at most once every 10 seconds, even if you
// specify different address ranges each time. To see an AWS CLI example of
// withdrawing an address range for BYOIP so it will no longer be advertised
// by AWS, scroll down to Example.
//
// It can take a few minutes before traffic to the specified addresses stops
// routing to AWS because of propagation delays.
//
// For more information, see Bring Your Own IP Addresses (BYOIP) (https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html)
// in the AWS Global Accelerator Developer Guide.
//
//    // Example sending a request using WithdrawByoipCidrRequest.
//    req := client.WithdrawByoipCidrRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/WithdrawByoipCidr
func (c *Client) WithdrawByoipCidrRequest(input *WithdrawByoipCidrInput) WithdrawByoipCidrRequest {
	op := &aws.Operation{
		Name:       opWithdrawByoipCidr,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &WithdrawByoipCidrInput{}
	}

	req := c.newRequest(op, input, &WithdrawByoipCidrOutput{})
	return WithdrawByoipCidrRequest{Request: req, Input: input, Copy: c.WithdrawByoipCidrRequest}
}

// WithdrawByoipCidrRequest is the request type for the
// WithdrawByoipCidr API operation.
type WithdrawByoipCidrRequest struct {
	*aws.Request
	Input *WithdrawByoipCidrInput
	Copy  func(*WithdrawByoipCidrInput) WithdrawByoipCidrRequest
}

// Send marshals and sends the WithdrawByoipCidr API request.
func (r WithdrawByoipCidrRequest) Send(ctx context.Context) (*WithdrawByoipCidrResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &WithdrawByoipCidrResponse{
		WithdrawByoipCidrOutput: r.Request.Data.(*WithdrawByoipCidrOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// WithdrawByoipCidrResponse is the response type for the
// WithdrawByoipCidr API operation.
type WithdrawByoipCidrResponse struct {
	*WithdrawByoipCidrOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// WithdrawByoipCidr request.
func (r *WithdrawByoipCidrResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
