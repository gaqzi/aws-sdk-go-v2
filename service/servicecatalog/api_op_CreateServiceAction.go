// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateServiceActionInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The self-service action definition. Can be one of the following:
	//
	// Name
	//
	// The name of the AWS Systems Manager Document. For example, AWS-RestartEC2Instance.
	//
	// Version
	//
	// The AWS Systems Manager automation document version. For example, "Version":
	// "1"
	//
	// AssumeRole
	//
	// The Amazon Resource Name (ARN) of the role that performs the self-service
	// actions on your behalf. For example, "AssumeRole": "arn:aws:iam::12345678910:role/ActionRole".
	//
	// To reuse the provisioned product launch role, set to "AssumeRole": "LAUNCH_ROLE".
	//
	// Parameters
	//
	// The list of parameters in JSON format.
	//
	// For example: [{\"Name\":\"InstanceId\",\"Type\":\"TARGET\"}].
	//
	// Definition is a required field
	Definition map[string]string `min:"1" type:"map" required:"true"`

	// The service action definition type. For example, SSM_AUTOMATION.
	//
	// DefinitionType is a required field
	DefinitionType ServiceActionDefinitionType `type:"string" required:"true" enum:"true"`

	// The self-service action description.
	Description *string `type:"string"`

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The self-service action name.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateServiceActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServiceActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateServiceActionInput"}

	if s.Definition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Definition"))
	}
	if s.Definition != nil && len(s.Definition) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Definition", 1))
	}
	if len(s.DefinitionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DefinitionType"))
	}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateServiceActionOutput struct {
	_ struct{} `type:"structure"`

	// An object containing information about the self-service action.
	ServiceActionDetail *ServiceActionDetail `type:"structure"`
}

// String returns the string representation
func (s CreateServiceActionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateServiceAction = "CreateServiceAction"

// CreateServiceActionRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Creates a self-service action.
//
//    // Example sending a request using CreateServiceActionRequest.
//    req := client.CreateServiceActionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateServiceAction
func (c *Client) CreateServiceActionRequest(input *CreateServiceActionInput) CreateServiceActionRequest {
	op := &aws.Operation{
		Name:       opCreateServiceAction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServiceActionInput{}
	}

	req := c.newRequest(op, input, &CreateServiceActionOutput{})
	return CreateServiceActionRequest{Request: req, Input: input, Copy: c.CreateServiceActionRequest}
}

// CreateServiceActionRequest is the request type for the
// CreateServiceAction API operation.
type CreateServiceActionRequest struct {
	*aws.Request
	Input *CreateServiceActionInput
	Copy  func(*CreateServiceActionInput) CreateServiceActionRequest
}

// Send marshals and sends the CreateServiceAction API request.
func (r CreateServiceActionRequest) Send(ctx context.Context) (*CreateServiceActionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateServiceActionResponse{
		CreateServiceActionOutput: r.Request.Data.(*CreateServiceActionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateServiceActionResponse is the response type for the
// CreateServiceAction API operation.
type CreateServiceActionResponse struct {
	*CreateServiceActionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateServiceAction request.
func (r *CreateServiceActionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
