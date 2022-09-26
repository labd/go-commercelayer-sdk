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

// GeocodersApiService GeocodersApi service
type GeocodersApiService service

type GeocodersApiGETAddressIdGeocoderRequest struct {
	ctx        context.Context
	ApiService *GeocodersApiService
	addressId  string
}

func (r GeocodersApiGETAddressIdGeocoderRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETAddressIdGeocoderExecute(r)
}

/*
GETAddressIdGeocoder Retrieve the geocoder associated to the address

Retrieve the geocoder associated to the address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param addressId The resource's id
 @return GeocodersApiGETAddressIdGeocoderRequest
*/
func (a *GeocodersApiService) GETAddressIdGeocoder(ctx context.Context, addressId string) GeocodersApiGETAddressIdGeocoderRequest {
	return GeocodersApiGETAddressIdGeocoderRequest{
		ApiService: a,
		ctx:        ctx,
		addressId:  addressId,
	}
}

// Execute executes the request
func (a *GeocodersApiService) GETAddressIdGeocoderExecute(r GeocodersApiGETAddressIdGeocoderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GeocodersApiService.GETAddressIdGeocoder")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/addresses/{addressId}/geocoder"
	localVarPath = strings.Replace(localVarPath, "{"+"addressId"+"}", url.PathEscape(parameterToString(r.addressId, "")), -1)

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

type GeocodersApiGETGeocodersRequest struct {
	ctx        context.Context
	ApiService *GeocodersApiService
}

func (r GeocodersApiGETGeocodersRequest) Execute() (*GeocoderResponseList, *http.Response, error) {
	return r.ApiService.GETGeocodersExecute(r)
}

/*
GETGeocoders List all geocoders

List all geocoders

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GeocodersApiGETGeocodersRequest
*/
func (a *GeocodersApiService) GETGeocoders(ctx context.Context) GeocodersApiGETGeocodersRequest {
	return GeocodersApiGETGeocodersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GeocoderResponseList
func (a *GeocodersApiService) GETGeocodersExecute(r GeocodersApiGETGeocodersRequest) (*GeocoderResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GeocoderResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GeocodersApiService.GETGeocoders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/geocoders"

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

type GeocodersApiGETGeocodersGeocoderIdRequest struct {
	ctx        context.Context
	ApiService *GeocodersApiService
	geocoderId string
}

func (r GeocodersApiGETGeocodersGeocoderIdRequest) Execute() (*GeocoderResponse, *http.Response, error) {
	return r.ApiService.GETGeocodersGeocoderIdExecute(r)
}

/*
GETGeocodersGeocoderId Retrieve a geocoder

Retrieve a geocoder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param geocoderId The resource's id
 @return GeocodersApiGETGeocodersGeocoderIdRequest
*/
func (a *GeocodersApiService) GETGeocodersGeocoderId(ctx context.Context, geocoderId string) GeocodersApiGETGeocodersGeocoderIdRequest {
	return GeocodersApiGETGeocodersGeocoderIdRequest{
		ApiService: a,
		ctx:        ctx,
		geocoderId: geocoderId,
	}
}

// Execute executes the request
//  @return GeocoderResponse
func (a *GeocodersApiService) GETGeocodersGeocoderIdExecute(r GeocodersApiGETGeocodersGeocoderIdRequest) (*GeocoderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GeocoderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GeocodersApiService.GETGeocodersGeocoderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/geocoders/{geocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"geocoderId"+"}", url.PathEscape(parameterToString(r.geocoderId, "")), -1)

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
