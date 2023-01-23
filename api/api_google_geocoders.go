/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
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

// GoogleGeocodersApiService GoogleGeocodersApi service
type GoogleGeocodersApiService service

type GoogleGeocodersApiDELETEGoogleGeocodersGoogleGeocoderIdRequest struct {
	ctx              context.Context
	ApiService       *GoogleGeocodersApiService
	googleGeocoderId string
}

func (r GoogleGeocodersApiDELETEGoogleGeocodersGoogleGeocoderIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEGoogleGeocodersGoogleGeocoderIdExecute(r)
}

/*
DELETEGoogleGeocodersGoogleGeocoderId Delete a google geocoder

Delete a google geocoder

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param googleGeocoderId The resource's id
	@return GoogleGeocodersApiDELETEGoogleGeocodersGoogleGeocoderIdRequest
*/
func (a *GoogleGeocodersApiService) DELETEGoogleGeocodersGoogleGeocoderId(ctx context.Context, googleGeocoderId string) GoogleGeocodersApiDELETEGoogleGeocodersGoogleGeocoderIdRequest {
	return GoogleGeocodersApiDELETEGoogleGeocodersGoogleGeocoderIdRequest{
		ApiService:       a,
		ctx:              ctx,
		googleGeocoderId: googleGeocoderId,
	}
}

// Execute executes the request
func (a *GoogleGeocodersApiService) DELETEGoogleGeocodersGoogleGeocoderIdExecute(r GoogleGeocodersApiDELETEGoogleGeocodersGoogleGeocoderIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.DELETEGoogleGeocodersGoogleGeocoderId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders/{googleGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"googleGeocoderId"+"}", url.PathEscape(parameterToString(r.googleGeocoderId, "")), -1)

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

type GoogleGeocodersApiGETGoogleGeocodersRequest struct {
	ctx        context.Context
	ApiService *GoogleGeocodersApiService
}

func (r GoogleGeocodersApiGETGoogleGeocodersRequest) Execute() (*GETGoogleGeocoders200Response, *http.Response, error) {
	return r.ApiService.GETGoogleGeocodersExecute(r)
}

/*
GETGoogleGeocoders List all google geocoders

List all google geocoders

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GoogleGeocodersApiGETGoogleGeocodersRequest
*/
func (a *GoogleGeocodersApiService) GETGoogleGeocoders(ctx context.Context) GoogleGeocodersApiGETGoogleGeocodersRequest {
	return GoogleGeocodersApiGETGoogleGeocodersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETGoogleGeocoders200Response
func (a *GoogleGeocodersApiService) GETGoogleGeocodersExecute(r GoogleGeocodersApiGETGoogleGeocodersRequest) (*GETGoogleGeocoders200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETGoogleGeocoders200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.GETGoogleGeocoders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders"

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

type GoogleGeocodersApiGETGoogleGeocodersGoogleGeocoderIdRequest struct {
	ctx              context.Context
	ApiService       *GoogleGeocodersApiService
	googleGeocoderId string
}

func (r GoogleGeocodersApiGETGoogleGeocodersGoogleGeocoderIdRequest) Execute() (*GETGoogleGeocodersGoogleGeocoderId200Response, *http.Response, error) {
	return r.ApiService.GETGoogleGeocodersGoogleGeocoderIdExecute(r)
}

/*
GETGoogleGeocodersGoogleGeocoderId Retrieve a google geocoder

Retrieve a google geocoder

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param googleGeocoderId The resource's id
	@return GoogleGeocodersApiGETGoogleGeocodersGoogleGeocoderIdRequest
*/
func (a *GoogleGeocodersApiService) GETGoogleGeocodersGoogleGeocoderId(ctx context.Context, googleGeocoderId string) GoogleGeocodersApiGETGoogleGeocodersGoogleGeocoderIdRequest {
	return GoogleGeocodersApiGETGoogleGeocodersGoogleGeocoderIdRequest{
		ApiService:       a,
		ctx:              ctx,
		googleGeocoderId: googleGeocoderId,
	}
}

// Execute executes the request
//
//	@return GETGoogleGeocodersGoogleGeocoderId200Response
func (a *GoogleGeocodersApiService) GETGoogleGeocodersGoogleGeocoderIdExecute(r GoogleGeocodersApiGETGoogleGeocodersGoogleGeocoderIdRequest) (*GETGoogleGeocodersGoogleGeocoderId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETGoogleGeocodersGoogleGeocoderId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.GETGoogleGeocodersGoogleGeocoderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders/{googleGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"googleGeocoderId"+"}", url.PathEscape(parameterToString(r.googleGeocoderId, "")), -1)

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

type GoogleGeocodersApiPATCHGoogleGeocodersGoogleGeocoderIdRequest struct {
	ctx                  context.Context
	ApiService           *GoogleGeocodersApiService
	googleGeocoderUpdate *GoogleGeocoderUpdate
	googleGeocoderId     string
}

func (r GoogleGeocodersApiPATCHGoogleGeocodersGoogleGeocoderIdRequest) GoogleGeocoderUpdate(googleGeocoderUpdate GoogleGeocoderUpdate) GoogleGeocodersApiPATCHGoogleGeocodersGoogleGeocoderIdRequest {
	r.googleGeocoderUpdate = &googleGeocoderUpdate
	return r
}

func (r GoogleGeocodersApiPATCHGoogleGeocodersGoogleGeocoderIdRequest) Execute() (*PATCHGoogleGeocodersGoogleGeocoderId200Response, *http.Response, error) {
	return r.ApiService.PATCHGoogleGeocodersGoogleGeocoderIdExecute(r)
}

/*
PATCHGoogleGeocodersGoogleGeocoderId Update a google geocoder

Update a google geocoder

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param googleGeocoderId The resource's id
	@return GoogleGeocodersApiPATCHGoogleGeocodersGoogleGeocoderIdRequest
*/
func (a *GoogleGeocodersApiService) PATCHGoogleGeocodersGoogleGeocoderId(ctx context.Context, googleGeocoderId string) GoogleGeocodersApiPATCHGoogleGeocodersGoogleGeocoderIdRequest {
	return GoogleGeocodersApiPATCHGoogleGeocodersGoogleGeocoderIdRequest{
		ApiService:       a,
		ctx:              ctx,
		googleGeocoderId: googleGeocoderId,
	}
}

// Execute executes the request
//
//	@return PATCHGoogleGeocodersGoogleGeocoderId200Response
func (a *GoogleGeocodersApiService) PATCHGoogleGeocodersGoogleGeocoderIdExecute(r GoogleGeocodersApiPATCHGoogleGeocodersGoogleGeocoderIdRequest) (*PATCHGoogleGeocodersGoogleGeocoderId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHGoogleGeocodersGoogleGeocoderId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.PATCHGoogleGeocodersGoogleGeocoderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders/{googleGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"googleGeocoderId"+"}", url.PathEscape(parameterToString(r.googleGeocoderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.googleGeocoderUpdate == nil {
		return localVarReturnValue, nil, reportError("googleGeocoderUpdate is required and must be specified")
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
	localVarPostBody = r.googleGeocoderUpdate
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

type GoogleGeocodersApiPOSTGoogleGeocodersRequest struct {
	ctx                  context.Context
	ApiService           *GoogleGeocodersApiService
	googleGeocoderCreate *GoogleGeocoderCreate
}

func (r GoogleGeocodersApiPOSTGoogleGeocodersRequest) GoogleGeocoderCreate(googleGeocoderCreate GoogleGeocoderCreate) GoogleGeocodersApiPOSTGoogleGeocodersRequest {
	r.googleGeocoderCreate = &googleGeocoderCreate
	return r
}

func (r GoogleGeocodersApiPOSTGoogleGeocodersRequest) Execute() (*POSTGoogleGeocoders201Response, *http.Response, error) {
	return r.ApiService.POSTGoogleGeocodersExecute(r)
}

/*
POSTGoogleGeocoders Create a google geocoder

Create a google geocoder

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GoogleGeocodersApiPOSTGoogleGeocodersRequest
*/
func (a *GoogleGeocodersApiService) POSTGoogleGeocoders(ctx context.Context) GoogleGeocodersApiPOSTGoogleGeocodersRequest {
	return GoogleGeocodersApiPOSTGoogleGeocodersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTGoogleGeocoders201Response
func (a *GoogleGeocodersApiService) POSTGoogleGeocodersExecute(r GoogleGeocodersApiPOSTGoogleGeocodersRequest) (*POSTGoogleGeocoders201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTGoogleGeocoders201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.POSTGoogleGeocoders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.googleGeocoderCreate == nil {
		return localVarReturnValue, nil, reportError("googleGeocoderCreate is required and must be specified")
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
	localVarPostBody = r.googleGeocoderCreate
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
