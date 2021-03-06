// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteAlgorithmInput struct {
	_ struct{} `type:"structure"`

	// The name of the algorithm to delete.
	//
	// AlgorithmName is a required field
	AlgorithmName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAlgorithmInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAlgorithmInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAlgorithmInput"}

	if s.AlgorithmName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AlgorithmName"))
	}
	if s.AlgorithmName != nil && len(*s.AlgorithmName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AlgorithmName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAlgorithmOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAlgorithmOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAlgorithm = "DeleteAlgorithm"

// DeleteAlgorithmRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Removes the specified algorithm from your account.
//
//    // Example sending a request using DeleteAlgorithmRequest.
//    req := client.DeleteAlgorithmRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteAlgorithm
func (c *Client) DeleteAlgorithmRequest(input *DeleteAlgorithmInput) DeleteAlgorithmRequest {
	op := &aws.Operation{
		Name:       opDeleteAlgorithm,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAlgorithmInput{}
	}

	req := c.newRequest(op, input, &DeleteAlgorithmOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAlgorithmRequest{Request: req, Input: input, Copy: c.DeleteAlgorithmRequest}
}

// DeleteAlgorithmRequest is the request type for the
// DeleteAlgorithm API operation.
type DeleteAlgorithmRequest struct {
	*aws.Request
	Input *DeleteAlgorithmInput
	Copy  func(*DeleteAlgorithmInput) DeleteAlgorithmRequest
}

// Send marshals and sends the DeleteAlgorithm API request.
func (r DeleteAlgorithmRequest) Send(ctx context.Context) (*DeleteAlgorithmResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAlgorithmResponse{
		DeleteAlgorithmOutput: r.Request.Data.(*DeleteAlgorithmOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAlgorithmResponse is the response type for the
// DeleteAlgorithm API operation.
type DeleteAlgorithmResponse struct {
	*DeleteAlgorithmOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAlgorithm request.
func (r *DeleteAlgorithmResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
