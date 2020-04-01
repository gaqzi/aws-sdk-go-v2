// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListDataflowEndpointGroupsInput struct {
	_ struct{} `type:"structure"`

	// Maximum number of dataflow endpoint groups returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// Next token returned in the request of a previous ListDataflowEndpointGroups
	// call. Used to get the next page of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDataflowEndpointGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDataflowEndpointGroupsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListDataflowEndpointGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of dataflow endpoint groups.
	DataflowEndpointGroupList []DataflowEndpointListItem `locationName:"dataflowEndpointGroupList" type:"list"`

	// Next token returned in the response of a previous ListDataflowEndpointGroups
	// call. Used to get the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDataflowEndpointGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDataflowEndpointGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DataflowEndpointGroupList != nil {
		v := s.DataflowEndpointGroupList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "dataflowEndpointGroupList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListDataflowEndpointGroups = "ListDataflowEndpointGroups"

// ListDataflowEndpointGroupsRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns a list of DataflowEndpoint groups.
//
//    // Example sending a request using ListDataflowEndpointGroupsRequest.
//    req := client.ListDataflowEndpointGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListDataflowEndpointGroups
func (c *Client) ListDataflowEndpointGroupsRequest(input *ListDataflowEndpointGroupsInput) ListDataflowEndpointGroupsRequest {
	op := &aws.Operation{
		Name:       opListDataflowEndpointGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/dataflowEndpointGroup",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDataflowEndpointGroupsInput{}
	}

	req := c.newRequest(op, input, &ListDataflowEndpointGroupsOutput{})
	return ListDataflowEndpointGroupsRequest{Request: req, Input: input, Copy: c.ListDataflowEndpointGroupsRequest}
}

// ListDataflowEndpointGroupsRequest is the request type for the
// ListDataflowEndpointGroups API operation.
type ListDataflowEndpointGroupsRequest struct {
	*aws.Request
	Input *ListDataflowEndpointGroupsInput
	Copy  func(*ListDataflowEndpointGroupsInput) ListDataflowEndpointGroupsRequest
}

// Send marshals and sends the ListDataflowEndpointGroups API request.
func (r ListDataflowEndpointGroupsRequest) Send(ctx context.Context) (*ListDataflowEndpointGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDataflowEndpointGroupsResponse{
		ListDataflowEndpointGroupsOutput: r.Request.Data.(*ListDataflowEndpointGroupsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDataflowEndpointGroupsRequestPaginator returns a paginator for ListDataflowEndpointGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDataflowEndpointGroupsRequest(input)
//   p := groundstation.NewListDataflowEndpointGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDataflowEndpointGroupsPaginator(req ListDataflowEndpointGroupsRequest) ListDataflowEndpointGroupsPaginator {
	return ListDataflowEndpointGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDataflowEndpointGroupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListDataflowEndpointGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDataflowEndpointGroupsPaginator struct {
	aws.Pager
}

func (p *ListDataflowEndpointGroupsPaginator) CurrentPage() *ListDataflowEndpointGroupsOutput {
	return p.Pager.CurrentPage().(*ListDataflowEndpointGroupsOutput)
}

// ListDataflowEndpointGroupsResponse is the response type for the
// ListDataflowEndpointGroups API operation.
type ListDataflowEndpointGroupsResponse struct {
	*ListDataflowEndpointGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDataflowEndpointGroups request.
func (r *ListDataflowEndpointGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
