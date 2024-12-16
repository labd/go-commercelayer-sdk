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

// SkuOptionsApiService SkuOptionsApi service
type SkuOptionsApiService service

type SkuOptionsApiDELETESkuOptionsSkuOptionIdRequest struct {
	ctx         context.Context
	ApiService  *SkuOptionsApiService
	skuOptionId interface{}
}

func (r SkuOptionsApiDELETESkuOptionsSkuOptionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETESkuOptionsSkuOptionIdExecute(r)
}

/*
DELETESkuOptionsSkuOptionId Delete a SKU option

Delete a SKU option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param skuOptionId The resource's id
	@return SkuOptionsApiDELETESkuOptionsSkuOptionIdRequest
*/
func (a *SkuOptionsApiService) DELETESkuOptionsSkuOptionId(ctx context.Context, skuOptionId interface{}) SkuOptionsApiDELETESkuOptionsSkuOptionIdRequest {
	return SkuOptionsApiDELETESkuOptionsSkuOptionIdRequest{
		ApiService:  a,
		ctx:         ctx,
		skuOptionId: skuOptionId,
	}
}

// Execute executes the request
func (a *SkuOptionsApiService) DELETESkuOptionsSkuOptionIdExecute(r SkuOptionsApiDELETESkuOptionsSkuOptionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkuOptionsApiService.DELETESkuOptionsSkuOptionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sku_options/{skuOptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skuOptionId"+"}", url.PathEscape(parameterValueToString(r.skuOptionId, "skuOptionId")), -1)

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

type SkuOptionsApiGETLineItemOptionIdSkuOptionRequest struct {
	ctx              context.Context
	ApiService       *SkuOptionsApiService
	lineItemOptionId interface{}
}

func (r SkuOptionsApiGETLineItemOptionIdSkuOptionRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETLineItemOptionIdSkuOptionExecute(r)
}

/*
GETLineItemOptionIdSkuOption Retrieve the sku option associated to the line item option

Retrieve the sku option associated to the line item option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemOptionId The resource's id
	@return SkuOptionsApiGETLineItemOptionIdSkuOptionRequest
*/
func (a *SkuOptionsApiService) GETLineItemOptionIdSkuOption(ctx context.Context, lineItemOptionId interface{}) SkuOptionsApiGETLineItemOptionIdSkuOptionRequest {
	return SkuOptionsApiGETLineItemOptionIdSkuOptionRequest{
		ApiService:       a,
		ctx:              ctx,
		lineItemOptionId: lineItemOptionId,
	}
}

// Execute executes the request
func (a *SkuOptionsApiService) GETLineItemOptionIdSkuOptionExecute(r SkuOptionsApiGETLineItemOptionIdSkuOptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkuOptionsApiService.GETLineItemOptionIdSkuOption")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_item_options/{lineItemOptionId}/sku_option"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemOptionId"+"}", url.PathEscape(parameterValueToString(r.lineItemOptionId, "lineItemOptionId")), -1)

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

type SkuOptionsApiGETSkuIdSkuOptionsRequest struct {
	ctx        context.Context
	ApiService *SkuOptionsApiService
	skuId      interface{}
}

func (r SkuOptionsApiGETSkuIdSkuOptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETSkuIdSkuOptionsExecute(r)
}

/*
GETSkuIdSkuOptions Retrieve the sku options associated to the SKU

Retrieve the sku options associated to the SKU

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param skuId The resource's id
	@return SkuOptionsApiGETSkuIdSkuOptionsRequest
*/
func (a *SkuOptionsApiService) GETSkuIdSkuOptions(ctx context.Context, skuId interface{}) SkuOptionsApiGETSkuIdSkuOptionsRequest {
	return SkuOptionsApiGETSkuIdSkuOptionsRequest{
		ApiService: a,
		ctx:        ctx,
		skuId:      skuId,
	}
}

// Execute executes the request
func (a *SkuOptionsApiService) GETSkuIdSkuOptionsExecute(r SkuOptionsApiGETSkuIdSkuOptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkuOptionsApiService.GETSkuIdSkuOptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/skus/{skuId}/sku_options"
	localVarPath = strings.Replace(localVarPath, "{"+"skuId"+"}", url.PathEscape(parameterValueToString(r.skuId, "skuId")), -1)

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

type SkuOptionsApiGETSkuOptionsRequest struct {
	ctx        context.Context
	ApiService *SkuOptionsApiService
}

func (r SkuOptionsApiGETSkuOptionsRequest) Execute() (*GETSkuOptions200Response, *http.Response, error) {
	return r.ApiService.GETSkuOptionsExecute(r)
}

/*
GETSkuOptions List all SKU options

List all SKU options

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SkuOptionsApiGETSkuOptionsRequest
*/
func (a *SkuOptionsApiService) GETSkuOptions(ctx context.Context) SkuOptionsApiGETSkuOptionsRequest {
	return SkuOptionsApiGETSkuOptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETSkuOptions200Response
func (a *SkuOptionsApiService) GETSkuOptionsExecute(r SkuOptionsApiGETSkuOptionsRequest) (*GETSkuOptions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETSkuOptions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkuOptionsApiService.GETSkuOptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sku_options"

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

type SkuOptionsApiGETSkuOptionsSkuOptionIdRequest struct {
	ctx         context.Context
	ApiService  *SkuOptionsApiService
	skuOptionId interface{}
}

func (r SkuOptionsApiGETSkuOptionsSkuOptionIdRequest) Execute() (*GETSkuOptionsSkuOptionId200Response, *http.Response, error) {
	return r.ApiService.GETSkuOptionsSkuOptionIdExecute(r)
}

/*
GETSkuOptionsSkuOptionId Retrieve a SKU option

Retrieve a SKU option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param skuOptionId The resource's id
	@return SkuOptionsApiGETSkuOptionsSkuOptionIdRequest
*/
func (a *SkuOptionsApiService) GETSkuOptionsSkuOptionId(ctx context.Context, skuOptionId interface{}) SkuOptionsApiGETSkuOptionsSkuOptionIdRequest {
	return SkuOptionsApiGETSkuOptionsSkuOptionIdRequest{
		ApiService:  a,
		ctx:         ctx,
		skuOptionId: skuOptionId,
	}
}

// Execute executes the request
//
//	@return GETSkuOptionsSkuOptionId200Response
func (a *SkuOptionsApiService) GETSkuOptionsSkuOptionIdExecute(r SkuOptionsApiGETSkuOptionsSkuOptionIdRequest) (*GETSkuOptionsSkuOptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETSkuOptionsSkuOptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkuOptionsApiService.GETSkuOptionsSkuOptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sku_options/{skuOptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skuOptionId"+"}", url.PathEscape(parameterValueToString(r.skuOptionId, "skuOptionId")), -1)

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

type SkuOptionsApiPATCHSkuOptionsSkuOptionIdRequest struct {
	ctx             context.Context
	ApiService      *SkuOptionsApiService
	skuOptionUpdate *SkuOptionUpdate
	skuOptionId     interface{}
}

func (r SkuOptionsApiPATCHSkuOptionsSkuOptionIdRequest) SkuOptionUpdate(skuOptionUpdate SkuOptionUpdate) SkuOptionsApiPATCHSkuOptionsSkuOptionIdRequest {
	r.skuOptionUpdate = &skuOptionUpdate
	return r
}

func (r SkuOptionsApiPATCHSkuOptionsSkuOptionIdRequest) Execute() (*PATCHSkuOptionsSkuOptionId200Response, *http.Response, error) {
	return r.ApiService.PATCHSkuOptionsSkuOptionIdExecute(r)
}

/*
PATCHSkuOptionsSkuOptionId Update a SKU option

Update a SKU option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param skuOptionId The resource's id
	@return SkuOptionsApiPATCHSkuOptionsSkuOptionIdRequest
*/
func (a *SkuOptionsApiService) PATCHSkuOptionsSkuOptionId(ctx context.Context, skuOptionId interface{}) SkuOptionsApiPATCHSkuOptionsSkuOptionIdRequest {
	return SkuOptionsApiPATCHSkuOptionsSkuOptionIdRequest{
		ApiService:  a,
		ctx:         ctx,
		skuOptionId: skuOptionId,
	}
}

// Execute executes the request
//
//	@return PATCHSkuOptionsSkuOptionId200Response
func (a *SkuOptionsApiService) PATCHSkuOptionsSkuOptionIdExecute(r SkuOptionsApiPATCHSkuOptionsSkuOptionIdRequest) (*PATCHSkuOptionsSkuOptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHSkuOptionsSkuOptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkuOptionsApiService.PATCHSkuOptionsSkuOptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sku_options/{skuOptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"skuOptionId"+"}", url.PathEscape(parameterValueToString(r.skuOptionId, "skuOptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.skuOptionUpdate == nil {
		return localVarReturnValue, nil, reportError("skuOptionUpdate is required and must be specified")
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
	localVarPostBody = r.skuOptionUpdate
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

type SkuOptionsApiPOSTSkuOptionsRequest struct {
	ctx             context.Context
	ApiService      *SkuOptionsApiService
	skuOptionCreate *SkuOptionCreate
}

func (r SkuOptionsApiPOSTSkuOptionsRequest) SkuOptionCreate(skuOptionCreate SkuOptionCreate) SkuOptionsApiPOSTSkuOptionsRequest {
	r.skuOptionCreate = &skuOptionCreate
	return r
}

func (r SkuOptionsApiPOSTSkuOptionsRequest) Execute() (*POSTSkuOptions201Response, *http.Response, error) {
	return r.ApiService.POSTSkuOptionsExecute(r)
}

/*
POSTSkuOptions Create a SKU option

Create a SKU option

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SkuOptionsApiPOSTSkuOptionsRequest
*/
func (a *SkuOptionsApiService) POSTSkuOptions(ctx context.Context) SkuOptionsApiPOSTSkuOptionsRequest {
	return SkuOptionsApiPOSTSkuOptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTSkuOptions201Response
func (a *SkuOptionsApiService) POSTSkuOptionsExecute(r SkuOptionsApiPOSTSkuOptionsRequest) (*POSTSkuOptions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTSkuOptions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SkuOptionsApiService.POSTSkuOptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sku_options"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.skuOptionCreate == nil {
		return localVarReturnValue, nil, reportError("skuOptionCreate is required and must be specified")
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
	localVarPostBody = r.skuOptionCreate
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
