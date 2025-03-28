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

package asset

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1p2beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateFeed []gax.CallOption
	GetFeed    []gax.CallOption
	ListFeeds  []gax.CallOption
	UpdateFeed []gax.CallOption
	DeleteFeed []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudasset.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudasset.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudasset.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateFeed: []gax.CallOption{},
		GetFeed: []gax.CallOption{
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
		ListFeeds: []gax.CallOption{
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
		UpdateFeed: []gax.CallOption{},
		DeleteFeed: []gax.CallOption{
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
	}
}

// internalClient is an interface that defines the methods available from Cloud Asset API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateFeed(context.Context, *assetpb.CreateFeedRequest, ...gax.CallOption) (*assetpb.Feed, error)
	GetFeed(context.Context, *assetpb.GetFeedRequest, ...gax.CallOption) (*assetpb.Feed, error)
	ListFeeds(context.Context, *assetpb.ListFeedsRequest, ...gax.CallOption) (*assetpb.ListFeedsResponse, error)
	UpdateFeed(context.Context, *assetpb.UpdateFeedRequest, ...gax.CallOption) (*assetpb.Feed, error)
	DeleteFeed(context.Context, *assetpb.DeleteFeedRequest, ...gax.CallOption) error
}

// Client is a client for interacting with Cloud Asset API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Asset service definition.
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

// CreateFeed creates a feed in a parent project/folder/organization to listen to its
// asset updates.
func (c *Client) CreateFeed(ctx context.Context, req *assetpb.CreateFeedRequest, opts ...gax.CallOption) (*assetpb.Feed, error) {
	return c.internalClient.CreateFeed(ctx, req, opts...)
}

// GetFeed gets details about an asset feed.
func (c *Client) GetFeed(ctx context.Context, req *assetpb.GetFeedRequest, opts ...gax.CallOption) (*assetpb.Feed, error) {
	return c.internalClient.GetFeed(ctx, req, opts...)
}

// ListFeeds lists all asset feeds in a parent project/folder/organization.
func (c *Client) ListFeeds(ctx context.Context, req *assetpb.ListFeedsRequest, opts ...gax.CallOption) (*assetpb.ListFeedsResponse, error) {
	return c.internalClient.ListFeeds(ctx, req, opts...)
}

// UpdateFeed updates an asset feed configuration.
func (c *Client) UpdateFeed(ctx context.Context, req *assetpb.UpdateFeedRequest, opts ...gax.CallOption) (*assetpb.Feed, error) {
	return c.internalClient.UpdateFeed(ctx, req, opts...)
}

// DeleteFeed deletes an asset feed.
func (c *Client) DeleteFeed(ctx context.Context, req *assetpb.DeleteFeedRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteFeed(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Cloud Asset API over gRPC transport.
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
	client assetpb.AssetServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new asset service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Asset service definition.
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
		client:           assetpb.NewAssetServiceClient(connPool),
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

func (c *gRPCClient) CreateFeed(ctx context.Context, req *assetpb.CreateFeedRequest, opts ...gax.CallOption) (*assetpb.Feed, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateFeed[0:len((*c.CallOptions).CreateFeed):len((*c.CallOptions).CreateFeed)], opts...)
	var resp *assetpb.Feed
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateFeed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetFeed(ctx context.Context, req *assetpb.GetFeedRequest, opts ...gax.CallOption) (*assetpb.Feed, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFeed[0:len((*c.CallOptions).GetFeed):len((*c.CallOptions).GetFeed)], opts...)
	var resp *assetpb.Feed
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetFeed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListFeeds(ctx context.Context, req *assetpb.ListFeedsRequest, opts ...gax.CallOption) (*assetpb.ListFeedsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListFeeds[0:len((*c.CallOptions).ListFeeds):len((*c.CallOptions).ListFeeds)], opts...)
	var resp *assetpb.ListFeedsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ListFeeds(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateFeed(ctx context.Context, req *assetpb.UpdateFeedRequest, opts ...gax.CallOption) (*assetpb.Feed, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "feed.name", url.QueryEscape(req.GetFeed().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateFeed[0:len((*c.CallOptions).UpdateFeed):len((*c.CallOptions).UpdateFeed)], opts...)
	var resp *assetpb.Feed
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateFeed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteFeed(ctx context.Context, req *assetpb.DeleteFeedRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteFeed[0:len((*c.CallOptions).DeleteFeed):len((*c.CallOptions).DeleteFeed)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteFeed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}
