/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
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

// StockItemsApiService StockItemsApi service
type StockItemsApiService service

type StockItemsApiDELETEStockItemsStockItemIdRequest struct {
	ctx         context.Context
	ApiService  *StockItemsApiService
	stockItemId string
}

func (r StockItemsApiDELETEStockItemsStockItemIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEStockItemsStockItemIdExecute(r)
}

/*
DELETEStockItemsStockItemId Delete a stock item

Delete a stock item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockItemId The resource's id
 @return StockItemsApiDELETEStockItemsStockItemIdRequest
*/
func (a *StockItemsApiService) DELETEStockItemsStockItemId(ctx context.Context, stockItemId string) StockItemsApiDELETEStockItemsStockItemIdRequest {
	return StockItemsApiDELETEStockItemsStockItemIdRequest{
		ApiService:  a,
		ctx:         ctx,
		stockItemId: stockItemId,
	}
}

// Execute executes the request
func (a *StockItemsApiService) DELETEStockItemsStockItemIdExecute(r StockItemsApiDELETEStockItemsStockItemIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockItemsApiService.DELETEStockItemsStockItemId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_items/{stockItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockItemId"+"}", url.PathEscape(parameterToString(r.stockItemId, "")), -1)

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

type StockItemsApiGETSkuIdStockItemsRequest struct {
	ctx        context.Context
	ApiService *StockItemsApiService
	skuId      string
}

func (r StockItemsApiGETSkuIdStockItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETSkuIdStockItemsExecute(r)
}

/*
GETSkuIdStockItems Retrieve the stock items associated to the SKU

Retrieve the stock items associated to the SKU

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param skuId The resource's id
 @return StockItemsApiGETSkuIdStockItemsRequest
*/
func (a *StockItemsApiService) GETSkuIdStockItems(ctx context.Context, skuId string) StockItemsApiGETSkuIdStockItemsRequest {
	return StockItemsApiGETSkuIdStockItemsRequest{
		ApiService: a,
		ctx:        ctx,
		skuId:      skuId,
	}
}

// Execute executes the request
func (a *StockItemsApiService) GETSkuIdStockItemsExecute(r StockItemsApiGETSkuIdStockItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockItemsApiService.GETSkuIdStockItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/skus/{skuId}/stock_items"
	localVarPath = strings.Replace(localVarPath, "{"+"skuId"+"}", url.PathEscape(parameterToString(r.skuId, "")), -1)

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

type StockItemsApiGETStockItemsRequest struct {
	ctx        context.Context
	ApiService *StockItemsApiService
}

func (r StockItemsApiGETStockItemsRequest) Execute() (*StockItemResponseList, *http.Response, error) {
	return r.ApiService.GETStockItemsExecute(r)
}

/*
GETStockItems List all stock items

List all stock items

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return StockItemsApiGETStockItemsRequest
*/
func (a *StockItemsApiService) GETStockItems(ctx context.Context) StockItemsApiGETStockItemsRequest {
	return StockItemsApiGETStockItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return StockItemResponseList
func (a *StockItemsApiService) GETStockItemsExecute(r StockItemsApiGETStockItemsRequest) (*StockItemResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StockItemResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockItemsApiService.GETStockItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_items"

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

type StockItemsApiGETStockItemsStockItemIdRequest struct {
	ctx         context.Context
	ApiService  *StockItemsApiService
	stockItemId string
}

func (r StockItemsApiGETStockItemsStockItemIdRequest) Execute() (*StockItemResponse, *http.Response, error) {
	return r.ApiService.GETStockItemsStockItemIdExecute(r)
}

/*
GETStockItemsStockItemId Retrieve a stock item

Retrieve a stock item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockItemId The resource's id
 @return StockItemsApiGETStockItemsStockItemIdRequest
*/
func (a *StockItemsApiService) GETStockItemsStockItemId(ctx context.Context, stockItemId string) StockItemsApiGETStockItemsStockItemIdRequest {
	return StockItemsApiGETStockItemsStockItemIdRequest{
		ApiService:  a,
		ctx:         ctx,
		stockItemId: stockItemId,
	}
}

// Execute executes the request
//  @return StockItemResponse
func (a *StockItemsApiService) GETStockItemsStockItemIdExecute(r StockItemsApiGETStockItemsStockItemIdRequest) (*StockItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StockItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockItemsApiService.GETStockItemsStockItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_items/{stockItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockItemId"+"}", url.PathEscape(parameterToString(r.stockItemId, "")), -1)

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

type StockItemsApiGETStockLineItemIdStockItemRequest struct {
	ctx             context.Context
	ApiService      *StockItemsApiService
	stockLineItemId string
}

func (r StockItemsApiGETStockLineItemIdStockItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockLineItemIdStockItemExecute(r)
}

/*
GETStockLineItemIdStockItem Retrieve the stock item associated to the stock line item

Retrieve the stock item associated to the stock line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockLineItemId The resource's id
 @return StockItemsApiGETStockLineItemIdStockItemRequest
*/
func (a *StockItemsApiService) GETStockLineItemIdStockItem(ctx context.Context, stockLineItemId string) StockItemsApiGETStockLineItemIdStockItemRequest {
	return StockItemsApiGETStockLineItemIdStockItemRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLineItemId: stockLineItemId,
	}
}

// Execute executes the request
func (a *StockItemsApiService) GETStockLineItemIdStockItemExecute(r StockItemsApiGETStockLineItemIdStockItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockItemsApiService.GETStockLineItemIdStockItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items/{stockLineItemId}/stock_item"
	localVarPath = strings.Replace(localVarPath, "{"+"stockLineItemId"+"}", url.PathEscape(parameterToString(r.stockLineItemId, "")), -1)

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

type StockItemsApiGETStockLocationIdStockItemsRequest struct {
	ctx             context.Context
	ApiService      *StockItemsApiService
	stockLocationId string
}

func (r StockItemsApiGETStockLocationIdStockItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockLocationIdStockItemsExecute(r)
}

/*
GETStockLocationIdStockItems Retrieve the stock items associated to the stock location

Retrieve the stock items associated to the stock location

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockLocationId The resource's id
 @return StockItemsApiGETStockLocationIdStockItemsRequest
*/
func (a *StockItemsApiService) GETStockLocationIdStockItems(ctx context.Context, stockLocationId string) StockItemsApiGETStockLocationIdStockItemsRequest {
	return StockItemsApiGETStockLocationIdStockItemsRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLocationId: stockLocationId,
	}
}

// Execute executes the request
func (a *StockItemsApiService) GETStockLocationIdStockItemsExecute(r StockItemsApiGETStockLocationIdStockItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockItemsApiService.GETStockLocationIdStockItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_locations/{stockLocationId}/stock_items"
	localVarPath = strings.Replace(localVarPath, "{"+"stockLocationId"+"}", url.PathEscape(parameterToString(r.stockLocationId, "")), -1)

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

type StockItemsApiPATCHStockItemsStockItemIdRequest struct {
	ctx             context.Context
	ApiService      *StockItemsApiService
	stockItemUpdate *StockItemUpdate
	stockItemId     string
}

func (r StockItemsApiPATCHStockItemsStockItemIdRequest) StockItemUpdate(stockItemUpdate StockItemUpdate) StockItemsApiPATCHStockItemsStockItemIdRequest {
	r.stockItemUpdate = &stockItemUpdate
	return r
}

func (r StockItemsApiPATCHStockItemsStockItemIdRequest) Execute() (*StockItemResponse, *http.Response, error) {
	return r.ApiService.PATCHStockItemsStockItemIdExecute(r)
}

/*
PATCHStockItemsStockItemId Update a stock item

Update a stock item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockItemId The resource's id
 @return StockItemsApiPATCHStockItemsStockItemIdRequest
*/
func (a *StockItemsApiService) PATCHStockItemsStockItemId(ctx context.Context, stockItemId string) StockItemsApiPATCHStockItemsStockItemIdRequest {
	return StockItemsApiPATCHStockItemsStockItemIdRequest{
		ApiService:  a,
		ctx:         ctx,
		stockItemId: stockItemId,
	}
}

// Execute executes the request
//  @return StockItemResponse
func (a *StockItemsApiService) PATCHStockItemsStockItemIdExecute(r StockItemsApiPATCHStockItemsStockItemIdRequest) (*StockItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StockItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockItemsApiService.PATCHStockItemsStockItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_items/{stockItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockItemId"+"}", url.PathEscape(parameterToString(r.stockItemId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stockItemUpdate == nil {
		return localVarReturnValue, nil, reportError("stockItemUpdate is required and must be specified")
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
	localVarPostBody = r.stockItemUpdate
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

type StockItemsApiPOSTStockItemsRequest struct {
	ctx             context.Context
	ApiService      *StockItemsApiService
	stockItemCreate *StockItemCreate
}

func (r StockItemsApiPOSTStockItemsRequest) StockItemCreate(stockItemCreate StockItemCreate) StockItemsApiPOSTStockItemsRequest {
	r.stockItemCreate = &stockItemCreate
	return r
}

func (r StockItemsApiPOSTStockItemsRequest) Execute() (*StockItemResponse, *http.Response, error) {
	return r.ApiService.POSTStockItemsExecute(r)
}

/*
POSTStockItems Create a stock item

Create a stock item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return StockItemsApiPOSTStockItemsRequest
*/
func (a *StockItemsApiService) POSTStockItems(ctx context.Context) StockItemsApiPOSTStockItemsRequest {
	return StockItemsApiPOSTStockItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return StockItemResponse
func (a *StockItemsApiService) POSTStockItemsExecute(r StockItemsApiPOSTStockItemsRequest) (*StockItemResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StockItemResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockItemsApiService.POSTStockItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_items"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stockItemCreate == nil {
		return localVarReturnValue, nil, reportError("stockItemCreate is required and must be specified")
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
	localVarPostBody = r.stockItemCreate
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
