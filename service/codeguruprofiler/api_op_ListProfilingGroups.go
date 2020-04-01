// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The structure representing the listProfilingGroupsRequest.
type ListProfilingGroupsInput struct {
	_ struct{} `type:"structure"`

	// A Boolean value indicating whether to include a description.
	IncludeDescription *bool `location:"querystring" locationName:"includeDescription" type:"boolean"`

	// The maximum number of profiling groups results returned by ListProfilingGroups
	// in paginated output. When this parameter is used, ListProfilingGroups only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListProfilingGroups request with the returned nextToken value.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The nextToken value returned from a previous paginated ListProfilingGroups
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `location:"querystring" locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListProfilingGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProfilingGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProfilingGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProfilingGroupsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.IncludeDescription != nil {
		v := *s.IncludeDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeDescription", protocol.BoolValue(v), metadata)
	}
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

// The structure representing the listProfilingGroupsResponse.
type ListProfilingGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListProfilingGroups request. When
	// the results of a ListProfilingGroups request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// Information about profiling group names.
	//
	// ProfilingGroupNames is a required field
	ProfilingGroupNames []string `locationName:"profilingGroupNames" type:"list" required:"true"`

	// Information about profiling groups.
	ProfilingGroups []ProfilingGroupDescription `locationName:"profilingGroups" type:"list"`
}

// String returns the string representation
func (s ListProfilingGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProfilingGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProfilingGroupNames != nil {
		v := s.ProfilingGroupNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "profilingGroupNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ProfilingGroups != nil {
		v := s.ProfilingGroups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "profilingGroups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListProfilingGroups = "ListProfilingGroups"

// ListProfilingGroupsRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
// Lists profiling groups.
//
//    // Example sending a request using ListProfilingGroupsRequest.
//    req := client.ListProfilingGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/ListProfilingGroups
func (c *Client) ListProfilingGroupsRequest(input *ListProfilingGroupsInput) ListProfilingGroupsRequest {
	op := &aws.Operation{
		Name:       opListProfilingGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/profilingGroups",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListProfilingGroupsInput{}
	}

	req := c.newRequest(op, input, &ListProfilingGroupsOutput{})
	return ListProfilingGroupsRequest{Request: req, Input: input, Copy: c.ListProfilingGroupsRequest}
}

// ListProfilingGroupsRequest is the request type for the
// ListProfilingGroups API operation.
type ListProfilingGroupsRequest struct {
	*aws.Request
	Input *ListProfilingGroupsInput
	Copy  func(*ListProfilingGroupsInput) ListProfilingGroupsRequest
}

// Send marshals and sends the ListProfilingGroups API request.
func (r ListProfilingGroupsRequest) Send(ctx context.Context) (*ListProfilingGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProfilingGroupsResponse{
		ListProfilingGroupsOutput: r.Request.Data.(*ListProfilingGroupsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProfilingGroupsRequestPaginator returns a paginator for ListProfilingGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProfilingGroupsRequest(input)
//   p := codeguruprofiler.NewListProfilingGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProfilingGroupsPaginator(req ListProfilingGroupsRequest) ListProfilingGroupsPaginator {
	return ListProfilingGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListProfilingGroupsInput
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

// ListProfilingGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProfilingGroupsPaginator struct {
	aws.Pager
}

func (p *ListProfilingGroupsPaginator) CurrentPage() *ListProfilingGroupsOutput {
	return p.Pager.CurrentPage().(*ListProfilingGroupsOutput)
}

// ListProfilingGroupsResponse is the response type for the
// ListProfilingGroups API operation.
type ListProfilingGroupsResponse struct {
	*ListProfilingGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProfilingGroups request.
func (r *ListProfilingGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
