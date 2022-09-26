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

// ExternalPaymentsApiService ExternalPaymentsApi service
type ExternalPaymentsApiService service

type ExternalPaymentsApiDELETEExternalPaymentsExternalPaymentIdRequest struct {
	ctx               context.Context
	ApiService        *ExternalPaymentsApiService
	externalPaymentId string
}

func (r ExternalPaymentsApiDELETEExternalPaymentsExternalPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEExternalPaymentsExternalPaymentIdExecute(r)
}

/*
DELETEExternalPaymentsExternalPaymentId Delete an external payment

Delete an external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPaymentId The resource's id
 @return ExternalPaymentsApiDELETEExternalPaymentsExternalPaymentIdRequest
*/
func (a *ExternalPaymentsApiService) DELETEExternalPaymentsExternalPaymentId(ctx context.Context, externalPaymentId string) ExternalPaymentsApiDELETEExternalPaymentsExternalPaymentIdRequest {
	return ExternalPaymentsApiDELETEExternalPaymentsExternalPaymentIdRequest{
		ApiService:        a,
		ctx:               ctx,
		externalPaymentId: externalPaymentId,
	}
}

// Execute executes the request
func (a *ExternalPaymentsApiService) DELETEExternalPaymentsExternalPaymentIdExecute(r ExternalPaymentsApiDELETEExternalPaymentsExternalPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.DELETEExternalPaymentsExternalPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments/{externalPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPaymentId"+"}", url.PathEscape(parameterToString(r.externalPaymentId, "")), -1)

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

type ExternalPaymentsApiGETExternalGatewayIdExternalPaymentsRequest struct {
	ctx               context.Context
	ApiService        *ExternalPaymentsApiService
	externalGatewayId string
}

func (r ExternalPaymentsApiGETExternalGatewayIdExternalPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETExternalGatewayIdExternalPaymentsExecute(r)
}

/*
GETExternalGatewayIdExternalPayments Retrieve the external payments associated to the external gateway

Retrieve the external payments associated to the external gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGatewayId The resource's id
 @return ExternalPaymentsApiGETExternalGatewayIdExternalPaymentsRequest
*/
func (a *ExternalPaymentsApiService) GETExternalGatewayIdExternalPayments(ctx context.Context, externalGatewayId string) ExternalPaymentsApiGETExternalGatewayIdExternalPaymentsRequest {
	return ExternalPaymentsApiGETExternalGatewayIdExternalPaymentsRequest{
		ApiService:        a,
		ctx:               ctx,
		externalGatewayId: externalGatewayId,
	}
}

// Execute executes the request
func (a *ExternalPaymentsApiService) GETExternalGatewayIdExternalPaymentsExecute(r ExternalPaymentsApiGETExternalGatewayIdExternalPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.GETExternalGatewayIdExternalPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_gateways/{externalGatewayId}/external_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGatewayId"+"}", url.PathEscape(parameterToString(r.externalGatewayId, "")), -1)

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

type ExternalPaymentsApiGETExternalPaymentsRequest struct {
	ctx        context.Context
	ApiService *ExternalPaymentsApiService
}

func (r ExternalPaymentsApiGETExternalPaymentsRequest) Execute() (*ExternalPaymentResponseList, *http.Response, error) {
	return r.ApiService.GETExternalPaymentsExecute(r)
}

/*
GETExternalPayments List all external payments

List all external payments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ExternalPaymentsApiGETExternalPaymentsRequest
*/
func (a *ExternalPaymentsApiService) GETExternalPayments(ctx context.Context) ExternalPaymentsApiGETExternalPaymentsRequest {
	return ExternalPaymentsApiGETExternalPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return ExternalPaymentResponseList
func (a *ExternalPaymentsApiService) GETExternalPaymentsExecute(r ExternalPaymentsApiGETExternalPaymentsRequest) (*ExternalPaymentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalPaymentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.GETExternalPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments"

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

type ExternalPaymentsApiGETExternalPaymentsExternalPaymentIdRequest struct {
	ctx               context.Context
	ApiService        *ExternalPaymentsApiService
	externalPaymentId string
}

func (r ExternalPaymentsApiGETExternalPaymentsExternalPaymentIdRequest) Execute() (*ExternalPaymentResponse, *http.Response, error) {
	return r.ApiService.GETExternalPaymentsExternalPaymentIdExecute(r)
}

/*
GETExternalPaymentsExternalPaymentId Retrieve an external payment

Retrieve an external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPaymentId The resource's id
 @return ExternalPaymentsApiGETExternalPaymentsExternalPaymentIdRequest
*/
func (a *ExternalPaymentsApiService) GETExternalPaymentsExternalPaymentId(ctx context.Context, externalPaymentId string) ExternalPaymentsApiGETExternalPaymentsExternalPaymentIdRequest {
	return ExternalPaymentsApiGETExternalPaymentsExternalPaymentIdRequest{
		ApiService:        a,
		ctx:               ctx,
		externalPaymentId: externalPaymentId,
	}
}

// Execute executes the request
//  @return ExternalPaymentResponse
func (a *ExternalPaymentsApiService) GETExternalPaymentsExternalPaymentIdExecute(r ExternalPaymentsApiGETExternalPaymentsExternalPaymentIdRequest) (*ExternalPaymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalPaymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.GETExternalPaymentsExternalPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments/{externalPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPaymentId"+"}", url.PathEscape(parameterToString(r.externalPaymentId, "")), -1)

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

type ExternalPaymentsApiPATCHExternalPaymentsExternalPaymentIdRequest struct {
	ctx                   context.Context
	ApiService            *ExternalPaymentsApiService
	externalPaymentUpdate *ExternalPaymentUpdate
	externalPaymentId     string
}

func (r ExternalPaymentsApiPATCHExternalPaymentsExternalPaymentIdRequest) ExternalPaymentUpdate(externalPaymentUpdate ExternalPaymentUpdate) ExternalPaymentsApiPATCHExternalPaymentsExternalPaymentIdRequest {
	r.externalPaymentUpdate = &externalPaymentUpdate
	return r
}

func (r ExternalPaymentsApiPATCHExternalPaymentsExternalPaymentIdRequest) Execute() (*ExternalPaymentResponse, *http.Response, error) {
	return r.ApiService.PATCHExternalPaymentsExternalPaymentIdExecute(r)
}

/*
PATCHExternalPaymentsExternalPaymentId Update an external payment

Update an external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalPaymentId The resource's id
 @return ExternalPaymentsApiPATCHExternalPaymentsExternalPaymentIdRequest
*/
func (a *ExternalPaymentsApiService) PATCHExternalPaymentsExternalPaymentId(ctx context.Context, externalPaymentId string) ExternalPaymentsApiPATCHExternalPaymentsExternalPaymentIdRequest {
	return ExternalPaymentsApiPATCHExternalPaymentsExternalPaymentIdRequest{
		ApiService:        a,
		ctx:               ctx,
		externalPaymentId: externalPaymentId,
	}
}

// Execute executes the request
//  @return ExternalPaymentResponse
func (a *ExternalPaymentsApiService) PATCHExternalPaymentsExternalPaymentIdExecute(r ExternalPaymentsApiPATCHExternalPaymentsExternalPaymentIdRequest) (*ExternalPaymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalPaymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.PATCHExternalPaymentsExternalPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments/{externalPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalPaymentId"+"}", url.PathEscape(parameterToString(r.externalPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalPaymentUpdate == nil {
		return localVarReturnValue, nil, reportError("externalPaymentUpdate is required and must be specified")
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
	localVarPostBody = r.externalPaymentUpdate
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

type ExternalPaymentsApiPOSTExternalPaymentsRequest struct {
	ctx                   context.Context
	ApiService            *ExternalPaymentsApiService
	externalPaymentCreate *ExternalPaymentCreate
}

func (r ExternalPaymentsApiPOSTExternalPaymentsRequest) ExternalPaymentCreate(externalPaymentCreate ExternalPaymentCreate) ExternalPaymentsApiPOSTExternalPaymentsRequest {
	r.externalPaymentCreate = &externalPaymentCreate
	return r
}

func (r ExternalPaymentsApiPOSTExternalPaymentsRequest) Execute() (*ExternalPaymentResponse, *http.Response, error) {
	return r.ApiService.POSTExternalPaymentsExecute(r)
}

/*
POSTExternalPayments Create an external payment

Create an external payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ExternalPaymentsApiPOSTExternalPaymentsRequest
*/
func (a *ExternalPaymentsApiService) POSTExternalPayments(ctx context.Context) ExternalPaymentsApiPOSTExternalPaymentsRequest {
	return ExternalPaymentsApiPOSTExternalPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return ExternalPaymentResponse
func (a *ExternalPaymentsApiService) POSTExternalPaymentsExecute(r ExternalPaymentsApiPOSTExternalPaymentsRequest) (*ExternalPaymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalPaymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalPaymentsApiService.POSTExternalPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalPaymentCreate == nil {
		return localVarReturnValue, nil, reportError("externalPaymentCreate is required and must be specified")
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
	localVarPostBody = r.externalPaymentCreate
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
