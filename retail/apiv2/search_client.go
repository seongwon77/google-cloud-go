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

package retail

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
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newSearchClientHook clientHook

// SearchCallOptions contains the retry settings for each method of SearchClient.
type SearchCallOptions struct {
	Search []gax.CallOption
}

func defaultSearchGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("retail.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("retail.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://retail.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSearchCallOptions() *SearchCallOptions {
	return &SearchCallOptions{
		Search: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        5000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalSearchClient is an interface that defines the methods available from Retail API.
type internalSearchClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Search(context.Context, *retailpb.SearchRequest, ...gax.CallOption) *SearchResponse_SearchResultIterator
}

// SearchClient is a client for interacting with Retail API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for search.
//
// This feature is only available for users who have Retail Search enabled.
// Please enable Retail Search on Cloud Console before using this feature.
type SearchClient struct {
	// The internal transport-dependent client.
	internalClient internalSearchClient

	// The call options for this service.
	CallOptions *SearchCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SearchClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SearchClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SearchClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Search performs a search.
//
// This feature is only available for users who have Retail Search enabled.
// Please enable Retail Search on Cloud Console before using this feature.
func (c *SearchClient) Search(ctx context.Context, req *retailpb.SearchRequest, opts ...gax.CallOption) *SearchResponse_SearchResultIterator {
	return c.internalClient.Search(ctx, req, opts...)
}

// searchGRPCClient is a client for interacting with Retail API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type searchGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SearchClient
	CallOptions **SearchCallOptions

	// The gRPC API client.
	searchClient retailpb.SearchServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSearchClient creates a new search service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for search.
//
// This feature is only available for users who have Retail Search enabled.
// Please enable Retail Search on Cloud Console before using this feature.
func NewSearchClient(ctx context.Context, opts ...option.ClientOption) (*SearchClient, error) {
	clientOpts := defaultSearchGRPCClientOptions()
	if newSearchClientHook != nil {
		hookOpts, err := newSearchClientHook(ctx, clientHookParams{})
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
	client := SearchClient{CallOptions: defaultSearchCallOptions()}

	c := &searchGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		searchClient:     retailpb.NewSearchServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *searchGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *searchGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *searchGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *searchGRPCClient) Search(ctx context.Context, req *retailpb.SearchRequest, opts ...gax.CallOption) *SearchResponse_SearchResultIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "placement", url.QueryEscape(req.GetPlacement())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).Search[0:len((*c.CallOptions).Search):len((*c.CallOptions).Search)], opts...)
	it := &SearchResponse_SearchResultIterator{}
	req = proto.Clone(req).(*retailpb.SearchRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*retailpb.SearchResponse_SearchResult, string, error) {
		resp := &retailpb.SearchResponse{}
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
			resp, err = c.searchClient.Search(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResults(), resp.GetNextPageToken(), nil
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

// SearchResponse_SearchResultIterator manages a stream of *retailpb.SearchResponse_SearchResult.
type SearchResponse_SearchResultIterator struct {
	items    []*retailpb.SearchResponse_SearchResult
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
	InternalFetch func(pageSize int, pageToken string) (results []*retailpb.SearchResponse_SearchResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SearchResponse_SearchResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SearchResponse_SearchResultIterator) Next() (*retailpb.SearchResponse_SearchResult, error) {
	var item *retailpb.SearchResponse_SearchResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SearchResponse_SearchResultIterator) bufLen() int {
	return len(it.items)
}

func (it *SearchResponse_SearchResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
