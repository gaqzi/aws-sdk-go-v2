// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetVpcLinkInput struct {
	_ struct{} `type:"structure"`

	// VpcLinkId is a required field
	VpcLinkId *string `location:"uri" locationName:"vpcLinkId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVpcLinkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVpcLinkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVpcLinkInput"}

	if s.VpcLinkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcLinkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVpcLinkInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VpcLinkId != nil {
		v := *s.VpcLinkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vpcLinkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetVpcLinkOutput struct {
	_ struct{} `type:"structure"`

	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"iso8601"`

	// A string with a length between [1-128].
	Name *string `locationName:"name" type:"string"`

	// A list of security group IDs for the VPC link.
	SecurityGroupIds []string `locationName:"securityGroupIds" type:"list"`

	// A list of subnet IDs to include in the VPC link.
	SubnetIds []string `locationName:"subnetIds" type:"list"`

	// Represents a collection of tags associated with the resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The identifier.
	VpcLinkId *string `locationName:"vpcLinkId" type:"string"`

	// The status of the VPC link.
	VpcLinkStatus VpcLinkStatus `locationName:"vpcLinkStatus" type:"string" enum:"true"`

	// A string with a length between [0-1024].
	VpcLinkStatusMessage *string `locationName:"vpcLinkStatusMessage" type:"string"`

	// The version of the VPC link.
	VpcLinkVersion VpcLinkVersion `locationName:"vpcLinkVersion" type:"string" enum:"true"`
}

// String returns the string representation
func (s GetVpcLinkOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetVpcLinkOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityGroupIds != nil {
		v := s.SecurityGroupIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "securityGroupIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SubnetIds != nil {
		v := s.SubnetIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "subnetIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.VpcLinkId != nil {
		v := *s.VpcLinkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "vpcLinkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.VpcLinkStatus) > 0 {
		v := s.VpcLinkStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "vpcLinkStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.VpcLinkStatusMessage != nil {
		v := *s.VpcLinkStatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "vpcLinkStatusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.VpcLinkVersion) > 0 {
		v := s.VpcLinkVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "vpcLinkVersion", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opGetVpcLink = "GetVpcLink"

// GetVpcLinkRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets a VPC link.
//
//    // Example sending a request using GetVpcLinkRequest.
//    req := client.GetVpcLinkRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetVpcLink
func (c *Client) GetVpcLinkRequest(input *GetVpcLinkInput) GetVpcLinkRequest {
	op := &aws.Operation{
		Name:       opGetVpcLink,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/vpclinks/{vpcLinkId}",
	}

	if input == nil {
		input = &GetVpcLinkInput{}
	}

	req := c.newRequest(op, input, &GetVpcLinkOutput{})
	return GetVpcLinkRequest{Request: req, Input: input, Copy: c.GetVpcLinkRequest}
}

// GetVpcLinkRequest is the request type for the
// GetVpcLink API operation.
type GetVpcLinkRequest struct {
	*aws.Request
	Input *GetVpcLinkInput
	Copy  func(*GetVpcLinkInput) GetVpcLinkRequest
}

// Send marshals and sends the GetVpcLink API request.
func (r GetVpcLinkRequest) Send(ctx context.Context) (*GetVpcLinkResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVpcLinkResponse{
		GetVpcLinkOutput: r.Request.Data.(*GetVpcLinkOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVpcLinkResponse is the response type for the
// GetVpcLink API operation.
type GetVpcLinkResponse struct {
	*GetVpcLinkOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVpcLink request.
func (r *GetVpcLinkResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
