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

// ParcelsApiService ParcelsApi service
type ParcelsApiService service

type ParcelsApiDELETEParcelsParcelIdRequest struct {
	ctx        context.Context
	ApiService *ParcelsApiService
	parcelId   string
}

func (r ParcelsApiDELETEParcelsParcelIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEParcelsParcelIdExecute(r)
}

/*
DELETEParcelsParcelId Delete a parcel

Delete a parcel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parcelId The resource's id
 @return ParcelsApiDELETEParcelsParcelIdRequest
*/
func (a *ParcelsApiService) DELETEParcelsParcelId(ctx context.Context, parcelId string) ParcelsApiDELETEParcelsParcelIdRequest {
	return ParcelsApiDELETEParcelsParcelIdRequest{
		ApiService: a,
		ctx:        ctx,
		parcelId:   parcelId,
	}
}

// Execute executes the request
func (a *ParcelsApiService) DELETEParcelsParcelIdExecute(r ParcelsApiDELETEParcelsParcelIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelsApiService.DELETEParcelsParcelId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcels/{parcelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelId"+"}", url.PathEscape(parameterToString(r.parcelId, "")), -1)

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

type ParcelsApiGETPackageIdParcelsRequest struct {
	ctx        context.Context
	ApiService *ParcelsApiService
	packageId  string
}

func (r ParcelsApiGETPackageIdParcelsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPackageIdParcelsExecute(r)
}

/*
GETPackageIdParcels Retrieve the parcels associated to the package

Retrieve the parcels associated to the package

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param packageId The resource's id
 @return ParcelsApiGETPackageIdParcelsRequest
*/
func (a *ParcelsApiService) GETPackageIdParcels(ctx context.Context, packageId string) ParcelsApiGETPackageIdParcelsRequest {
	return ParcelsApiGETPackageIdParcelsRequest{
		ApiService: a,
		ctx:        ctx,
		packageId:  packageId,
	}
}

// Execute executes the request
func (a *ParcelsApiService) GETPackageIdParcelsExecute(r ParcelsApiGETPackageIdParcelsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelsApiService.GETPackageIdParcels")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages/{packageId}/parcels"
	localVarPath = strings.Replace(localVarPath, "{"+"packageId"+"}", url.PathEscape(parameterToString(r.packageId, "")), -1)

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

type ParcelsApiGETParcelLineItemIdParcelRequest struct {
	ctx              context.Context
	ApiService       *ParcelsApiService
	parcelLineItemId string
}

func (r ParcelsApiGETParcelLineItemIdParcelRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETParcelLineItemIdParcelExecute(r)
}

/*
GETParcelLineItemIdParcel Retrieve the parcel associated to the parcel line item

Retrieve the parcel associated to the parcel line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parcelLineItemId The resource's id
 @return ParcelsApiGETParcelLineItemIdParcelRequest
*/
func (a *ParcelsApiService) GETParcelLineItemIdParcel(ctx context.Context, parcelLineItemId string) ParcelsApiGETParcelLineItemIdParcelRequest {
	return ParcelsApiGETParcelLineItemIdParcelRequest{
		ApiService:       a,
		ctx:              ctx,
		parcelLineItemId: parcelLineItemId,
	}
}

// Execute executes the request
func (a *ParcelsApiService) GETParcelLineItemIdParcelExecute(r ParcelsApiGETParcelLineItemIdParcelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelsApiService.GETParcelLineItemIdParcel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcel_line_items/{parcelLineItemId}/parcel"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelLineItemId"+"}", url.PathEscape(parameterToString(r.parcelLineItemId, "")), -1)

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

type ParcelsApiGETParcelsRequest struct {
	ctx        context.Context
	ApiService *ParcelsApiService
}

func (r ParcelsApiGETParcelsRequest) Execute() (*ParcelResponseList, *http.Response, error) {
	return r.ApiService.GETParcelsExecute(r)
}

/*
GETParcels List all parcels

List all parcels

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ParcelsApiGETParcelsRequest
*/
func (a *ParcelsApiService) GETParcels(ctx context.Context) ParcelsApiGETParcelsRequest {
	return ParcelsApiGETParcelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return ParcelResponseList
func (a *ParcelsApiService) GETParcelsExecute(r ParcelsApiGETParcelsRequest) (*ParcelResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ParcelResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelsApiService.GETParcels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcels"

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

type ParcelsApiGETParcelsParcelIdRequest struct {
	ctx        context.Context
	ApiService *ParcelsApiService
	parcelId   string
}

func (r ParcelsApiGETParcelsParcelIdRequest) Execute() (*ParcelResponse, *http.Response, error) {
	return r.ApiService.GETParcelsParcelIdExecute(r)
}

/*
GETParcelsParcelId Retrieve a parcel

Retrieve a parcel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parcelId The resource's id
 @return ParcelsApiGETParcelsParcelIdRequest
*/
func (a *ParcelsApiService) GETParcelsParcelId(ctx context.Context, parcelId string) ParcelsApiGETParcelsParcelIdRequest {
	return ParcelsApiGETParcelsParcelIdRequest{
		ApiService: a,
		ctx:        ctx,
		parcelId:   parcelId,
	}
}

// Execute executes the request
//  @return ParcelResponse
func (a *ParcelsApiService) GETParcelsParcelIdExecute(r ParcelsApiGETParcelsParcelIdRequest) (*ParcelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ParcelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelsApiService.GETParcelsParcelId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcels/{parcelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelId"+"}", url.PathEscape(parameterToString(r.parcelId, "")), -1)

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

type ParcelsApiGETShipmentIdParcelsRequest struct {
	ctx        context.Context
	ApiService *ParcelsApiService
	shipmentId string
}

func (r ParcelsApiGETShipmentIdParcelsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShipmentIdParcelsExecute(r)
}

/*
GETShipmentIdParcels Retrieve the parcels associated to the shipment

Retrieve the parcels associated to the shipment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shipmentId The resource's id
 @return ParcelsApiGETShipmentIdParcelsRequest
*/
func (a *ParcelsApiService) GETShipmentIdParcels(ctx context.Context, shipmentId string) ParcelsApiGETShipmentIdParcelsRequest {
	return ParcelsApiGETShipmentIdParcelsRequest{
		ApiService: a,
		ctx:        ctx,
		shipmentId: shipmentId,
	}
}

// Execute executes the request
func (a *ParcelsApiService) GETShipmentIdParcelsExecute(r ParcelsApiGETShipmentIdParcelsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelsApiService.GETShipmentIdParcels")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipments/{shipmentId}/parcels"
	localVarPath = strings.Replace(localVarPath, "{"+"shipmentId"+"}", url.PathEscape(parameterToString(r.shipmentId, "")), -1)

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

type ParcelsApiPATCHParcelsParcelIdRequest struct {
	ctx          context.Context
	ApiService   *ParcelsApiService
	parcelUpdate *ParcelUpdate
	parcelId     string
}

func (r ParcelsApiPATCHParcelsParcelIdRequest) ParcelUpdate(parcelUpdate ParcelUpdate) ParcelsApiPATCHParcelsParcelIdRequest {
	r.parcelUpdate = &parcelUpdate
	return r
}

func (r ParcelsApiPATCHParcelsParcelIdRequest) Execute() (*ParcelResponse, *http.Response, error) {
	return r.ApiService.PATCHParcelsParcelIdExecute(r)
}

/*
PATCHParcelsParcelId Update a parcel

Update a parcel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parcelId The resource's id
 @return ParcelsApiPATCHParcelsParcelIdRequest
*/
func (a *ParcelsApiService) PATCHParcelsParcelId(ctx context.Context, parcelId string) ParcelsApiPATCHParcelsParcelIdRequest {
	return ParcelsApiPATCHParcelsParcelIdRequest{
		ApiService: a,
		ctx:        ctx,
		parcelId:   parcelId,
	}
}

// Execute executes the request
//  @return ParcelResponse
func (a *ParcelsApiService) PATCHParcelsParcelIdExecute(r ParcelsApiPATCHParcelsParcelIdRequest) (*ParcelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ParcelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelsApiService.PATCHParcelsParcelId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcels/{parcelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"parcelId"+"}", url.PathEscape(parameterToString(r.parcelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.parcelUpdate == nil {
		return localVarReturnValue, nil, reportError("parcelUpdate is required and must be specified")
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
	localVarPostBody = r.parcelUpdate
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

type ParcelsApiPOSTParcelsRequest struct {
	ctx          context.Context
	ApiService   *ParcelsApiService
	parcelCreate *ParcelCreate
}

func (r ParcelsApiPOSTParcelsRequest) ParcelCreate(parcelCreate ParcelCreate) ParcelsApiPOSTParcelsRequest {
	r.parcelCreate = &parcelCreate
	return r
}

func (r ParcelsApiPOSTParcelsRequest) Execute() (*ParcelResponse, *http.Response, error) {
	return r.ApiService.POSTParcelsExecute(r)
}

/*
POSTParcels Create a parcel

Create a parcel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ParcelsApiPOSTParcelsRequest
*/
func (a *ParcelsApiService) POSTParcels(ctx context.Context) ParcelsApiPOSTParcelsRequest {
	return ParcelsApiPOSTParcelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return ParcelResponse
func (a *ParcelsApiService) POSTParcelsExecute(r ParcelsApiPOSTParcelsRequest) (*ParcelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ParcelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParcelsApiService.POSTParcels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parcels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.parcelCreate == nil {
		return localVarReturnValue, nil, reportError("parcelCreate is required and must be specified")
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
	localVarPostBody = r.parcelCreate
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
