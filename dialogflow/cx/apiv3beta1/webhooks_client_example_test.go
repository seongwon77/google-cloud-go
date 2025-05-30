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

package cx_test

import (
	"context"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"google.golang.org/api/iterator"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func ExampleNewWebhooksClient() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleWebhooksClient_ListWebhooks() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.ListWebhooksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1#ListWebhooksRequest.
	}
	it := c.ListWebhooks(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleWebhooksClient_GetWebhook() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.GetWebhookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1#GetWebhookRequest.
	}
	resp, err := c.GetWebhook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleWebhooksClient_CreateWebhook() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.CreateWebhookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1#CreateWebhookRequest.
	}
	resp, err := c.CreateWebhook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleWebhooksClient_UpdateWebhook() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.UpdateWebhookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1#UpdateWebhookRequest.
	}
	resp, err := c.UpdateWebhook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleWebhooksClient_DeleteWebhook() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.DeleteWebhookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1#DeleteWebhookRequest.
	}
	err = c.DeleteWebhook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleWebhooksClient_GetLocation() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.GetLocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#GetLocationRequest.
	}
	resp, err := c.GetLocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleWebhooksClient_ListLocations() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	}
	it := c.ListLocations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleWebhooksClient_CancelOperation() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.CancelOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#CancelOperationRequest.
	}
	err = c.CancelOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleWebhooksClient_GetOperation() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleWebhooksClient_ListOperations() {
	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#ListOperationsRequest.
	}
	it := c.ListOperations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
