/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ResourceErrorsApiService ResourceErrorsApi service
type ResourceErrorsApiService service

type ResourceErrorsApiGETOrderIdResourceErrorsRequest struct {
	ctx        context.Context
	ApiService *ResourceErrorsApiService
	orderId    interface{}
}

func (r ResourceErrorsApiGETOrderIdResourceErrorsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdResourceErrorsExecute(r)
}

/*
GETOrderIdResourceErrors Retrieve the resource errors associated to the order

Retrieve the resource errors associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return ResourceErrorsApiGETOrderIdResourceErrorsRequest
*/
func (a *ResourceErrorsApiService) GETOrderIdResourceErrors(ctx context.Context, orderId interface{}) ResourceErrorsApiGETOrderIdResourceErrorsRequest {
	return ResourceErrorsApiGETOrderIdResourceErrorsRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *ResourceErrorsApiService) GETOrderIdResourceErrorsExecute(r ResourceErrorsApiGETOrderIdResourceErrorsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceErrorsApiService.GETOrderIdResourceErrors")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/resource_errors"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ResourceErrorsApiGETResourceErrorsRequest struct {
	ctx        context.Context
	ApiService *ResourceErrorsApiService
}

func (r ResourceErrorsApiGETResourceErrorsRequest) Execute() (*GETResourceErrors200Response, *http.Response, error) {
	return r.ApiService.GETResourceErrorsExecute(r)
}

/*
GETResourceErrors List all resource errors

List all resource errors

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ResourceErrorsApiGETResourceErrorsRequest
*/
func (a *ResourceErrorsApiService) GETResourceErrors(ctx context.Context) ResourceErrorsApiGETResourceErrorsRequest {
	return ResourceErrorsApiGETResourceErrorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETResourceErrors200Response
func (a *ResourceErrorsApiService) GETResourceErrorsExecute(r ResourceErrorsApiGETResourceErrorsRequest) (*GETResourceErrors200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETResourceErrors200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceErrorsApiService.GETResourceErrors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resource_errors"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ResourceErrorsApiGETResourceErrorsResourceErrorIdRequest struct {
	ctx             context.Context
	ApiService      *ResourceErrorsApiService
	resourceErrorId interface{}
}

func (r ResourceErrorsApiGETResourceErrorsResourceErrorIdRequest) Execute() (*GETResourceErrorsResourceErrorId200Response, *http.Response, error) {
	return r.ApiService.GETResourceErrorsResourceErrorIdExecute(r)
}

/*
GETResourceErrorsResourceErrorId Retrieve a resource error

Retrieve a resource error

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param resourceErrorId The resource's id
	@return ResourceErrorsApiGETResourceErrorsResourceErrorIdRequest
*/
func (a *ResourceErrorsApiService) GETResourceErrorsResourceErrorId(ctx context.Context, resourceErrorId interface{}) ResourceErrorsApiGETResourceErrorsResourceErrorIdRequest {
	return ResourceErrorsApiGETResourceErrorsResourceErrorIdRequest{
		ApiService:      a,
		ctx:             ctx,
		resourceErrorId: resourceErrorId,
	}
}

// Execute executes the request
//
//	@return GETResourceErrorsResourceErrorId200Response
func (a *ResourceErrorsApiService) GETResourceErrorsResourceErrorIdExecute(r ResourceErrorsApiGETResourceErrorsResourceErrorIdRequest) (*GETResourceErrorsResourceErrorId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETResourceErrorsResourceErrorId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceErrorsApiService.GETResourceErrorsResourceErrorId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resource_errors/{resourceErrorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceErrorId"+"}", url.PathEscape(parameterValueToString(r.resourceErrorId, "resourceErrorId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ResourceErrorsApiGETReturnIdResourceErrorsRequest struct {
	ctx        context.Context
	ApiService *ResourceErrorsApiService
	returnId   interface{}
}

func (r ResourceErrorsApiGETReturnIdResourceErrorsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETReturnIdResourceErrorsExecute(r)
}

/*
GETReturnIdResourceErrors Retrieve the resource errors associated to the return

Retrieve the resource errors associated to the return

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param returnId The resource's id
	@return ResourceErrorsApiGETReturnIdResourceErrorsRequest
*/
func (a *ResourceErrorsApiService) GETReturnIdResourceErrors(ctx context.Context, returnId interface{}) ResourceErrorsApiGETReturnIdResourceErrorsRequest {
	return ResourceErrorsApiGETReturnIdResourceErrorsRequest{
		ApiService: a,
		ctx:        ctx,
		returnId:   returnId,
	}
}

// Execute executes the request
func (a *ResourceErrorsApiService) GETReturnIdResourceErrorsExecute(r ResourceErrorsApiGETReturnIdResourceErrorsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourceErrorsApiService.GETReturnIdResourceErrors")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/returns/{returnId}/resource_errors"
	localVarPath = strings.Replace(localVarPath, "{"+"returnId"+"}", url.PathEscape(parameterValueToString(r.returnId, "returnId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
