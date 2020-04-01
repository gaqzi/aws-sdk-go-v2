// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutJobTaggingInput struct {
	_ struct{} `locationName:"PutJobTaggingRequest" type:"structure" xmlURI:"http://awss3control.amazonaws.com/doc/2018-08-20/"`

	// The account ID for the Amazon Web Services account associated with the Amazon
	// S3 batch operations job you want to replace tags on.
	//
	// AccountId is a required field
	AccountId *string `location:"header" locationName:"x-amz-account-id" type:"string" required:"true"`

	// The ID for the job whose tags you want to replace.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"id" min:"5" type:"string" required:"true"`

	// The set of tags to associate with the job.
	//
	// Tags is a required field
	Tags []S3Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s PutJobTaggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutJobTaggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutJobTaggingInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 5))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutJobTaggingInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "PutJobTaggingRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
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
	}), protocol.Metadata{XMLNamespaceURI: "http://awss3control.amazonaws.com/doc/2018-08-20/"})
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

type PutJobTaggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutJobTaggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutJobTaggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutJobTagging = "PutJobTagging"

// PutJobTaggingRequest returns a request value for making API operation for
// AWS S3 Control.
//
// Replace the set of tags on a Amazon S3 batch operations job.
//
//    // Example sending a request using PutJobTaggingRequest.
//    req := client.PutJobTaggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3control-2018-08-20/PutJobTagging
func (c *Client) PutJobTaggingRequest(input *PutJobTaggingInput) PutJobTaggingRequest {
	op := &aws.Operation{
		Name:       opPutJobTagging,
		HTTPMethod: "PUT",
		HTTPPath:   "/v20180820/jobs/{id}/tagging",
	}

	if input == nil {
		input = &PutJobTaggingInput{}
	}

	req := c.newRequest(op, input, &PutJobTaggingOutput{})
	req.Handlers.Build.PushBackNamed(buildPrefixHostHandler("AccountID", aws.StringValue(input.AccountId)))
	req.Handlers.Build.PushBackNamed(buildRemoveHeaderHandler("X-Amz-Account-Id"))
	return PutJobTaggingRequest{Request: req, Input: input, Copy: c.PutJobTaggingRequest}
}

// PutJobTaggingRequest is the request type for the
// PutJobTagging API operation.
type PutJobTaggingRequest struct {
	*aws.Request
	Input *PutJobTaggingInput
	Copy  func(*PutJobTaggingInput) PutJobTaggingRequest
}

// Send marshals and sends the PutJobTagging API request.
func (r PutJobTaggingRequest) Send(ctx context.Context) (*PutJobTaggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutJobTaggingResponse{
		PutJobTaggingOutput: r.Request.Data.(*PutJobTaggingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutJobTaggingResponse is the response type for the
// PutJobTagging API operation.
type PutJobTaggingResponse struct {
	*PutJobTaggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutJobTagging request.
func (r *PutJobTaggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
