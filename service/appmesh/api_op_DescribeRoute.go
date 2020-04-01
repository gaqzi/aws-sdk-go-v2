// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeRouteInput struct {
	_ struct{} `type:"structure"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	MeshOwner *string `location:"querystring" locationName:"meshOwner" min:"12" type:"string"`

	// RouteName is a required field
	RouteName *string `location:"uri" locationName:"routeName" min:"1" type:"string" required:"true"`

	// VirtualRouterName is a required field
	VirtualRouterName *string `location:"uri" locationName:"virtualRouterName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRouteInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}
	if s.MeshOwner != nil && len(*s.MeshOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshOwner", 12))
	}

	if s.RouteName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteName"))
	}
	if s.RouteName != nil && len(*s.RouteName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RouteName", 1))
	}

	if s.VirtualRouterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualRouterName"))
	}
	if s.VirtualRouterName != nil && len(*s.VirtualRouterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualRouterName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRouteInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteName != nil {
		v := *s.RouteName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualRouterName != nil {
		v := *s.VirtualRouterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualRouterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeshOwner != nil {
		v := *s.MeshOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "meshOwner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeRouteOutput struct {
	_ struct{} `type:"structure" payload:"Route"`

	// An object that represents a route returned by a describe operation.
	//
	// Route is a required field
	Route *RouteData `locationName:"route" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeRouteOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeRouteOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Route != nil {
		v := s.Route

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "route", v, metadata)
	}
	return nil
}

const opDescribeRoute = "DescribeRoute"

// DescribeRouteRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Describes an existing route.
//
//    // Example sending a request using DescribeRouteRequest.
//    req := client.DescribeRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DescribeRoute
func (c *Client) DescribeRouteRequest(input *DescribeRouteInput) DescribeRouteRequest {
	op := &aws.Operation{
		Name:       opDescribeRoute,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualRouter/{virtualRouterName}/routes/{routeName}",
	}

	if input == nil {
		input = &DescribeRouteInput{}
	}

	req := c.newRequest(op, input, &DescribeRouteOutput{})
	return DescribeRouteRequest{Request: req, Input: input, Copy: c.DescribeRouteRequest}
}

// DescribeRouteRequest is the request type for the
// DescribeRoute API operation.
type DescribeRouteRequest struct {
	*aws.Request
	Input *DescribeRouteInput
	Copy  func(*DescribeRouteInput) DescribeRouteRequest
}

// Send marshals and sends the DescribeRoute API request.
func (r DescribeRouteRequest) Send(ctx context.Context) (*DescribeRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRouteResponse{
		DescribeRouteOutput: r.Request.Data.(*DescribeRouteOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRouteResponse is the response type for the
// DescribeRoute API operation.
type DescribeRouteResponse struct {
	*DescribeRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRoute request.
func (r *DescribeRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
