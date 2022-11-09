/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
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

// BraintreeGatewaysApiService BraintreeGatewaysApi service
type BraintreeGatewaysApiService service

type BraintreeGatewaysApiDELETEBraintreeGatewaysBraintreeGatewayIdRequest struct {
	ctx                context.Context
	ApiService         *BraintreeGatewaysApiService
	braintreeGatewayId string
}

func (r BraintreeGatewaysApiDELETEBraintreeGatewaysBraintreeGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEBraintreeGatewaysBraintreeGatewayIdExecute(r)
}

/*
DELETEBraintreeGatewaysBraintreeGatewayId Delete a braintree gateway

Delete a braintree gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param braintreeGatewayId The resource's id
 @return BraintreeGatewaysApiDELETEBraintreeGatewaysBraintreeGatewayIdRequest
*/
func (a *BraintreeGatewaysApiService) DELETEBraintreeGatewaysBraintreeGatewayId(ctx context.Context, braintreeGatewayId string) BraintreeGatewaysApiDELETEBraintreeGatewaysBraintreeGatewayIdRequest {
	return BraintreeGatewaysApiDELETEBraintreeGatewaysBraintreeGatewayIdRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreeGatewayId: braintreeGatewayId,
	}
}

// Execute executes the request
func (a *BraintreeGatewaysApiService) DELETEBraintreeGatewaysBraintreeGatewayIdExecute(r BraintreeGatewaysApiDELETEBraintreeGatewaysBraintreeGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreeGatewaysApiService.DELETEBraintreeGatewaysBraintreeGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_gateways/{braintreeGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreeGatewayId"+"}", url.PathEscape(parameterToString(r.braintreeGatewayId, "")), -1)

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

type BraintreeGatewaysApiGETBraintreeGatewaysRequest struct {
	ctx        context.Context
	ApiService *BraintreeGatewaysApiService
}

func (r BraintreeGatewaysApiGETBraintreeGatewaysRequest) Execute() (*GETBraintreeGateways200Response, *http.Response, error) {
	return r.ApiService.GETBraintreeGatewaysExecute(r)
}

/*
GETBraintreeGateways List all braintree gateways

List all braintree gateways

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BraintreeGatewaysApiGETBraintreeGatewaysRequest
*/
func (a *BraintreeGatewaysApiService) GETBraintreeGateways(ctx context.Context) BraintreeGatewaysApiGETBraintreeGatewaysRequest {
	return BraintreeGatewaysApiGETBraintreeGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETBraintreeGateways200Response
func (a *BraintreeGatewaysApiService) GETBraintreeGatewaysExecute(r BraintreeGatewaysApiGETBraintreeGatewaysRequest) (*GETBraintreeGateways200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBraintreeGateways200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreeGatewaysApiService.GETBraintreeGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_gateways"

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

type BraintreeGatewaysApiGETBraintreeGatewaysBraintreeGatewayIdRequest struct {
	ctx                context.Context
	ApiService         *BraintreeGatewaysApiService
	braintreeGatewayId string
}

func (r BraintreeGatewaysApiGETBraintreeGatewaysBraintreeGatewayIdRequest) Execute() (*GETBraintreeGatewaysBraintreeGatewayId200Response, *http.Response, error) {
	return r.ApiService.GETBraintreeGatewaysBraintreeGatewayIdExecute(r)
}

/*
GETBraintreeGatewaysBraintreeGatewayId Retrieve a braintree gateway

Retrieve a braintree gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param braintreeGatewayId The resource's id
 @return BraintreeGatewaysApiGETBraintreeGatewaysBraintreeGatewayIdRequest
*/
func (a *BraintreeGatewaysApiService) GETBraintreeGatewaysBraintreeGatewayId(ctx context.Context, braintreeGatewayId string) BraintreeGatewaysApiGETBraintreeGatewaysBraintreeGatewayIdRequest {
	return BraintreeGatewaysApiGETBraintreeGatewaysBraintreeGatewayIdRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreeGatewayId: braintreeGatewayId,
	}
}

// Execute executes the request
//  @return GETBraintreeGatewaysBraintreeGatewayId200Response
func (a *BraintreeGatewaysApiService) GETBraintreeGatewaysBraintreeGatewayIdExecute(r BraintreeGatewaysApiGETBraintreeGatewaysBraintreeGatewayIdRequest) (*GETBraintreeGatewaysBraintreeGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBraintreeGatewaysBraintreeGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreeGatewaysApiService.GETBraintreeGatewaysBraintreeGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_gateways/{braintreeGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreeGatewayId"+"}", url.PathEscape(parameterToString(r.braintreeGatewayId, "")), -1)

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

type BraintreeGatewaysApiPATCHBraintreeGatewaysBraintreeGatewayIdRequest struct {
	ctx                    context.Context
	ApiService             *BraintreeGatewaysApiService
	braintreeGatewayUpdate *BraintreeGatewayUpdate
	braintreeGatewayId     string
}

func (r BraintreeGatewaysApiPATCHBraintreeGatewaysBraintreeGatewayIdRequest) BraintreeGatewayUpdate(braintreeGatewayUpdate BraintreeGatewayUpdate) BraintreeGatewaysApiPATCHBraintreeGatewaysBraintreeGatewayIdRequest {
	r.braintreeGatewayUpdate = &braintreeGatewayUpdate
	return r
}

func (r BraintreeGatewaysApiPATCHBraintreeGatewaysBraintreeGatewayIdRequest) Execute() (*PATCHBraintreeGatewaysBraintreeGatewayId200Response, *http.Response, error) {
	return r.ApiService.PATCHBraintreeGatewaysBraintreeGatewayIdExecute(r)
}

/*
PATCHBraintreeGatewaysBraintreeGatewayId Update a braintree gateway

Update a braintree gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param braintreeGatewayId The resource's id
 @return BraintreeGatewaysApiPATCHBraintreeGatewaysBraintreeGatewayIdRequest
*/
func (a *BraintreeGatewaysApiService) PATCHBraintreeGatewaysBraintreeGatewayId(ctx context.Context, braintreeGatewayId string) BraintreeGatewaysApiPATCHBraintreeGatewaysBraintreeGatewayIdRequest {
	return BraintreeGatewaysApiPATCHBraintreeGatewaysBraintreeGatewayIdRequest{
		ApiService:         a,
		ctx:                ctx,
		braintreeGatewayId: braintreeGatewayId,
	}
}

// Execute executes the request
//  @return PATCHBraintreeGatewaysBraintreeGatewayId200Response
func (a *BraintreeGatewaysApiService) PATCHBraintreeGatewaysBraintreeGatewayIdExecute(r BraintreeGatewaysApiPATCHBraintreeGatewaysBraintreeGatewayIdRequest) (*PATCHBraintreeGatewaysBraintreeGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHBraintreeGatewaysBraintreeGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreeGatewaysApiService.PATCHBraintreeGatewaysBraintreeGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_gateways/{braintreeGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"braintreeGatewayId"+"}", url.PathEscape(parameterToString(r.braintreeGatewayId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.braintreeGatewayUpdate == nil {
		return localVarReturnValue, nil, reportError("braintreeGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.braintreeGatewayUpdate
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

type BraintreeGatewaysApiPOSTBraintreeGatewaysRequest struct {
	ctx                    context.Context
	ApiService             *BraintreeGatewaysApiService
	braintreeGatewayCreate *BraintreeGatewayCreate
}

func (r BraintreeGatewaysApiPOSTBraintreeGatewaysRequest) BraintreeGatewayCreate(braintreeGatewayCreate BraintreeGatewayCreate) BraintreeGatewaysApiPOSTBraintreeGatewaysRequest {
	r.braintreeGatewayCreate = &braintreeGatewayCreate
	return r
}

func (r BraintreeGatewaysApiPOSTBraintreeGatewaysRequest) Execute() (*POSTBraintreeGateways201Response, *http.Response, error) {
	return r.ApiService.POSTBraintreeGatewaysExecute(r)
}

/*
POSTBraintreeGateways Create a braintree gateway

Create a braintree gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BraintreeGatewaysApiPOSTBraintreeGatewaysRequest
*/
func (a *BraintreeGatewaysApiService) POSTBraintreeGateways(ctx context.Context) BraintreeGatewaysApiPOSTBraintreeGatewaysRequest {
	return BraintreeGatewaysApiPOSTBraintreeGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTBraintreeGateways201Response
func (a *BraintreeGatewaysApiService) POSTBraintreeGatewaysExecute(r BraintreeGatewaysApiPOSTBraintreeGatewaysRequest) (*POSTBraintreeGateways201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTBraintreeGateways201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BraintreeGatewaysApiService.POSTBraintreeGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/braintree_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.braintreeGatewayCreate == nil {
		return localVarReturnValue, nil, reportError("braintreeGatewayCreate is required and must be specified")
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
	localVarPostBody = r.braintreeGatewayCreate
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
