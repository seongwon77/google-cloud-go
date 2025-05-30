// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package recommender

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListInsights                []gax.CallOption
	GetInsight                  []gax.CallOption
	MarkInsightAccepted         []gax.CallOption
	ListRecommendations         []gax.CallOption
	GetRecommendation           []gax.CallOption
	MarkRecommendationClaimed   []gax.CallOption
	MarkRecommendationSucceeded []gax.CallOption
	MarkRecommendationFailed    []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("recommender.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("recommender.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://recommender.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListInsights: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetInsight: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		MarkInsightAccepted: []gax.CallOption{},
		ListRecommendations: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetRecommendation: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		MarkRecommendationClaimed:   []gax.CallOption{},
		MarkRecommendationSucceeded: []gax.CallOption{},
		MarkRecommendationFailed:    []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Recommender API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListInsights(context.Context, *recommenderpb.ListInsightsRequest, ...gax.CallOption) *InsightIterator
	GetInsight(context.Context, *recommenderpb.GetInsightRequest, ...gax.CallOption) (*recommenderpb.Insight, error)
	MarkInsightAccepted(context.Context, *recommenderpb.MarkInsightAcceptedRequest, ...gax.CallOption) (*recommenderpb.Insight, error)
	ListRecommendations(context.Context, *recommenderpb.ListRecommendationsRequest, ...gax.CallOption) *RecommendationIterator
	GetRecommendation(context.Context, *recommenderpb.GetRecommendationRequest, ...gax.CallOption) (*recommenderpb.Recommendation, error)
	MarkRecommendationClaimed(context.Context, *recommenderpb.MarkRecommendationClaimedRequest, ...gax.CallOption) (*recommenderpb.Recommendation, error)
	MarkRecommendationSucceeded(context.Context, *recommenderpb.MarkRecommendationSucceededRequest, ...gax.CallOption) (*recommenderpb.Recommendation, error)
	MarkRecommendationFailed(context.Context, *recommenderpb.MarkRecommendationFailedRequest, ...gax.CallOption) (*recommenderpb.Recommendation, error)
}

// Client is a client for interacting with Recommender API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides insights and recommendations for cloud customers for various
// categories like performance optimization, cost savings, reliability, feature
// discovery, etc. Insights and recommendations are generated automatically
// based on analysis of user resources, configuration and monitoring metrics.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListInsights lists insights for the specified Cloud Resource. Requires the
// recommender.*.list IAM permission for the specified insight type.
func (c *Client) ListInsights(ctx context.Context, req *recommenderpb.ListInsightsRequest, opts ...gax.CallOption) *InsightIterator {
	return c.internalClient.ListInsights(ctx, req, opts...)
}

// GetInsight gets the requested insight. Requires the recommender.*.get IAM permission
// for the specified insight type.
func (c *Client) GetInsight(ctx context.Context, req *recommenderpb.GetInsightRequest, opts ...gax.CallOption) (*recommenderpb.Insight, error) {
	return c.internalClient.GetInsight(ctx, req, opts...)
}

// MarkInsightAccepted marks the Insight State as Accepted. Users can use this method to
// indicate to the Recommender API that they have applied some action based
// on the insight. This stops the insight content from being updated.
//
// MarkInsightAccepted can be applied to insights in ACTIVE state. Requires
// the recommender.*.update IAM permission for the specified insight.
func (c *Client) MarkInsightAccepted(ctx context.Context, req *recommenderpb.MarkInsightAcceptedRequest, opts ...gax.CallOption) (*recommenderpb.Insight, error) {
	return c.internalClient.MarkInsightAccepted(ctx, req, opts...)
}

// ListRecommendations lists recommendations for the specified Cloud Resource. Requires the
// recommender.*.list IAM permission for the specified recommender.
func (c *Client) ListRecommendations(ctx context.Context, req *recommenderpb.ListRecommendationsRequest, opts ...gax.CallOption) *RecommendationIterator {
	return c.internalClient.ListRecommendations(ctx, req, opts...)
}

// GetRecommendation gets the requested recommendation. Requires the recommender.*.get
// IAM permission for the specified recommender.
func (c *Client) GetRecommendation(ctx context.Context, req *recommenderpb.GetRecommendationRequest, opts ...gax.CallOption) (*recommenderpb.Recommendation, error) {
	return c.internalClient.GetRecommendation(ctx, req, opts...)
}

// MarkRecommendationClaimed marks the Recommendation State as Claimed. Users can use this method to
// indicate to the Recommender API that they are starting to apply the
// recommendation themselves. This stops the recommendation content from being
// updated. Associated insights are frozen and placed in the ACCEPTED state.
//
// MarkRecommendationClaimed can be applied to recommendations in CLAIMED,
// SUCCEEDED, FAILED, or ACTIVE state.
//
// Requires the recommender.*.update IAM permission for the specified
// recommender.
func (c *Client) MarkRecommendationClaimed(ctx context.Context, req *recommenderpb.MarkRecommendationClaimedRequest, opts ...gax.CallOption) (*recommenderpb.Recommendation, error) {
	return c.internalClient.MarkRecommendationClaimed(ctx, req, opts...)
}

// MarkRecommendationSucceeded marks the Recommendation State as Succeeded. Users can use this method to
// indicate to the Recommender API that they have applied the recommendation
// themselves, and the operation was successful. This stops the recommendation
// content from being updated. Associated insights are frozen and placed in
// the ACCEPTED state.
//
// MarkRecommendationSucceeded can be applied to recommendations in ACTIVE,
// CLAIMED, SUCCEEDED, or FAILED state.
//
// Requires the recommender.*.update IAM permission for the specified
// recommender.
func (c *Client) MarkRecommendationSucceeded(ctx context.Context, req *recommenderpb.MarkRecommendationSucceededRequest, opts ...gax.CallOption) (*recommenderpb.Recommendation, error) {
	return c.internalClient.MarkRecommendationSucceeded(ctx, req, opts...)
}

// MarkRecommendationFailed marks the Recommendation State as Failed. Users can use this method to
// indicate to the Recommender API that they have applied the recommendation
// themselves, and the operation failed. This stops the recommendation content
// from being updated. Associated insights are frozen and placed in the
// ACCEPTED state.
//
// MarkRecommendationFailed can be applied to recommendations in ACTIVE,
// CLAIMED, SUCCEEDED, or FAILED state.
//
// Requires the recommender.*.update IAM permission for the specified
// recommender.
func (c *Client) MarkRecommendationFailed(ctx context.Context, req *recommenderpb.MarkRecommendationFailedRequest, opts ...gax.CallOption) (*recommenderpb.Recommendation, error) {
	return c.internalClient.MarkRecommendationFailed(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Recommender API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client recommenderpb.RecommenderClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new recommender client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides insights and recommendations for cloud customers for various
// categories like performance optimization, cost savings, reliability, feature
// discovery, etc. Insights and recommendations are generated automatically
// based on analysis of user resources, configuration and monitoring metrics.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           recommenderpb.NewRecommenderClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) ListInsights(ctx context.Context, req *recommenderpb.ListInsightsRequest, opts ...gax.CallOption) *InsightIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListInsights[0:len((*c.CallOptions).ListInsights):len((*c.CallOptions).ListInsights)], opts...)
	it := &InsightIterator{}
	req = proto.Clone(req).(*recommenderpb.ListInsightsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recommenderpb.Insight, string, error) {
		resp := &recommenderpb.ListInsightsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListInsights(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetInsights(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetInsight(ctx context.Context, req *recommenderpb.GetInsightRequest, opts ...gax.CallOption) (*recommenderpb.Insight, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetInsight[0:len((*c.CallOptions).GetInsight):len((*c.CallOptions).GetInsight)], opts...)
	var resp *recommenderpb.Insight
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetInsight(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) MarkInsightAccepted(ctx context.Context, req *recommenderpb.MarkInsightAcceptedRequest, opts ...gax.CallOption) (*recommenderpb.Insight, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MarkInsightAccepted[0:len((*c.CallOptions).MarkInsightAccepted):len((*c.CallOptions).MarkInsightAccepted)], opts...)
	var resp *recommenderpb.Insight
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.MarkInsightAccepted(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListRecommendations(ctx context.Context, req *recommenderpb.ListRecommendationsRequest, opts ...gax.CallOption) *RecommendationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListRecommendations[0:len((*c.CallOptions).ListRecommendations):len((*c.CallOptions).ListRecommendations)], opts...)
	it := &RecommendationIterator{}
	req = proto.Clone(req).(*recommenderpb.ListRecommendationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recommenderpb.Recommendation, string, error) {
		resp := &recommenderpb.ListRecommendationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListRecommendations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetRecommendations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetRecommendation(ctx context.Context, req *recommenderpb.GetRecommendationRequest, opts ...gax.CallOption) (*recommenderpb.Recommendation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetRecommendation[0:len((*c.CallOptions).GetRecommendation):len((*c.CallOptions).GetRecommendation)], opts...)
	var resp *recommenderpb.Recommendation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetRecommendation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) MarkRecommendationClaimed(ctx context.Context, req *recommenderpb.MarkRecommendationClaimedRequest, opts ...gax.CallOption) (*recommenderpb.Recommendation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MarkRecommendationClaimed[0:len((*c.CallOptions).MarkRecommendationClaimed):len((*c.CallOptions).MarkRecommendationClaimed)], opts...)
	var resp *recommenderpb.Recommendation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.MarkRecommendationClaimed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) MarkRecommendationSucceeded(ctx context.Context, req *recommenderpb.MarkRecommendationSucceededRequest, opts ...gax.CallOption) (*recommenderpb.Recommendation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MarkRecommendationSucceeded[0:len((*c.CallOptions).MarkRecommendationSucceeded):len((*c.CallOptions).MarkRecommendationSucceeded)], opts...)
	var resp *recommenderpb.Recommendation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.MarkRecommendationSucceeded(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) MarkRecommendationFailed(ctx context.Context, req *recommenderpb.MarkRecommendationFailedRequest, opts ...gax.CallOption) (*recommenderpb.Recommendation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MarkRecommendationFailed[0:len((*c.CallOptions).MarkRecommendationFailed):len((*c.CallOptions).MarkRecommendationFailed)], opts...)
	var resp *recommenderpb.Recommendation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.MarkRecommendationFailed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// InsightIterator manages a stream of *recommenderpb.Insight.
type InsightIterator struct {
	items    []*recommenderpb.Insight
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*recommenderpb.Insight, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InsightIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InsightIterator) Next() (*recommenderpb.Insight, error) {
	var item *recommenderpb.Insight
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InsightIterator) bufLen() int {
	return len(it.items)
}

func (it *InsightIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RecommendationIterator manages a stream of *recommenderpb.Recommendation.
type RecommendationIterator struct {
	items    []*recommenderpb.Recommendation
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*recommenderpb.Recommendation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *RecommendationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RecommendationIterator) Next() (*recommenderpb.Recommendation, error) {
	var item *recommenderpb.Recommendation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RecommendationIterator) bufLen() int {
	return len(it.items)
}

func (it *RecommendationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
