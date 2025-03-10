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

// FixedPricePromotionsApiService FixedPricePromotionsApi service
type FixedPricePromotionsApiService service

type FixedPricePromotionsApiDELETEFixedPricePromotionsFixedPricePromotionIdRequest struct {
	ctx                   context.Context
	ApiService            *FixedPricePromotionsApiService
	fixedPricePromotionId interface{}
}

func (r FixedPricePromotionsApiDELETEFixedPricePromotionsFixedPricePromotionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEFixedPricePromotionsFixedPricePromotionIdExecute(r)
}

/*
DELETEFixedPricePromotionsFixedPricePromotionId Delete a fixed price promotion

Delete a fixed price promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fixedPricePromotionId The resource's id
	@return FixedPricePromotionsApiDELETEFixedPricePromotionsFixedPricePromotionIdRequest
*/
func (a *FixedPricePromotionsApiService) DELETEFixedPricePromotionsFixedPricePromotionId(ctx context.Context, fixedPricePromotionId interface{}) FixedPricePromotionsApiDELETEFixedPricePromotionsFixedPricePromotionIdRequest {
	return FixedPricePromotionsApiDELETEFixedPricePromotionsFixedPricePromotionIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		fixedPricePromotionId: fixedPricePromotionId,
	}
}

// Execute executes the request
func (a *FixedPricePromotionsApiService) DELETEFixedPricePromotionsFixedPricePromotionIdExecute(r FixedPricePromotionsApiDELETEFixedPricePromotionsFixedPricePromotionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedPricePromotionsApiService.DELETEFixedPricePromotionsFixedPricePromotionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_price_promotions/{fixedPricePromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fixedPricePromotionId"+"}", url.PathEscape(parameterValueToString(r.fixedPricePromotionId, "fixedPricePromotionId")), -1)

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

type FixedPricePromotionsApiGETFixedPricePromotionsRequest struct {
	ctx        context.Context
	ApiService *FixedPricePromotionsApiService
}

func (r FixedPricePromotionsApiGETFixedPricePromotionsRequest) Execute() (*GETFixedPricePromotions200Response, *http.Response, error) {
	return r.ApiService.GETFixedPricePromotionsExecute(r)
}

/*
GETFixedPricePromotions List all fixed price promotions

List all fixed price promotions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FixedPricePromotionsApiGETFixedPricePromotionsRequest
*/
func (a *FixedPricePromotionsApiService) GETFixedPricePromotions(ctx context.Context) FixedPricePromotionsApiGETFixedPricePromotionsRequest {
	return FixedPricePromotionsApiGETFixedPricePromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETFixedPricePromotions200Response
func (a *FixedPricePromotionsApiService) GETFixedPricePromotionsExecute(r FixedPricePromotionsApiGETFixedPricePromotionsRequest) (*GETFixedPricePromotions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFixedPricePromotions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedPricePromotionsApiService.GETFixedPricePromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_price_promotions"

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

type FixedPricePromotionsApiGETFixedPricePromotionsFixedPricePromotionIdRequest struct {
	ctx                   context.Context
	ApiService            *FixedPricePromotionsApiService
	fixedPricePromotionId interface{}
}

func (r FixedPricePromotionsApiGETFixedPricePromotionsFixedPricePromotionIdRequest) Execute() (*GETFixedPricePromotionsFixedPricePromotionId200Response, *http.Response, error) {
	return r.ApiService.GETFixedPricePromotionsFixedPricePromotionIdExecute(r)
}

/*
GETFixedPricePromotionsFixedPricePromotionId Retrieve a fixed price promotion

Retrieve a fixed price promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fixedPricePromotionId The resource's id
	@return FixedPricePromotionsApiGETFixedPricePromotionsFixedPricePromotionIdRequest
*/
func (a *FixedPricePromotionsApiService) GETFixedPricePromotionsFixedPricePromotionId(ctx context.Context, fixedPricePromotionId interface{}) FixedPricePromotionsApiGETFixedPricePromotionsFixedPricePromotionIdRequest {
	return FixedPricePromotionsApiGETFixedPricePromotionsFixedPricePromotionIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		fixedPricePromotionId: fixedPricePromotionId,
	}
}

// Execute executes the request
//
//	@return GETFixedPricePromotionsFixedPricePromotionId200Response
func (a *FixedPricePromotionsApiService) GETFixedPricePromotionsFixedPricePromotionIdExecute(r FixedPricePromotionsApiGETFixedPricePromotionsFixedPricePromotionIdRequest) (*GETFixedPricePromotionsFixedPricePromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFixedPricePromotionsFixedPricePromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedPricePromotionsApiService.GETFixedPricePromotionsFixedPricePromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_price_promotions/{fixedPricePromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fixedPricePromotionId"+"}", url.PathEscape(parameterValueToString(r.fixedPricePromotionId, "fixedPricePromotionId")), -1)

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

type FixedPricePromotionsApiPATCHFixedPricePromotionsFixedPricePromotionIdRequest struct {
	ctx                       context.Context
	ApiService                *FixedPricePromotionsApiService
	fixedPricePromotionUpdate *FixedPricePromotionUpdate
	fixedPricePromotionId     interface{}
}

func (r FixedPricePromotionsApiPATCHFixedPricePromotionsFixedPricePromotionIdRequest) FixedPricePromotionUpdate(fixedPricePromotionUpdate FixedPricePromotionUpdate) FixedPricePromotionsApiPATCHFixedPricePromotionsFixedPricePromotionIdRequest {
	r.fixedPricePromotionUpdate = &fixedPricePromotionUpdate
	return r
}

func (r FixedPricePromotionsApiPATCHFixedPricePromotionsFixedPricePromotionIdRequest) Execute() (*PATCHFixedPricePromotionsFixedPricePromotionId200Response, *http.Response, error) {
	return r.ApiService.PATCHFixedPricePromotionsFixedPricePromotionIdExecute(r)
}

/*
PATCHFixedPricePromotionsFixedPricePromotionId Update a fixed price promotion

Update a fixed price promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param fixedPricePromotionId The resource's id
	@return FixedPricePromotionsApiPATCHFixedPricePromotionsFixedPricePromotionIdRequest
*/
func (a *FixedPricePromotionsApiService) PATCHFixedPricePromotionsFixedPricePromotionId(ctx context.Context, fixedPricePromotionId interface{}) FixedPricePromotionsApiPATCHFixedPricePromotionsFixedPricePromotionIdRequest {
	return FixedPricePromotionsApiPATCHFixedPricePromotionsFixedPricePromotionIdRequest{
		ApiService:            a,
		ctx:                   ctx,
		fixedPricePromotionId: fixedPricePromotionId,
	}
}

// Execute executes the request
//
//	@return PATCHFixedPricePromotionsFixedPricePromotionId200Response
func (a *FixedPricePromotionsApiService) PATCHFixedPricePromotionsFixedPricePromotionIdExecute(r FixedPricePromotionsApiPATCHFixedPricePromotionsFixedPricePromotionIdRequest) (*PATCHFixedPricePromotionsFixedPricePromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHFixedPricePromotionsFixedPricePromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedPricePromotionsApiService.PATCHFixedPricePromotionsFixedPricePromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_price_promotions/{fixedPricePromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fixedPricePromotionId"+"}", url.PathEscape(parameterValueToString(r.fixedPricePromotionId, "fixedPricePromotionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fixedPricePromotionUpdate == nil {
		return localVarReturnValue, nil, reportError("fixedPricePromotionUpdate is required and must be specified")
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
	localVarPostBody = r.fixedPricePromotionUpdate
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

type FixedPricePromotionsApiPOSTFixedPricePromotionsRequest struct {
	ctx                       context.Context
	ApiService                *FixedPricePromotionsApiService
	fixedPricePromotionCreate *FixedPricePromotionCreate
}

func (r FixedPricePromotionsApiPOSTFixedPricePromotionsRequest) FixedPricePromotionCreate(fixedPricePromotionCreate FixedPricePromotionCreate) FixedPricePromotionsApiPOSTFixedPricePromotionsRequest {
	r.fixedPricePromotionCreate = &fixedPricePromotionCreate
	return r
}

func (r FixedPricePromotionsApiPOSTFixedPricePromotionsRequest) Execute() (*POSTFixedPricePromotions201Response, *http.Response, error) {
	return r.ApiService.POSTFixedPricePromotionsExecute(r)
}

/*
POSTFixedPricePromotions Create a fixed price promotion

Create a fixed price promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FixedPricePromotionsApiPOSTFixedPricePromotionsRequest
*/
func (a *FixedPricePromotionsApiService) POSTFixedPricePromotions(ctx context.Context) FixedPricePromotionsApiPOSTFixedPricePromotionsRequest {
	return FixedPricePromotionsApiPOSTFixedPricePromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTFixedPricePromotions201Response
func (a *FixedPricePromotionsApiService) POSTFixedPricePromotionsExecute(r FixedPricePromotionsApiPOSTFixedPricePromotionsRequest) (*POSTFixedPricePromotions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTFixedPricePromotions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FixedPricePromotionsApiService.POSTFixedPricePromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fixed_price_promotions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fixedPricePromotionCreate == nil {
		return localVarReturnValue, nil, reportError("fixedPricePromotionCreate is required and must be specified")
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
	localVarPostBody = r.fixedPricePromotionCreate
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
