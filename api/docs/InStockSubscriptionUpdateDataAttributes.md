# InStockSubscriptionUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **string** | The code of the associated SKU, replace the relationship | [optional] 
**StockThreshold** | Pointer to **int32** | The threshold at which to trigger the back in stock notification, default to 1. | [optional] 
**Activate** | Pointer to **bool** | Send this attribute if you want to activate an inactive subscription. | [optional] 
**Deactivate** | Pointer to **bool** | Send this attribute if you want to dactivate an active subscription. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewInStockSubscriptionUpdateDataAttributes

`func NewInStockSubscriptionUpdateDataAttributes() *InStockSubscriptionUpdateDataAttributes`

NewInStockSubscriptionUpdateDataAttributes instantiates a new InStockSubscriptionUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInStockSubscriptionUpdateDataAttributesWithDefaults

`func NewInStockSubscriptionUpdateDataAttributesWithDefaults() *InStockSubscriptionUpdateDataAttributes`

NewInStockSubscriptionUpdateDataAttributesWithDefaults instantiates a new InStockSubscriptionUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *InStockSubscriptionUpdateDataAttributes) GetSkuCode() string`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *InStockSubscriptionUpdateDataAttributes) GetSkuCodeOk() (*string, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *InStockSubscriptionUpdateDataAttributes) SetSkuCode(v string)`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *InStockSubscriptionUpdateDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### GetStockThreshold

`func (o *InStockSubscriptionUpdateDataAttributes) GetStockThreshold() int32`

GetStockThreshold returns the StockThreshold field if non-nil, zero value otherwise.

### GetStockThresholdOk

`func (o *InStockSubscriptionUpdateDataAttributes) GetStockThresholdOk() (*int32, bool)`

GetStockThresholdOk returns a tuple with the StockThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockThreshold

`func (o *InStockSubscriptionUpdateDataAttributes) SetStockThreshold(v int32)`

SetStockThreshold sets StockThreshold field to given value.

### HasStockThreshold

`func (o *InStockSubscriptionUpdateDataAttributes) HasStockThreshold() bool`

HasStockThreshold returns a boolean if a field has been set.

### GetActivate

`func (o *InStockSubscriptionUpdateDataAttributes) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *InStockSubscriptionUpdateDataAttributes) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *InStockSubscriptionUpdateDataAttributes) SetActivate(v bool)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *InStockSubscriptionUpdateDataAttributes) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetDeactivate

`func (o *InStockSubscriptionUpdateDataAttributes) GetDeactivate() bool`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *InStockSubscriptionUpdateDataAttributes) GetDeactivateOk() (*bool, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *InStockSubscriptionUpdateDataAttributes) SetDeactivate(v bool)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *InStockSubscriptionUpdateDataAttributes) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetReference

`func (o *InStockSubscriptionUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *InStockSubscriptionUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *InStockSubscriptionUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *InStockSubscriptionUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *InStockSubscriptionUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *InStockSubscriptionUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *InStockSubscriptionUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *InStockSubscriptionUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *InStockSubscriptionUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InStockSubscriptionUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InStockSubscriptionUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InStockSubscriptionUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


