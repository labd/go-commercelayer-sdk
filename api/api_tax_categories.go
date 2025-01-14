/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
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

// TaxCategoriesApiService TaxCategoriesApi service
type TaxCategoriesApiService service

type TaxCategoriesApiDELETETaxCategoriesTaxCategoryIdRequest struct {
	ctx           context.Context
	ApiService    *TaxCategoriesApiService
	taxCategoryId interface{}
}

func (r TaxCategoriesApiDELETETaxCategoriesTaxCategoryIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETETaxCategoriesTaxCategoryIdExecute(r)
}

/*
DELETETaxCategoriesTaxCategoryId Delete a tax category

Delete a tax category

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxCategoryId The resource's id
	@return TaxCategoriesApiDELETETaxCategoriesTaxCategoryIdRequest
*/
func (a *TaxCategoriesApiService) DELETETaxCategoriesTaxCategoryId(ctx context.Context, taxCategoryId interface{}) TaxCategoriesApiDELETETaxCategoriesTaxCategoryIdRequest {
	return TaxCategoriesApiDELETETaxCategoriesTaxCategoryIdRequest{
		ApiService:    a,
		ctx:           ctx,
		taxCategoryId: taxCategoryId,
	}
}

// Execute executes the request
func (a *TaxCategoriesApiService) DELETETaxCategoriesTaxCategoryIdExecute(r TaxCategoriesApiDELETETaxCategoriesTaxCategoryIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCategoriesApiService.DELETETaxCategoriesTaxCategoryId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_categories/{taxCategoryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxCategoryId"+"}", url.PathEscape(parameterValueToString(r.taxCategoryId, "taxCategoryId")), -1)

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

type TaxCategoriesApiGETAvalaraAccountIdTaxCategoriesRequest struct {
	ctx              context.Context
	ApiService       *TaxCategoriesApiService
	avalaraAccountId interface{}
}

func (r TaxCategoriesApiGETAvalaraAccountIdTaxCategoriesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETAvalaraAccountIdTaxCategoriesExecute(r)
}

/*
GETAvalaraAccountIdTaxCategories Retrieve the tax categories associated to the avalara account

Retrieve the tax categories associated to the avalara account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param avalaraAccountId The resource's id
	@return TaxCategoriesApiGETAvalaraAccountIdTaxCategoriesRequest
*/
func (a *TaxCategoriesApiService) GETAvalaraAccountIdTaxCategories(ctx context.Context, avalaraAccountId interface{}) TaxCategoriesApiGETAvalaraAccountIdTaxCategoriesRequest {
	return TaxCategoriesApiGETAvalaraAccountIdTaxCategoriesRequest{
		ApiService:       a,
		ctx:              ctx,
		avalaraAccountId: avalaraAccountId,
	}
}

// Execute executes the request
func (a *TaxCategoriesApiService) GETAvalaraAccountIdTaxCategoriesExecute(r TaxCategoriesApiGETAvalaraAccountIdTaxCategoriesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCategoriesApiService.GETAvalaraAccountIdTaxCategories")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/avalara_accounts/{avalaraAccountId}/tax_categories"
	localVarPath = strings.Replace(localVarPath, "{"+"avalaraAccountId"+"}", url.PathEscape(parameterValueToString(r.avalaraAccountId, "avalaraAccountId")), -1)

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

type TaxCategoriesApiGETTaxCategoriesRequest struct {
	ctx        context.Context
	ApiService *TaxCategoriesApiService
}

func (r TaxCategoriesApiGETTaxCategoriesRequest) Execute() (*GETTaxCategories200Response, *http.Response, error) {
	return r.ApiService.GETTaxCategoriesExecute(r)
}

/*
GETTaxCategories List all tax categories

List all tax categories

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TaxCategoriesApiGETTaxCategoriesRequest
*/
func (a *TaxCategoriesApiService) GETTaxCategories(ctx context.Context) TaxCategoriesApiGETTaxCategoriesRequest {
	return TaxCategoriesApiGETTaxCategoriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETTaxCategories200Response
func (a *TaxCategoriesApiService) GETTaxCategoriesExecute(r TaxCategoriesApiGETTaxCategoriesRequest) (*GETTaxCategories200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETTaxCategories200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCategoriesApiService.GETTaxCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_categories"

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

type TaxCategoriesApiGETTaxCategoriesTaxCategoryIdRequest struct {
	ctx           context.Context
	ApiService    *TaxCategoriesApiService
	taxCategoryId interface{}
}

func (r TaxCategoriesApiGETTaxCategoriesTaxCategoryIdRequest) Execute() (*GETTaxCategoriesTaxCategoryId200Response, *http.Response, error) {
	return r.ApiService.GETTaxCategoriesTaxCategoryIdExecute(r)
}

/*
GETTaxCategoriesTaxCategoryId Retrieve a tax category

Retrieve a tax category

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxCategoryId The resource's id
	@return TaxCategoriesApiGETTaxCategoriesTaxCategoryIdRequest
*/
func (a *TaxCategoriesApiService) GETTaxCategoriesTaxCategoryId(ctx context.Context, taxCategoryId interface{}) TaxCategoriesApiGETTaxCategoriesTaxCategoryIdRequest {
	return TaxCategoriesApiGETTaxCategoriesTaxCategoryIdRequest{
		ApiService:    a,
		ctx:           ctx,
		taxCategoryId: taxCategoryId,
	}
}

// Execute executes the request
//
//	@return GETTaxCategoriesTaxCategoryId200Response
func (a *TaxCategoriesApiService) GETTaxCategoriesTaxCategoryIdExecute(r TaxCategoriesApiGETTaxCategoriesTaxCategoryIdRequest) (*GETTaxCategoriesTaxCategoryId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETTaxCategoriesTaxCategoryId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCategoriesApiService.GETTaxCategoriesTaxCategoryId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_categories/{taxCategoryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxCategoryId"+"}", url.PathEscape(parameterValueToString(r.taxCategoryId, "taxCategoryId")), -1)

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

type TaxCategoriesApiGETTaxjarAccountIdTaxCategoriesRequest struct {
	ctx             context.Context
	ApiService      *TaxCategoriesApiService
	taxjarAccountId interface{}
}

func (r TaxCategoriesApiGETTaxjarAccountIdTaxCategoriesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETTaxjarAccountIdTaxCategoriesExecute(r)
}

/*
GETTaxjarAccountIdTaxCategories Retrieve the tax categories associated to the taxjar account

Retrieve the tax categories associated to the taxjar account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxjarAccountId The resource's id
	@return TaxCategoriesApiGETTaxjarAccountIdTaxCategoriesRequest
*/
func (a *TaxCategoriesApiService) GETTaxjarAccountIdTaxCategories(ctx context.Context, taxjarAccountId interface{}) TaxCategoriesApiGETTaxjarAccountIdTaxCategoriesRequest {
	return TaxCategoriesApiGETTaxjarAccountIdTaxCategoriesRequest{
		ApiService:      a,
		ctx:             ctx,
		taxjarAccountId: taxjarAccountId,
	}
}

// Execute executes the request
func (a *TaxCategoriesApiService) GETTaxjarAccountIdTaxCategoriesExecute(r TaxCategoriesApiGETTaxjarAccountIdTaxCategoriesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCategoriesApiService.GETTaxjarAccountIdTaxCategories")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/taxjar_accounts/{taxjarAccountId}/tax_categories"
	localVarPath = strings.Replace(localVarPath, "{"+"taxjarAccountId"+"}", url.PathEscape(parameterValueToString(r.taxjarAccountId, "taxjarAccountId")), -1)

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

type TaxCategoriesApiPATCHTaxCategoriesTaxCategoryIdRequest struct {
	ctx               context.Context
	ApiService        *TaxCategoriesApiService
	taxCategoryUpdate *TaxCategoryUpdate
	taxCategoryId     interface{}
}

func (r TaxCategoriesApiPATCHTaxCategoriesTaxCategoryIdRequest) TaxCategoryUpdate(taxCategoryUpdate TaxCategoryUpdate) TaxCategoriesApiPATCHTaxCategoriesTaxCategoryIdRequest {
	r.taxCategoryUpdate = &taxCategoryUpdate
	return r
}

func (r TaxCategoriesApiPATCHTaxCategoriesTaxCategoryIdRequest) Execute() (*PATCHTaxCategoriesTaxCategoryId200Response, *http.Response, error) {
	return r.ApiService.PATCHTaxCategoriesTaxCategoryIdExecute(r)
}

/*
PATCHTaxCategoriesTaxCategoryId Update a tax category

Update a tax category

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxCategoryId The resource's id
	@return TaxCategoriesApiPATCHTaxCategoriesTaxCategoryIdRequest
*/
func (a *TaxCategoriesApiService) PATCHTaxCategoriesTaxCategoryId(ctx context.Context, taxCategoryId interface{}) TaxCategoriesApiPATCHTaxCategoriesTaxCategoryIdRequest {
	return TaxCategoriesApiPATCHTaxCategoriesTaxCategoryIdRequest{
		ApiService:    a,
		ctx:           ctx,
		taxCategoryId: taxCategoryId,
	}
}

// Execute executes the request
//
//	@return PATCHTaxCategoriesTaxCategoryId200Response
func (a *TaxCategoriesApiService) PATCHTaxCategoriesTaxCategoryIdExecute(r TaxCategoriesApiPATCHTaxCategoriesTaxCategoryIdRequest) (*PATCHTaxCategoriesTaxCategoryId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHTaxCategoriesTaxCategoryId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCategoriesApiService.PATCHTaxCategoriesTaxCategoryId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_categories/{taxCategoryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxCategoryId"+"}", url.PathEscape(parameterValueToString(r.taxCategoryId, "taxCategoryId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.taxCategoryUpdate == nil {
		return localVarReturnValue, nil, reportError("taxCategoryUpdate is required and must be specified")
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
	localVarPostBody = r.taxCategoryUpdate
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

type TaxCategoriesApiPOSTTaxCategoriesRequest struct {
	ctx               context.Context
	ApiService        *TaxCategoriesApiService
	taxCategoryCreate *TaxCategoryCreate
}

func (r TaxCategoriesApiPOSTTaxCategoriesRequest) TaxCategoryCreate(taxCategoryCreate TaxCategoryCreate) TaxCategoriesApiPOSTTaxCategoriesRequest {
	r.taxCategoryCreate = &taxCategoryCreate
	return r
}

func (r TaxCategoriesApiPOSTTaxCategoriesRequest) Execute() (*POSTTaxCategories201Response, *http.Response, error) {
	return r.ApiService.POSTTaxCategoriesExecute(r)
}

/*
POSTTaxCategories Create a tax category

Create a tax category

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TaxCategoriesApiPOSTTaxCategoriesRequest
*/
func (a *TaxCategoriesApiService) POSTTaxCategories(ctx context.Context) TaxCategoriesApiPOSTTaxCategoriesRequest {
	return TaxCategoriesApiPOSTTaxCategoriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTTaxCategories201Response
func (a *TaxCategoriesApiService) POSTTaxCategoriesExecute(r TaxCategoriesApiPOSTTaxCategoriesRequest) (*POSTTaxCategories201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTTaxCategories201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxCategoriesApiService.POSTTaxCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_categories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.taxCategoryCreate == nil {
		return localVarReturnValue, nil, reportError("taxCategoryCreate is required and must be specified")
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
	localVarPostBody = r.taxCategoryCreate
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
