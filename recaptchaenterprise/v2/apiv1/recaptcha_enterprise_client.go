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

package recaptchaenterprise

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
	recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateAssessment                     []gax.CallOption
	AnnotateAssessment                   []gax.CallOption
	CreateKey                            []gax.CallOption
	ListKeys                             []gax.CallOption
	GetKey                               []gax.CallOption
	UpdateKey                            []gax.CallOption
	DeleteKey                            []gax.CallOption
	MigrateKey                           []gax.CallOption
	GetMetrics                           []gax.CallOption
	ListRelatedAccountGroups             []gax.CallOption
	ListRelatedAccountGroupMemberships   []gax.CallOption
	SearchRelatedAccountGroupMemberships []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("recaptchaenterprise.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("recaptchaenterprise.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://recaptchaenterprise.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateAssessment:                     []gax.CallOption{},
		AnnotateAssessment:                   []gax.CallOption{},
		CreateKey:                            []gax.CallOption{},
		ListKeys:                             []gax.CallOption{},
		GetKey:                               []gax.CallOption{},
		UpdateKey:                            []gax.CallOption{},
		DeleteKey:                            []gax.CallOption{},
		MigrateKey:                           []gax.CallOption{},
		GetMetrics:                           []gax.CallOption{},
		ListRelatedAccountGroups:             []gax.CallOption{},
		ListRelatedAccountGroupMemberships:   []gax.CallOption{},
		SearchRelatedAccountGroupMemberships: []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from reCAPTCHA Enterprise API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateAssessment(context.Context, *recaptchaenterprisepb.CreateAssessmentRequest, ...gax.CallOption) (*recaptchaenterprisepb.Assessment, error)
	AnnotateAssessment(context.Context, *recaptchaenterprisepb.AnnotateAssessmentRequest, ...gax.CallOption) (*recaptchaenterprisepb.AnnotateAssessmentResponse, error)
	CreateKey(context.Context, *recaptchaenterprisepb.CreateKeyRequest, ...gax.CallOption) (*recaptchaenterprisepb.Key, error)
	ListKeys(context.Context, *recaptchaenterprisepb.ListKeysRequest, ...gax.CallOption) *KeyIterator
	GetKey(context.Context, *recaptchaenterprisepb.GetKeyRequest, ...gax.CallOption) (*recaptchaenterprisepb.Key, error)
	UpdateKey(context.Context, *recaptchaenterprisepb.UpdateKeyRequest, ...gax.CallOption) (*recaptchaenterprisepb.Key, error)
	DeleteKey(context.Context, *recaptchaenterprisepb.DeleteKeyRequest, ...gax.CallOption) error
	MigrateKey(context.Context, *recaptchaenterprisepb.MigrateKeyRequest, ...gax.CallOption) (*recaptchaenterprisepb.Key, error)
	GetMetrics(context.Context, *recaptchaenterprisepb.GetMetricsRequest, ...gax.CallOption) (*recaptchaenterprisepb.Metrics, error)
	ListRelatedAccountGroups(context.Context, *recaptchaenterprisepb.ListRelatedAccountGroupsRequest, ...gax.CallOption) *RelatedAccountGroupIterator
	ListRelatedAccountGroupMemberships(context.Context, *recaptchaenterprisepb.ListRelatedAccountGroupMembershipsRequest, ...gax.CallOption) *RelatedAccountGroupMembershipIterator
	SearchRelatedAccountGroupMemberships(context.Context, *recaptchaenterprisepb.SearchRelatedAccountGroupMembershipsRequest, ...gax.CallOption) *RelatedAccountGroupMembershipIterator
}

// Client is a client for interacting with reCAPTCHA Enterprise API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to determine the likelihood an event is legitimate.
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

// CreateAssessment creates an Assessment of the likelihood an event is legitimate.
func (c *Client) CreateAssessment(ctx context.Context, req *recaptchaenterprisepb.CreateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Assessment, error) {
	return c.internalClient.CreateAssessment(ctx, req, opts...)
}

// AnnotateAssessment annotates a previously created Assessment to provide additional information
// on whether the event turned out to be authentic or fraudulent.
func (c *Client) AnnotateAssessment(ctx context.Context, req *recaptchaenterprisepb.AnnotateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.AnnotateAssessmentResponse, error) {
	return c.internalClient.AnnotateAssessment(ctx, req, opts...)
}

// CreateKey creates a new reCAPTCHA Enterprise key.
func (c *Client) CreateKey(ctx context.Context, req *recaptchaenterprisepb.CreateKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	return c.internalClient.CreateKey(ctx, req, opts...)
}

// ListKeys returns the list of all keys that belong to a project.
func (c *Client) ListKeys(ctx context.Context, req *recaptchaenterprisepb.ListKeysRequest, opts ...gax.CallOption) *KeyIterator {
	return c.internalClient.ListKeys(ctx, req, opts...)
}

// GetKey returns the specified key.
func (c *Client) GetKey(ctx context.Context, req *recaptchaenterprisepb.GetKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	return c.internalClient.GetKey(ctx, req, opts...)
}

// UpdateKey updates the specified key.
func (c *Client) UpdateKey(ctx context.Context, req *recaptchaenterprisepb.UpdateKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	return c.internalClient.UpdateKey(ctx, req, opts...)
}

// DeleteKey deletes the specified key.
func (c *Client) DeleteKey(ctx context.Context, req *recaptchaenterprisepb.DeleteKeyRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteKey(ctx, req, opts...)
}

// MigrateKey migrates an existing key from reCAPTCHA to reCAPTCHA Enterprise.
// Once a key is migrated, it can be used from either product. SiteVerify
// requests are billed as CreateAssessment calls. You must be
// authenticated as one of the current owners of the reCAPTCHA Site Key, and
// your user must have the reCAPTCHA Enterprise Admin IAM role in the
// destination project.
func (c *Client) MigrateKey(ctx context.Context, req *recaptchaenterprisepb.MigrateKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	return c.internalClient.MigrateKey(ctx, req, opts...)
}

// GetMetrics get some aggregated metrics for a Key. This data can be used to build
// dashboards.
func (c *Client) GetMetrics(ctx context.Context, req *recaptchaenterprisepb.GetMetricsRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Metrics, error) {
	return c.internalClient.GetMetrics(ctx, req, opts...)
}

// ListRelatedAccountGroups list groups of related accounts.
func (c *Client) ListRelatedAccountGroups(ctx context.Context, req *recaptchaenterprisepb.ListRelatedAccountGroupsRequest, opts ...gax.CallOption) *RelatedAccountGroupIterator {
	return c.internalClient.ListRelatedAccountGroups(ctx, req, opts...)
}

// ListRelatedAccountGroupMemberships get the memberships in a group of related accounts.
func (c *Client) ListRelatedAccountGroupMemberships(ctx context.Context, req *recaptchaenterprisepb.ListRelatedAccountGroupMembershipsRequest, opts ...gax.CallOption) *RelatedAccountGroupMembershipIterator {
	return c.internalClient.ListRelatedAccountGroupMemberships(ctx, req, opts...)
}

// SearchRelatedAccountGroupMemberships search group memberships related to a given account.
func (c *Client) SearchRelatedAccountGroupMemberships(ctx context.Context, req *recaptchaenterprisepb.SearchRelatedAccountGroupMembershipsRequest, opts ...gax.CallOption) *RelatedAccountGroupMembershipIterator {
	return c.internalClient.SearchRelatedAccountGroupMemberships(ctx, req, opts...)
}

// gRPCClient is a client for interacting with reCAPTCHA Enterprise API over gRPC transport.
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
	client recaptchaenterprisepb.RecaptchaEnterpriseServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new recaptcha enterprise service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to determine the likelihood an event is legitimate.
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
		client:           recaptchaenterprisepb.NewRecaptchaEnterpriseServiceClient(connPool),
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

func (c *gRPCClient) CreateAssessment(ctx context.Context, req *recaptchaenterprisepb.CreateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Assessment, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateAssessment[0:len((*c.CallOptions).CreateAssessment):len((*c.CallOptions).CreateAssessment)], opts...)
	var resp *recaptchaenterprisepb.Assessment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateAssessment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) AnnotateAssessment(ctx context.Context, req *recaptchaenterprisepb.AnnotateAssessmentRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.AnnotateAssessmentResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AnnotateAssessment[0:len((*c.CallOptions).AnnotateAssessment):len((*c.CallOptions).AnnotateAssessment)], opts...)
	var resp *recaptchaenterprisepb.AnnotateAssessmentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.AnnotateAssessment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateKey(ctx context.Context, req *recaptchaenterprisepb.CreateKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateKey[0:len((*c.CallOptions).CreateKey):len((*c.CallOptions).CreateKey)], opts...)
	var resp *recaptchaenterprisepb.Key
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListKeys(ctx context.Context, req *recaptchaenterprisepb.ListKeysRequest, opts ...gax.CallOption) *KeyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListKeys[0:len((*c.CallOptions).ListKeys):len((*c.CallOptions).ListKeys)], opts...)
	it := &KeyIterator{}
	req = proto.Clone(req).(*recaptchaenterprisepb.ListKeysRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recaptchaenterprisepb.Key, string, error) {
		resp := &recaptchaenterprisepb.ListKeysResponse{}
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
			resp, err = c.client.ListKeys(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetKeys(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetKey(ctx context.Context, req *recaptchaenterprisepb.GetKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetKey[0:len((*c.CallOptions).GetKey):len((*c.CallOptions).GetKey)], opts...)
	var resp *recaptchaenterprisepb.Key
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateKey(ctx context.Context, req *recaptchaenterprisepb.UpdateKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "key.name", url.QueryEscape(req.GetKey().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateKey[0:len((*c.CallOptions).UpdateKey):len((*c.CallOptions).UpdateKey)], opts...)
	var resp *recaptchaenterprisepb.Key
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteKey(ctx context.Context, req *recaptchaenterprisepb.DeleteKeyRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteKey[0:len((*c.CallOptions).DeleteKey):len((*c.CallOptions).DeleteKey)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) MigrateKey(ctx context.Context, req *recaptchaenterprisepb.MigrateKeyRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Key, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MigrateKey[0:len((*c.CallOptions).MigrateKey):len((*c.CallOptions).MigrateKey)], opts...)
	var resp *recaptchaenterprisepb.Key
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.MigrateKey(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetMetrics(ctx context.Context, req *recaptchaenterprisepb.GetMetricsRequest, opts ...gax.CallOption) (*recaptchaenterprisepb.Metrics, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetMetrics[0:len((*c.CallOptions).GetMetrics):len((*c.CallOptions).GetMetrics)], opts...)
	var resp *recaptchaenterprisepb.Metrics
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetMetrics(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListRelatedAccountGroups(ctx context.Context, req *recaptchaenterprisepb.ListRelatedAccountGroupsRequest, opts ...gax.CallOption) *RelatedAccountGroupIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListRelatedAccountGroups[0:len((*c.CallOptions).ListRelatedAccountGroups):len((*c.CallOptions).ListRelatedAccountGroups)], opts...)
	it := &RelatedAccountGroupIterator{}
	req = proto.Clone(req).(*recaptchaenterprisepb.ListRelatedAccountGroupsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recaptchaenterprisepb.RelatedAccountGroup, string, error) {
		resp := &recaptchaenterprisepb.ListRelatedAccountGroupsResponse{}
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
			resp, err = c.client.ListRelatedAccountGroups(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetRelatedAccountGroups(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) ListRelatedAccountGroupMemberships(ctx context.Context, req *recaptchaenterprisepb.ListRelatedAccountGroupMembershipsRequest, opts ...gax.CallOption) *RelatedAccountGroupMembershipIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListRelatedAccountGroupMemberships[0:len((*c.CallOptions).ListRelatedAccountGroupMemberships):len((*c.CallOptions).ListRelatedAccountGroupMemberships)], opts...)
	it := &RelatedAccountGroupMembershipIterator{}
	req = proto.Clone(req).(*recaptchaenterprisepb.ListRelatedAccountGroupMembershipsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recaptchaenterprisepb.RelatedAccountGroupMembership, string, error) {
		resp := &recaptchaenterprisepb.ListRelatedAccountGroupMembershipsResponse{}
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
			resp, err = c.client.ListRelatedAccountGroupMemberships(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetRelatedAccountGroupMemberships(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) SearchRelatedAccountGroupMemberships(ctx context.Context, req *recaptchaenterprisepb.SearchRelatedAccountGroupMembershipsRequest, opts ...gax.CallOption) *RelatedAccountGroupMembershipIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project", url.QueryEscape(req.GetProject())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SearchRelatedAccountGroupMemberships[0:len((*c.CallOptions).SearchRelatedAccountGroupMemberships):len((*c.CallOptions).SearchRelatedAccountGroupMemberships)], opts...)
	it := &RelatedAccountGroupMembershipIterator{}
	req = proto.Clone(req).(*recaptchaenterprisepb.SearchRelatedAccountGroupMembershipsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*recaptchaenterprisepb.RelatedAccountGroupMembership, string, error) {
		resp := &recaptchaenterprisepb.SearchRelatedAccountGroupMembershipsResponse{}
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
			resp, err = c.client.SearchRelatedAccountGroupMemberships(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetRelatedAccountGroupMemberships(), resp.GetNextPageToken(), nil
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

// KeyIterator manages a stream of *recaptchaenterprisepb.Key.
type KeyIterator struct {
	items    []*recaptchaenterprisepb.Key
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
	InternalFetch func(pageSize int, pageToken string) (results []*recaptchaenterprisepb.Key, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *KeyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *KeyIterator) Next() (*recaptchaenterprisepb.Key, error) {
	var item *recaptchaenterprisepb.Key
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *KeyIterator) bufLen() int {
	return len(it.items)
}

func (it *KeyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RelatedAccountGroupIterator manages a stream of *recaptchaenterprisepb.RelatedAccountGroup.
type RelatedAccountGroupIterator struct {
	items    []*recaptchaenterprisepb.RelatedAccountGroup
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
	InternalFetch func(pageSize int, pageToken string) (results []*recaptchaenterprisepb.RelatedAccountGroup, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *RelatedAccountGroupIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RelatedAccountGroupIterator) Next() (*recaptchaenterprisepb.RelatedAccountGroup, error) {
	var item *recaptchaenterprisepb.RelatedAccountGroup
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RelatedAccountGroupIterator) bufLen() int {
	return len(it.items)
}

func (it *RelatedAccountGroupIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RelatedAccountGroupMembershipIterator manages a stream of *recaptchaenterprisepb.RelatedAccountGroupMembership.
type RelatedAccountGroupMembershipIterator struct {
	items    []*recaptchaenterprisepb.RelatedAccountGroupMembership
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
	InternalFetch func(pageSize int, pageToken string) (results []*recaptchaenterprisepb.RelatedAccountGroupMembership, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *RelatedAccountGroupMembershipIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RelatedAccountGroupMembershipIterator) Next() (*recaptchaenterprisepb.RelatedAccountGroupMembership, error) {
	var item *recaptchaenterprisepb.RelatedAccountGroupMembership
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RelatedAccountGroupMembershipIterator) bufLen() int {
	return len(it.items)
}

func (it *RelatedAccountGroupMembershipIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
