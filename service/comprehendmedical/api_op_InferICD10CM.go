// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type InferICD10CMInput struct {
	_ struct{} `type:"structure"`

	// The input text used for analysis. The input for InferICD10CM is a string
	// from 1 to 10000 characters.
	//
	// Text is a required field
	Text *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s InferICD10CMInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InferICD10CMInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InferICD10CMInput"}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if s.Text != nil && len(*s.Text) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Text", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type InferICD10CMOutput struct {
	_ struct{} `type:"structure"`

	// The medical conditions detected in the text linked to ICD-10-CM concepts.
	// If the action is successful, the service sends back an HTTP 200 response,
	// as well as the entities detected.
	//
	// Entities is a required field
	Entities []ICD10CMEntity `type:"list" required:"true"`

	// The version of the model used to analyze the documents, in the format n.n.n
	// You can use this information to track the model used for a particular batch
	// of documents.
	ModelVersion *string `min:"1" type:"string"`

	// If the result of the previous request to InferICD10CM was truncated, include
	// the PaginationToken to fetch the next page of medical condition entities.
	PaginationToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s InferICD10CMOutput) String() string {
	return awsutil.Prettify(s)
}

const opInferICD10CM = "InferICD10CM"

// InferICD10CMRequest returns a request value for making API operation for
// AWS Comprehend Medical.
//
// InferICD10CM detects medical conditions as entities listed in a patient record
// and links those entities to normalized concept identifiers in the ICD-10-CM
// knowledge base from the Centers for Disease Control. Amazon Comprehend Medical
// only detects medical entities in English language texts.
//
//    // Example sending a request using InferICD10CMRequest.
//    req := client.InferICD10CMRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/InferICD10CM
func (c *Client) InferICD10CMRequest(input *InferICD10CMInput) InferICD10CMRequest {
	op := &aws.Operation{
		Name:       opInferICD10CM,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InferICD10CMInput{}
	}

	req := c.newRequest(op, input, &InferICD10CMOutput{})
	return InferICD10CMRequest{Request: req, Input: input, Copy: c.InferICD10CMRequest}
}

// InferICD10CMRequest is the request type for the
// InferICD10CM API operation.
type InferICD10CMRequest struct {
	*aws.Request
	Input *InferICD10CMInput
	Copy  func(*InferICD10CMInput) InferICD10CMRequest
}

// Send marshals and sends the InferICD10CM API request.
func (r InferICD10CMRequest) Send(ctx context.Context) (*InferICD10CMResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InferICD10CMResponse{
		InferICD10CMOutput: r.Request.Data.(*InferICD10CMOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InferICD10CMResponse is the response type for the
// InferICD10CM API operation.
type InferICD10CMResponse struct {
	*InferICD10CMOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InferICD10CM request.
func (r *InferICD10CMResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
