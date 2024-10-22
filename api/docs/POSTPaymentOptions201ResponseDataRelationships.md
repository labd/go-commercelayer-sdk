# POSTPaymentOptions201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsOrder**](POSTAdyenPayments201ResponseDataRelationshipsOrder.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewPOSTPaymentOptions201ResponseDataRelationships

`func NewPOSTPaymentOptions201ResponseDataRelationships() *POSTPaymentOptions201ResponseDataRelationships`

NewPOSTPaymentOptions201ResponseDataRelationships instantiates a new POSTPaymentOptions201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaymentOptions201ResponseDataRelationshipsWithDefaults

`func NewPOSTPaymentOptions201ResponseDataRelationshipsWithDefaults() *POSTPaymentOptions201ResponseDataRelationships`

NewPOSTPaymentOptions201ResponseDataRelationshipsWithDefaults instantiates a new POSTPaymentOptions201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *POSTPaymentOptions201ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *POSTPaymentOptions201ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *POSTPaymentOptions201ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *POSTPaymentOptions201ResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTPaymentOptions201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTPaymentOptions201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTPaymentOptions201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTPaymentOptions201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


