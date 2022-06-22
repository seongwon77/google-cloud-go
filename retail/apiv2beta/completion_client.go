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

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	retailpb "google.golang.org/genproto/googleapis/cloud/retail/v2beta"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCompletionClientHook clientHook

// CompletionCallOptions contains the retry settings for each method of CompletionClient.
type CompletionCallOptions struct {
	CompleteQuery        []gax.CallOption
	ImportCompletionData []gax.CallOption
}

func defaultCompletionGRPCClientOptions() []option.ClientOption {
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

func defaultCompletionCallOptions() *CompletionCallOptions {
	return &CompletionCallOptions{
		CompleteQuery: []gax.CallOption{
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
		ImportCompletionData: []gax.CallOption{
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

// internalCompletionClient is an interface that defines the methods available from Retail API.
type internalCompletionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CompleteQuery(context.Context, *retailpb.CompleteQueryRequest, ...gax.CallOption) (*retailpb.CompleteQueryResponse, error)
	ImportCompletionData(context.Context, *retailpb.ImportCompletionDataRequest, ...gax.CallOption) (*ImportCompletionDataOperation, error)
	ImportCompletionDataOperation(name string) *ImportCompletionDataOperation
}

// CompletionClient is a client for interacting with Retail API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Auto-completion service for retail.
//
// This feature is only available for users who have Retail Search enabled.
// Please enable Retail Search on Cloud Console before using this feature.
type CompletionClient struct {
	// The internal transport-dependent client.
	internalClient internalCompletionClient

	// The call options for this service.
	CallOptions *CompletionCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CompletionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CompletionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CompletionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CompleteQuery completes the specified prefix with keyword suggestions.
//
// This feature is only available for users who have Retail Search enabled.
// Please enable Retail Search on Cloud Console before using this feature.
func (c *CompletionClient) CompleteQuery(ctx context.Context, req *retailpb.CompleteQueryRequest, opts ...gax.CallOption) (*retailpb.CompleteQueryResponse, error) {
	return c.internalClient.CompleteQuery(ctx, req, opts...)
}

// ImportCompletionData bulk import of processed completion dataset.
//
// Request processing is asynchronous. Partial updating is not supported.
//
// The operation is successfully finished only after the imported suggestions
// are indexed successfully and ready for serving. The process takes hours.
//
// This feature is only available for users who have Retail Search enabled.
// Please enable Retail Search on Cloud Console before using this feature.
func (c *CompletionClient) ImportCompletionData(ctx context.Context, req *retailpb.ImportCompletionDataRequest, opts ...gax.CallOption) (*ImportCompletionDataOperation, error) {
	return c.internalClient.ImportCompletionData(ctx, req, opts...)
}

// ImportCompletionDataOperation returns a new ImportCompletionDataOperation from a given name.
// The name must be that of a previously created ImportCompletionDataOperation, possibly from a different process.
func (c *CompletionClient) ImportCompletionDataOperation(name string) *ImportCompletionDataOperation {
	return c.internalClient.ImportCompletionDataOperation(name)
}

// completionGRPCClient is a client for interacting with Retail API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type completionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CompletionClient
	CallOptions **CompletionCallOptions

	// The gRPC API client.
	completionClient retailpb.CompletionServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCompletionClient creates a new completion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Auto-completion service for retail.
//
// This feature is only available for users who have Retail Search enabled.
// Please enable Retail Search on Cloud Console before using this feature.
func NewCompletionClient(ctx context.Context, opts ...option.ClientOption) (*CompletionClient, error) {
	clientOpts := defaultCompletionGRPCClientOptions()
	if newCompletionClientHook != nil {
		hookOpts, err := newCompletionClientHook(ctx, clientHookParams{})
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
	client := CompletionClient{CallOptions: defaultCompletionCallOptions()}

	c := &completionGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		completionClient: retailpb.NewCompletionServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *completionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *completionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *completionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *completionGRPCClient) CompleteQuery(ctx context.Context, req *retailpb.CompleteQueryRequest, opts ...gax.CallOption) (*retailpb.CompleteQueryResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "catalog", url.QueryEscape(req.GetCatalog())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CompleteQuery[0:len((*c.CallOptions).CompleteQuery):len((*c.CallOptions).CompleteQuery)], opts...)
	var resp *retailpb.CompleteQueryResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.completionClient.CompleteQuery(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *completionGRPCClient) ImportCompletionData(ctx context.Context, req *retailpb.ImportCompletionDataRequest, opts ...gax.CallOption) (*ImportCompletionDataOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ImportCompletionData[0:len((*c.CallOptions).ImportCompletionData):len((*c.CallOptions).ImportCompletionData)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.completionClient.ImportCompletionData(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportCompletionDataOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// ImportCompletionDataOperation manages a long-running operation from ImportCompletionData.
type ImportCompletionDataOperation struct {
	lro *longrunning.Operation
}

// ImportCompletionDataOperation returns a new ImportCompletionDataOperation from a given name.
// The name must be that of a previously created ImportCompletionDataOperation, possibly from a different process.
func (c *completionGRPCClient) ImportCompletionDataOperation(name string) *ImportCompletionDataOperation {
	return &ImportCompletionDataOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ImportCompletionDataOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*retailpb.ImportCompletionDataResponse, error) {
	var resp retailpb.ImportCompletionDataResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ImportCompletionDataOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*retailpb.ImportCompletionDataResponse, error) {
	var resp retailpb.ImportCompletionDataResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ImportCompletionDataOperation) Metadata() (*retailpb.ImportMetadata, error) {
	var meta retailpb.ImportMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ImportCompletionDataOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ImportCompletionDataOperation) Name() string {
	return op.lro.Name()
}
