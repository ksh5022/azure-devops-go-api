// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package audit

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("94ff054d-5ee1-413d-9341-3f4a7827de2e")

type Client interface {
	// [Preview API] Downloads audit log entries.
	DownloadLog(context.Context, DownloadLogArgs) (io.ReadCloser, error)
	// [Preview API] Queries audit log entries
	QueryLog(context.Context, QueryLogArgs) (*AuditLogQueryResult, error)
	// [Preview API] Creates new Audit Stream
	CreateAuditStream(context.Context, CreateAuditStreamArgs) (*AuditStream, error)
	// [Preview API] Delete Audit Stream
	DeleteAuditStream(context.Context, DeleteAuditStreamArgs) error
	// [Preview API] Return all Audit Streams scoped to an organization
	QueryAllAuditStreams(context.Context) (*[]AuditStream, error)
	// [Preview API] Return Audit Stream with id of streamId if one exists otherwise throw
	QueryAuditStreamById(context.Context, QueryStreamByIdArgs) (*AuditStream, error)
	// [Preview API] Update existing Audit Stream status
	UpdateAuditStreamStatus(context.Context, UpdateAuditStreamStatusArgs) (*AuditStream, error)
	// [Preview API] Update existing Audit Stream
	UpdateAuditStream(context.Context, UpdateAuditStreamArgs) (*AuditStream, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &ClientImpl{
		Client: *client,
	}, nil
}

// [Preview API] Downloads audit log entries.
func (client *ClientImpl) DownloadLog(ctx context.Context, args DownloadLogArgs) (io.ReadCloser, error) {
	queryParams := url.Values{}
	if args.Format == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "format"}
	}
	queryParams.Add("format", *args.Format)
	if args.StartTime != nil {
		queryParams.Add("startTime", (*args.StartTime).AsQueryParameter())
	}
	if args.EndTime != nil {
		queryParams.Add("endTime", (*args.EndTime).AsQueryParameter())
	}
	locationId, _ := uuid.Parse("b7b98a76-04e8-4f4d-ac72-9d46492caaac")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the DownloadLog function
type DownloadLogArgs struct {
	// (required) File format for download. Can be "json" or "csv".
	Format *string
	// (optional) Start time of download window. Optional
	StartTime *azuredevops.Time
	// (optional) End time of download window. Optional
	EndTime *azuredevops.Time
}

// [Preview API] Queries audit log entries
func (client *ClientImpl) QueryLog(ctx context.Context, args QueryLogArgs) (*AuditLogQueryResult, error) {
	queryParams := url.Values{}
	if args.StartTime != nil {
		queryParams.Add("startTime", (*args.StartTime).AsQueryParameter())
	}
	if args.EndTime != nil {
		queryParams.Add("endTime", (*args.EndTime).AsQueryParameter())
	}
	if args.BatchSize != nil {
		queryParams.Add("batchSize", strconv.Itoa(*args.BatchSize))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.SkipAggregation != nil {
		queryParams.Add("skipAggregation", strconv.FormatBool(*args.SkipAggregation))
	}
	locationId, _ := uuid.Parse("4e5fa14f-7097-4b73-9c85-00abc7353c61")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue AuditLogQueryResult
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the QueryLog function
type QueryLogArgs struct {
	// (optional) Start time of download window. Optional
	StartTime *azuredevops.Time
	// (optional) End time of download window. Optional
	EndTime *azuredevops.Time
	// (optional) Max number of results to return. Optional
	BatchSize *int
	// (optional) Token used for returning next set of results from previous query. Optional
	ContinuationToken *string
	// (optional) Skips aggregating events and leaves them as individual entries instead. By default events are aggregated. Event types that are aggregated: AuditLog.AccessLog.
	SkipAggregation *bool
}

// [Preview API] Creates new Audit Stream
func (client *ClientImpl) CreateAuditStream(ctx context.Context, args CreateAuditStreamArgs) (*AuditStream, error) {
	queryParams := url.Values{}
	if args.DaysToBackfill == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DaysToBackfill"}
	}
	queryParams.Add("status", strconv.Itoa(*args.DaysToBackfill))

	body, marshalErr := json.Marshal(*args.AuditStreamToCreate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("77d60bf9-1882-41c5-a90d-3a6d3c13fd3b")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue AuditStream
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

type CreateAuditStreamArgs struct {
	// (required) The number of days of previously recorded audit data that will be replayed into the stream
	DaysToBackfill *int
	// (required) The audit stream to create.
	AuditStreamToCreate *AuditStream
}

// [Preview API] Deletes the Audit Stream
func (client *ClientImpl) DeleteAuditStream(ctx context.Context, args DeleteAuditStreamArgs) error {
	routeValues := make(map[string]string)
	if args.StreamId == nil {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.StreamId"}
	}
	routeValues["streamId"] = strconv.Itoa(*args.StreamId)

	locationId, _ := uuid.Parse("77d60bf9-1882-41c5-a90d-3a6d3c13fd3b")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

type DeleteAuditStreamArgs struct {
	// (required) The ID of stream entry to delete
	StreamId *int
}

// [Preview API] Return all Audit Streams scoped to an organization
func (client *ClientImpl) QueryAllAuditStreams(ctx context.Context) (*[]AuditStream, error) {
	locationId, _ := uuid.Parse("77d60bf9-1882-41c5-a90d-3a6d3c13fd3b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []AuditStream
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// [Preview API] Return Audit Stream with id of streamId if one exists otherwise throw
func (client *ClientImpl) QueryAuditStreamById(ctx context.Context, args QueryStreamByIdArgs) (*AuditStream, error) {
	routeValues := make(map[string]string)
	if args.StreamId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.StreamId"}
	}
	routeValues["streamId"] = strconv.Itoa(*args.StreamId)

	locationId, _ := uuid.Parse("77d60bf9-1882-41c5-a90d-3a6d3c13fd3b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue AuditStream
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

type QueryStreamByIdArgs struct {
	// (required) The ID or name of the audit stream.
	StreamId *int
}

// [Preview API] Update existing Audit Stream status
func (client *ClientImpl) UpdateAuditStreamStatus(ctx context.Context, args UpdateAuditStreamStatusArgs) (*AuditStream, error) {
	routeValues := make(map[string]string)
	if args.StreamId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.StreamId"}
	}
	routeValues["streamId"] = strconv.Itoa(*args.StreamId)

	queryParams := url.Values{}
	if args.Status == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Status"}
	}
	queryParams.Add("status", string(*args.Status))

	locationId, _ := uuid.Parse("77d60bf9-1882-41c5-a90d-3a6d3c13fd3b")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue AuditStream
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

type UpdateAuditStreamStatusArgs struct {
	// Status of the audit stream
	Status *AuditStreamStatus
	// (required) The ID or name of the audit stream
	StreamId *int
}

// [Preview API] Update existing Audit Stream
func (client *ClientImpl) UpdateAuditStream(ctx context.Context, args UpdateAuditStreamArgs) (*AuditStream, error) {
	if args.AuditStreamUpdate == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AuditStreamUpdate"}
	}

	body, marshalErr := json.Marshal(*args.AuditStreamUpdate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("77d60bf9-1882-41c5-a90d-3a6d3c13fd3b")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "6.0-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue AuditStream
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

type UpdateAuditStreamArgs struct {
	// (required) The updates for the audit stream.
	AuditStreamUpdate *AuditStream
}
