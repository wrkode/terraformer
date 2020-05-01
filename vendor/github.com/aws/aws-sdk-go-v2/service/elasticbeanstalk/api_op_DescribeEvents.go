// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to retrieve a list of events for an environment.
type DescribeEventsInput struct {
	_ struct{} `type:"structure"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those associated with this application.
	ApplicationName *string `min:"1" type:"string"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// those that occur up to, but not including, the EndTime.
	EndTime *time.Time `type:"timestamp"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// those associated with this environment.
	EnvironmentId *string `type:"string"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// those associated with this environment.
	EnvironmentName *string `min:"4" type:"string"`

	// Specifies the maximum number of events that can be returned, beginning with
	// the most recent event.
	MaxRecords *int64 `min:"1" type:"integer"`

	// Pagination token. If specified, the events return the next batch of results.
	NextToken *string `type:"string"`

	// The ARN of a custom platform version. If specified, AWS Elastic Beanstalk
	// restricts the returned descriptions to those associated with this custom
	// platform version.
	PlatformArn *string `type:"string"`

	// If specified, AWS Elastic Beanstalk restricts the described events to include
	// only those associated with this request ID.
	RequestId *string `type:"string"`

	// If specified, limits the events returned from this call to include only those
	// with the specified severity or higher.
	Severity EventSeverity `type:"string" enum:"true"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// those that occur on or after this time.
	StartTime *time.Time `type:"timestamp"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// those that are associated with this environment configuration.
	TemplateName *string `min:"1" type:"string"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// those associated with this application version.
	VersionLabel *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventsInput"}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.EnvironmentName != nil && len(*s.EnvironmentName) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("EnvironmentName", 4))
	}
	if s.MaxRecords != nil && *s.MaxRecords < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxRecords", 1))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateName", 1))
	}
	if s.VersionLabel != nil && len(*s.VersionLabel) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionLabel", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result message wrapping a list of event descriptions.
type DescribeEventsOutput struct {
	_ struct{} `type:"structure"`

	// A list of EventDescription.
	Events []EventDescription `type:"list"`

	// If returned, this indicates that there are more results to obtain. Use this
	// token in the next DescribeEvents call to get the next batch of events.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEvents = "DescribeEvents"

// DescribeEventsRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Returns list of event descriptions matching criteria up to the last 6 weeks.
//
// This action returns the most recent 1,000 events from the specified NextToken.
//
//    // Example sending a request using DescribeEventsRequest.
//    req := client.DescribeEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeEvents
func (c *Client) DescribeEventsRequest(input *DescribeEventsInput) DescribeEventsRequest {
	op := &aws.Operation{
		Name:       opDescribeEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEventsInput{}
	}

	req := c.newRequest(op, input, &DescribeEventsOutput{})
	return DescribeEventsRequest{Request: req, Input: input, Copy: c.DescribeEventsRequest}
}

// DescribeEventsRequest is the request type for the
// DescribeEvents API operation.
type DescribeEventsRequest struct {
	*aws.Request
	Input *DescribeEventsInput
	Copy  func(*DescribeEventsInput) DescribeEventsRequest
}

// Send marshals and sends the DescribeEvents API request.
func (r DescribeEventsRequest) Send(ctx context.Context) (*DescribeEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventsResponse{
		DescribeEventsOutput: r.Request.Data.(*DescribeEventsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEventsRequestPaginator returns a paginator for DescribeEvents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEventsRequest(input)
//   p := elasticbeanstalk.NewDescribeEventsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEventsPaginator(req DescribeEventsRequest) DescribeEventsPaginator {
	return DescribeEventsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEventsInput
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

// DescribeEventsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEventsPaginator struct {
	aws.Pager
}

func (p *DescribeEventsPaginator) CurrentPage() *DescribeEventsOutput {
	return p.Pager.CurrentPage().(*DescribeEventsOutput)
}

// DescribeEventsResponse is the response type for the
// DescribeEvents API operation.
type DescribeEventsResponse struct {
	*DescribeEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEvents request.
func (r *DescribeEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
