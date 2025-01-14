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

// OrderCopiesApiService OrderCopiesApi service
type OrderCopiesApiService service

type OrderCopiesApiDELETEOrderCopiesOrderCopyIdRequest struct {
	ctx         context.Context
	ApiService  *OrderCopiesApiService
	orderCopyId interface{}
}

func (r OrderCopiesApiDELETEOrderCopiesOrderCopyIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEOrderCopiesOrderCopyIdExecute(r)
}

/*
DELETEOrderCopiesOrderCopyId Delete an order copy

Delete an order copy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderCopyId The resource's id
	@return OrderCopiesApiDELETEOrderCopiesOrderCopyIdRequest
*/
func (a *OrderCopiesApiService) DELETEOrderCopiesOrderCopyId(ctx context.Context, orderCopyId interface{}) OrderCopiesApiDELETEOrderCopiesOrderCopyIdRequest {
	return OrderCopiesApiDELETEOrderCopiesOrderCopyIdRequest{
		ApiService:  a,
		ctx:         ctx,
		orderCopyId: orderCopyId,
	}
}

// Execute executes the request
func (a *OrderCopiesApiService) DELETEOrderCopiesOrderCopyIdExecute(r OrderCopiesApiDELETEOrderCopiesOrderCopyIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderCopiesApiService.DELETEOrderCopiesOrderCopyId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_copies/{orderCopyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderCopyId"+"}", url.PathEscape(parameterValueToString(r.orderCopyId, "orderCopyId")), -1)

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

type OrderCopiesApiGETOrderCopiesRequest struct {
	ctx        context.Context
	ApiService *OrderCopiesApiService
}

func (r OrderCopiesApiGETOrderCopiesRequest) Execute() (*GETOrderCopies200Response, *http.Response, error) {
	return r.ApiService.GETOrderCopiesExecute(r)
}

/*
GETOrderCopies List all order copies

List all order copies

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OrderCopiesApiGETOrderCopiesRequest
*/
func (a *OrderCopiesApiService) GETOrderCopies(ctx context.Context) OrderCopiesApiGETOrderCopiesRequest {
	return OrderCopiesApiGETOrderCopiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETOrderCopies200Response
func (a *OrderCopiesApiService) GETOrderCopiesExecute(r OrderCopiesApiGETOrderCopiesRequest) (*GETOrderCopies200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETOrderCopies200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderCopiesApiService.GETOrderCopies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_copies"

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

type OrderCopiesApiGETOrderCopiesOrderCopyIdRequest struct {
	ctx         context.Context
	ApiService  *OrderCopiesApiService
	orderCopyId interface{}
}

func (r OrderCopiesApiGETOrderCopiesOrderCopyIdRequest) Execute() (*GETOrderCopiesOrderCopyId200Response, *http.Response, error) {
	return r.ApiService.GETOrderCopiesOrderCopyIdExecute(r)
}

/*
GETOrderCopiesOrderCopyId Retrieve an order copy

Retrieve an order copy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderCopyId The resource's id
	@return OrderCopiesApiGETOrderCopiesOrderCopyIdRequest
*/
func (a *OrderCopiesApiService) GETOrderCopiesOrderCopyId(ctx context.Context, orderCopyId interface{}) OrderCopiesApiGETOrderCopiesOrderCopyIdRequest {
	return OrderCopiesApiGETOrderCopiesOrderCopyIdRequest{
		ApiService:  a,
		ctx:         ctx,
		orderCopyId: orderCopyId,
	}
}

// Execute executes the request
//
//	@return GETOrderCopiesOrderCopyId200Response
func (a *OrderCopiesApiService) GETOrderCopiesOrderCopyIdExecute(r OrderCopiesApiGETOrderCopiesOrderCopyIdRequest) (*GETOrderCopiesOrderCopyId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETOrderCopiesOrderCopyId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderCopiesApiService.GETOrderCopiesOrderCopyId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_copies/{orderCopyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderCopyId"+"}", url.PathEscape(parameterValueToString(r.orderCopyId, "orderCopyId")), -1)

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

type OrderCopiesApiGETOrderIdOrderCopiesRequest struct {
	ctx        context.Context
	ApiService *OrderCopiesApiService
	orderId    interface{}
}

func (r OrderCopiesApiGETOrderIdOrderCopiesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdOrderCopiesExecute(r)
}

/*
GETOrderIdOrderCopies Retrieve the order copies associated to the order

Retrieve the order copies associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return OrderCopiesApiGETOrderIdOrderCopiesRequest
*/
func (a *OrderCopiesApiService) GETOrderIdOrderCopies(ctx context.Context, orderId interface{}) OrderCopiesApiGETOrderIdOrderCopiesRequest {
	return OrderCopiesApiGETOrderIdOrderCopiesRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *OrderCopiesApiService) GETOrderIdOrderCopiesExecute(r OrderCopiesApiGETOrderIdOrderCopiesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderCopiesApiService.GETOrderIdOrderCopies")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/order_copies"
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

type OrderCopiesApiPATCHOrderCopiesOrderCopyIdRequest struct {
	ctx             context.Context
	ApiService      *OrderCopiesApiService
	orderCopyUpdate *OrderCopyUpdate
	orderCopyId     interface{}
}

func (r OrderCopiesApiPATCHOrderCopiesOrderCopyIdRequest) OrderCopyUpdate(orderCopyUpdate OrderCopyUpdate) OrderCopiesApiPATCHOrderCopiesOrderCopyIdRequest {
	r.orderCopyUpdate = &orderCopyUpdate
	return r
}

func (r OrderCopiesApiPATCHOrderCopiesOrderCopyIdRequest) Execute() (*PATCHOrderCopiesOrderCopyId200Response, *http.Response, error) {
	return r.ApiService.PATCHOrderCopiesOrderCopyIdExecute(r)
}

/*
PATCHOrderCopiesOrderCopyId Update an order copy

Update an order copy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderCopyId The resource's id
	@return OrderCopiesApiPATCHOrderCopiesOrderCopyIdRequest
*/
func (a *OrderCopiesApiService) PATCHOrderCopiesOrderCopyId(ctx context.Context, orderCopyId interface{}) OrderCopiesApiPATCHOrderCopiesOrderCopyIdRequest {
	return OrderCopiesApiPATCHOrderCopiesOrderCopyIdRequest{
		ApiService:  a,
		ctx:         ctx,
		orderCopyId: orderCopyId,
	}
}

// Execute executes the request
//
//	@return PATCHOrderCopiesOrderCopyId200Response
func (a *OrderCopiesApiService) PATCHOrderCopiesOrderCopyIdExecute(r OrderCopiesApiPATCHOrderCopiesOrderCopyIdRequest) (*PATCHOrderCopiesOrderCopyId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHOrderCopiesOrderCopyId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderCopiesApiService.PATCHOrderCopiesOrderCopyId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_copies/{orderCopyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orderCopyId"+"}", url.PathEscape(parameterValueToString(r.orderCopyId, "orderCopyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderCopyUpdate == nil {
		return localVarReturnValue, nil, reportError("orderCopyUpdate is required and must be specified")
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
	localVarPostBody = r.orderCopyUpdate
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

type OrderCopiesApiPOSTOrderCopiesRequest struct {
	ctx             context.Context
	ApiService      *OrderCopiesApiService
	orderCopyCreate *OrderCopyCreate
}

func (r OrderCopiesApiPOSTOrderCopiesRequest) OrderCopyCreate(orderCopyCreate OrderCopyCreate) OrderCopiesApiPOSTOrderCopiesRequest {
	r.orderCopyCreate = &orderCopyCreate
	return r
}

func (r OrderCopiesApiPOSTOrderCopiesRequest) Execute() (*POSTOrderCopies201Response, *http.Response, error) {
	return r.ApiService.POSTOrderCopiesExecute(r)
}

/*
POSTOrderCopies Create an order copy

Create an order copy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OrderCopiesApiPOSTOrderCopiesRequest
*/
func (a *OrderCopiesApiService) POSTOrderCopies(ctx context.Context) OrderCopiesApiPOSTOrderCopiesRequest {
	return OrderCopiesApiPOSTOrderCopiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTOrderCopies201Response
func (a *OrderCopiesApiService) POSTOrderCopiesExecute(r OrderCopiesApiPOSTOrderCopiesRequest) (*POSTOrderCopies201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTOrderCopies201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderCopiesApiService.POSTOrderCopies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order_copies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderCopyCreate == nil {
		return localVarReturnValue, nil, reportError("orderCopyCreate is required and must be specified")
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
	localVarPostBody = r.orderCopyCreate
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
