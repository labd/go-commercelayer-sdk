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

// ShippingMethodTiersApiService ShippingMethodTiersApi service
type ShippingMethodTiersApiService service

type ShippingMethodTiersApiGETShippingMethodIdShippingMethodTiersRequest struct {
	ctx              context.Context
	ApiService       *ShippingMethodTiersApiService
	shippingMethodId interface{}
}

func (r ShippingMethodTiersApiGETShippingMethodIdShippingMethodTiersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingMethodIdShippingMethodTiersExecute(r)
}

/*
GETShippingMethodIdShippingMethodTiers Retrieve the shipping method tiers associated to the shipping method

Retrieve the shipping method tiers associated to the shipping method

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shippingMethodId The resource's id
	@return ShippingMethodTiersApiGETShippingMethodIdShippingMethodTiersRequest
*/
func (a *ShippingMethodTiersApiService) GETShippingMethodIdShippingMethodTiers(ctx context.Context, shippingMethodId interface{}) ShippingMethodTiersApiGETShippingMethodIdShippingMethodTiersRequest {
	return ShippingMethodTiersApiGETShippingMethodIdShippingMethodTiersRequest{
		ApiService:       a,
		ctx:              ctx,
		shippingMethodId: shippingMethodId,
	}
}

// Execute executes the request
func (a *ShippingMethodTiersApiService) GETShippingMethodIdShippingMethodTiersExecute(r ShippingMethodTiersApiGETShippingMethodIdShippingMethodTiersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodTiersApiService.GETShippingMethodIdShippingMethodTiers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods/{shippingMethodId}/shipping_method_tiers"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodId"+"}", url.PathEscape(parameterValueToString(r.shippingMethodId, "shippingMethodId")), -1)

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

type ShippingMethodTiersApiGETShippingMethodTiersRequest struct {
	ctx        context.Context
	ApiService *ShippingMethodTiersApiService
}

func (r ShippingMethodTiersApiGETShippingMethodTiersRequest) Execute() (*GETShippingMethodTiers200Response, *http.Response, error) {
	return r.ApiService.GETShippingMethodTiersExecute(r)
}

/*
GETShippingMethodTiers List all shipping method tiers

List all shipping method tiers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ShippingMethodTiersApiGETShippingMethodTiersRequest
*/
func (a *ShippingMethodTiersApiService) GETShippingMethodTiers(ctx context.Context) ShippingMethodTiersApiGETShippingMethodTiersRequest {
	return ShippingMethodTiersApiGETShippingMethodTiersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETShippingMethodTiers200Response
func (a *ShippingMethodTiersApiService) GETShippingMethodTiersExecute(r ShippingMethodTiersApiGETShippingMethodTiersRequest) (*GETShippingMethodTiers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETShippingMethodTiers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodTiersApiService.GETShippingMethodTiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_method_tiers"

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

type ShippingMethodTiersApiGETShippingMethodTiersShippingMethodTierIdRequest struct {
	ctx                  context.Context
	ApiService           *ShippingMethodTiersApiService
	shippingMethodTierId interface{}
}

func (r ShippingMethodTiersApiGETShippingMethodTiersShippingMethodTierIdRequest) Execute() (*GETShippingMethodTiersShippingMethodTierId200Response, *http.Response, error) {
	return r.ApiService.GETShippingMethodTiersShippingMethodTierIdExecute(r)
}

/*
GETShippingMethodTiersShippingMethodTierId Retrieve a shipping method tier

Retrieve a shipping method tier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shippingMethodTierId The resource's id
	@return ShippingMethodTiersApiGETShippingMethodTiersShippingMethodTierIdRequest
*/
func (a *ShippingMethodTiersApiService) GETShippingMethodTiersShippingMethodTierId(ctx context.Context, shippingMethodTierId interface{}) ShippingMethodTiersApiGETShippingMethodTiersShippingMethodTierIdRequest {
	return ShippingMethodTiersApiGETShippingMethodTiersShippingMethodTierIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		shippingMethodTierId: shippingMethodTierId,
	}
}

// Execute executes the request
//
//	@return GETShippingMethodTiersShippingMethodTierId200Response
func (a *ShippingMethodTiersApiService) GETShippingMethodTiersShippingMethodTierIdExecute(r ShippingMethodTiersApiGETShippingMethodTiersShippingMethodTierIdRequest) (*GETShippingMethodTiersShippingMethodTierId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETShippingMethodTiersShippingMethodTierId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingMethodTiersApiService.GETShippingMethodTiersShippingMethodTierId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_method_tiers/{shippingMethodTierId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodTierId"+"}", url.PathEscape(parameterValueToString(r.shippingMethodTierId, "shippingMethodTierId")), -1)

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
