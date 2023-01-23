/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// EventCallbacksApiService EventCallbacksApi service
type EventCallbacksApiService service

type EventCallbacksApiGETEventCallbacksRequest struct {
	ctx        context.Context
	ApiService *EventCallbacksApiService
}

func (r EventCallbacksApiGETEventCallbacksRequest) Execute() (*GETEventCallbacks200Response, *http.Response, error) {
	return r.ApiService.GETEventCallbacksExecute(r)
}

/*
GETEventCallbacks List all event callbacks

List all event callbacks

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return EventCallbacksApiGETEventCallbacksRequest
*/
func (a *EventCallbacksApiService) GETEventCallbacks(ctx context.Context) EventCallbacksApiGETEventCallbacksRequest {
	return EventCallbacksApiGETEventCallbacksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETEventCallbacks200Response
func (a *EventCallbacksApiService) GETEventCallbacksExecute(r EventCallbacksApiGETEventCallbacksRequest) (*GETEventCallbacks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETEventCallbacks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventCallbacksApiService.GETEventCallbacks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event_callbacks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type EventCallbacksApiGETEventCallbacksEventCallbackIdRequest struct {
	ctx             context.Context
	ApiService      *EventCallbacksApiService
	eventCallbackId string
}

func (r EventCallbacksApiGETEventCallbacksEventCallbackIdRequest) Execute() (*GETEventCallbacksEventCallbackId200Response, *http.Response, error) {
	return r.ApiService.GETEventCallbacksEventCallbackIdExecute(r)
}

/*
GETEventCallbacksEventCallbackId Retrieve an event callback

Retrieve an event callback

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventCallbackId The resource's id
	@return EventCallbacksApiGETEventCallbacksEventCallbackIdRequest
*/
func (a *EventCallbacksApiService) GETEventCallbacksEventCallbackId(ctx context.Context, eventCallbackId string) EventCallbacksApiGETEventCallbacksEventCallbackIdRequest {
	return EventCallbacksApiGETEventCallbacksEventCallbackIdRequest{
		ApiService:      a,
		ctx:             ctx,
		eventCallbackId: eventCallbackId,
	}
}

// Execute executes the request
//
//	@return GETEventCallbacksEventCallbackId200Response
func (a *EventCallbacksApiService) GETEventCallbacksEventCallbackIdExecute(r EventCallbacksApiGETEventCallbacksEventCallbackIdRequest) (*GETEventCallbacksEventCallbackId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETEventCallbacksEventCallbackId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventCallbacksApiService.GETEventCallbacksEventCallbackId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event_callbacks/{eventCallbackId}"
	localVarPath = strings.Replace(localVarPath, "{"+"eventCallbackId"+"}", url.PathEscape(parameterToString(r.eventCallbackId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type EventCallbacksApiGETEventIdLastEventCallbacksRequest struct {
	ctx        context.Context
	ApiService *EventCallbacksApiService
	eventId    string
}

func (r EventCallbacksApiGETEventIdLastEventCallbacksRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETEventIdLastEventCallbacksExecute(r)
}

/*
GETEventIdLastEventCallbacks Retrieve the last event callbacks associated to the event

Retrieve the last event callbacks associated to the event

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventId The resource's id
	@return EventCallbacksApiGETEventIdLastEventCallbacksRequest
*/
func (a *EventCallbacksApiService) GETEventIdLastEventCallbacks(ctx context.Context, eventId string) EventCallbacksApiGETEventIdLastEventCallbacksRequest {
	return EventCallbacksApiGETEventIdLastEventCallbacksRequest{
		ApiService: a,
		ctx:        ctx,
		eventId:    eventId,
	}
}

// Execute executes the request
func (a *EventCallbacksApiService) GETEventIdLastEventCallbacksExecute(r EventCallbacksApiGETEventIdLastEventCallbacksRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventCallbacksApiService.GETEventIdLastEventCallbacks")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{eventId}/last_event_callbacks"
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterToString(r.eventId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type EventCallbacksApiGETWebhookIdLastEventCallbacksRequest struct {
	ctx        context.Context
	ApiService *EventCallbacksApiService
	webhookId  string
}

func (r EventCallbacksApiGETWebhookIdLastEventCallbacksRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETWebhookIdLastEventCallbacksExecute(r)
}

/*
GETWebhookIdLastEventCallbacks Retrieve the last event callbacks associated to the webhook

Retrieve the last event callbacks associated to the webhook

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param webhookId The resource's id
	@return EventCallbacksApiGETWebhookIdLastEventCallbacksRequest
*/
func (a *EventCallbacksApiService) GETWebhookIdLastEventCallbacks(ctx context.Context, webhookId string) EventCallbacksApiGETWebhookIdLastEventCallbacksRequest {
	return EventCallbacksApiGETWebhookIdLastEventCallbacksRequest{
		ApiService: a,
		ctx:        ctx,
		webhookId:  webhookId,
	}
}

// Execute executes the request
func (a *EventCallbacksApiService) GETWebhookIdLastEventCallbacksExecute(r EventCallbacksApiGETWebhookIdLastEventCallbacksRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventCallbacksApiService.GETWebhookIdLastEventCallbacks")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookId}/last_event_callbacks"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
