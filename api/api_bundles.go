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

// BundlesApiService BundlesApi service
type BundlesApiService service

type BundlesApiDELETEBundlesBundleIdRequest struct {
	ctx        context.Context
	ApiService *BundlesApiService
	bundleId   interface{}
}

func (r BundlesApiDELETEBundlesBundleIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEBundlesBundleIdExecute(r)
}

/*
DELETEBundlesBundleId Delete a bundle

Delete a bundle

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bundleId The resource's id
	@return BundlesApiDELETEBundlesBundleIdRequest
*/
func (a *BundlesApiService) DELETEBundlesBundleId(ctx context.Context, bundleId interface{}) BundlesApiDELETEBundlesBundleIdRequest {
	return BundlesApiDELETEBundlesBundleIdRequest{
		ApiService: a,
		ctx:        ctx,
		bundleId:   bundleId,
	}
}

// Execute executes the request
func (a *BundlesApiService) DELETEBundlesBundleIdExecute(r BundlesApiDELETEBundlesBundleIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.DELETEBundlesBundleId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles/{bundleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundleId"+"}", url.PathEscape(parameterValueToString(r.bundleId, "bundleId")), -1)

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

type BundlesApiGETBundlesRequest struct {
	ctx        context.Context
	ApiService *BundlesApiService
}

func (r BundlesApiGETBundlesRequest) Execute() (*GETBundles200Response, *http.Response, error) {
	return r.ApiService.GETBundlesExecute(r)
}

/*
GETBundles List all bundles

List all bundles

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BundlesApiGETBundlesRequest
*/
func (a *BundlesApiService) GETBundles(ctx context.Context) BundlesApiGETBundlesRequest {
	return BundlesApiGETBundlesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETBundles200Response
func (a *BundlesApiService) GETBundlesExecute(r BundlesApiGETBundlesRequest) (*GETBundles200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBundles200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.GETBundles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles"

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

type BundlesApiGETBundlesBundleIdRequest struct {
	ctx        context.Context
	ApiService *BundlesApiService
	bundleId   interface{}
}

func (r BundlesApiGETBundlesBundleIdRequest) Execute() (*GETBundlesBundleId200Response, *http.Response, error) {
	return r.ApiService.GETBundlesBundleIdExecute(r)
}

/*
GETBundlesBundleId Retrieve a bundle

Retrieve a bundle

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bundleId The resource's id
	@return BundlesApiGETBundlesBundleIdRequest
*/
func (a *BundlesApiService) GETBundlesBundleId(ctx context.Context, bundleId interface{}) BundlesApiGETBundlesBundleIdRequest {
	return BundlesApiGETBundlesBundleIdRequest{
		ApiService: a,
		ctx:        ctx,
		bundleId:   bundleId,
	}
}

// Execute executes the request
//
//	@return GETBundlesBundleId200Response
func (a *BundlesApiService) GETBundlesBundleIdExecute(r BundlesApiGETBundlesBundleIdRequest) (*GETBundlesBundleId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETBundlesBundleId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.GETBundlesBundleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles/{bundleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundleId"+"}", url.PathEscape(parameterValueToString(r.bundleId, "bundleId")), -1)

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

type BundlesApiGETOrderIdAvailableFreeBundlesRequest struct {
	ctx        context.Context
	ApiService *BundlesApiService
	orderId    interface{}
}

func (r BundlesApiGETOrderIdAvailableFreeBundlesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdAvailableFreeBundlesExecute(r)
}

/*
GETOrderIdAvailableFreeBundles Retrieve the available free bundles associated to the order

Retrieve the available free bundles associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return BundlesApiGETOrderIdAvailableFreeBundlesRequest
*/
func (a *BundlesApiService) GETOrderIdAvailableFreeBundles(ctx context.Context, orderId interface{}) BundlesApiGETOrderIdAvailableFreeBundlesRequest {
	return BundlesApiGETOrderIdAvailableFreeBundlesRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *BundlesApiService) GETOrderIdAvailableFreeBundlesExecute(r BundlesApiGETOrderIdAvailableFreeBundlesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.GETOrderIdAvailableFreeBundles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/available_free_bundles"
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

type BundlesApiGETSkuListIdBundlesRequest struct {
	ctx        context.Context
	ApiService *BundlesApiService
	skuListId  interface{}
}

func (r BundlesApiGETSkuListIdBundlesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETSkuListIdBundlesExecute(r)
}

/*
GETSkuListIdBundles Retrieve the bundles associated to the SKU list

Retrieve the bundles associated to the SKU list

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param skuListId The resource's id
	@return BundlesApiGETSkuListIdBundlesRequest
*/
func (a *BundlesApiService) GETSkuListIdBundles(ctx context.Context, skuListId interface{}) BundlesApiGETSkuListIdBundlesRequest {
	return BundlesApiGETSkuListIdBundlesRequest{
		ApiService: a,
		ctx:        ctx,
		skuListId:  skuListId,
	}
}

// Execute executes the request
func (a *BundlesApiService) GETSkuListIdBundlesExecute(r BundlesApiGETSkuListIdBundlesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.GETSkuListIdBundles")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sku_lists/{skuListId}/bundles"
	localVarPath = strings.Replace(localVarPath, "{"+"skuListId"+"}", url.PathEscape(parameterValueToString(r.skuListId, "skuListId")), -1)

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

type BundlesApiPATCHBundlesBundleIdRequest struct {
	ctx          context.Context
	ApiService   *BundlesApiService
	bundleUpdate *BundleUpdate
	bundleId     interface{}
}

func (r BundlesApiPATCHBundlesBundleIdRequest) BundleUpdate(bundleUpdate BundleUpdate) BundlesApiPATCHBundlesBundleIdRequest {
	r.bundleUpdate = &bundleUpdate
	return r
}

func (r BundlesApiPATCHBundlesBundleIdRequest) Execute() (*PATCHBundlesBundleId200Response, *http.Response, error) {
	return r.ApiService.PATCHBundlesBundleIdExecute(r)
}

/*
PATCHBundlesBundleId Update a bundle

Update a bundle

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bundleId The resource's id
	@return BundlesApiPATCHBundlesBundleIdRequest
*/
func (a *BundlesApiService) PATCHBundlesBundleId(ctx context.Context, bundleId interface{}) BundlesApiPATCHBundlesBundleIdRequest {
	return BundlesApiPATCHBundlesBundleIdRequest{
		ApiService: a,
		ctx:        ctx,
		bundleId:   bundleId,
	}
}

// Execute executes the request
//
//	@return PATCHBundlesBundleId200Response
func (a *BundlesApiService) PATCHBundlesBundleIdExecute(r BundlesApiPATCHBundlesBundleIdRequest) (*PATCHBundlesBundleId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHBundlesBundleId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.PATCHBundlesBundleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles/{bundleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundleId"+"}", url.PathEscape(parameterValueToString(r.bundleId, "bundleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bundleUpdate == nil {
		return localVarReturnValue, nil, reportError("bundleUpdate is required and must be specified")
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
	localVarPostBody = r.bundleUpdate
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

type BundlesApiPOSTBundlesRequest struct {
	ctx          context.Context
	ApiService   *BundlesApiService
	bundleCreate *BundleCreate
}

func (r BundlesApiPOSTBundlesRequest) BundleCreate(bundleCreate BundleCreate) BundlesApiPOSTBundlesRequest {
	r.bundleCreate = &bundleCreate
	return r
}

func (r BundlesApiPOSTBundlesRequest) Execute() (*POSTBundles201Response, *http.Response, error) {
	return r.ApiService.POSTBundlesExecute(r)
}

/*
POSTBundles Create a bundle

Create a bundle

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BundlesApiPOSTBundlesRequest
*/
func (a *BundlesApiService) POSTBundles(ctx context.Context) BundlesApiPOSTBundlesRequest {
	return BundlesApiPOSTBundlesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTBundles201Response
func (a *BundlesApiService) POSTBundlesExecute(r BundlesApiPOSTBundlesRequest) (*POSTBundles201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTBundles201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BundlesApiService.POSTBundles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bundles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bundleCreate == nil {
		return localVarReturnValue, nil, reportError("bundleCreate is required and must be specified")
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
	localVarPostBody = r.bundleCreate
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
