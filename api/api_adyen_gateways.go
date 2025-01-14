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

// AdyenGatewaysApiService AdyenGatewaysApi service
type AdyenGatewaysApiService service

type AdyenGatewaysApiDELETEAdyenGatewaysAdyenGatewayIdRequest struct {
	ctx            context.Context
	ApiService     *AdyenGatewaysApiService
	adyenGatewayId interface{}
}

func (r AdyenGatewaysApiDELETEAdyenGatewaysAdyenGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEAdyenGatewaysAdyenGatewayIdExecute(r)
}

/*
DELETEAdyenGatewaysAdyenGatewayId Delete an adyen gateway

Delete an adyen gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param adyenGatewayId The resource's id
	@return AdyenGatewaysApiDELETEAdyenGatewaysAdyenGatewayIdRequest
*/
func (a *AdyenGatewaysApiService) DELETEAdyenGatewaysAdyenGatewayId(ctx context.Context, adyenGatewayId interface{}) AdyenGatewaysApiDELETEAdyenGatewaysAdyenGatewayIdRequest {
	return AdyenGatewaysApiDELETEAdyenGatewaysAdyenGatewayIdRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenGatewayId: adyenGatewayId,
	}
}

// Execute executes the request
func (a *AdyenGatewaysApiService) DELETEAdyenGatewaysAdyenGatewayIdExecute(r AdyenGatewaysApiDELETEAdyenGatewaysAdyenGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenGatewaysApiService.DELETEAdyenGatewaysAdyenGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_gateways/{adyenGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenGatewayId"+"}", url.PathEscape(parameterValueToString(r.adyenGatewayId, "adyenGatewayId")), -1)

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

type AdyenGatewaysApiGETAdyenGatewaysRequest struct {
	ctx        context.Context
	ApiService *AdyenGatewaysApiService
}

func (r AdyenGatewaysApiGETAdyenGatewaysRequest) Execute() (*GETAdyenGateways200Response, *http.Response, error) {
	return r.ApiService.GETAdyenGatewaysExecute(r)
}

/*
GETAdyenGateways List all adyen gateways

List all adyen gateways

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AdyenGatewaysApiGETAdyenGatewaysRequest
*/
func (a *AdyenGatewaysApiService) GETAdyenGateways(ctx context.Context) AdyenGatewaysApiGETAdyenGatewaysRequest {
	return AdyenGatewaysApiGETAdyenGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETAdyenGateways200Response
func (a *AdyenGatewaysApiService) GETAdyenGatewaysExecute(r AdyenGatewaysApiGETAdyenGatewaysRequest) (*GETAdyenGateways200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAdyenGateways200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenGatewaysApiService.GETAdyenGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_gateways"

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

type AdyenGatewaysApiGETAdyenGatewaysAdyenGatewayIdRequest struct {
	ctx            context.Context
	ApiService     *AdyenGatewaysApiService
	adyenGatewayId interface{}
}

func (r AdyenGatewaysApiGETAdyenGatewaysAdyenGatewayIdRequest) Execute() (*GETAdyenGatewaysAdyenGatewayId200Response, *http.Response, error) {
	return r.ApiService.GETAdyenGatewaysAdyenGatewayIdExecute(r)
}

/*
GETAdyenGatewaysAdyenGatewayId Retrieve an adyen gateway

Retrieve an adyen gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param adyenGatewayId The resource's id
	@return AdyenGatewaysApiGETAdyenGatewaysAdyenGatewayIdRequest
*/
func (a *AdyenGatewaysApiService) GETAdyenGatewaysAdyenGatewayId(ctx context.Context, adyenGatewayId interface{}) AdyenGatewaysApiGETAdyenGatewaysAdyenGatewayIdRequest {
	return AdyenGatewaysApiGETAdyenGatewaysAdyenGatewayIdRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenGatewayId: adyenGatewayId,
	}
}

// Execute executes the request
//
//	@return GETAdyenGatewaysAdyenGatewayId200Response
func (a *AdyenGatewaysApiService) GETAdyenGatewaysAdyenGatewayIdExecute(r AdyenGatewaysApiGETAdyenGatewaysAdyenGatewayIdRequest) (*GETAdyenGatewaysAdyenGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAdyenGatewaysAdyenGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenGatewaysApiService.GETAdyenGatewaysAdyenGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_gateways/{adyenGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenGatewayId"+"}", url.PathEscape(parameterValueToString(r.adyenGatewayId, "adyenGatewayId")), -1)

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

type AdyenGatewaysApiPATCHAdyenGatewaysAdyenGatewayIdRequest struct {
	ctx                context.Context
	ApiService         *AdyenGatewaysApiService
	adyenGatewayUpdate *AdyenGatewayUpdate
	adyenGatewayId     interface{}
}

func (r AdyenGatewaysApiPATCHAdyenGatewaysAdyenGatewayIdRequest) AdyenGatewayUpdate(adyenGatewayUpdate AdyenGatewayUpdate) AdyenGatewaysApiPATCHAdyenGatewaysAdyenGatewayIdRequest {
	r.adyenGatewayUpdate = &adyenGatewayUpdate
	return r
}

func (r AdyenGatewaysApiPATCHAdyenGatewaysAdyenGatewayIdRequest) Execute() (*PATCHAdyenGatewaysAdyenGatewayId200Response, *http.Response, error) {
	return r.ApiService.PATCHAdyenGatewaysAdyenGatewayIdExecute(r)
}

/*
PATCHAdyenGatewaysAdyenGatewayId Update an adyen gateway

Update an adyen gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param adyenGatewayId The resource's id
	@return AdyenGatewaysApiPATCHAdyenGatewaysAdyenGatewayIdRequest
*/
func (a *AdyenGatewaysApiService) PATCHAdyenGatewaysAdyenGatewayId(ctx context.Context, adyenGatewayId interface{}) AdyenGatewaysApiPATCHAdyenGatewaysAdyenGatewayIdRequest {
	return AdyenGatewaysApiPATCHAdyenGatewaysAdyenGatewayIdRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenGatewayId: adyenGatewayId,
	}
}

// Execute executes the request
//
//	@return PATCHAdyenGatewaysAdyenGatewayId200Response
func (a *AdyenGatewaysApiService) PATCHAdyenGatewaysAdyenGatewayIdExecute(r AdyenGatewaysApiPATCHAdyenGatewaysAdyenGatewayIdRequest) (*PATCHAdyenGatewaysAdyenGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHAdyenGatewaysAdyenGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenGatewaysApiService.PATCHAdyenGatewaysAdyenGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_gateways/{adyenGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenGatewayId"+"}", url.PathEscape(parameterValueToString(r.adyenGatewayId, "adyenGatewayId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adyenGatewayUpdate == nil {
		return localVarReturnValue, nil, reportError("adyenGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.adyenGatewayUpdate
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

type AdyenGatewaysApiPOSTAdyenGatewaysRequest struct {
	ctx                context.Context
	ApiService         *AdyenGatewaysApiService
	adyenGatewayCreate *AdyenGatewayCreate
}

func (r AdyenGatewaysApiPOSTAdyenGatewaysRequest) AdyenGatewayCreate(adyenGatewayCreate AdyenGatewayCreate) AdyenGatewaysApiPOSTAdyenGatewaysRequest {
	r.adyenGatewayCreate = &adyenGatewayCreate
	return r
}

func (r AdyenGatewaysApiPOSTAdyenGatewaysRequest) Execute() (*POSTAdyenGateways201Response, *http.Response, error) {
	return r.ApiService.POSTAdyenGatewaysExecute(r)
}

/*
POSTAdyenGateways Create an adyen gateway

Create an adyen gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AdyenGatewaysApiPOSTAdyenGatewaysRequest
*/
func (a *AdyenGatewaysApiService) POSTAdyenGateways(ctx context.Context) AdyenGatewaysApiPOSTAdyenGatewaysRequest {
	return AdyenGatewaysApiPOSTAdyenGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTAdyenGateways201Response
func (a *AdyenGatewaysApiService) POSTAdyenGatewaysExecute(r AdyenGatewaysApiPOSTAdyenGatewaysRequest) (*POSTAdyenGateways201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTAdyenGateways201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenGatewaysApiService.POSTAdyenGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adyenGatewayCreate == nil {
		return localVarReturnValue, nil, reportError("adyenGatewayCreate is required and must be specified")
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
	localVarPostBody = r.adyenGatewayCreate
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
