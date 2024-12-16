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

// InventoryReturnLocationsApiService InventoryReturnLocationsApi service
type InventoryReturnLocationsApiService service

type InventoryReturnLocationsApiDELETEInventoryReturnLocationsInventoryReturnLocationIdRequest struct {
	ctx                       context.Context
	ApiService                *InventoryReturnLocationsApiService
	inventoryReturnLocationId interface{}
}

func (r InventoryReturnLocationsApiDELETEInventoryReturnLocationsInventoryReturnLocationIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEInventoryReturnLocationsInventoryReturnLocationIdExecute(r)
}

/*
DELETEInventoryReturnLocationsInventoryReturnLocationId Delete an inventory return location

Delete an inventory return location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryReturnLocationId The resource's id
	@return InventoryReturnLocationsApiDELETEInventoryReturnLocationsInventoryReturnLocationIdRequest
*/
func (a *InventoryReturnLocationsApiService) DELETEInventoryReturnLocationsInventoryReturnLocationId(ctx context.Context, inventoryReturnLocationId interface{}) InventoryReturnLocationsApiDELETEInventoryReturnLocationsInventoryReturnLocationIdRequest {
	return InventoryReturnLocationsApiDELETEInventoryReturnLocationsInventoryReturnLocationIdRequest{
		ApiService:                a,
		ctx:                       ctx,
		inventoryReturnLocationId: inventoryReturnLocationId,
	}
}

// Execute executes the request
func (a *InventoryReturnLocationsApiService) DELETEInventoryReturnLocationsInventoryReturnLocationIdExecute(r InventoryReturnLocationsApiDELETEInventoryReturnLocationsInventoryReturnLocationIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryReturnLocationsApiService.DELETEInventoryReturnLocationsInventoryReturnLocationId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_return_locations/{inventoryReturnLocationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryReturnLocationId"+"}", url.PathEscape(parameterValueToString(r.inventoryReturnLocationId, "inventoryReturnLocationId")), -1)

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

type InventoryReturnLocationsApiGETInventoryModelIdInventoryReturnLocationsRequest struct {
	ctx              context.Context
	ApiService       *InventoryReturnLocationsApiService
	inventoryModelId interface{}
}

func (r InventoryReturnLocationsApiGETInventoryModelIdInventoryReturnLocationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETInventoryModelIdInventoryReturnLocationsExecute(r)
}

/*
GETInventoryModelIdInventoryReturnLocations Retrieve the inventory return locations associated to the inventory model

Retrieve the inventory return locations associated to the inventory model

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryModelId The resource's id
	@return InventoryReturnLocationsApiGETInventoryModelIdInventoryReturnLocationsRequest
*/
func (a *InventoryReturnLocationsApiService) GETInventoryModelIdInventoryReturnLocations(ctx context.Context, inventoryModelId interface{}) InventoryReturnLocationsApiGETInventoryModelIdInventoryReturnLocationsRequest {
	return InventoryReturnLocationsApiGETInventoryModelIdInventoryReturnLocationsRequest{
		ApiService:       a,
		ctx:              ctx,
		inventoryModelId: inventoryModelId,
	}
}

// Execute executes the request
func (a *InventoryReturnLocationsApiService) GETInventoryModelIdInventoryReturnLocationsExecute(r InventoryReturnLocationsApiGETInventoryModelIdInventoryReturnLocationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryReturnLocationsApiService.GETInventoryModelIdInventoryReturnLocations")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_models/{inventoryModelId}/inventory_return_locations"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryModelId"+"}", url.PathEscape(parameterValueToString(r.inventoryModelId, "inventoryModelId")), -1)

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

type InventoryReturnLocationsApiGETInventoryReturnLocationsRequest struct {
	ctx        context.Context
	ApiService *InventoryReturnLocationsApiService
}

func (r InventoryReturnLocationsApiGETInventoryReturnLocationsRequest) Execute() (*GETInventoryReturnLocations200Response, *http.Response, error) {
	return r.ApiService.GETInventoryReturnLocationsExecute(r)
}

/*
GETInventoryReturnLocations List all inventory return locations

List all inventory return locations

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InventoryReturnLocationsApiGETInventoryReturnLocationsRequest
*/
func (a *InventoryReturnLocationsApiService) GETInventoryReturnLocations(ctx context.Context) InventoryReturnLocationsApiGETInventoryReturnLocationsRequest {
	return InventoryReturnLocationsApiGETInventoryReturnLocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETInventoryReturnLocations200Response
func (a *InventoryReturnLocationsApiService) GETInventoryReturnLocationsExecute(r InventoryReturnLocationsApiGETInventoryReturnLocationsRequest) (*GETInventoryReturnLocations200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETInventoryReturnLocations200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryReturnLocationsApiService.GETInventoryReturnLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_return_locations"

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

type InventoryReturnLocationsApiGETInventoryReturnLocationsInventoryReturnLocationIdRequest struct {
	ctx                       context.Context
	ApiService                *InventoryReturnLocationsApiService
	inventoryReturnLocationId interface{}
}

func (r InventoryReturnLocationsApiGETInventoryReturnLocationsInventoryReturnLocationIdRequest) Execute() (*GETInventoryReturnLocationsInventoryReturnLocationId200Response, *http.Response, error) {
	return r.ApiService.GETInventoryReturnLocationsInventoryReturnLocationIdExecute(r)
}

/*
GETInventoryReturnLocationsInventoryReturnLocationId Retrieve an inventory return location

Retrieve an inventory return location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryReturnLocationId The resource's id
	@return InventoryReturnLocationsApiGETInventoryReturnLocationsInventoryReturnLocationIdRequest
*/
func (a *InventoryReturnLocationsApiService) GETInventoryReturnLocationsInventoryReturnLocationId(ctx context.Context, inventoryReturnLocationId interface{}) InventoryReturnLocationsApiGETInventoryReturnLocationsInventoryReturnLocationIdRequest {
	return InventoryReturnLocationsApiGETInventoryReturnLocationsInventoryReturnLocationIdRequest{
		ApiService:                a,
		ctx:                       ctx,
		inventoryReturnLocationId: inventoryReturnLocationId,
	}
}

// Execute executes the request
//
//	@return GETInventoryReturnLocationsInventoryReturnLocationId200Response
func (a *InventoryReturnLocationsApiService) GETInventoryReturnLocationsInventoryReturnLocationIdExecute(r InventoryReturnLocationsApiGETInventoryReturnLocationsInventoryReturnLocationIdRequest) (*GETInventoryReturnLocationsInventoryReturnLocationId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETInventoryReturnLocationsInventoryReturnLocationId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryReturnLocationsApiService.GETInventoryReturnLocationsInventoryReturnLocationId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_return_locations/{inventoryReturnLocationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryReturnLocationId"+"}", url.PathEscape(parameterValueToString(r.inventoryReturnLocationId, "inventoryReturnLocationId")), -1)

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

type InventoryReturnLocationsApiGETStockLocationIdInventoryReturnLocationsRequest struct {
	ctx             context.Context
	ApiService      *InventoryReturnLocationsApiService
	stockLocationId interface{}
}

func (r InventoryReturnLocationsApiGETStockLocationIdInventoryReturnLocationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockLocationIdInventoryReturnLocationsExecute(r)
}

/*
GETStockLocationIdInventoryReturnLocations Retrieve the inventory return locations associated to the stock location

Retrieve the inventory return locations associated to the stock location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockLocationId The resource's id
	@return InventoryReturnLocationsApiGETStockLocationIdInventoryReturnLocationsRequest
*/
func (a *InventoryReturnLocationsApiService) GETStockLocationIdInventoryReturnLocations(ctx context.Context, stockLocationId interface{}) InventoryReturnLocationsApiGETStockLocationIdInventoryReturnLocationsRequest {
	return InventoryReturnLocationsApiGETStockLocationIdInventoryReturnLocationsRequest{
		ApiService:      a,
		ctx:             ctx,
		stockLocationId: stockLocationId,
	}
}

// Execute executes the request
func (a *InventoryReturnLocationsApiService) GETStockLocationIdInventoryReturnLocationsExecute(r InventoryReturnLocationsApiGETStockLocationIdInventoryReturnLocationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryReturnLocationsApiService.GETStockLocationIdInventoryReturnLocations")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_locations/{stockLocationId}/inventory_return_locations"
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

type InventoryReturnLocationsApiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest struct {
	ctx                           context.Context
	ApiService                    *InventoryReturnLocationsApiService
	inventoryReturnLocationUpdate *InventoryReturnLocationUpdate
	inventoryReturnLocationId     interface{}
}

func (r InventoryReturnLocationsApiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest) InventoryReturnLocationUpdate(inventoryReturnLocationUpdate InventoryReturnLocationUpdate) InventoryReturnLocationsApiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest {
	r.inventoryReturnLocationUpdate = &inventoryReturnLocationUpdate
	return r
}

func (r InventoryReturnLocationsApiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest) Execute() (*PATCHInventoryReturnLocationsInventoryReturnLocationId200Response, *http.Response, error) {
	return r.ApiService.PATCHInventoryReturnLocationsInventoryReturnLocationIdExecute(r)
}

/*
PATCHInventoryReturnLocationsInventoryReturnLocationId Update an inventory return location

Update an inventory return location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param inventoryReturnLocationId The resource's id
	@return InventoryReturnLocationsApiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest
*/
func (a *InventoryReturnLocationsApiService) PATCHInventoryReturnLocationsInventoryReturnLocationId(ctx context.Context, inventoryReturnLocationId interface{}) InventoryReturnLocationsApiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest {
	return InventoryReturnLocationsApiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest{
		ApiService:                a,
		ctx:                       ctx,
		inventoryReturnLocationId: inventoryReturnLocationId,
	}
}

// Execute executes the request
//
//	@return PATCHInventoryReturnLocationsInventoryReturnLocationId200Response
func (a *InventoryReturnLocationsApiService) PATCHInventoryReturnLocationsInventoryReturnLocationIdExecute(r InventoryReturnLocationsApiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest) (*PATCHInventoryReturnLocationsInventoryReturnLocationId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHInventoryReturnLocationsInventoryReturnLocationId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryReturnLocationsApiService.PATCHInventoryReturnLocationsInventoryReturnLocationId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_return_locations/{inventoryReturnLocationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryReturnLocationId"+"}", url.PathEscape(parameterValueToString(r.inventoryReturnLocationId, "inventoryReturnLocationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inventoryReturnLocationUpdate == nil {
		return localVarReturnValue, nil, reportError("inventoryReturnLocationUpdate is required and must be specified")
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
	localVarPostBody = r.inventoryReturnLocationUpdate
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

type InventoryReturnLocationsApiPOSTInventoryReturnLocationsRequest struct {
	ctx                           context.Context
	ApiService                    *InventoryReturnLocationsApiService
	inventoryReturnLocationCreate *InventoryReturnLocationCreate
}

func (r InventoryReturnLocationsApiPOSTInventoryReturnLocationsRequest) InventoryReturnLocationCreate(inventoryReturnLocationCreate InventoryReturnLocationCreate) InventoryReturnLocationsApiPOSTInventoryReturnLocationsRequest {
	r.inventoryReturnLocationCreate = &inventoryReturnLocationCreate
	return r
}

func (r InventoryReturnLocationsApiPOSTInventoryReturnLocationsRequest) Execute() (*POSTInventoryReturnLocations201Response, *http.Response, error) {
	return r.ApiService.POSTInventoryReturnLocationsExecute(r)
}

/*
POSTInventoryReturnLocations Create an inventory return location

Create an inventory return location

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return InventoryReturnLocationsApiPOSTInventoryReturnLocationsRequest
*/
func (a *InventoryReturnLocationsApiService) POSTInventoryReturnLocations(ctx context.Context) InventoryReturnLocationsApiPOSTInventoryReturnLocationsRequest {
	return InventoryReturnLocationsApiPOSTInventoryReturnLocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTInventoryReturnLocations201Response
func (a *InventoryReturnLocationsApiService) POSTInventoryReturnLocationsExecute(r InventoryReturnLocationsApiPOSTInventoryReturnLocationsRequest) (*POSTInventoryReturnLocations201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTInventoryReturnLocations201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryReturnLocationsApiService.POSTInventoryReturnLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory_return_locations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inventoryReturnLocationCreate == nil {
		return localVarReturnValue, nil, reportError("inventoryReturnLocationCreate is required and must be specified")
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
	localVarPostBody = r.inventoryReturnLocationCreate
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
