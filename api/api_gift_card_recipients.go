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

// GiftCardRecipientsApiService GiftCardRecipientsApi service
type GiftCardRecipientsApiService service

type GiftCardRecipientsApiDELETEGiftCardRecipientsGiftCardRecipientIdRequest struct {
	ctx                 context.Context
	ApiService          *GiftCardRecipientsApiService
	giftCardRecipientId string
}

func (r GiftCardRecipientsApiDELETEGiftCardRecipientsGiftCardRecipientIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEGiftCardRecipientsGiftCardRecipientIdExecute(r)
}

/*
DELETEGiftCardRecipientsGiftCardRecipientId Delete a gift card recipient

Delete a gift card recipient

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param giftCardRecipientId The resource's id
 @return GiftCardRecipientsApiDELETEGiftCardRecipientsGiftCardRecipientIdRequest
*/
func (a *GiftCardRecipientsApiService) DELETEGiftCardRecipientsGiftCardRecipientId(ctx context.Context, giftCardRecipientId string) GiftCardRecipientsApiDELETEGiftCardRecipientsGiftCardRecipientIdRequest {
	return GiftCardRecipientsApiDELETEGiftCardRecipientsGiftCardRecipientIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		giftCardRecipientId: giftCardRecipientId,
	}
}

// Execute executes the request
func (a *GiftCardRecipientsApiService) DELETEGiftCardRecipientsGiftCardRecipientIdExecute(r GiftCardRecipientsApiDELETEGiftCardRecipientsGiftCardRecipientIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardRecipientsApiService.DELETEGiftCardRecipientsGiftCardRecipientId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_card_recipients/{giftCardRecipientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"giftCardRecipientId"+"}", url.PathEscape(parameterToString(r.giftCardRecipientId, "")), -1)

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

type GiftCardRecipientsApiGETGiftCardIdGiftCardRecipientRequest struct {
	ctx        context.Context
	ApiService *GiftCardRecipientsApiService
	giftCardId string
}

func (r GiftCardRecipientsApiGETGiftCardIdGiftCardRecipientRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETGiftCardIdGiftCardRecipientExecute(r)
}

/*
GETGiftCardIdGiftCardRecipient Retrieve the gift card recipient associated to the gift card

Retrieve the gift card recipient associated to the gift card

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param giftCardId The resource's id
 @return GiftCardRecipientsApiGETGiftCardIdGiftCardRecipientRequest
*/
func (a *GiftCardRecipientsApiService) GETGiftCardIdGiftCardRecipient(ctx context.Context, giftCardId string) GiftCardRecipientsApiGETGiftCardIdGiftCardRecipientRequest {
	return GiftCardRecipientsApiGETGiftCardIdGiftCardRecipientRequest{
		ApiService: a,
		ctx:        ctx,
		giftCardId: giftCardId,
	}
}

// Execute executes the request
func (a *GiftCardRecipientsApiService) GETGiftCardIdGiftCardRecipientExecute(r GiftCardRecipientsApiGETGiftCardIdGiftCardRecipientRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardRecipientsApiService.GETGiftCardIdGiftCardRecipient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_cards/{giftCardId}/gift_card_recipient"
	localVarPath = strings.Replace(localVarPath, "{"+"giftCardId"+"}", url.PathEscape(parameterToString(r.giftCardId, "")), -1)

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

type GiftCardRecipientsApiGETGiftCardRecipientsRequest struct {
	ctx        context.Context
	ApiService *GiftCardRecipientsApiService
}

func (r GiftCardRecipientsApiGETGiftCardRecipientsRequest) Execute() (*GiftCardRecipientResponseList, *http.Response, error) {
	return r.ApiService.GETGiftCardRecipientsExecute(r)
}

/*
GETGiftCardRecipients List all gift card recipients

List all gift card recipients

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GiftCardRecipientsApiGETGiftCardRecipientsRequest
*/
func (a *GiftCardRecipientsApiService) GETGiftCardRecipients(ctx context.Context) GiftCardRecipientsApiGETGiftCardRecipientsRequest {
	return GiftCardRecipientsApiGETGiftCardRecipientsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GiftCardRecipientResponseList
func (a *GiftCardRecipientsApiService) GETGiftCardRecipientsExecute(r GiftCardRecipientsApiGETGiftCardRecipientsRequest) (*GiftCardRecipientResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GiftCardRecipientResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardRecipientsApiService.GETGiftCardRecipients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_card_recipients"

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

type GiftCardRecipientsApiGETGiftCardRecipientsGiftCardRecipientIdRequest struct {
	ctx                 context.Context
	ApiService          *GiftCardRecipientsApiService
	giftCardRecipientId string
}

func (r GiftCardRecipientsApiGETGiftCardRecipientsGiftCardRecipientIdRequest) Execute() (*GiftCardRecipientResponse, *http.Response, error) {
	return r.ApiService.GETGiftCardRecipientsGiftCardRecipientIdExecute(r)
}

/*
GETGiftCardRecipientsGiftCardRecipientId Retrieve a gift card recipient

Retrieve a gift card recipient

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param giftCardRecipientId The resource's id
 @return GiftCardRecipientsApiGETGiftCardRecipientsGiftCardRecipientIdRequest
*/
func (a *GiftCardRecipientsApiService) GETGiftCardRecipientsGiftCardRecipientId(ctx context.Context, giftCardRecipientId string) GiftCardRecipientsApiGETGiftCardRecipientsGiftCardRecipientIdRequest {
	return GiftCardRecipientsApiGETGiftCardRecipientsGiftCardRecipientIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		giftCardRecipientId: giftCardRecipientId,
	}
}

// Execute executes the request
//  @return GiftCardRecipientResponse
func (a *GiftCardRecipientsApiService) GETGiftCardRecipientsGiftCardRecipientIdExecute(r GiftCardRecipientsApiGETGiftCardRecipientsGiftCardRecipientIdRequest) (*GiftCardRecipientResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GiftCardRecipientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardRecipientsApiService.GETGiftCardRecipientsGiftCardRecipientId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_card_recipients/{giftCardRecipientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"giftCardRecipientId"+"}", url.PathEscape(parameterToString(r.giftCardRecipientId, "")), -1)

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

type GiftCardRecipientsApiPATCHGiftCardRecipientsGiftCardRecipientIdRequest struct {
	ctx                     context.Context
	ApiService              *GiftCardRecipientsApiService
	giftCardRecipientUpdate *GiftCardRecipientUpdate
	giftCardRecipientId     string
}

func (r GiftCardRecipientsApiPATCHGiftCardRecipientsGiftCardRecipientIdRequest) GiftCardRecipientUpdate(giftCardRecipientUpdate GiftCardRecipientUpdate) GiftCardRecipientsApiPATCHGiftCardRecipientsGiftCardRecipientIdRequest {
	r.giftCardRecipientUpdate = &giftCardRecipientUpdate
	return r
}

func (r GiftCardRecipientsApiPATCHGiftCardRecipientsGiftCardRecipientIdRequest) Execute() (*GiftCardRecipientResponse, *http.Response, error) {
	return r.ApiService.PATCHGiftCardRecipientsGiftCardRecipientIdExecute(r)
}

/*
PATCHGiftCardRecipientsGiftCardRecipientId Update a gift card recipient

Update a gift card recipient

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param giftCardRecipientId The resource's id
 @return GiftCardRecipientsApiPATCHGiftCardRecipientsGiftCardRecipientIdRequest
*/
func (a *GiftCardRecipientsApiService) PATCHGiftCardRecipientsGiftCardRecipientId(ctx context.Context, giftCardRecipientId string) GiftCardRecipientsApiPATCHGiftCardRecipientsGiftCardRecipientIdRequest {
	return GiftCardRecipientsApiPATCHGiftCardRecipientsGiftCardRecipientIdRequest{
		ApiService:          a,
		ctx:                 ctx,
		giftCardRecipientId: giftCardRecipientId,
	}
}

// Execute executes the request
//  @return GiftCardRecipientResponse
func (a *GiftCardRecipientsApiService) PATCHGiftCardRecipientsGiftCardRecipientIdExecute(r GiftCardRecipientsApiPATCHGiftCardRecipientsGiftCardRecipientIdRequest) (*GiftCardRecipientResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GiftCardRecipientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardRecipientsApiService.PATCHGiftCardRecipientsGiftCardRecipientId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_card_recipients/{giftCardRecipientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"giftCardRecipientId"+"}", url.PathEscape(parameterToString(r.giftCardRecipientId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.giftCardRecipientUpdate == nil {
		return localVarReturnValue, nil, reportError("giftCardRecipientUpdate is required and must be specified")
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
	localVarPostBody = r.giftCardRecipientUpdate
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

type GiftCardRecipientsApiPOSTGiftCardRecipientsRequest struct {
	ctx                     context.Context
	ApiService              *GiftCardRecipientsApiService
	giftCardRecipientCreate *GiftCardRecipientCreate
}

func (r GiftCardRecipientsApiPOSTGiftCardRecipientsRequest) GiftCardRecipientCreate(giftCardRecipientCreate GiftCardRecipientCreate) GiftCardRecipientsApiPOSTGiftCardRecipientsRequest {
	r.giftCardRecipientCreate = &giftCardRecipientCreate
	return r
}

func (r GiftCardRecipientsApiPOSTGiftCardRecipientsRequest) Execute() (*GiftCardRecipientResponse, *http.Response, error) {
	return r.ApiService.POSTGiftCardRecipientsExecute(r)
}

/*
POSTGiftCardRecipients Create a gift card recipient

Create a gift card recipient

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GiftCardRecipientsApiPOSTGiftCardRecipientsRequest
*/
func (a *GiftCardRecipientsApiService) POSTGiftCardRecipients(ctx context.Context) GiftCardRecipientsApiPOSTGiftCardRecipientsRequest {
	return GiftCardRecipientsApiPOSTGiftCardRecipientsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GiftCardRecipientResponse
func (a *GiftCardRecipientsApiService) POSTGiftCardRecipientsExecute(r GiftCardRecipientsApiPOSTGiftCardRecipientsRequest) (*GiftCardRecipientResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GiftCardRecipientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardRecipientsApiService.POSTGiftCardRecipients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_card_recipients"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.giftCardRecipientCreate == nil {
		return localVarReturnValue, nil, reportError("giftCardRecipientCreate is required and must be specified")
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
	localVarPostBody = r.giftCardRecipientCreate
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
