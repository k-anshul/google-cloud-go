// Copyright 2024 Google LLC
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
	cxpb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func ExampleNewPlaybooksClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewPlaybooksRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExamplePlaybooksClient_CreatePlaybook() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.CreatePlaybookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb#CreatePlaybookRequest.
	}
	resp, err := c.CreatePlaybook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePlaybooksClient_CreatePlaybookVersion() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.CreatePlaybookVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb#CreatePlaybookVersionRequest.
	}
	resp, err := c.CreatePlaybookVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePlaybooksClient_DeletePlaybook() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.DeletePlaybookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb#DeletePlaybookRequest.
	}
	err = c.DeletePlaybook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExamplePlaybooksClient_DeletePlaybookVersion() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.DeletePlaybookVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb#DeletePlaybookVersionRequest.
	}
	err = c.DeletePlaybookVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExamplePlaybooksClient_GetPlaybook() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.GetPlaybookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb#GetPlaybookRequest.
	}
	resp, err := c.GetPlaybook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePlaybooksClient_GetPlaybookVersion() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.GetPlaybookVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb#GetPlaybookVersionRequest.
	}
	resp, err := c.GetPlaybookVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePlaybooksClient_ListPlaybookVersions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.ListPlaybookVersionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb#ListPlaybookVersionsRequest.
	}
	it := c.ListPlaybookVersions(ctx, req)
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

func ExamplePlaybooksClient_ListPlaybooks() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.ListPlaybooksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb#ListPlaybooksRequest.
	}
	it := c.ListPlaybooks(ctx, req)
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

func ExamplePlaybooksClient_UpdatePlaybook() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.UpdatePlaybookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb#UpdatePlaybookRequest.
	}
	resp, err := c.UpdatePlaybook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePlaybooksClient_GetLocation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
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

func ExamplePlaybooksClient_ListLocations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
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

func ExamplePlaybooksClient_CancelOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.CancelOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#CancelOperationRequest.
	}
	err = c.CancelOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExamplePlaybooksClient_GetOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExamplePlaybooksClient_ListOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := cx.NewPlaybooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
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
