// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kendra

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDataSourcesInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the index that contains the data source.
	//
	// IndexId is a required field
	IndexId *string `min:"36" type:"string" required:"true"`

	// The maximum number of data sources to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Kendra returns a pagination token in the response. You can use this
	// pagination token to retrieve the next set of data sources (DataSourceSummaryItems).
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDataSourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDataSourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDataSourcesInput"}

	if s.IndexId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IndexId"))
	}
	if s.IndexId != nil && len(*s.IndexId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexId", 36))
	}
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

type ListDataSourcesOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, Amazon Kendra returns this token that you can
	// use in the subsequent request to retrieve the next set of data sources.
	NextToken *string `min:"1" type:"string"`

	// An array of summary information for one or more data sources.
	SummaryItems []DataSourceSummary `type:"list"`
}

// String returns the string representation
func (s ListDataSourcesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDataSources = "ListDataSources"

// ListDataSourcesRequest returns a request value for making API operation for
// AWSKendraFrontendService.
//
// Lists the data sources that you have created.
//
//    // Example sending a request using ListDataSourcesRequest.
//    req := client.ListDataSourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kendra-2019-02-03/ListDataSources
func (c *Client) ListDataSourcesRequest(input *ListDataSourcesInput) ListDataSourcesRequest {
	op := &aws.Operation{
		Name:       opListDataSources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDataSourcesInput{}
	}

	req := c.newRequest(op, input, &ListDataSourcesOutput{})
	return ListDataSourcesRequest{Request: req, Input: input, Copy: c.ListDataSourcesRequest}
}

// ListDataSourcesRequest is the request type for the
// ListDataSources API operation.
type ListDataSourcesRequest struct {
	*aws.Request
	Input *ListDataSourcesInput
	Copy  func(*ListDataSourcesInput) ListDataSourcesRequest
}

// Send marshals and sends the ListDataSources API request.
func (r ListDataSourcesRequest) Send(ctx context.Context) (*ListDataSourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDataSourcesResponse{
		ListDataSourcesOutput: r.Request.Data.(*ListDataSourcesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDataSourcesRequestPaginator returns a paginator for ListDataSources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDataSourcesRequest(input)
//   p := kendra.NewListDataSourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDataSourcesPaginator(req ListDataSourcesRequest) ListDataSourcesPaginator {
	return ListDataSourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDataSourcesInput
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

// ListDataSourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDataSourcesPaginator struct {
	aws.Pager
}

func (p *ListDataSourcesPaginator) CurrentPage() *ListDataSourcesOutput {
	return p.Pager.CurrentPage().(*ListDataSourcesOutput)
}

// ListDataSourcesResponse is the response type for the
// ListDataSources API operation.
type ListDataSourcesResponse struct {
	*ListDataSourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDataSources request.
func (r *ListDataSourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
