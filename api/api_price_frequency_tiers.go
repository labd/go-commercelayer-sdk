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

// PriceFrequencyTiersApiService PriceFrequencyTiersApi service
type PriceFrequencyTiersApiService service

type PriceFrequencyTiersApiDELETEPriceFrequencyTiersPriceFrequencyTierIdRequest struct {
	ctx                  context.Context
	ApiService           *PriceFrequencyTiersApiService
	priceFrequencyTierId interface{}
}

func (r PriceFrequencyTiersApiDELETEPriceFrequencyTiersPriceFrequencyTierIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEPriceFrequencyTiersPriceFrequencyTierIdExecute(r)
}

/*
DELETEPriceFrequencyTiersPriceFrequencyTierId Delete a price frequency tier

Delete a price frequency tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceFrequencyTierId The resource's id
	@return PriceFrequencyTiersApiDELETEPriceFrequencyTiersPriceFrequencyTierIdRequest
*/
func (a *PriceFrequencyTiersApiService) DELETEPriceFrequencyTiersPriceFrequencyTierId(ctx context.Context, priceFrequencyTierId interface{}) PriceFrequencyTiersApiDELETEPriceFrequencyTiersPriceFrequencyTierIdRequest {
	return PriceFrequencyTiersApiDELETEPriceFrequencyTiersPriceFrequencyTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		priceFrequencyTierId: priceFrequencyTierId,
	}
}

// Execute executes the request
func (a *PriceFrequencyTiersApiService) DELETEPriceFrequencyTiersPriceFrequencyTierIdExecute(r PriceFrequencyTiersApiDELETEPriceFrequencyTiersPriceFrequencyTierIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceFrequencyTiersApiService.DELETEPriceFrequencyTiersPriceFrequencyTierId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_frequency_tiers/{priceFrequencyTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceFrequencyTierId"+"}", url.PathEscape(parameterValueToString(r.priceFrequencyTierId, "priceFrequencyTierId")), -1)

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

type PriceFrequencyTiersApiGETPriceFrequencyTiersRequest struct {
	ctx        context.Context
	ApiService *PriceFrequencyTiersApiService
}

func (r PriceFrequencyTiersApiGETPriceFrequencyTiersRequest) Execute() (*GETPriceFrequencyTiers200Response, *http.Response, error) {
	return r.ApiService.GETPriceFrequencyTiersExecute(r)
}

/*
GETPriceFrequencyTiers List all price frequency tiers

List all price frequency tiers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PriceFrequencyTiersApiGETPriceFrequencyTiersRequest
*/
func (a *PriceFrequencyTiersApiService) GETPriceFrequencyTiers(ctx context.Context) PriceFrequencyTiersApiGETPriceFrequencyTiersRequest {
	return PriceFrequencyTiersApiGETPriceFrequencyTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETPriceFrequencyTiers200Response
func (a *PriceFrequencyTiersApiService) GETPriceFrequencyTiersExecute(r PriceFrequencyTiersApiGETPriceFrequencyTiersRequest) (*GETPriceFrequencyTiers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceFrequencyTiers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceFrequencyTiersApiService.GETPriceFrequencyTiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_frequency_tiers"

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

type PriceFrequencyTiersApiGETPriceFrequencyTiersPriceFrequencyTierIdRequest struct {
	ctx                  context.Context
	ApiService           *PriceFrequencyTiersApiService
	priceFrequencyTierId interface{}
}

func (r PriceFrequencyTiersApiGETPriceFrequencyTiersPriceFrequencyTierIdRequest) Execute() (*GETPriceFrequencyTiersPriceFrequencyTierId200Response, *http.Response, error) {
	return r.ApiService.GETPriceFrequencyTiersPriceFrequencyTierIdExecute(r)
}

/*
GETPriceFrequencyTiersPriceFrequencyTierId Retrieve a price frequency tier

Retrieve a price frequency tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceFrequencyTierId The resource's id
	@return PriceFrequencyTiersApiGETPriceFrequencyTiersPriceFrequencyTierIdRequest
*/
func (a *PriceFrequencyTiersApiService) GETPriceFrequencyTiersPriceFrequencyTierId(ctx context.Context, priceFrequencyTierId interface{}) PriceFrequencyTiersApiGETPriceFrequencyTiersPriceFrequencyTierIdRequest {
	return PriceFrequencyTiersApiGETPriceFrequencyTiersPriceFrequencyTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		priceFrequencyTierId: priceFrequencyTierId,
	}
}

// Execute executes the request
//
//	@return GETPriceFrequencyTiersPriceFrequencyTierId200Response
func (a *PriceFrequencyTiersApiService) GETPriceFrequencyTiersPriceFrequencyTierIdExecute(r PriceFrequencyTiersApiGETPriceFrequencyTiersPriceFrequencyTierIdRequest) (*GETPriceFrequencyTiersPriceFrequencyTierId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETPriceFrequencyTiersPriceFrequencyTierId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceFrequencyTiersApiService.GETPriceFrequencyTiersPriceFrequencyTierId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_frequency_tiers/{priceFrequencyTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceFrequencyTierId"+"}", url.PathEscape(parameterValueToString(r.priceFrequencyTierId, "priceFrequencyTierId")), -1)

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

type PriceFrequencyTiersApiGETPriceIdPriceFrequencyTiersRequest struct {
	ctx        context.Context
	ApiService *PriceFrequencyTiersApiService
	priceId    interface{}
}

func (r PriceFrequencyTiersApiGETPriceIdPriceFrequencyTiersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPriceIdPriceFrequencyTiersExecute(r)
}

/*
GETPriceIdPriceFrequencyTiers Retrieve the price frequency tiers associated to the price

Retrieve the price frequency tiers associated to the price

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceId The resource's id
	@return PriceFrequencyTiersApiGETPriceIdPriceFrequencyTiersRequest
*/
func (a *PriceFrequencyTiersApiService) GETPriceIdPriceFrequencyTiers(ctx context.Context, priceId interface{}) PriceFrequencyTiersApiGETPriceIdPriceFrequencyTiersRequest {
	return PriceFrequencyTiersApiGETPriceIdPriceFrequencyTiersRequest{
		ApiService: a,
		ctx:        ctx,
		priceId:    priceId,
	}
}

// Execute executes the request
func (a *PriceFrequencyTiersApiService) GETPriceIdPriceFrequencyTiersExecute(r PriceFrequencyTiersApiGETPriceIdPriceFrequencyTiersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceFrequencyTiersApiService.GETPriceIdPriceFrequencyTiers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prices/{priceId}/price_frequency_tiers"
	localVarPath = strings.Replace(localVarPath, "{"+"priceId"+"}", url.PathEscape(parameterValueToString(r.priceId, "priceId")), -1)

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

type PriceFrequencyTiersApiPATCHPriceFrequencyTiersPriceFrequencyTierIdRequest struct {
	ctx                      context.Context
	ApiService               *PriceFrequencyTiersApiService
	priceFrequencyTierUpdate *PriceFrequencyTierUpdate
	priceFrequencyTierId     interface{}
}

func (r PriceFrequencyTiersApiPATCHPriceFrequencyTiersPriceFrequencyTierIdRequest) PriceFrequencyTierUpdate(priceFrequencyTierUpdate PriceFrequencyTierUpdate) PriceFrequencyTiersApiPATCHPriceFrequencyTiersPriceFrequencyTierIdRequest {
	r.priceFrequencyTierUpdate = &priceFrequencyTierUpdate
	return r
}

func (r PriceFrequencyTiersApiPATCHPriceFrequencyTiersPriceFrequencyTierIdRequest) Execute() (*PATCHPriceFrequencyTiersPriceFrequencyTierId200Response, *http.Response, error) {
	return r.ApiService.PATCHPriceFrequencyTiersPriceFrequencyTierIdExecute(r)
}

/*
PATCHPriceFrequencyTiersPriceFrequencyTierId Update a price frequency tier

Update a price frequency tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param priceFrequencyTierId The resource's id
	@return PriceFrequencyTiersApiPATCHPriceFrequencyTiersPriceFrequencyTierIdRequest
*/
func (a *PriceFrequencyTiersApiService) PATCHPriceFrequencyTiersPriceFrequencyTierId(ctx context.Context, priceFrequencyTierId interface{}) PriceFrequencyTiersApiPATCHPriceFrequencyTiersPriceFrequencyTierIdRequest {
	return PriceFrequencyTiersApiPATCHPriceFrequencyTiersPriceFrequencyTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		priceFrequencyTierId: priceFrequencyTierId,
	}
}

// Execute executes the request
//
//	@return PATCHPriceFrequencyTiersPriceFrequencyTierId200Response
func (a *PriceFrequencyTiersApiService) PATCHPriceFrequencyTiersPriceFrequencyTierIdExecute(r PriceFrequencyTiersApiPATCHPriceFrequencyTiersPriceFrequencyTierIdRequest) (*PATCHPriceFrequencyTiersPriceFrequencyTierId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHPriceFrequencyTiersPriceFrequencyTierId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceFrequencyTiersApiService.PATCHPriceFrequencyTiersPriceFrequencyTierId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_frequency_tiers/{priceFrequencyTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"priceFrequencyTierId"+"}", url.PathEscape(parameterValueToString(r.priceFrequencyTierId, "priceFrequencyTierId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceFrequencyTierUpdate == nil {
		return localVarReturnValue, nil, reportError("priceFrequencyTierUpdate is required and must be specified")
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
	localVarPostBody = r.priceFrequencyTierUpdate
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

type PriceFrequencyTiersApiPOSTPriceFrequencyTiersRequest struct {
	ctx                      context.Context
	ApiService               *PriceFrequencyTiersApiService
	priceFrequencyTierCreate *PriceFrequencyTierCreate
}

func (r PriceFrequencyTiersApiPOSTPriceFrequencyTiersRequest) PriceFrequencyTierCreate(priceFrequencyTierCreate PriceFrequencyTierCreate) PriceFrequencyTiersApiPOSTPriceFrequencyTiersRequest {
	r.priceFrequencyTierCreate = &priceFrequencyTierCreate
	return r
}

func (r PriceFrequencyTiersApiPOSTPriceFrequencyTiersRequest) Execute() (*POSTPriceFrequencyTiers201Response, *http.Response, error) {
	return r.ApiService.POSTPriceFrequencyTiersExecute(r)
}

/*
POSTPriceFrequencyTiers Create a price frequency tier

Create a price frequency tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PriceFrequencyTiersApiPOSTPriceFrequencyTiersRequest
*/
func (a *PriceFrequencyTiersApiService) POSTPriceFrequencyTiers(ctx context.Context) PriceFrequencyTiersApiPOSTPriceFrequencyTiersRequest {
	return PriceFrequencyTiersApiPOSTPriceFrequencyTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTPriceFrequencyTiers201Response
func (a *PriceFrequencyTiersApiService) POSTPriceFrequencyTiersExecute(r PriceFrequencyTiersApiPOSTPriceFrequencyTiersRequest) (*POSTPriceFrequencyTiers201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTPriceFrequencyTiers201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PriceFrequencyTiersApiService.POSTPriceFrequencyTiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/price_frequency_tiers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.priceFrequencyTierCreate == nil {
		return localVarReturnValue, nil, reportError("priceFrequencyTierCreate is required and must be specified")
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
	localVarPostBody = r.priceFrequencyTierCreate
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
