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

// StripeGatewaysApiService StripeGatewaysApi service
type StripeGatewaysApiService service

type StripeGatewaysApiDELETEStripeGatewaysStripeGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *StripeGatewaysApiService
	stripeGatewayId interface{}
}

func (r StripeGatewaysApiDELETEStripeGatewaysStripeGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEStripeGatewaysStripeGatewayIdExecute(r)
}

/*
DELETEStripeGatewaysStripeGatewayId Delete a stripe gateway

Delete a stripe gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stripeGatewayId The resource's id
	@return StripeGatewaysApiDELETEStripeGatewaysStripeGatewayIdRequest
*/
func (a *StripeGatewaysApiService) DELETEStripeGatewaysStripeGatewayId(ctx context.Context, stripeGatewayId interface{}) StripeGatewaysApiDELETEStripeGatewaysStripeGatewayIdRequest {
	return StripeGatewaysApiDELETEStripeGatewaysStripeGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stripeGatewayId: stripeGatewayId,
	}
}

// Execute executes the request
func (a *StripeGatewaysApiService) DELETEStripeGatewaysStripeGatewayIdExecute(r StripeGatewaysApiDELETEStripeGatewaysStripeGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.DELETEStripeGatewaysStripeGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways/{stripeGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stripeGatewayId"+"}", url.PathEscape(parameterValueToString(r.stripeGatewayId, "stripeGatewayId")), -1)

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

type StripeGatewaysApiGETStripeGatewaysRequest struct {
	ctx        context.Context
	ApiService *StripeGatewaysApiService
}

func (r StripeGatewaysApiGETStripeGatewaysRequest) Execute() (*GETStripeGateways200Response, *http.Response, error) {
	return r.ApiService.GETStripeGatewaysExecute(r)
}

/*
GETStripeGateways List all stripe gateways

List all stripe gateways

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return StripeGatewaysApiGETStripeGatewaysRequest
*/
func (a *StripeGatewaysApiService) GETStripeGateways(ctx context.Context) StripeGatewaysApiGETStripeGatewaysRequest {
	return StripeGatewaysApiGETStripeGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETStripeGateways200Response
func (a *StripeGatewaysApiService) GETStripeGatewaysExecute(r StripeGatewaysApiGETStripeGatewaysRequest) (*GETStripeGateways200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETStripeGateways200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.GETStripeGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways"

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

type StripeGatewaysApiGETStripeGatewaysStripeGatewayIdRequest struct {
	ctx             context.Context
	ApiService      *StripeGatewaysApiService
	stripeGatewayId interface{}
}

func (r StripeGatewaysApiGETStripeGatewaysStripeGatewayIdRequest) Execute() (*GETStripeGatewaysStripeGatewayId200Response, *http.Response, error) {
	return r.ApiService.GETStripeGatewaysStripeGatewayIdExecute(r)
}

/*
GETStripeGatewaysStripeGatewayId Retrieve a stripe gateway

Retrieve a stripe gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stripeGatewayId The resource's id
	@return StripeGatewaysApiGETStripeGatewaysStripeGatewayIdRequest
*/
func (a *StripeGatewaysApiService) GETStripeGatewaysStripeGatewayId(ctx context.Context, stripeGatewayId interface{}) StripeGatewaysApiGETStripeGatewaysStripeGatewayIdRequest {
	return StripeGatewaysApiGETStripeGatewaysStripeGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stripeGatewayId: stripeGatewayId,
	}
}

// Execute executes the request
//
//	@return GETStripeGatewaysStripeGatewayId200Response
func (a *StripeGatewaysApiService) GETStripeGatewaysStripeGatewayIdExecute(r StripeGatewaysApiGETStripeGatewaysStripeGatewayIdRequest) (*GETStripeGatewaysStripeGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETStripeGatewaysStripeGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.GETStripeGatewaysStripeGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways/{stripeGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stripeGatewayId"+"}", url.PathEscape(parameterValueToString(r.stripeGatewayId, "stripeGatewayId")), -1)

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

type StripeGatewaysApiPATCHStripeGatewaysStripeGatewayIdRequest struct {
	ctx                 context.Context
	ApiService          *StripeGatewaysApiService
	stripeGatewayUpdate *StripeGatewayUpdate
	stripeGatewayId     interface{}
}

func (r StripeGatewaysApiPATCHStripeGatewaysStripeGatewayIdRequest) StripeGatewayUpdate(stripeGatewayUpdate StripeGatewayUpdate) StripeGatewaysApiPATCHStripeGatewaysStripeGatewayIdRequest {
	r.stripeGatewayUpdate = &stripeGatewayUpdate
	return r
}

func (r StripeGatewaysApiPATCHStripeGatewaysStripeGatewayIdRequest) Execute() (*PATCHStripeGatewaysStripeGatewayId200Response, *http.Response, error) {
	return r.ApiService.PATCHStripeGatewaysStripeGatewayIdExecute(r)
}

/*
PATCHStripeGatewaysStripeGatewayId Update a stripe gateway

Update a stripe gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stripeGatewayId The resource's id
	@return StripeGatewaysApiPATCHStripeGatewaysStripeGatewayIdRequest
*/
func (a *StripeGatewaysApiService) PATCHStripeGatewaysStripeGatewayId(ctx context.Context, stripeGatewayId interface{}) StripeGatewaysApiPATCHStripeGatewaysStripeGatewayIdRequest {
	return StripeGatewaysApiPATCHStripeGatewaysStripeGatewayIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stripeGatewayId: stripeGatewayId,
	}
}

// Execute executes the request
//
//	@return PATCHStripeGatewaysStripeGatewayId200Response
func (a *StripeGatewaysApiService) PATCHStripeGatewaysStripeGatewayIdExecute(r StripeGatewaysApiPATCHStripeGatewaysStripeGatewayIdRequest) (*PATCHStripeGatewaysStripeGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHStripeGatewaysStripeGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.PATCHStripeGatewaysStripeGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways/{stripeGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stripeGatewayId"+"}", url.PathEscape(parameterValueToString(r.stripeGatewayId, "stripeGatewayId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stripeGatewayUpdate == nil {
		return localVarReturnValue, nil, reportError("stripeGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.stripeGatewayUpdate
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

type StripeGatewaysApiPOSTStripeGatewaysRequest struct {
	ctx                 context.Context
	ApiService          *StripeGatewaysApiService
	stripeGatewayCreate *StripeGatewayCreate
}

func (r StripeGatewaysApiPOSTStripeGatewaysRequest) StripeGatewayCreate(stripeGatewayCreate StripeGatewayCreate) StripeGatewaysApiPOSTStripeGatewaysRequest {
	r.stripeGatewayCreate = &stripeGatewayCreate
	return r
}

func (r StripeGatewaysApiPOSTStripeGatewaysRequest) Execute() (*POSTStripeGateways201Response, *http.Response, error) {
	return r.ApiService.POSTStripeGatewaysExecute(r)
}

/*
POSTStripeGateways Create a stripe gateway

Create a stripe gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return StripeGatewaysApiPOSTStripeGatewaysRequest
*/
func (a *StripeGatewaysApiService) POSTStripeGateways(ctx context.Context) StripeGatewaysApiPOSTStripeGatewaysRequest {
	return StripeGatewaysApiPOSTStripeGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTStripeGateways201Response
func (a *StripeGatewaysApiService) POSTStripeGatewaysExecute(r StripeGatewaysApiPOSTStripeGatewaysRequest) (*POSTStripeGateways201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTStripeGateways201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripeGatewaysApiService.POSTStripeGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stripeGatewayCreate == nil {
		return localVarReturnValue, nil, reportError("stripeGatewayCreate is required and must be specified")
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
	localVarPostBody = r.stripeGatewayCreate
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
