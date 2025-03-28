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

package containeranalysis

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
	containeranalysispb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newContainerAnalysisV1Beta1ClientHook clientHook

// ContainerAnalysisV1Beta1CallOptions contains the retry settings for each method of ContainerAnalysisV1Beta1Client.
type ContainerAnalysisV1Beta1CallOptions struct {
	SetIamPolicy       []gax.CallOption
	GetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
	GetScanConfig      []gax.CallOption
	ListScanConfigs    []gax.CallOption
	UpdateScanConfig   []gax.CallOption
}

func defaultContainerAnalysisV1Beta1GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("containeranalysis.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("containeranalysis.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://containeranalysis.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultContainerAnalysisV1Beta1CallOptions() *ContainerAnalysisV1Beta1CallOptions {
	return &ContainerAnalysisV1Beta1CallOptions{
		SetIamPolicy:       []gax.CallOption{},
		GetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
		GetScanConfig: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListScanConfigs: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateScanConfig: []gax.CallOption{},
	}
}

// internalContainerAnalysisV1Beta1Client is an interface that defines the methods available from Container Analysis API.
type internalContainerAnalysisV1Beta1Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	GetScanConfig(context.Context, *containeranalysispb.GetScanConfigRequest, ...gax.CallOption) (*containeranalysispb.ScanConfig, error)
	ListScanConfigs(context.Context, *containeranalysispb.ListScanConfigsRequest, ...gax.CallOption) *ScanConfigIterator
	UpdateScanConfig(context.Context, *containeranalysispb.UpdateScanConfigRequest, ...gax.CallOption) (*containeranalysispb.ScanConfig, error)
}

// ContainerAnalysisV1Beta1Client is a client for interacting with Container Analysis API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Retrieves analysis results of Cloud components such as Docker container
// images. The Container Analysis API is an implementation of the
// Grafeas (at grafeas.io) API.
//
// Analysis results are stored as a series of occurrences. An Occurrence
// contains information about a specific analysis instance on a resource. An
// occurrence refers to a Note. A note contains details describing the
// analysis and is generally stored in a separate project, called a Provider.
// Multiple occurrences can refer to the same note.
//
// For example, an SSL vulnerability could affect multiple images. In this case,
// there would be one note for the vulnerability and an occurrence for each
// image with the vulnerability referring to that note.
type ContainerAnalysisV1Beta1Client struct {
	// The internal transport-dependent client.
	internalClient internalContainerAnalysisV1Beta1Client

	// The call options for this service.
	CallOptions *ContainerAnalysisV1Beta1CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ContainerAnalysisV1Beta1Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ContainerAnalysisV1Beta1Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ContainerAnalysisV1Beta1Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SetIamPolicy sets the access control policy on the specified note or occurrence.
// Requires containeranalysis.notes.setIamPolicy or
// containeranalysis.occurrences.setIamPolicy permission if the resource is
// a note or an occurrence, respectively.
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *ContainerAnalysisV1Beta1Client) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a note or an occurrence resource.
// Requires containeranalysis.notes.setIamPolicy or
// containeranalysis.occurrences.setIamPolicy permission if the resource is
// a note or occurrence, respectively.
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *ContainerAnalysisV1Beta1Client) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns the permissions that a caller has on the specified note or
// occurrence. Requires list permission on the project (for example,
// containeranalysis.notes.list).
//
// The resource takes the format projects/[PROJECT_ID]/notes/[NOTE_ID] for
// notes and projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID] for
// occurrences.
func (c *ContainerAnalysisV1Beta1Client) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// GetScanConfig gets the specified scan configuration.
func (c *ContainerAnalysisV1Beta1Client) GetScanConfig(ctx context.Context, req *containeranalysispb.GetScanConfigRequest, opts ...gax.CallOption) (*containeranalysispb.ScanConfig, error) {
	return c.internalClient.GetScanConfig(ctx, req, opts...)
}

// ListScanConfigs lists scan configurations for the specified project.
func (c *ContainerAnalysisV1Beta1Client) ListScanConfigs(ctx context.Context, req *containeranalysispb.ListScanConfigsRequest, opts ...gax.CallOption) *ScanConfigIterator {
	return c.internalClient.ListScanConfigs(ctx, req, opts...)
}

// UpdateScanConfig updates the specified scan configuration.
func (c *ContainerAnalysisV1Beta1Client) UpdateScanConfig(ctx context.Context, req *containeranalysispb.UpdateScanConfigRequest, opts ...gax.CallOption) (*containeranalysispb.ScanConfig, error) {
	return c.internalClient.UpdateScanConfig(ctx, req, opts...)
}

// containerAnalysisV1Beta1GRPCClient is a client for interacting with Container Analysis API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type containerAnalysisV1Beta1GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ContainerAnalysisV1Beta1Client
	CallOptions **ContainerAnalysisV1Beta1CallOptions

	// The gRPC API client.
	containerAnalysisV1Beta1Client containeranalysispb.ContainerAnalysisV1Beta1Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewContainerAnalysisV1Beta1Client creates a new container analysis v1 beta1 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Retrieves analysis results of Cloud components such as Docker container
// images. The Container Analysis API is an implementation of the
// Grafeas (at grafeas.io) API.
//
// Analysis results are stored as a series of occurrences. An Occurrence
// contains information about a specific analysis instance on a resource. An
// occurrence refers to a Note. A note contains details describing the
// analysis and is generally stored in a separate project, called a Provider.
// Multiple occurrences can refer to the same note.
//
// For example, an SSL vulnerability could affect multiple images. In this case,
// there would be one note for the vulnerability and an occurrence for each
// image with the vulnerability referring to that note.
func NewContainerAnalysisV1Beta1Client(ctx context.Context, opts ...option.ClientOption) (*ContainerAnalysisV1Beta1Client, error) {
	clientOpts := defaultContainerAnalysisV1Beta1GRPCClientOptions()
	if newContainerAnalysisV1Beta1ClientHook != nil {
		hookOpts, err := newContainerAnalysisV1Beta1ClientHook(ctx, clientHookParams{})
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
	client := ContainerAnalysisV1Beta1Client{CallOptions: defaultContainerAnalysisV1Beta1CallOptions()}

	c := &containerAnalysisV1Beta1GRPCClient{
		connPool:                       connPool,
		disableDeadlines:               disableDeadlines,
		containerAnalysisV1Beta1Client: containeranalysispb.NewContainerAnalysisV1Beta1Client(connPool),
		CallOptions:                    &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *containerAnalysisV1Beta1GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *containerAnalysisV1Beta1GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *containerAnalysisV1Beta1GRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *containerAnalysisV1Beta1GRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.containerAnalysisV1Beta1Client.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *containerAnalysisV1Beta1GRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.containerAnalysisV1Beta1Client.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *containerAnalysisV1Beta1GRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.containerAnalysisV1Beta1Client.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *containerAnalysisV1Beta1GRPCClient) GetScanConfig(ctx context.Context, req *containeranalysispb.GetScanConfigRequest, opts ...gax.CallOption) (*containeranalysispb.ScanConfig, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetScanConfig[0:len((*c.CallOptions).GetScanConfig):len((*c.CallOptions).GetScanConfig)], opts...)
	var resp *containeranalysispb.ScanConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.containerAnalysisV1Beta1Client.GetScanConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *containerAnalysisV1Beta1GRPCClient) ListScanConfigs(ctx context.Context, req *containeranalysispb.ListScanConfigsRequest, opts ...gax.CallOption) *ScanConfigIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListScanConfigs[0:len((*c.CallOptions).ListScanConfigs):len((*c.CallOptions).ListScanConfigs)], opts...)
	it := &ScanConfigIterator{}
	req = proto.Clone(req).(*containeranalysispb.ListScanConfigsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*containeranalysispb.ScanConfig, string, error) {
		resp := &containeranalysispb.ListScanConfigsResponse{}
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
			resp, err = c.containerAnalysisV1Beta1Client.ListScanConfigs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetScanConfigs(), resp.GetNextPageToken(), nil
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

func (c *containerAnalysisV1Beta1GRPCClient) UpdateScanConfig(ctx context.Context, req *containeranalysispb.UpdateScanConfigRequest, opts ...gax.CallOption) (*containeranalysispb.ScanConfig, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateScanConfig[0:len((*c.CallOptions).UpdateScanConfig):len((*c.CallOptions).UpdateScanConfig)], opts...)
	var resp *containeranalysispb.ScanConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.containerAnalysisV1Beta1Client.UpdateScanConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ScanConfigIterator manages a stream of *containeranalysispb.ScanConfig.
type ScanConfigIterator struct {
	items    []*containeranalysispb.ScanConfig
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
	InternalFetch func(pageSize int, pageToken string) (results []*containeranalysispb.ScanConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ScanConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ScanConfigIterator) Next() (*containeranalysispb.ScanConfig, error) {
	var item *containeranalysispb.ScanConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ScanConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *ScanConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
