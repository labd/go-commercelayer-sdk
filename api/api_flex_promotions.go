/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
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

// FlexPromotionsApiService FlexPromotionsApi service
type FlexPromotionsApiService service

type FlexPromotionsApiDELETEFlexPromotionsFlexPromotionIdRequest struct {
	ctx             context.Context
	ApiService      *FlexPromotionsApiService
	flexPromotionId interface{}
}

func (r FlexPromotionsApiDELETEFlexPromotionsFlexPromotionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEFlexPromotionsFlexPromotionIdExecute(r)
}

/*
DELETEFlexPromotionsFlexPromotionId Delete a flex promotion

Delete a flex promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param flexPromotionId The resource's id
	@return FlexPromotionsApiDELETEFlexPromotionsFlexPromotionIdRequest
*/
func (a *FlexPromotionsApiService) DELETEFlexPromotionsFlexPromotionId(ctx context.Context, flexPromotionId interface{}) FlexPromotionsApiDELETEFlexPromotionsFlexPromotionIdRequest {
	return FlexPromotionsApiDELETEFlexPromotionsFlexPromotionIdRequest{
		ApiService:      a,
		ctx:             ctx,
		flexPromotionId: flexPromotionId,
	}
}

// Execute executes the request
func (a *FlexPromotionsApiService) DELETEFlexPromotionsFlexPromotionIdExecute(r FlexPromotionsApiDELETEFlexPromotionsFlexPromotionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexPromotionsApiService.DELETEFlexPromotionsFlexPromotionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/flex_promotions/{flexPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"flexPromotionId"+"}", url.PathEscape(parameterValueToString(r.flexPromotionId, "flexPromotionId")), -1)

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

type FlexPromotionsApiGETFlexPromotionsRequest struct {
	ctx        context.Context
	ApiService *FlexPromotionsApiService
}

func (r FlexPromotionsApiGETFlexPromotionsRequest) Execute() (*GETFlexPromotions200Response, *http.Response, error) {
	return r.ApiService.GETFlexPromotionsExecute(r)
}

/*
GETFlexPromotions List all flex promotions

List all flex promotions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FlexPromotionsApiGETFlexPromotionsRequest
*/
func (a *FlexPromotionsApiService) GETFlexPromotions(ctx context.Context) FlexPromotionsApiGETFlexPromotionsRequest {
	return FlexPromotionsApiGETFlexPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETFlexPromotions200Response
func (a *FlexPromotionsApiService) GETFlexPromotionsExecute(r FlexPromotionsApiGETFlexPromotionsRequest) (*GETFlexPromotions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFlexPromotions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexPromotionsApiService.GETFlexPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/flex_promotions"

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

type FlexPromotionsApiGETFlexPromotionsFlexPromotionIdRequest struct {
	ctx             context.Context
	ApiService      *FlexPromotionsApiService
	flexPromotionId interface{}
}

func (r FlexPromotionsApiGETFlexPromotionsFlexPromotionIdRequest) Execute() (*GETFlexPromotionsFlexPromotionId200Response, *http.Response, error) {
	return r.ApiService.GETFlexPromotionsFlexPromotionIdExecute(r)
}

/*
GETFlexPromotionsFlexPromotionId Retrieve a flex promotion

Retrieve a flex promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param flexPromotionId The resource's id
	@return FlexPromotionsApiGETFlexPromotionsFlexPromotionIdRequest
*/
func (a *FlexPromotionsApiService) GETFlexPromotionsFlexPromotionId(ctx context.Context, flexPromotionId interface{}) FlexPromotionsApiGETFlexPromotionsFlexPromotionIdRequest {
	return FlexPromotionsApiGETFlexPromotionsFlexPromotionIdRequest{
		ApiService:      a,
		ctx:             ctx,
		flexPromotionId: flexPromotionId,
	}
}

// Execute executes the request
//
//	@return GETFlexPromotionsFlexPromotionId200Response
func (a *FlexPromotionsApiService) GETFlexPromotionsFlexPromotionIdExecute(r FlexPromotionsApiGETFlexPromotionsFlexPromotionIdRequest) (*GETFlexPromotionsFlexPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETFlexPromotionsFlexPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexPromotionsApiService.GETFlexPromotionsFlexPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/flex_promotions/{flexPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"flexPromotionId"+"}", url.PathEscape(parameterValueToString(r.flexPromotionId, "flexPromotionId")), -1)

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

type FlexPromotionsApiPATCHFlexPromotionsFlexPromotionIdRequest struct {
	ctx                 context.Context
	ApiService          *FlexPromotionsApiService
	flexPromotionUpdate *FlexPromotionUpdate
	flexPromotionId     interface{}
}

func (r FlexPromotionsApiPATCHFlexPromotionsFlexPromotionIdRequest) FlexPromotionUpdate(flexPromotionUpdate FlexPromotionUpdate) FlexPromotionsApiPATCHFlexPromotionsFlexPromotionIdRequest {
	r.flexPromotionUpdate = &flexPromotionUpdate
	return r
}

func (r FlexPromotionsApiPATCHFlexPromotionsFlexPromotionIdRequest) Execute() (*PATCHFlexPromotionsFlexPromotionId200Response, *http.Response, error) {
	return r.ApiService.PATCHFlexPromotionsFlexPromotionIdExecute(r)
}

/*
PATCHFlexPromotionsFlexPromotionId Update a flex promotion

Update a flex promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param flexPromotionId The resource's id
	@return FlexPromotionsApiPATCHFlexPromotionsFlexPromotionIdRequest
*/
func (a *FlexPromotionsApiService) PATCHFlexPromotionsFlexPromotionId(ctx context.Context, flexPromotionId interface{}) FlexPromotionsApiPATCHFlexPromotionsFlexPromotionIdRequest {
	return FlexPromotionsApiPATCHFlexPromotionsFlexPromotionIdRequest{
		ApiService:      a,
		ctx:             ctx,
		flexPromotionId: flexPromotionId,
	}
}

// Execute executes the request
//
//	@return PATCHFlexPromotionsFlexPromotionId200Response
func (a *FlexPromotionsApiService) PATCHFlexPromotionsFlexPromotionIdExecute(r FlexPromotionsApiPATCHFlexPromotionsFlexPromotionIdRequest) (*PATCHFlexPromotionsFlexPromotionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHFlexPromotionsFlexPromotionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexPromotionsApiService.PATCHFlexPromotionsFlexPromotionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/flex_promotions/{flexPromotionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"flexPromotionId"+"}", url.PathEscape(parameterValueToString(r.flexPromotionId, "flexPromotionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flexPromotionUpdate == nil {
		return localVarReturnValue, nil, reportError("flexPromotionUpdate is required and must be specified")
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
	localVarPostBody = r.flexPromotionUpdate
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

type FlexPromotionsApiPOSTFlexPromotionsRequest struct {
	ctx                 context.Context
	ApiService          *FlexPromotionsApiService
	flexPromotionCreate *FlexPromotionCreate
}

func (r FlexPromotionsApiPOSTFlexPromotionsRequest) FlexPromotionCreate(flexPromotionCreate FlexPromotionCreate) FlexPromotionsApiPOSTFlexPromotionsRequest {
	r.flexPromotionCreate = &flexPromotionCreate
	return r
}

func (r FlexPromotionsApiPOSTFlexPromotionsRequest) Execute() (*POSTFlexPromotions201Response, *http.Response, error) {
	return r.ApiService.POSTFlexPromotionsExecute(r)
}

/*
POSTFlexPromotions Create a flex promotion

Create a flex promotion

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FlexPromotionsApiPOSTFlexPromotionsRequest
*/
func (a *FlexPromotionsApiService) POSTFlexPromotions(ctx context.Context) FlexPromotionsApiPOSTFlexPromotionsRequest {
	return FlexPromotionsApiPOSTFlexPromotionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTFlexPromotions201Response
func (a *FlexPromotionsApiService) POSTFlexPromotionsExecute(r FlexPromotionsApiPOSTFlexPromotionsRequest) (*POSTFlexPromotions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTFlexPromotions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FlexPromotionsApiService.POSTFlexPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/flex_promotions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.flexPromotionCreate == nil {
		return localVarReturnValue, nil, reportError("flexPromotionCreate is required and must be specified")
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
	localVarPostBody = r.flexPromotionCreate
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
