// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Container for request parameters to DescribePackage operation.
type DescribePackagesInput struct {
	_ struct{} `type:"structure"`

	// Only returns packages that match the DescribePackagesFilterList values.
	Filters []DescribePackagesFilter `type:"list"`

	// Limits results to a maximum number of packages.
	MaxResults *int64 `type:"integer"`

	// Used for pagination. Only necessary if a previous API call includes a non-null
	// NextToken value. If provided, returns results for the next page.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribePackagesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePackagesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Container for response returned by DescribePackages operation.
type DescribePackagesOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `type:"string"`

	// List of PackageDetails objects.
	PackageDetailsList []PackageDetails `type:"list"`
}

// String returns the string representation
func (s DescribePackagesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribePackagesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PackageDetailsList != nil {
		v := s.PackageDetailsList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "PackageDetailsList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribePackages = "DescribePackages"

// DescribePackagesRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Describes all packages available to Amazon ES. Includes options for filtering,
// limiting the number of results, and pagination.
//
//    // Example sending a request using DescribePackagesRequest.
//    req := client.DescribePackagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribePackagesRequest(input *DescribePackagesInput) DescribePackagesRequest {
	op := &aws.Operation{
		Name:       opDescribePackages,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/packages/describe",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribePackagesInput{}
	}

	req := c.newRequest(op, input, &DescribePackagesOutput{})
	return DescribePackagesRequest{Request: req, Input: input, Copy: c.DescribePackagesRequest}
}

// DescribePackagesRequest is the request type for the
// DescribePackages API operation.
type DescribePackagesRequest struct {
	*aws.Request
	Input *DescribePackagesInput
	Copy  func(*DescribePackagesInput) DescribePackagesRequest
}

// Send marshals and sends the DescribePackages API request.
func (r DescribePackagesRequest) Send(ctx context.Context) (*DescribePackagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePackagesResponse{
		DescribePackagesOutput: r.Request.Data.(*DescribePackagesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribePackagesRequestPaginator returns a paginator for DescribePackages.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribePackagesRequest(input)
//   p := elasticsearchservice.NewDescribePackagesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribePackagesPaginator(req DescribePackagesRequest) DescribePackagesPaginator {
	return DescribePackagesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribePackagesInput
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

// DescribePackagesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribePackagesPaginator struct {
	aws.Pager
}

func (p *DescribePackagesPaginator) CurrentPage() *DescribePackagesOutput {
	return p.Pager.CurrentPage().(*DescribePackagesOutput)
}

// DescribePackagesResponse is the response type for the
// DescribePackages API operation.
type DescribePackagesResponse struct {
	*DescribePackagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePackages request.
func (r *DescribePackagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
