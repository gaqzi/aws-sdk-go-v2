// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetJobTaggingInput struct {
	_ struct{} `type:"structure"`

	// The account ID for the Amazon Web Services account associated with the Amazon
	// S3 batch operations job you want to retrieve tags for.
	//
	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`

	// The ID for the job whose tags you want to retrieve.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"id" min:"5" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJobTaggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJobTaggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJobTaggingInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobTaggingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-account-id", protocol.StringValue(v), metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.StringValue(v), metadata)
	}
	return nil
}

type GetJobTaggingOutput struct {
	_ struct{} `type:"structure"`

	// The set of tags associated with the job.
	Tags []S3Tag `type:"list"`
}

// String returns the string representation
func (s GetJobTaggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobTaggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetJobTagging = "GetJobTagging"

// GetJobTaggingRequest returns a request value for making API operation for
// AWS S3 Control.
//
// Retrieve the tags on a Amazon S3 batch operations job.
//
//    // Example sending a request using GetJobTaggingRequest.
//    req := client.GetJobTaggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/GetJobTagging
func (c *Client) GetJobTaggingRequest(input *GetJobTaggingInput) GetJobTaggingRequest {
	op := &aws.Operation{
		Name:       opGetJobTagging,
		HTTPMethod: "GET",
		HTTPPath:   "/v20180820/jobs/{id}/tagging",
	}

	if input == nil {
		input = &GetJobTaggingInput{}
	}

	req := c.newRequest(op, input, &GetJobTaggingOutput{})
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))
	return GetJobTaggingRequest{Request: req, Input: input, Copy: c.GetJobTaggingRequest}
}

// GetJobTaggingRequest is the request type for the
// GetJobTagging API operation.
type GetJobTaggingRequest struct {
	*aws.Request
	Input *GetJobTaggingInput
	Copy  func(*GetJobTaggingInput) GetJobTaggingRequest
}

// Send marshals and sends the GetJobTagging API request.
func (r GetJobTaggingRequest) Send(ctx context.Context) (*GetJobTaggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobTaggingResponse{
		GetJobTaggingOutput: r.Request.Data.(*GetJobTaggingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobTaggingResponse is the response type for the
// GetJobTagging API operation.
type GetJobTaggingResponse struct {
	*GetJobTaggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJobTagging request.
func (r *GetJobTaggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
