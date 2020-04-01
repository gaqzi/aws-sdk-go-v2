// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DetachStaticIpInput struct {
	_ struct{} `type:"structure"`

	// The name of the static IP to detach from the instance.
	//
	// StaticIpName is a required field
	StaticIpName *string `locationName:"staticIpName" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachStaticIpInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachStaticIpInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachStaticIpInput"}

	if s.StaticIpName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StaticIpName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachStaticIpOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the time stamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DetachStaticIpOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachStaticIp = "DetachStaticIp"

// DetachStaticIpRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Detaches a static IP from the Amazon Lightsail instance to which it is attached.
//
//    // Example sending a request using DetachStaticIpRequest.
//    req := client.DetachStaticIpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DetachStaticIp
func (c *Client) DetachStaticIpRequest(input *DetachStaticIpInput) DetachStaticIpRequest {
	op := &aws.Operation{
		Name:       opDetachStaticIp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachStaticIpInput{}
	}

	req := c.newRequest(op, input, &DetachStaticIpOutput{})
	return DetachStaticIpRequest{Request: req, Input: input, Copy: c.DetachStaticIpRequest}
}

// DetachStaticIpRequest is the request type for the
// DetachStaticIp API operation.
type DetachStaticIpRequest struct {
	*aws.Request
	Input *DetachStaticIpInput
	Copy  func(*DetachStaticIpInput) DetachStaticIpRequest
}

// Send marshals and sends the DetachStaticIp API request.
func (r DetachStaticIpRequest) Send(ctx context.Context) (*DetachStaticIpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachStaticIpResponse{
		DetachStaticIpOutput: r.Request.Data.(*DetachStaticIpOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachStaticIpResponse is the response type for the
// DetachStaticIp API operation.
type DetachStaticIpResponse struct {
	*DetachStaticIpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachStaticIp request.
func (r *DetachStaticIpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
