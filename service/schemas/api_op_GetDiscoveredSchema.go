// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetDiscoveredSchemaInput struct {
	_ struct{} `type:"structure"`

	// Events is a required field
	Events []string `min:"1" type:"list" required:"true"`

	// Type is a required field
	Type Type `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetDiscoveredSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDiscoveredSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDiscoveredSchemaInput"}

	if s.Events == nil {
		invalidParams.Add(aws.NewErrParamRequired("Events"))
	}
	if s.Events != nil && len(s.Events) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Events", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDiscoveredSchemaInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Events != nil {
		v := s.Events

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Events", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type GetDiscoveredSchemaOutput struct {
	_ struct{} `type:"structure"`

	Content *string `type:"string"`
}

// String returns the string representation
func (s GetDiscoveredSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDiscoveredSchemaOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Content != nil {
		v := *s.Content

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Content", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDiscoveredSchema = "GetDiscoveredSchema"

// GetDiscoveredSchemaRequest returns a request value for making API operation for
// Schemas.
//
// Get the discovered schema that was generated based on sampled events.
//
//    // Example sending a request using GetDiscoveredSchemaRequest.
//    req := client.GetDiscoveredSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/GetDiscoveredSchema
func (c *Client) GetDiscoveredSchemaRequest(input *GetDiscoveredSchemaInput) GetDiscoveredSchemaRequest {
	op := &aws.Operation{
		Name:       opGetDiscoveredSchema,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/discover",
	}

	if input == nil {
		input = &GetDiscoveredSchemaInput{}
	}

	req := c.newRequest(op, input, &GetDiscoveredSchemaOutput{})
	return GetDiscoveredSchemaRequest{Request: req, Input: input, Copy: c.GetDiscoveredSchemaRequest}
}

// GetDiscoveredSchemaRequest is the request type for the
// GetDiscoveredSchema API operation.
type GetDiscoveredSchemaRequest struct {
	*aws.Request
	Input *GetDiscoveredSchemaInput
	Copy  func(*GetDiscoveredSchemaInput) GetDiscoveredSchemaRequest
}

// Send marshals and sends the GetDiscoveredSchema API request.
func (r GetDiscoveredSchemaRequest) Send(ctx context.Context) (*GetDiscoveredSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDiscoveredSchemaResponse{
		GetDiscoveredSchemaOutput: r.Request.Data.(*GetDiscoveredSchemaOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDiscoveredSchemaResponse is the response type for the
// GetDiscoveredSchema API operation.
type GetDiscoveredSchemaResponse struct {
	*GetDiscoveredSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDiscoveredSchema request.
func (r *GetDiscoveredSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
