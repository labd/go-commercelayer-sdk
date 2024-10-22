# WireTransferDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewWireTransferDataRelationships

`func NewWireTransferDataRelationships() *WireTransferDataRelationships`

NewWireTransferDataRelationships instantiates a new WireTransferDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireTransferDataRelationshipsWithDefaults

`func NewWireTransferDataRelationshipsWithDefaults() *WireTransferDataRelationships`

NewWireTransferDataRelationshipsWithDefaults instantiates a new WireTransferDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *WireTransferDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WireTransferDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WireTransferDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WireTransferDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetVersions

`func (o *WireTransferDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *WireTransferDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *WireTransferDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *WireTransferDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


