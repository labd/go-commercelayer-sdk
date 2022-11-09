/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
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

// AdyenPaymentsApiService AdyenPaymentsApi service
type AdyenPaymentsApiService service

type AdyenPaymentsApiDELETEAdyenPaymentsAdyenPaymentIdRequest struct {
	ctx            context.Context
	ApiService     *AdyenPaymentsApiService
	adyenPaymentId string
}

func (r AdyenPaymentsApiDELETEAdyenPaymentsAdyenPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEAdyenPaymentsAdyenPaymentIdExecute(r)
}

/*
DELETEAdyenPaymentsAdyenPaymentId Delete an adyen payment

Delete an adyen payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param adyenPaymentId The resource's id
 @return AdyenPaymentsApiDELETEAdyenPaymentsAdyenPaymentIdRequest
*/
func (a *AdyenPaymentsApiService) DELETEAdyenPaymentsAdyenPaymentId(ctx context.Context, adyenPaymentId string) AdyenPaymentsApiDELETEAdyenPaymentsAdyenPaymentIdRequest {
	return AdyenPaymentsApiDELETEAdyenPaymentsAdyenPaymentIdRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenPaymentId: adyenPaymentId,
	}
}

// Execute executes the request
func (a *AdyenPaymentsApiService) DELETEAdyenPaymentsAdyenPaymentIdExecute(r AdyenPaymentsApiDELETEAdyenPaymentsAdyenPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.DELETEAdyenPaymentsAdyenPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments/{adyenPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenPaymentId"+"}", url.PathEscape(parameterToString(r.adyenPaymentId, "")), -1)

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

type AdyenPaymentsApiGETAdyenGatewayIdAdyenPaymentsRequest struct {
	ctx            context.Context
	ApiService     *AdyenPaymentsApiService
	adyenGatewayId string
}

func (r AdyenPaymentsApiGETAdyenGatewayIdAdyenPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETAdyenGatewayIdAdyenPaymentsExecute(r)
}

/*
GETAdyenGatewayIdAdyenPayments Retrieve the adyen payments associated to the adyen gateway

Retrieve the adyen payments associated to the adyen gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param adyenGatewayId The resource's id
 @return AdyenPaymentsApiGETAdyenGatewayIdAdyenPaymentsRequest
*/
func (a *AdyenPaymentsApiService) GETAdyenGatewayIdAdyenPayments(ctx context.Context, adyenGatewayId string) AdyenPaymentsApiGETAdyenGatewayIdAdyenPaymentsRequest {
	return AdyenPaymentsApiGETAdyenGatewayIdAdyenPaymentsRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenGatewayId: adyenGatewayId,
	}
}

// Execute executes the request
func (a *AdyenPaymentsApiService) GETAdyenGatewayIdAdyenPaymentsExecute(r AdyenPaymentsApiGETAdyenGatewayIdAdyenPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.GETAdyenGatewayIdAdyenPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_gateways/{adyenGatewayId}/adyen_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenGatewayId"+"}", url.PathEscape(parameterToString(r.adyenGatewayId, "")), -1)

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

type AdyenPaymentsApiGETAdyenPaymentsRequest struct {
	ctx        context.Context
	ApiService *AdyenPaymentsApiService
}

func (r AdyenPaymentsApiGETAdyenPaymentsRequest) Execute() (*GETAdyenPayments200Response, *http.Response, error) {
	return r.ApiService.GETAdyenPaymentsExecute(r)
}

/*
GETAdyenPayments List all adyen payments

List all adyen payments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AdyenPaymentsApiGETAdyenPaymentsRequest
*/
func (a *AdyenPaymentsApiService) GETAdyenPayments(ctx context.Context) AdyenPaymentsApiGETAdyenPaymentsRequest {
	return AdyenPaymentsApiGETAdyenPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETAdyenPayments200Response
func (a *AdyenPaymentsApiService) GETAdyenPaymentsExecute(r AdyenPaymentsApiGETAdyenPaymentsRequest) (*GETAdyenPayments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAdyenPayments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.GETAdyenPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments"

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

type AdyenPaymentsApiGETAdyenPaymentsAdyenPaymentIdRequest struct {
	ctx            context.Context
	ApiService     *AdyenPaymentsApiService
	adyenPaymentId string
}

func (r AdyenPaymentsApiGETAdyenPaymentsAdyenPaymentIdRequest) Execute() (*GETAdyenPaymentsAdyenPaymentId200Response, *http.Response, error) {
	return r.ApiService.GETAdyenPaymentsAdyenPaymentIdExecute(r)
}

/*
GETAdyenPaymentsAdyenPaymentId Retrieve an adyen payment

Retrieve an adyen payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param adyenPaymentId The resource's id
 @return AdyenPaymentsApiGETAdyenPaymentsAdyenPaymentIdRequest
*/
func (a *AdyenPaymentsApiService) GETAdyenPaymentsAdyenPaymentId(ctx context.Context, adyenPaymentId string) AdyenPaymentsApiGETAdyenPaymentsAdyenPaymentIdRequest {
	return AdyenPaymentsApiGETAdyenPaymentsAdyenPaymentIdRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenPaymentId: adyenPaymentId,
	}
}

// Execute executes the request
//  @return GETAdyenPaymentsAdyenPaymentId200Response
func (a *AdyenPaymentsApiService) GETAdyenPaymentsAdyenPaymentIdExecute(r AdyenPaymentsApiGETAdyenPaymentsAdyenPaymentIdRequest) (*GETAdyenPaymentsAdyenPaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAdyenPaymentsAdyenPaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.GETAdyenPaymentsAdyenPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments/{adyenPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenPaymentId"+"}", url.PathEscape(parameterToString(r.adyenPaymentId, "")), -1)

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

type AdyenPaymentsApiPATCHAdyenPaymentsAdyenPaymentIdRequest struct {
	ctx                context.Context
	ApiService         *AdyenPaymentsApiService
	adyenPaymentUpdate *AdyenPaymentUpdate
	adyenPaymentId     string
}

func (r AdyenPaymentsApiPATCHAdyenPaymentsAdyenPaymentIdRequest) AdyenPaymentUpdate(adyenPaymentUpdate AdyenPaymentUpdate) AdyenPaymentsApiPATCHAdyenPaymentsAdyenPaymentIdRequest {
	r.adyenPaymentUpdate = &adyenPaymentUpdate
	return r
}

func (r AdyenPaymentsApiPATCHAdyenPaymentsAdyenPaymentIdRequest) Execute() (*PATCHAdyenPaymentsAdyenPaymentId200Response, *http.Response, error) {
	return r.ApiService.PATCHAdyenPaymentsAdyenPaymentIdExecute(r)
}

/*
PATCHAdyenPaymentsAdyenPaymentId Update an adyen payment

Update an adyen payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param adyenPaymentId The resource's id
 @return AdyenPaymentsApiPATCHAdyenPaymentsAdyenPaymentIdRequest
*/
func (a *AdyenPaymentsApiService) PATCHAdyenPaymentsAdyenPaymentId(ctx context.Context, adyenPaymentId string) AdyenPaymentsApiPATCHAdyenPaymentsAdyenPaymentIdRequest {
	return AdyenPaymentsApiPATCHAdyenPaymentsAdyenPaymentIdRequest{
		ApiService:     a,
		ctx:            ctx,
		adyenPaymentId: adyenPaymentId,
	}
}

// Execute executes the request
//  @return PATCHAdyenPaymentsAdyenPaymentId200Response
func (a *AdyenPaymentsApiService) PATCHAdyenPaymentsAdyenPaymentIdExecute(r AdyenPaymentsApiPATCHAdyenPaymentsAdyenPaymentIdRequest) (*PATCHAdyenPaymentsAdyenPaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHAdyenPaymentsAdyenPaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.PATCHAdyenPaymentsAdyenPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments/{adyenPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adyenPaymentId"+"}", url.PathEscape(parameterToString(r.adyenPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adyenPaymentUpdate == nil {
		return localVarReturnValue, nil, reportError("adyenPaymentUpdate is required and must be specified")
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
	localVarPostBody = r.adyenPaymentUpdate
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

type AdyenPaymentsApiPOSTAdyenPaymentsRequest struct {
	ctx                context.Context
	ApiService         *AdyenPaymentsApiService
	adyenPaymentCreate *AdyenPaymentCreate
}

func (r AdyenPaymentsApiPOSTAdyenPaymentsRequest) AdyenPaymentCreate(adyenPaymentCreate AdyenPaymentCreate) AdyenPaymentsApiPOSTAdyenPaymentsRequest {
	r.adyenPaymentCreate = &adyenPaymentCreate
	return r
}

func (r AdyenPaymentsApiPOSTAdyenPaymentsRequest) Execute() (*POSTAdyenPayments201Response, *http.Response, error) {
	return r.ApiService.POSTAdyenPaymentsExecute(r)
}

/*
POSTAdyenPayments Create an adyen payment

Create an adyen payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AdyenPaymentsApiPOSTAdyenPaymentsRequest
*/
func (a *AdyenPaymentsApiService) POSTAdyenPayments(ctx context.Context) AdyenPaymentsApiPOSTAdyenPaymentsRequest {
	return AdyenPaymentsApiPOSTAdyenPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTAdyenPayments201Response
func (a *AdyenPaymentsApiService) POSTAdyenPaymentsExecute(r AdyenPaymentsApiPOSTAdyenPaymentsRequest) (*POSTAdyenPayments201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTAdyenPayments201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdyenPaymentsApiService.POSTAdyenPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adyen_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adyenPaymentCreate == nil {
		return localVarReturnValue, nil, reportError("adyenPaymentCreate is required and must be specified")
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
	localVarPostBody = r.adyenPaymentCreate
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
