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

package resourcesettings

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
	resourcesettingspb "google.golang.org/genproto/googleapis/cloud/resourcesettings/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ListSettings  []gax.CallOption
	GetSetting    []gax.CallOption
	UpdateSetting []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("resourcesettings.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("resourcesettings.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://resourcesettings.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ListSettings: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetSetting: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateSetting: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from Resource Settings API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListSettings(context.Context, *resourcesettingspb.ListSettingsRequest, ...gax.CallOption) *SettingIterator
	GetSetting(context.Context, *resourcesettingspb.GetSettingRequest, ...gax.CallOption) (*resourcesettingspb.Setting, error)
	UpdateSetting(context.Context, *resourcesettingspb.UpdateSettingRequest, ...gax.CallOption) (*resourcesettingspb.Setting, error)
}

// Client is a client for interacting with Resource Settings API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// An interface to interact with resource settings and setting values throughout
// the resource hierarchy.
//
// Services may surface a number of settings for users to control how their
// resources behave. Values of settings applied on a given Cloud resource are
// evaluated hierarchically and inherited by all descendants of that resource.
//
// For all requests, returns a google.rpc.Status with
// google.rpc.Code.PERMISSION_DENIED if the IAM check fails or the parent
// resource is not in a Cloud Organization.
// For all requests, returns a google.rpc.Status with
// google.rpc.Code.INVALID_ARGUMENT if the request is malformed.
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

// ListSettings lists all the settings that are available on the Cloud resource parent.
func (c *Client) ListSettings(ctx context.Context, req *resourcesettingspb.ListSettingsRequest, opts ...gax.CallOption) *SettingIterator {
	return c.internalClient.ListSettings(ctx, req, opts...)
}

// GetSetting gets a setting.
//
// Returns a google.rpc.Status with google.rpc.Code.NOT_FOUND if the
// setting does not exist.
func (c *Client) GetSetting(ctx context.Context, req *resourcesettingspb.GetSettingRequest, opts ...gax.CallOption) (*resourcesettingspb.Setting, error) {
	return c.internalClient.GetSetting(ctx, req, opts...)
}

// UpdateSetting updates a setting.
//
// Returns a google.rpc.Status with google.rpc.Code.NOT_FOUND if the
// setting does not exist.
// Returns a google.rpc.Status with google.rpc.Code.FAILED_PRECONDITION if
// the setting is flagged as read only.
// Returns a google.rpc.Status with google.rpc.Code.ABORTED if the etag
// supplied in the request does not match the persisted etag of the setting
// value.
//
// On success, the response will contain only name, local_value and
// etag.  The metadata and effective_value cannot be updated through
// this API.
//
// Note: the supplied setting will perform a full overwrite of the
// local_value field.
func (c *Client) UpdateSetting(ctx context.Context, req *resourcesettingspb.UpdateSettingRequest, opts ...gax.CallOption) (*resourcesettingspb.Setting, error) {
	return c.internalClient.UpdateSetting(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Resource Settings API over gRPC transport.
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
	client resourcesettingspb.ResourceSettingsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new resource settings service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// An interface to interact with resource settings and setting values throughout
// the resource hierarchy.
//
// Services may surface a number of settings for users to control how their
// resources behave. Values of settings applied on a given Cloud resource are
// evaluated hierarchically and inherited by all descendants of that resource.
//
// For all requests, returns a google.rpc.Status with
// google.rpc.Code.PERMISSION_DENIED if the IAM check fails or the parent
// resource is not in a Cloud Organization.
// For all requests, returns a google.rpc.Status with
// google.rpc.Code.INVALID_ARGUMENT if the request is malformed.
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
		client:           resourcesettingspb.NewResourceSettingsServiceClient(connPool),
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

func (c *gRPCClient) ListSettings(ctx context.Context, req *resourcesettingspb.ListSettingsRequest, opts ...gax.CallOption) *SettingIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListSettings[0:len((*c.CallOptions).ListSettings):len((*c.CallOptions).ListSettings)], opts...)
	it := &SettingIterator{}
	req = proto.Clone(req).(*resourcesettingspb.ListSettingsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcesettingspb.Setting, string, error) {
		resp := &resourcesettingspb.ListSettingsResponse{}
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
			resp, err = c.client.ListSettings(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSettings(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetSetting(ctx context.Context, req *resourcesettingspb.GetSettingRequest, opts ...gax.CallOption) (*resourcesettingspb.Setting, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetSetting[0:len((*c.CallOptions).GetSetting):len((*c.CallOptions).GetSetting)], opts...)
	var resp *resourcesettingspb.Setting
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetSetting(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateSetting(ctx context.Context, req *resourcesettingspb.UpdateSettingRequest, opts ...gax.CallOption) (*resourcesettingspb.Setting, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "setting.name", url.QueryEscape(req.GetSetting().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateSetting[0:len((*c.CallOptions).UpdateSetting):len((*c.CallOptions).UpdateSetting)], opts...)
	var resp *resourcesettingspb.Setting
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateSetting(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SettingIterator manages a stream of *resourcesettingspb.Setting.
type SettingIterator struct {
	items    []*resourcesettingspb.Setting
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
	InternalFetch func(pageSize int, pageToken string) (results []*resourcesettingspb.Setting, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SettingIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SettingIterator) Next() (*resourcesettingspb.Setting, error) {
	var item *resourcesettingspb.Setting
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SettingIterator) bufLen() int {
	return len(it.items)
}

func (it *SettingIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
