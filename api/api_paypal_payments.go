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

// PaypalPaymentsApiService PaypalPaymentsApi service
type PaypalPaymentsApiService service

type PaypalPaymentsApiDELETEPaypalPaymentsPaypalPaymentIdRequest struct {
	ctx             context.Context
	ApiService      *PaypalPaymentsApiService
	paypalPaymentId string
}

func (r PaypalPaymentsApiDELETEPaypalPaymentsPaypalPaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEPaypalPaymentsPaypalPaymentIdExecute(r)
}

/*
DELETEPaypalPaymentsPaypalPaymentId Delete a paypal payment

Delete a paypal payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paypalPaymentId The resource's id
 @return PaypalPaymentsApiDELETEPaypalPaymentsPaypalPaymentIdRequest
*/
func (a *PaypalPaymentsApiService) DELETEPaypalPaymentsPaypalPaymentId(ctx context.Context, paypalPaymentId string) PaypalPaymentsApiDELETEPaypalPaymentsPaypalPaymentIdRequest {
	return PaypalPaymentsApiDELETEPaypalPaymentsPaypalPaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		paypalPaymentId: paypalPaymentId,
	}
}

// Execute executes the request
func (a *PaypalPaymentsApiService) DELETEPaypalPaymentsPaypalPaymentIdExecute(r PaypalPaymentsApiDELETEPaypalPaymentsPaypalPaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaypalPaymentsApiService.DELETEPaypalPaymentsPaypalPaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paypal_payments/{paypalPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"paypalPaymentId"+"}", url.PathEscape(parameterToString(r.paypalPaymentId, "")), -1)

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

type PaypalPaymentsApiGETPaypalGatewayIdPaypalPaymentsRequest struct {
	ctx             context.Context
	ApiService      *PaypalPaymentsApiService
	paypalGatewayId string
}

func (r PaypalPaymentsApiGETPaypalGatewayIdPaypalPaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETPaypalGatewayIdPaypalPaymentsExecute(r)
}

/*
GETPaypalGatewayIdPaypalPayments Retrieve the paypal payments associated to the paypal gateway

Retrieve the paypal payments associated to the paypal gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paypalGatewayId The resource's id
 @return PaypalPaymentsApiGETPaypalGatewayIdPaypalPaymentsRequest
*/
func (a *PaypalPaymentsApiService) GETPaypalGatewayIdPaypalPayments(ctx context.Context, paypalGatewayId string) PaypalPaymentsApiGETPaypalGatewayIdPaypalPaymentsRequest {
	return PaypalPaymentsApiGETPaypalGatewayIdPaypalPaymentsRequest{
		ApiService:      a,
		ctx:             ctx,
		paypalGatewayId: paypalGatewayId,
	}
}

// Execute executes the request
func (a *PaypalPaymentsApiService) GETPaypalGatewayIdPaypalPaymentsExecute(r PaypalPaymentsApiGETPaypalGatewayIdPaypalPaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaypalPaymentsApiService.GETPaypalGatewayIdPaypalPayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paypal_gateways/{paypalGatewayId}/paypal_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"paypalGatewayId"+"}", url.PathEscape(parameterToString(r.paypalGatewayId, "")), -1)

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

type PaypalPaymentsApiGETPaypalPaymentsRequest struct {
	ctx        context.Context
	ApiService *PaypalPaymentsApiService
}

func (r PaypalPaymentsApiGETPaypalPaymentsRequest) Execute() (*PaypalPaymentResponseList, *http.Response, error) {
	return r.ApiService.GETPaypalPaymentsExecute(r)
}

/*
GETPaypalPayments List all paypal payments

List all paypal payments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaypalPaymentsApiGETPaypalPaymentsRequest
*/
func (a *PaypalPaymentsApiService) GETPaypalPayments(ctx context.Context) PaypalPaymentsApiGETPaypalPaymentsRequest {
	return PaypalPaymentsApiGETPaypalPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PaypalPaymentResponseList
func (a *PaypalPaymentsApiService) GETPaypalPaymentsExecute(r PaypalPaymentsApiGETPaypalPaymentsRequest) (*PaypalPaymentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaypalPaymentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaypalPaymentsApiService.GETPaypalPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paypal_payments"

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

type PaypalPaymentsApiGETPaypalPaymentsPaypalPaymentIdRequest struct {
	ctx             context.Context
	ApiService      *PaypalPaymentsApiService
	paypalPaymentId string
}

func (r PaypalPaymentsApiGETPaypalPaymentsPaypalPaymentIdRequest) Execute() (*PaypalPaymentResponse, *http.Response, error) {
	return r.ApiService.GETPaypalPaymentsPaypalPaymentIdExecute(r)
}

/*
GETPaypalPaymentsPaypalPaymentId Retrieve a paypal payment

Retrieve a paypal payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paypalPaymentId The resource's id
 @return PaypalPaymentsApiGETPaypalPaymentsPaypalPaymentIdRequest
*/
func (a *PaypalPaymentsApiService) GETPaypalPaymentsPaypalPaymentId(ctx context.Context, paypalPaymentId string) PaypalPaymentsApiGETPaypalPaymentsPaypalPaymentIdRequest {
	return PaypalPaymentsApiGETPaypalPaymentsPaypalPaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		paypalPaymentId: paypalPaymentId,
	}
}

// Execute executes the request
//  @return PaypalPaymentResponse
func (a *PaypalPaymentsApiService) GETPaypalPaymentsPaypalPaymentIdExecute(r PaypalPaymentsApiGETPaypalPaymentsPaypalPaymentIdRequest) (*PaypalPaymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaypalPaymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaypalPaymentsApiService.GETPaypalPaymentsPaypalPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paypal_payments/{paypalPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"paypalPaymentId"+"}", url.PathEscape(parameterToString(r.paypalPaymentId, "")), -1)

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

type PaypalPaymentsApiPATCHPaypalPaymentsPaypalPaymentIdRequest struct {
	ctx                 context.Context
	ApiService          *PaypalPaymentsApiService
	paypalPaymentUpdate *PaypalPaymentUpdate
	paypalPaymentId     string
}

func (r PaypalPaymentsApiPATCHPaypalPaymentsPaypalPaymentIdRequest) PaypalPaymentUpdate(paypalPaymentUpdate PaypalPaymentUpdate) PaypalPaymentsApiPATCHPaypalPaymentsPaypalPaymentIdRequest {
	r.paypalPaymentUpdate = &paypalPaymentUpdate
	return r
}

func (r PaypalPaymentsApiPATCHPaypalPaymentsPaypalPaymentIdRequest) Execute() (*PaypalPaymentResponse, *http.Response, error) {
	return r.ApiService.PATCHPaypalPaymentsPaypalPaymentIdExecute(r)
}

/*
PATCHPaypalPaymentsPaypalPaymentId Update a paypal payment

Update a paypal payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paypalPaymentId The resource's id
 @return PaypalPaymentsApiPATCHPaypalPaymentsPaypalPaymentIdRequest
*/
func (a *PaypalPaymentsApiService) PATCHPaypalPaymentsPaypalPaymentId(ctx context.Context, paypalPaymentId string) PaypalPaymentsApiPATCHPaypalPaymentsPaypalPaymentIdRequest {
	return PaypalPaymentsApiPATCHPaypalPaymentsPaypalPaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		paypalPaymentId: paypalPaymentId,
	}
}

// Execute executes the request
//  @return PaypalPaymentResponse
func (a *PaypalPaymentsApiService) PATCHPaypalPaymentsPaypalPaymentIdExecute(r PaypalPaymentsApiPATCHPaypalPaymentsPaypalPaymentIdRequest) (*PaypalPaymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaypalPaymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaypalPaymentsApiService.PATCHPaypalPaymentsPaypalPaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paypal_payments/{paypalPaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"paypalPaymentId"+"}", url.PathEscape(parameterToString(r.paypalPaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.paypalPaymentUpdate == nil {
		return localVarReturnValue, nil, reportError("paypalPaymentUpdate is required and must be specified")
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
	localVarPostBody = r.paypalPaymentUpdate
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

type PaypalPaymentsApiPOSTPaypalPaymentsRequest struct {
	ctx                 context.Context
	ApiService          *PaypalPaymentsApiService
	paypalPaymentCreate *PaypalPaymentCreate
}

func (r PaypalPaymentsApiPOSTPaypalPaymentsRequest) PaypalPaymentCreate(paypalPaymentCreate PaypalPaymentCreate) PaypalPaymentsApiPOSTPaypalPaymentsRequest {
	r.paypalPaymentCreate = &paypalPaymentCreate
	return r
}

func (r PaypalPaymentsApiPOSTPaypalPaymentsRequest) Execute() (*PaypalPaymentResponse, *http.Response, error) {
	return r.ApiService.POSTPaypalPaymentsExecute(r)
}

/*
POSTPaypalPayments Create a paypal payment

Create a paypal payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PaypalPaymentsApiPOSTPaypalPaymentsRequest
*/
func (a *PaypalPaymentsApiService) POSTPaypalPayments(ctx context.Context) PaypalPaymentsApiPOSTPaypalPaymentsRequest {
	return PaypalPaymentsApiPOSTPaypalPaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return PaypalPaymentResponse
func (a *PaypalPaymentsApiService) POSTPaypalPaymentsExecute(r PaypalPaymentsApiPOSTPaypalPaymentsRequest) (*PaypalPaymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaypalPaymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaypalPaymentsApiService.POSTPaypalPayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/paypal_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.paypalPaymentCreate == nil {
		return localVarReturnValue, nil, reportError("paypalPaymentCreate is required and must be specified")
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
	localVarPostBody = r.paypalPaymentCreate
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
