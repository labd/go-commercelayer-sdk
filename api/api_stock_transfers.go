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

// StockTransfersApiService StockTransfersApi service
type StockTransfersApiService service

type StockTransfersApiDELETEStockTransfersStockTransferIdRequest struct {
	ctx             context.Context
	ApiService      *StockTransfersApiService
	stockTransferId interface{}
}

func (r StockTransfersApiDELETEStockTransfersStockTransferIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEStockTransfersStockTransferIdExecute(r)
}

/*
DELETEStockTransfersStockTransferId Delete a stock transfer

Delete a stock transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockTransferId The resource's id
	@return StockTransfersApiDELETEStockTransfersStockTransferIdRequest
*/
func (a *StockTransfersApiService) DELETEStockTransfersStockTransferId(ctx context.Context, stockTransferId interface{}) StockTransfersApiDELETEStockTransfersStockTransferIdRequest {
	return StockTransfersApiDELETEStockTransfersStockTransferIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stockTransferId: stockTransferId,
	}
}

// Execute executes the request
func (a *StockTransfersApiService) DELETEStockTransfersStockTransferIdExecute(r StockTransfersApiDELETEStockTransfersStockTransferIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.DELETEStockTransfersStockTransferId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_transfers/{stockTransferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockTransferId"+"}", url.PathEscape(parameterValueToString(r.stockTransferId, "stockTransferId")), -1)

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

type StockTransfersApiGETLineItemIdStockTransfersRequest struct {
	ctx        context.Context
	ApiService *StockTransfersApiService
	lineItemId interface{}
}

func (r StockTransfersApiGETLineItemIdStockTransfersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETLineItemIdStockTransfersExecute(r)
}

/*
GETLineItemIdStockTransfers Retrieve the stock transfers associated to the line item

Retrieve the stock transfers associated to the line item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param lineItemId The resource's id
	@return StockTransfersApiGETLineItemIdStockTransfersRequest
*/
func (a *StockTransfersApiService) GETLineItemIdStockTransfers(ctx context.Context, lineItemId interface{}) StockTransfersApiGETLineItemIdStockTransfersRequest {
	return StockTransfersApiGETLineItemIdStockTransfersRequest{
		ApiService: a,
		ctx:        ctx,
		lineItemId: lineItemId,
	}
}

// Execute executes the request
func (a *StockTransfersApiService) GETLineItemIdStockTransfersExecute(r StockTransfersApiGETLineItemIdStockTransfersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.GETLineItemIdStockTransfers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/line_items/{lineItemId}/stock_transfers"
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

type StockTransfersApiGETOrderIdStockTransfersRequest struct {
	ctx        context.Context
	ApiService *StockTransfersApiService
	orderId    interface{}
}

func (r StockTransfersApiGETOrderIdStockTransfersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdStockTransfersExecute(r)
}

/*
GETOrderIdStockTransfers Retrieve the stock transfers associated to the order

Retrieve the stock transfers associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return StockTransfersApiGETOrderIdStockTransfersRequest
*/
func (a *StockTransfersApiService) GETOrderIdStockTransfers(ctx context.Context, orderId interface{}) StockTransfersApiGETOrderIdStockTransfersRequest {
	return StockTransfersApiGETOrderIdStockTransfersRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *StockTransfersApiService) GETOrderIdStockTransfersExecute(r StockTransfersApiGETOrderIdStockTransfersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.GETOrderIdStockTransfers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/stock_transfers"
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

type StockTransfersApiGETShipmentIdStockTransfersRequest struct {
	ctx        context.Context
	ApiService *StockTransfersApiService
	shipmentId interface{}
}

func (r StockTransfersApiGETShipmentIdStockTransfersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentIdStockTransfersExecute(r)
}

/*
GETShipmentIdStockTransfers Retrieve the stock transfers associated to the shipment

Retrieve the stock transfers associated to the shipment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param shipmentId The resource's id
	@return StockTransfersApiGETShipmentIdStockTransfersRequest
*/
func (a *StockTransfersApiService) GETShipmentIdStockTransfers(ctx context.Context, shipmentId interface{}) StockTransfersApiGETShipmentIdStockTransfersRequest {
	return StockTransfersApiGETShipmentIdStockTransfersRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *StockTransfersApiService) GETShipmentIdStockTransfersExecute(r StockTransfersApiGETShipmentIdStockTransfersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.GETShipmentIdStockTransfers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}/stock_transfers"
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

type StockTransfersApiGETStockLocationIdStockTransfersRequest struct {
	ctx             context.Context
	ApiService      *StockTransfersApiService
	stockLocationId interface{}
}

func (r StockTransfersApiGETStockLocationIdStockTransfersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockLocationIdStockTransfersExecute(r)
}

/*
GETStockLocationIdStockTransfers Retrieve the stock transfers associated to the stock location

Retrieve the stock transfers associated to the stock location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockLocationId The resource's id
	@return StockTransfersApiGETStockLocationIdStockTransfersRequest
*/
func (a *StockTransfersApiService) GETStockLocationIdStockTransfers(ctx context.Context, stockLocationId interface{}) StockTransfersApiGETStockLocationIdStockTransfersRequest {
	return StockTransfersApiGETStockLocationIdStockTransfersRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLocationId: stockLocationId,
	}
}

// Execute executes the request
func (a *StockTransfersApiService) GETStockLocationIdStockTransfersExecute(r StockTransfersApiGETStockLocationIdStockTransfersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.GETStockLocationIdStockTransfers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_locations/{stockLocationId}/stock_transfers"
	localVarPath = strings.Replace(localVarPath, "{"+"stockLocationId"+"}", url.PathEscape(parameterValueToString(r.stockLocationId, "stockLocationId")), -1)

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

type StockTransfersApiGETStockReservationIdStockTransferRequest struct {
	ctx                context.Context
	ApiService         *StockTransfersApiService
	stockReservationId interface{}
}

func (r StockTransfersApiGETStockReservationIdStockTransferRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockReservationIdStockTransferExecute(r)
}

/*
GETStockReservationIdStockTransfer Retrieve the stock transfer associated to the stock reservation

Retrieve the stock transfer associated to the stock reservation

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockReservationId The resource's id
	@return StockTransfersApiGETStockReservationIdStockTransferRequest
*/
func (a *StockTransfersApiService) GETStockReservationIdStockTransfer(ctx context.Context, stockReservationId interface{}) StockTransfersApiGETStockReservationIdStockTransferRequest {
	return StockTransfersApiGETStockReservationIdStockTransferRequest{
		ApiService:         a,
		ctx:                ctx,
		stockReservationId: stockReservationId,
	}
}

// Execute executes the request
func (a *StockTransfersApiService) GETStockReservationIdStockTransferExecute(r StockTransfersApiGETStockReservationIdStockTransferRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.GETStockReservationIdStockTransfer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_reservations/{stockReservationId}/stock_transfer"
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

type StockTransfersApiGETStockTransfersRequest struct {
	ctx        context.Context
	ApiService *StockTransfersApiService
}

func (r StockTransfersApiGETStockTransfersRequest) Execute() (*GETStockTransfers200Response, *http.Response, error) {
	return r.ApiService.GETStockTransfersExecute(r)
}

/*
GETStockTransfers List all stock transfers

List all stock transfers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return StockTransfersApiGETStockTransfersRequest
*/
func (a *StockTransfersApiService) GETStockTransfers(ctx context.Context) StockTransfersApiGETStockTransfersRequest {
	return StockTransfersApiGETStockTransfersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETStockTransfers200Response
func (a *StockTransfersApiService) GETStockTransfersExecute(r StockTransfersApiGETStockTransfersRequest) (*GETStockTransfers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETStockTransfers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.GETStockTransfers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_transfers"

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

type StockTransfersApiGETStockTransfersStockTransferIdRequest struct {
	ctx             context.Context
	ApiService      *StockTransfersApiService
	stockTransferId interface{}
}

func (r StockTransfersApiGETStockTransfersStockTransferIdRequest) Execute() (*GETStockTransfersStockTransferId200Response, *http.Response, error) {
	return r.ApiService.GETStockTransfersStockTransferIdExecute(r)
}

/*
GETStockTransfersStockTransferId Retrieve a stock transfer

Retrieve a stock transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockTransferId The resource's id
	@return StockTransfersApiGETStockTransfersStockTransferIdRequest
*/
func (a *StockTransfersApiService) GETStockTransfersStockTransferId(ctx context.Context, stockTransferId interface{}) StockTransfersApiGETStockTransfersStockTransferIdRequest {
	return StockTransfersApiGETStockTransfersStockTransferIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stockTransferId: stockTransferId,
	}
}

// Execute executes the request
//
//	@return GETStockTransfersStockTransferId200Response
func (a *StockTransfersApiService) GETStockTransfersStockTransferIdExecute(r StockTransfersApiGETStockTransfersStockTransferIdRequest) (*GETStockTransfersStockTransferId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETStockTransfersStockTransferId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.GETStockTransfersStockTransferId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_transfers/{stockTransferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockTransferId"+"}", url.PathEscape(parameterValueToString(r.stockTransferId, "stockTransferId")), -1)

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

type StockTransfersApiPATCHStockTransfersStockTransferIdRequest struct {
	ctx                 context.Context
	ApiService          *StockTransfersApiService
	stockTransferUpdate *StockTransferUpdate
	stockTransferId     interface{}
}

func (r StockTransfersApiPATCHStockTransfersStockTransferIdRequest) StockTransferUpdate(stockTransferUpdate StockTransferUpdate) StockTransfersApiPATCHStockTransfersStockTransferIdRequest {
	r.stockTransferUpdate = &stockTransferUpdate
	return r
}

func (r StockTransfersApiPATCHStockTransfersStockTransferIdRequest) Execute() (*PATCHStockTransfersStockTransferId200Response, *http.Response, error) {
	return r.ApiService.PATCHStockTransfersStockTransferIdExecute(r)
}

/*
PATCHStockTransfersStockTransferId Update a stock transfer

Update a stock transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockTransferId The resource's id
	@return StockTransfersApiPATCHStockTransfersStockTransferIdRequest
*/
func (a *StockTransfersApiService) PATCHStockTransfersStockTransferId(ctx context.Context, stockTransferId interface{}) StockTransfersApiPATCHStockTransfersStockTransferIdRequest {
	return StockTransfersApiPATCHStockTransfersStockTransferIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stockTransferId: stockTransferId,
	}
}

// Execute executes the request
//
//	@return PATCHStockTransfersStockTransferId200Response
func (a *StockTransfersApiService) PATCHStockTransfersStockTransferIdExecute(r StockTransfersApiPATCHStockTransfersStockTransferIdRequest) (*PATCHStockTransfersStockTransferId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHStockTransfersStockTransferId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.PATCHStockTransfersStockTransferId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_transfers/{stockTransferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stockTransferId"+"}", url.PathEscape(parameterValueToString(r.stockTransferId, "stockTransferId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stockTransferUpdate == nil {
		return localVarReturnValue, nil, reportError("stockTransferUpdate is required and must be specified")
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
	localVarPostBody = r.stockTransferUpdate
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

type StockTransfersApiPOSTStockTransfersRequest struct {
	ctx                 context.Context
	ApiService          *StockTransfersApiService
	stockTransferCreate *StockTransferCreate
}

func (r StockTransfersApiPOSTStockTransfersRequest) StockTransferCreate(stockTransferCreate StockTransferCreate) StockTransfersApiPOSTStockTransfersRequest {
	r.stockTransferCreate = &stockTransferCreate
	return r
}

func (r StockTransfersApiPOSTStockTransfersRequest) Execute() (*POSTStockTransfers201Response, *http.Response, error) {
	return r.ApiService.POSTStockTransfersExecute(r)
}

/*
POSTStockTransfers Create a stock transfer

Create a stock transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return StockTransfersApiPOSTStockTransfersRequest
*/
func (a *StockTransfersApiService) POSTStockTransfers(ctx context.Context) StockTransfersApiPOSTStockTransfersRequest {
	return StockTransfersApiPOSTStockTransfersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTStockTransfers201Response
func (a *StockTransfersApiService) POSTStockTransfersExecute(r StockTransfersApiPOSTStockTransfersRequest) (*POSTStockTransfers201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTStockTransfers201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockTransfersApiService.POSTStockTransfers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_transfers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stockTransferCreate == nil {
		return localVarReturnValue, nil, reportError("stockTransferCreate is required and must be specified")
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
	localVarPostBody = r.stockTransferCreate
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
