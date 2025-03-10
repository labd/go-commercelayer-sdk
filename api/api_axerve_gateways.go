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

// AxerveGatewaysApiService AxerveGatewaysApi service
type AxerveGatewaysApiService service

type AxerveGatewaysApiDELETEAxerveGatewaysAxerveGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *AxerveGatewaysApiService
	axerveGatewayId interface{}
}

func (r AxerveGatewaysApiDELETEAxerveGatewaysAxerveGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEAxerveGatewaysAxerveGatewayIdExecute(r)
}

/*
DELETEAxerveGatewaysAxerveGatewayId Delete an axerve gateway

Delete an axerve gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param axerveGatewayId The resource's id
	@return AxerveGatewaysApiDELETEAxerveGatewaysAxerveGatewayIdRequest
*/
func (a *AxerveGatewaysApiService) DELETEAxerveGatewaysAxerveGatewayId(ctx context.Context, axerveGatewayId interface{}) AxerveGatewaysApiDELETEAxerveGatewaysAxerveGatewayIdRequest {
	return AxerveGatewaysApiDELETEAxerveGatewaysAxerveGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		axerveGatewayId: axerveGatewayId,
	}
}

// Execute executes the request
func (a *AxerveGatewaysApiService) DELETEAxerveGatewaysAxerveGatewayIdExecute(r AxerveGatewaysApiDELETEAxerveGatewaysAxerveGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxerveGatewaysApiService.DELETEAxerveGatewaysAxerveGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_gateways/{axerveGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"axerveGatewayId"+"}", url.PathEscape(parameterValueToString(r.axerveGatewayId, "axerveGatewayId")), -1)

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

type AxerveGatewaysApiGETAxerveGatewaysRequest struct {
	ctx        context.Context
	ApiService *AxerveGatewaysApiService
}

func (r AxerveGatewaysApiGETAxerveGatewaysRequest) Execute() (*GETAxerveGateways200Response, *http.Response, error) {
	return r.ApiService.GETAxerveGatewaysExecute(r)
}

/*
GETAxerveGateways List all axerve gateways

List all axerve gateways

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AxerveGatewaysApiGETAxerveGatewaysRequest
*/
func (a *AxerveGatewaysApiService) GETAxerveGateways(ctx context.Context) AxerveGatewaysApiGETAxerveGatewaysRequest {
	return AxerveGatewaysApiGETAxerveGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETAxerveGateways200Response
func (a *AxerveGatewaysApiService) GETAxerveGatewaysExecute(r AxerveGatewaysApiGETAxerveGatewaysRequest) (*GETAxerveGateways200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAxerveGateways200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxerveGatewaysApiService.GETAxerveGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_gateways"

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

type AxerveGatewaysApiGETAxerveGatewaysAxerveGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *AxerveGatewaysApiService
	axerveGatewayId interface{}
}

func (r AxerveGatewaysApiGETAxerveGatewaysAxerveGatewayIdRequest) Execute() (*GETAxerveGatewaysAxerveGatewayId200Response, *http.Response, error) {
	return r.ApiService.GETAxerveGatewaysAxerveGatewayIdExecute(r)
}

/*
GETAxerveGatewaysAxerveGatewayId Retrieve an axerve gateway

Retrieve an axerve gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param axerveGatewayId The resource's id
	@return AxerveGatewaysApiGETAxerveGatewaysAxerveGatewayIdRequest
*/
func (a *AxerveGatewaysApiService) GETAxerveGatewaysAxerveGatewayId(ctx context.Context, axerveGatewayId interface{}) AxerveGatewaysApiGETAxerveGatewaysAxerveGatewayIdRequest {
	return AxerveGatewaysApiGETAxerveGatewaysAxerveGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		axerveGatewayId: axerveGatewayId,
	}
}

// Execute executes the request
//
//	@return GETAxerveGatewaysAxerveGatewayId200Response
func (a *AxerveGatewaysApiService) GETAxerveGatewaysAxerveGatewayIdExecute(r AxerveGatewaysApiGETAxerveGatewaysAxerveGatewayIdRequest) (*GETAxerveGatewaysAxerveGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAxerveGatewaysAxerveGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxerveGatewaysApiService.GETAxerveGatewaysAxerveGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_gateways/{axerveGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"axerveGatewayId"+"}", url.PathEscape(parameterValueToString(r.axerveGatewayId, "axerveGatewayId")), -1)

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

type AxerveGatewaysApiPATCHAxerveGatewaysAxerveGatewayIdRequest struct {
	ctx                 context.Context
	ApiService          *AxerveGatewaysApiService
	axerveGatewayUpdate *AxerveGatewayUpdate
	axerveGatewayId     interface{}
}

func (r AxerveGatewaysApiPATCHAxerveGatewaysAxerveGatewayIdRequest) AxerveGatewayUpdate(axerveGatewayUpdate AxerveGatewayUpdate) AxerveGatewaysApiPATCHAxerveGatewaysAxerveGatewayIdRequest {
	r.axerveGatewayUpdate = &axerveGatewayUpdate
	return r
}

func (r AxerveGatewaysApiPATCHAxerveGatewaysAxerveGatewayIdRequest) Execute() (*PATCHAxerveGatewaysAxerveGatewayId200Response, *http.Response, error) {
	return r.ApiService.PATCHAxerveGatewaysAxerveGatewayIdExecute(r)
}

/*
PATCHAxerveGatewaysAxerveGatewayId Update an axerve gateway

Update an axerve gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param axerveGatewayId The resource's id
	@return AxerveGatewaysApiPATCHAxerveGatewaysAxerveGatewayIdRequest
*/
func (a *AxerveGatewaysApiService) PATCHAxerveGatewaysAxerveGatewayId(ctx context.Context, axerveGatewayId interface{}) AxerveGatewaysApiPATCHAxerveGatewaysAxerveGatewayIdRequest {
	return AxerveGatewaysApiPATCHAxerveGatewaysAxerveGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		axerveGatewayId: axerveGatewayId,
	}
}

// Execute executes the request
//
//	@return PATCHAxerveGatewaysAxerveGatewayId200Response
func (a *AxerveGatewaysApiService) PATCHAxerveGatewaysAxerveGatewayIdExecute(r AxerveGatewaysApiPATCHAxerveGatewaysAxerveGatewayIdRequest) (*PATCHAxerveGatewaysAxerveGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHAxerveGatewaysAxerveGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxerveGatewaysApiService.PATCHAxerveGatewaysAxerveGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_gateways/{axerveGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"axerveGatewayId"+"}", url.PathEscape(parameterValueToString(r.axerveGatewayId, "axerveGatewayId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.axerveGatewayUpdate == nil {
		return localVarReturnValue, nil, reportError("axerveGatewayUpdate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

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
	// body params
	localVarPostBody = r.axerveGatewayUpdate
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

type AxerveGatewaysApiPOSTAxerveGatewaysRequest struct {
	ctx                 context.Context
	ApiService          *AxerveGatewaysApiService
	axerveGatewayCreate *AxerveGatewayCreate
}

func (r AxerveGatewaysApiPOSTAxerveGatewaysRequest) AxerveGatewayCreate(axerveGatewayCreate AxerveGatewayCreate) AxerveGatewaysApiPOSTAxerveGatewaysRequest {
	r.axerveGatewayCreate = &axerveGatewayCreate
	return r
}

func (r AxerveGatewaysApiPOSTAxerveGatewaysRequest) Execute() (*POSTAxerveGateways201Response, *http.Response, error) {
	return r.ApiService.POSTAxerveGatewaysExecute(r)
}

/*
POSTAxerveGateways Create an axerve gateway

Create an axerve gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AxerveGatewaysApiPOSTAxerveGatewaysRequest
*/
func (a *AxerveGatewaysApiService) POSTAxerveGateways(ctx context.Context) AxerveGatewaysApiPOSTAxerveGatewaysRequest {
	return AxerveGatewaysApiPOSTAxerveGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTAxerveGateways201Response
func (a *AxerveGatewaysApiService) POSTAxerveGatewaysExecute(r AxerveGatewaysApiPOSTAxerveGatewaysRequest) (*POSTAxerveGateways201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTAxerveGateways201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxerveGatewaysApiService.POSTAxerveGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.axerveGatewayCreate == nil {
		return localVarReturnValue, nil, reportError("axerveGatewayCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

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
	// body params
	localVarPostBody = r.axerveGatewayCreate
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
