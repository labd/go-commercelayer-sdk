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

// StockLineItemsApiService StockLineItemsApi service
type StockLineItemsApiService service

type StockLineItemsApiDELETEStockLineItemsStockLineItemIdRequest struct {
	ctx             context.Context
	ApiService      *StockLineItemsApiService
	stockLineItemId interface{}
}

func (r StockLineItemsApiDELETEStockLineItemsStockLineItemIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEStockLineItemsStockLineItemIdExecute(r)
}

/*
DELETEStockLineItemsStockLineItemId Delete a stock line item

Delete a stock line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockLineItemId The resource's id
	@return StockLineItemsApiDELETEStockLineItemsStockLineItemIdRequest
*/
func (a *StockLineItemsApiService) DELETEStockLineItemsStockLineItemId(ctx context.Context, stockLineItemId interface{}) StockLineItemsApiDELETEStockLineItemsStockLineItemIdRequest {
	return StockLineItemsApiDELETEStockLineItemsStockLineItemIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLineItemId: stockLineItemId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) DELETEStockLineItemsStockLineItemIdExecute(r StockLineItemsApiDELETEStockLineItemsStockLineItemIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.DELETEStockLineItemsStockLineItemId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items/{stockLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockLineItemId"+"}", url.PathEscape(parameterValueToString(r.stockLineItemId, "stockLineItemId")), -1)

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

type StockLineItemsApiGETLineItemIdStockLineItemsRequest struct {
	ctx        context.Context
	ApiService *StockLineItemsApiService
	lineItemId interface{}
}

func (r StockLineItemsApiGETLineItemIdStockLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETLineItemIdStockLineItemsExecute(r)
}

/*
GETLineItemIdStockLineItems Retrieve the stock line items associated to the line item

Retrieve the stock line items associated to the line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemId The resource's id
	@return StockLineItemsApiGETLineItemIdStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETLineItemIdStockLineItems(ctx context.Context, lineItemId interface{}) StockLineItemsApiGETLineItemIdStockLineItemsRequest {
	return StockLineItemsApiGETLineItemIdStockLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
		lineItemId: lineItemId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETLineItemIdStockLineItemsExecute(r StockLineItemsApiGETLineItemIdStockLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETLineItemIdStockLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items/{lineItemId}/stock_line_items"
	localVarPath = strings.Replace(localVarPath, "{"+"lineItemId"+"}", url.PathEscape(parameterValueToString(r.lineItemId, "lineItemId")), -1)

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

type StockLineItemsApiGETOrderIdStockLineItemsRequest struct {
	ctx        context.Context
	ApiService *StockLineItemsApiService
	orderId    interface{}
}

func (r StockLineItemsApiGETOrderIdStockLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdStockLineItemsExecute(r)
}

/*
GETOrderIdStockLineItems Retrieve the stock line items associated to the order

Retrieve the stock line items associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return StockLineItemsApiGETOrderIdStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETOrderIdStockLineItems(ctx context.Context, orderId interface{}) StockLineItemsApiGETOrderIdStockLineItemsRequest {
	return StockLineItemsApiGETOrderIdStockLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETOrderIdStockLineItemsExecute(r StockLineItemsApiGETOrderIdStockLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETOrderIdStockLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/stock_line_items"
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

type StockLineItemsApiGETParcelLineItemIdStockLineItemRequest struct {
	ctx              context.Context
	ApiService       *StockLineItemsApiService
	parcelLineItemId interface{}
}

func (r StockLineItemsApiGETParcelLineItemIdStockLineItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETParcelLineItemIdStockLineItemExecute(r)
}

/*
GETParcelLineItemIdStockLineItem Retrieve the stock line item associated to the parcel line item

Retrieve the stock line item associated to the parcel line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parcelLineItemId The resource's id
	@return StockLineItemsApiGETParcelLineItemIdStockLineItemRequest
*/
func (a *StockLineItemsApiService) GETParcelLineItemIdStockLineItem(ctx context.Context, parcelLineItemId interface{}) StockLineItemsApiGETParcelLineItemIdStockLineItemRequest {
	return StockLineItemsApiGETParcelLineItemIdStockLineItemRequest{
		ApiService:       a,
		ctx:              ctx,
		parcelLineItemId: parcelLineItemId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETParcelLineItemIdStockLineItemExecute(r StockLineItemsApiGETParcelLineItemIdStockLineItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETParcelLineItemIdStockLineItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcel_line_items/{parcelLineItemId}/stock_line_item"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelLineItemId"+"}", url.PathEscape(parameterValueToString(r.parcelLineItemId, "parcelLineItemId")), -1)

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

type StockLineItemsApiGETShipmentIdStockLineItemsRequest struct {
	ctx        context.Context
	ApiService *StockLineItemsApiService
	shipmentId interface{}
}

func (r StockLineItemsApiGETShipmentIdStockLineItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentIdStockLineItemsExecute(r)
}

/*
GETShipmentIdStockLineItems Retrieve the stock line items associated to the shipment

Retrieve the stock line items associated to the shipment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId The resource's id
	@return StockLineItemsApiGETShipmentIdStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETShipmentIdStockLineItems(ctx context.Context, shipmentId interface{}) StockLineItemsApiGETShipmentIdStockLineItemsRequest {
	return StockLineItemsApiGETShipmentIdStockLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETShipmentIdStockLineItemsExecute(r StockLineItemsApiGETShipmentIdStockLineItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETShipmentIdStockLineItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}/stock_line_items"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterValueToString(r.shipmentId, "shipmentId")), -1)

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

type StockLineItemsApiGETStockLineItemsRequest struct {
	ctx        context.Context
	ApiService *StockLineItemsApiService
}

func (r StockLineItemsApiGETStockLineItemsRequest) Execute() (*GETStockLineItems200Response, *http.Response, error) {
	return r.ApiService.GETStockLineItemsExecute(r)
}

/*
GETStockLineItems List all stock line items

List all stock line items

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return StockLineItemsApiGETStockLineItemsRequest
*/
func (a *StockLineItemsApiService) GETStockLineItems(ctx context.Context) StockLineItemsApiGETStockLineItemsRequest {
	return StockLineItemsApiGETStockLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETStockLineItems200Response
func (a *StockLineItemsApiService) GETStockLineItemsExecute(r StockLineItemsApiGETStockLineItemsRequest) (*GETStockLineItems200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETStockLineItems200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETStockLineItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items"

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

type StockLineItemsApiGETStockLineItemsStockLineItemIdRequest struct {
	ctx             context.Context
	ApiService      *StockLineItemsApiService
	stockLineItemId interface{}
}

func (r StockLineItemsApiGETStockLineItemsStockLineItemIdRequest) Execute() (*GETStockLineItemsStockLineItemId200Response, *http.Response, error) {
	return r.ApiService.GETStockLineItemsStockLineItemIdExecute(r)
}

/*
GETStockLineItemsStockLineItemId Retrieve a stock line item

Retrieve a stock line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockLineItemId The resource's id
	@return StockLineItemsApiGETStockLineItemsStockLineItemIdRequest
*/
func (a *StockLineItemsApiService) GETStockLineItemsStockLineItemId(ctx context.Context, stockLineItemId interface{}) StockLineItemsApiGETStockLineItemsStockLineItemIdRequest {
	return StockLineItemsApiGETStockLineItemsStockLineItemIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLineItemId: stockLineItemId,
	}
}

// Execute executes the request
//
//	@return GETStockLineItemsStockLineItemId200Response
func (a *StockLineItemsApiService) GETStockLineItemsStockLineItemIdExecute(r StockLineItemsApiGETStockLineItemsStockLineItemIdRequest) (*GETStockLineItemsStockLineItemId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETStockLineItemsStockLineItemId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETStockLineItemsStockLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items/{stockLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockLineItemId"+"}", url.PathEscape(parameterValueToString(r.stockLineItemId, "stockLineItemId")), -1)

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

type StockLineItemsApiGETStockReservationIdStockLineItemRequest struct {
	ctx                context.Context
	ApiService         *StockLineItemsApiService
	stockReservationId interface{}
}

func (r StockLineItemsApiGETStockReservationIdStockLineItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockReservationIdStockLineItemExecute(r)
}

/*
GETStockReservationIdStockLineItem Retrieve the stock line item associated to the stock reservation

Retrieve the stock line item associated to the stock reservation

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockReservationId The resource's id
	@return StockLineItemsApiGETStockReservationIdStockLineItemRequest
*/
func (a *StockLineItemsApiService) GETStockReservationIdStockLineItem(ctx context.Context, stockReservationId interface{}) StockLineItemsApiGETStockReservationIdStockLineItemRequest {
	return StockLineItemsApiGETStockReservationIdStockLineItemRequest{
		ApiService:         a,
		ctx:                ctx,
		stockReservationId: stockReservationId,
	}
}

// Execute executes the request
func (a *StockLineItemsApiService) GETStockReservationIdStockLineItemExecute(r StockLineItemsApiGETStockReservationIdStockLineItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.GETStockReservationIdStockLineItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_reservations/{stockReservationId}/stock_line_item"
	localVarPath = strings.Replace(localVarPath, "{"+"stockReservationId"+"}", url.PathEscape(parameterValueToString(r.stockReservationId, "stockReservationId")), -1)

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

type StockLineItemsApiPATCHStockLineItemsStockLineItemIdRequest struct {
	ctx                 context.Context
	ApiService          *StockLineItemsApiService
	stockLineItemUpdate *StockLineItemUpdate
	stockLineItemId     interface{}
}

func (r StockLineItemsApiPATCHStockLineItemsStockLineItemIdRequest) StockLineItemUpdate(stockLineItemUpdate StockLineItemUpdate) StockLineItemsApiPATCHStockLineItemsStockLineItemIdRequest {
	r.stockLineItemUpdate = &stockLineItemUpdate
	return r
}

func (r StockLineItemsApiPATCHStockLineItemsStockLineItemIdRequest) Execute() (*PATCHStockLineItemsStockLineItemId200Response, *http.Response, error) {
	return r.ApiService.PATCHStockLineItemsStockLineItemIdExecute(r)
}

/*
PATCHStockLineItemsStockLineItemId Update a stock line item

Update a stock line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockLineItemId The resource's id
	@return StockLineItemsApiPATCHStockLineItemsStockLineItemIdRequest
*/
func (a *StockLineItemsApiService) PATCHStockLineItemsStockLineItemId(ctx context.Context, stockLineItemId interface{}) StockLineItemsApiPATCHStockLineItemsStockLineItemIdRequest {
	return StockLineItemsApiPATCHStockLineItemsStockLineItemIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLineItemId: stockLineItemId,
	}
}

// Execute executes the request
//
//	@return PATCHStockLineItemsStockLineItemId200Response
func (a *StockLineItemsApiService) PATCHStockLineItemsStockLineItemIdExecute(r StockLineItemsApiPATCHStockLineItemsStockLineItemIdRequest) (*PATCHStockLineItemsStockLineItemId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHStockLineItemsStockLineItemId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.PATCHStockLineItemsStockLineItemId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items/{stockLineItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockLineItemId"+"}", url.PathEscape(parameterValueToString(r.stockLineItemId, "stockLineItemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stockLineItemUpdate == nil {
		return localVarReturnValue, nil, reportError("stockLineItemUpdate is required and must be specified")
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
	localVarPostBody = r.stockLineItemUpdate
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

type StockLineItemsApiPOSTStockLineItemsRequest struct {
	ctx                 context.Context
	ApiService          *StockLineItemsApiService
	stockLineItemCreate *StockLineItemCreate
}

func (r StockLineItemsApiPOSTStockLineItemsRequest) StockLineItemCreate(stockLineItemCreate StockLineItemCreate) StockLineItemsApiPOSTStockLineItemsRequest {
	r.stockLineItemCreate = &stockLineItemCreate
	return r
}

func (r StockLineItemsApiPOSTStockLineItemsRequest) Execute() (*POSTStockLineItems201Response, *http.Response, error) {
	return r.ApiService.POSTStockLineItemsExecute(r)
}

/*
POSTStockLineItems Create a stock line item

Create a stock line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return StockLineItemsApiPOSTStockLineItemsRequest
*/
func (a *StockLineItemsApiService) POSTStockLineItems(ctx context.Context) StockLineItemsApiPOSTStockLineItemsRequest {
	return StockLineItemsApiPOSTStockLineItemsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTStockLineItems201Response
func (a *StockLineItemsApiService) POSTStockLineItemsExecute(r StockLineItemsApiPOSTStockLineItemsRequest) (*POSTStockLineItems201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTStockLineItems201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockLineItemsApiService.POSTStockLineItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_line_items"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stockLineItemCreate == nil {
		return localVarReturnValue, nil, reportError("stockLineItemCreate is required and must be specified")
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
	localVarPostBody = r.stockLineItemCreate
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
