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

// CheckoutComGatewaysApiService CheckoutComGatewaysApi service
type CheckoutComGatewaysApiService service

type CheckoutComGatewaysApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest struct {
	ctx                  context.Context
	ApiService           *CheckoutComGatewaysApiService
	checkoutComGatewayId string
}

func (r CheckoutComGatewaysApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECheckoutComGatewaysCheckoutComGatewayIdExecute(r)
}

/*
DELETECheckoutComGatewaysCheckoutComGatewayId Delete a checkout.com gateway

Delete a checkout.com gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComGatewayId The resource's id
 @return CheckoutComGatewaysApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest
*/
func (a *CheckoutComGatewaysApiService) DELETECheckoutComGatewaysCheckoutComGatewayId(ctx context.Context, checkoutComGatewayId string) CheckoutComGatewaysApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest {
	return CheckoutComGatewaysApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComGatewayId: checkoutComGatewayId,
	}
}

// Execute executes the request
func (a *CheckoutComGatewaysApiService) DELETECheckoutComGatewaysCheckoutComGatewayIdExecute(r CheckoutComGatewaysApiDELETECheckoutComGatewaysCheckoutComGatewayIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.DELETECheckoutComGatewaysCheckoutComGatewayId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways/{checkoutComGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComGatewayId"+"}", url.PathEscape(parameterToString(r.checkoutComGatewayId, "")), -1)

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

type CheckoutComGatewaysApiGETCheckoutComGatewaysRequest struct {
	ctx        context.Context
	ApiService *CheckoutComGatewaysApiService
}

func (r CheckoutComGatewaysApiGETCheckoutComGatewaysRequest) Execute() (*GETCheckoutComGateways200Response, *http.Response, error) {
	return r.ApiService.GETCheckoutComGatewaysExecute(r)
}

/*
GETCheckoutComGateways List all checkout.com gateways

List all checkout.com gateways

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CheckoutComGatewaysApiGETCheckoutComGatewaysRequest
*/
func (a *CheckoutComGatewaysApiService) GETCheckoutComGateways(ctx context.Context) CheckoutComGatewaysApiGETCheckoutComGatewaysRequest {
	return CheckoutComGatewaysApiGETCheckoutComGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETCheckoutComGateways200Response
func (a *CheckoutComGatewaysApiService) GETCheckoutComGatewaysExecute(r CheckoutComGatewaysApiGETCheckoutComGatewaysRequest) (*GETCheckoutComGateways200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCheckoutComGateways200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.GETCheckoutComGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways"

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

type CheckoutComGatewaysApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest struct {
	ctx                  context.Context
	ApiService           *CheckoutComGatewaysApiService
	checkoutComGatewayId string
}

func (r CheckoutComGatewaysApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest) Execute() (*GETCheckoutComGatewaysCheckoutComGatewayId200Response, *http.Response, error) {
	return r.ApiService.GETCheckoutComGatewaysCheckoutComGatewayIdExecute(r)
}

/*
GETCheckoutComGatewaysCheckoutComGatewayId Retrieve a checkout.com gateway

Retrieve a checkout.com gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComGatewayId The resource's id
 @return CheckoutComGatewaysApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest
*/
func (a *CheckoutComGatewaysApiService) GETCheckoutComGatewaysCheckoutComGatewayId(ctx context.Context, checkoutComGatewayId string) CheckoutComGatewaysApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest {
	return CheckoutComGatewaysApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComGatewayId: checkoutComGatewayId,
	}
}

// Execute executes the request
//  @return GETCheckoutComGatewaysCheckoutComGatewayId200Response
func (a *CheckoutComGatewaysApiService) GETCheckoutComGatewaysCheckoutComGatewayIdExecute(r CheckoutComGatewaysApiGETCheckoutComGatewaysCheckoutComGatewayIdRequest) (*GETCheckoutComGatewaysCheckoutComGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCheckoutComGatewaysCheckoutComGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.GETCheckoutComGatewaysCheckoutComGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways/{checkoutComGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComGatewayId"+"}", url.PathEscape(parameterToString(r.checkoutComGatewayId, "")), -1)

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

type CheckoutComGatewaysApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest struct {
	ctx                      context.Context
	ApiService               *CheckoutComGatewaysApiService
	checkoutComGatewayUpdate *CheckoutComGatewayUpdate
	checkoutComGatewayId     string
}

func (r CheckoutComGatewaysApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest) CheckoutComGatewayUpdate(checkoutComGatewayUpdate CheckoutComGatewayUpdate) CheckoutComGatewaysApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest {
	r.checkoutComGatewayUpdate = &checkoutComGatewayUpdate
	return r
}

func (r CheckoutComGatewaysApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest) Execute() (*PATCHCheckoutComGatewaysCheckoutComGatewayId200Response, *http.Response, error) {
	return r.ApiService.PATCHCheckoutComGatewaysCheckoutComGatewayIdExecute(r)
}

/*
PATCHCheckoutComGatewaysCheckoutComGatewayId Update a checkout.com gateway

Update a checkout.com gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param checkoutComGatewayId The resource's id
 @return CheckoutComGatewaysApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest
*/
func (a *CheckoutComGatewaysApiService) PATCHCheckoutComGatewaysCheckoutComGatewayId(ctx context.Context, checkoutComGatewayId string) CheckoutComGatewaysApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest {
	return CheckoutComGatewaysApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest{
		ApiService:           a,
		ctx:                  ctx,
		checkoutComGatewayId: checkoutComGatewayId,
	}
}

// Execute executes the request
//  @return PATCHCheckoutComGatewaysCheckoutComGatewayId200Response
func (a *CheckoutComGatewaysApiService) PATCHCheckoutComGatewaysCheckoutComGatewayIdExecute(r CheckoutComGatewaysApiPATCHCheckoutComGatewaysCheckoutComGatewayIdRequest) (*PATCHCheckoutComGatewaysCheckoutComGatewayId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHCheckoutComGatewaysCheckoutComGatewayId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.PATCHCheckoutComGatewaysCheckoutComGatewayId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways/{checkoutComGatewayId}"
	localVarPath = strings.Replace(localVarPath, "{"+"checkoutComGatewayId"+"}", url.PathEscape(parameterToString(r.checkoutComGatewayId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkoutComGatewayUpdate == nil {
		return localVarReturnValue, nil, reportError("checkoutComGatewayUpdate is required and must be specified")
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
	localVarPostBody = r.checkoutComGatewayUpdate
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

type CheckoutComGatewaysApiPOSTCheckoutComGatewaysRequest struct {
	ctx                      context.Context
	ApiService               *CheckoutComGatewaysApiService
	checkoutComGatewayCreate *CheckoutComGatewayCreate
}

func (r CheckoutComGatewaysApiPOSTCheckoutComGatewaysRequest) CheckoutComGatewayCreate(checkoutComGatewayCreate CheckoutComGatewayCreate) CheckoutComGatewaysApiPOSTCheckoutComGatewaysRequest {
	r.checkoutComGatewayCreate = &checkoutComGatewayCreate
	return r
}

func (r CheckoutComGatewaysApiPOSTCheckoutComGatewaysRequest) Execute() (*POSTCheckoutComGateways201Response, *http.Response, error) {
	return r.ApiService.POSTCheckoutComGatewaysExecute(r)
}

/*
POSTCheckoutComGateways Create a checkout.com gateway

Create a checkout.com gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CheckoutComGatewaysApiPOSTCheckoutComGatewaysRequest
*/
func (a *CheckoutComGatewaysApiService) POSTCheckoutComGateways(ctx context.Context) CheckoutComGatewaysApiPOSTCheckoutComGatewaysRequest {
	return CheckoutComGatewaysApiPOSTCheckoutComGatewaysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTCheckoutComGateways201Response
func (a *CheckoutComGatewaysApiService) POSTCheckoutComGatewaysExecute(r CheckoutComGatewaysApiPOSTCheckoutComGatewaysRequest) (*POSTCheckoutComGateways201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTCheckoutComGateways201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CheckoutComGatewaysApiService.POSTCheckoutComGateways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout_com_gateways"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkoutComGatewayCreate == nil {
		return localVarReturnValue, nil, reportError("checkoutComGatewayCreate is required and must be specified")
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
	localVarPostBody = r.checkoutComGatewayCreate
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
