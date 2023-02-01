# GETRefunds200ResponseDataInnerRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**GETAdyenPayments200ResponseDataInnerRelationshipsOrder**](GETAdyenPayments200ResponseDataInnerRelationshipsOrder.md) |  | [optional] 
**ReferenceCapture** | Pointer to [**GETRefunds200ResponseDataInnerRelationshipsReferenceCapture**](GETRefunds200ResponseDataInnerRelationshipsReferenceCapture.md) |  | [optional] 
**Events** | Pointer to [**GETAuthorizations200ResponseDataInnerRelationshipsEvents**](GETAuthorizations200ResponseDataInnerRelationshipsEvents.md) |  | [optional] 

## Methods

### NewGETRefunds200ResponseDataInnerRelationships

`func NewGETRefunds200ResponseDataInnerRelationships() *GETRefunds200ResponseDataInnerRelationships`

NewGETRefunds200ResponseDataInnerRelationships instantiates a new GETRefunds200ResponseDataInnerRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETRefunds200ResponseDataInnerRelationshipsWithDefaults

`func NewGETRefunds200ResponseDataInnerRelationshipsWithDefaults() *GETRefunds200ResponseDataInnerRelationships`

NewGETRefunds200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETRefunds200ResponseDataInnerRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GETRefunds200ResponseDataInnerRelationships) GetOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETRefunds200ResponseDataInnerRelationships) GetOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETRefunds200ResponseDataInnerRelationships) SetOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETRefunds200ResponseDataInnerRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceCapture

`func (o *GETRefunds200ResponseDataInnerRelationships) GetReferenceCapture() GETRefunds200ResponseDataInnerRelationshipsReferenceCapture`

GetReferenceCapture returns the ReferenceCapture field if non-nil, zero value otherwise.

### GetReferenceCaptureOk

`func (o *GETRefunds200ResponseDataInnerRelationships) GetReferenceCaptureOk() (*GETRefunds200ResponseDataInnerRelationshipsReferenceCapture, bool)`

GetReferenceCaptureOk returns a tuple with the ReferenceCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCapture

`func (o *GETRefunds200ResponseDataInnerRelationships) SetReferenceCapture(v GETRefunds200ResponseDataInnerRelationshipsReferenceCapture)`

SetReferenceCapture sets ReferenceCapture field to given value.

### HasReferenceCapture

`func (o *GETRefunds200ResponseDataInnerRelationships) HasReferenceCapture() bool`

HasReferenceCapture returns a boolean if a field has been set.

### GetEvents

`func (o *GETRefunds200ResponseDataInnerRelationships) GetEvents() GETAuthorizations200ResponseDataInnerRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETRefunds200ResponseDataInnerRelationships) GetEventsOk() (*GETAuthorizations200ResponseDataInnerRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETRefunds200ResponseDataInnerRelationships) SetEvents(v GETAuthorizations200ResponseDataInnerRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETRefunds200ResponseDataInnerRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


