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

// CustomerGroupsApiService CustomerGroupsApi service
type CustomerGroupsApiService service

type CustomerGroupsApiDELETECustomerGroupsCustomerGroupIdRequest struct {
	ctx             context.Context
	ApiService      *CustomerGroupsApiService
	customerGroupId string
}

func (r CustomerGroupsApiDELETECustomerGroupsCustomerGroupIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECustomerGroupsCustomerGroupIdExecute(r)
}

/*
DELETECustomerGroupsCustomerGroupId Delete a customer group

Delete a customer group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerGroupId The resource's id
 @return CustomerGroupsApiDELETECustomerGroupsCustomerGroupIdRequest
*/
func (a *CustomerGroupsApiService) DELETECustomerGroupsCustomerGroupId(ctx context.Context, customerGroupId string) CustomerGroupsApiDELETECustomerGroupsCustomerGroupIdRequest {
	return CustomerGroupsApiDELETECustomerGroupsCustomerGroupIdRequest{
		ApiService:      a,
		ctx:             ctx,
		customerGroupId: customerGroupId,
	}
}

// Execute executes the request
func (a *CustomerGroupsApiService) DELETECustomerGroupsCustomerGroupIdExecute(r CustomerGroupsApiDELETECustomerGroupsCustomerGroupIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerGroupsApiService.DELETECustomerGroupsCustomerGroupId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_groups/{customerGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerGroupId"+"}", url.PathEscape(parameterToString(r.customerGroupId, "")), -1)

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

type CustomerGroupsApiGETCustomerGroupsRequest struct {
	ctx        context.Context
	ApiService *CustomerGroupsApiService
}

func (r CustomerGroupsApiGETCustomerGroupsRequest) Execute() (*GETCustomerGroups200Response, *http.Response, error) {
	return r.ApiService.GETCustomerGroupsExecute(r)
}

/*
GETCustomerGroups List all customer groups

List all customer groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CustomerGroupsApiGETCustomerGroupsRequest
*/
func (a *CustomerGroupsApiService) GETCustomerGroups(ctx context.Context) CustomerGroupsApiGETCustomerGroupsRequest {
	return CustomerGroupsApiGETCustomerGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GETCustomerGroups200Response
func (a *CustomerGroupsApiService) GETCustomerGroupsExecute(r CustomerGroupsApiGETCustomerGroupsRequest) (*GETCustomerGroups200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCustomerGroups200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerGroupsApiService.GETCustomerGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_groups"

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

type CustomerGroupsApiGETCustomerGroupsCustomerGroupIdRequest struct {
	ctx             context.Context
	ApiService      *CustomerGroupsApiService
	customerGroupId string
}

func (r CustomerGroupsApiGETCustomerGroupsCustomerGroupIdRequest) Execute() (*GETCustomerGroupsCustomerGroupId200Response, *http.Response, error) {
	return r.ApiService.GETCustomerGroupsCustomerGroupIdExecute(r)
}

/*
GETCustomerGroupsCustomerGroupId Retrieve a customer group

Retrieve a customer group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerGroupId The resource's id
 @return CustomerGroupsApiGETCustomerGroupsCustomerGroupIdRequest
*/
func (a *CustomerGroupsApiService) GETCustomerGroupsCustomerGroupId(ctx context.Context, customerGroupId string) CustomerGroupsApiGETCustomerGroupsCustomerGroupIdRequest {
	return CustomerGroupsApiGETCustomerGroupsCustomerGroupIdRequest{
		ApiService:      a,
		ctx:             ctx,
		customerGroupId: customerGroupId,
	}
}

// Execute executes the request
//  @return GETCustomerGroupsCustomerGroupId200Response
func (a *CustomerGroupsApiService) GETCustomerGroupsCustomerGroupIdExecute(r CustomerGroupsApiGETCustomerGroupsCustomerGroupIdRequest) (*GETCustomerGroupsCustomerGroupId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCustomerGroupsCustomerGroupId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerGroupsApiService.GETCustomerGroupsCustomerGroupId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_groups/{customerGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerGroupId"+"}", url.PathEscape(parameterToString(r.customerGroupId, "")), -1)

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

type CustomerGroupsApiGETCustomerIdCustomerGroupRequest struct {
	ctx        context.Context
	ApiService *CustomerGroupsApiService
	customerId string
}

func (r CustomerGroupsApiGETCustomerIdCustomerGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerIdCustomerGroupExecute(r)
}

/*
GETCustomerIdCustomerGroup Retrieve the customer group associated to the customer

Retrieve the customer group associated to the customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The resource's id
 @return CustomerGroupsApiGETCustomerIdCustomerGroupRequest
*/
func (a *CustomerGroupsApiService) GETCustomerIdCustomerGroup(ctx context.Context, customerId string) CustomerGroupsApiGETCustomerIdCustomerGroupRequest {
	return CustomerGroupsApiGETCustomerIdCustomerGroupRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
	}
}

// Execute executes the request
func (a *CustomerGroupsApiService) GETCustomerIdCustomerGroupExecute(r CustomerGroupsApiGETCustomerIdCustomerGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerGroupsApiService.GETCustomerIdCustomerGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/customer_group"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

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

type CustomerGroupsApiGETMarketIdCustomerGroupRequest struct {
	ctx        context.Context
	ApiService *CustomerGroupsApiService
	marketId   string
}

func (r CustomerGroupsApiGETMarketIdCustomerGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETMarketIdCustomerGroupExecute(r)
}

/*
GETMarketIdCustomerGroup Retrieve the customer group associated to the market

Retrieve the customer group associated to the market

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketId The resource's id
 @return CustomerGroupsApiGETMarketIdCustomerGroupRequest
*/
func (a *CustomerGroupsApiService) GETMarketIdCustomerGroup(ctx context.Context, marketId string) CustomerGroupsApiGETMarketIdCustomerGroupRequest {
	return CustomerGroupsApiGETMarketIdCustomerGroupRequest{
		ApiService: a,
		ctx:        ctx,
		marketId:   marketId,
	}
}

// Execute executes the request
func (a *CustomerGroupsApiService) GETMarketIdCustomerGroupExecute(r CustomerGroupsApiGETMarketIdCustomerGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerGroupsApiService.GETMarketIdCustomerGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/markets/{marketId}/customer_group"
	localVarPath = strings.Replace(localVarPath, "{"+"marketId"+"}", url.PathEscape(parameterToString(r.marketId, "")), -1)

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

type CustomerGroupsApiPATCHCustomerGroupsCustomerGroupIdRequest struct {
	ctx                 context.Context
	ApiService          *CustomerGroupsApiService
	customerGroupUpdate *CustomerGroupUpdate
	customerGroupId     string
}

func (r CustomerGroupsApiPATCHCustomerGroupsCustomerGroupIdRequest) CustomerGroupUpdate(customerGroupUpdate CustomerGroupUpdate) CustomerGroupsApiPATCHCustomerGroupsCustomerGroupIdRequest {
	r.customerGroupUpdate = &customerGroupUpdate
	return r
}

func (r CustomerGroupsApiPATCHCustomerGroupsCustomerGroupIdRequest) Execute() (*PATCHCustomerGroupsCustomerGroupId200Response, *http.Response, error) {
	return r.ApiService.PATCHCustomerGroupsCustomerGroupIdExecute(r)
}

/*
PATCHCustomerGroupsCustomerGroupId Update a customer group

Update a customer group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerGroupId The resource's id
 @return CustomerGroupsApiPATCHCustomerGroupsCustomerGroupIdRequest
*/
func (a *CustomerGroupsApiService) PATCHCustomerGroupsCustomerGroupId(ctx context.Context, customerGroupId string) CustomerGroupsApiPATCHCustomerGroupsCustomerGroupIdRequest {
	return CustomerGroupsApiPATCHCustomerGroupsCustomerGroupIdRequest{
		ApiService:      a,
		ctx:             ctx,
		customerGroupId: customerGroupId,
	}
}

// Execute executes the request
//  @return PATCHCustomerGroupsCustomerGroupId200Response
func (a *CustomerGroupsApiService) PATCHCustomerGroupsCustomerGroupIdExecute(r CustomerGroupsApiPATCHCustomerGroupsCustomerGroupIdRequest) (*PATCHCustomerGroupsCustomerGroupId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHCustomerGroupsCustomerGroupId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerGroupsApiService.PATCHCustomerGroupsCustomerGroupId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_groups/{customerGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerGroupId"+"}", url.PathEscape(parameterToString(r.customerGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerGroupUpdate == nil {
		return localVarReturnValue, nil, reportError("customerGroupUpdate is required and must be specified")
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
	localVarPostBody = r.customerGroupUpdate
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

type CustomerGroupsApiPOSTCustomerGroupsRequest struct {
	ctx                 context.Context
	ApiService          *CustomerGroupsApiService
	customerGroupCreate *CustomerGroupCreate
}

func (r CustomerGroupsApiPOSTCustomerGroupsRequest) CustomerGroupCreate(customerGroupCreate CustomerGroupCreate) CustomerGroupsApiPOSTCustomerGroupsRequest {
	r.customerGroupCreate = &customerGroupCreate
	return r
}

func (r CustomerGroupsApiPOSTCustomerGroupsRequest) Execute() (*POSTCustomerGroups201Response, *http.Response, error) {
	return r.ApiService.POSTCustomerGroupsExecute(r)
}

/*
POSTCustomerGroups Create a customer group

Create a customer group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CustomerGroupsApiPOSTCustomerGroupsRequest
*/
func (a *CustomerGroupsApiService) POSTCustomerGroups(ctx context.Context) CustomerGroupsApiPOSTCustomerGroupsRequest {
	return CustomerGroupsApiPOSTCustomerGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return POSTCustomerGroups201Response
func (a *CustomerGroupsApiService) POSTCustomerGroupsExecute(r CustomerGroupsApiPOSTCustomerGroupsRequest) (*POSTCustomerGroups201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTCustomerGroups201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerGroupsApiService.POSTCustomerGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerGroupCreate == nil {
		return localVarReturnValue, nil, reportError("customerGroupCreate is required and must be specified")
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
	localVarPostBody = r.customerGroupCreate
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
