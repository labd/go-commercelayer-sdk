# InStockSubscriptionCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerEmail** | Pointer to **string** | The email of the associated customer, replace the relationship | [optional] 
**SkuCode** | Pointer to **string** | The code of the associated SKU, replace the relationship | [optional] 
**StockThreshold** | Pointer to **int32** | The threshold at which to trigger the back in stock notification, default to 1. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewInStockSubscriptionCreateDataAttributes

`func NewInStockSubscriptionCreateDataAttributes() *InStockSubscriptionCreateDataAttributes`

NewInStockSubscriptionCreateDataAttributes instantiates a new InStockSubscriptionCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInStockSubscriptionCreateDataAttributesWithDefaults

`func NewInStockSubscriptionCreateDataAttributesWithDefaults() *InStockSubscriptionCreateDataAttributes`

NewInStockSubscriptionCreateDataAttributesWithDefaults instantiates a new InStockSubscriptionCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerEmail

`func (o *InStockSubscriptionCreateDataAttributes) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *InStockSubscriptionCreateDataAttributes) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *InStockSubscriptionCreateDataAttributes) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *InStockSubscriptionCreateDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetSkuCode

`func (o *InStockSubscriptionCreateDataAttributes) GetSkuCode() string`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *InStockSubscriptionCreateDataAttributes) GetSkuCodeOk() (*string, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *InStockSubscriptionCreateDataAttributes) SetSkuCode(v string)`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *InStockSubscriptionCreateDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### GetStockThreshold

`func (o *InStockSubscriptionCreateDataAttributes) GetStockThreshold() int32`

GetStockThreshold returns the StockThreshold field if non-nil, zero value otherwise.

### GetStockThresholdOk

`func (o *InStockSubscriptionCreateDataAttributes) GetStockThresholdOk() (*int32, bool)`

GetStockThresholdOk returns a tuple with the StockThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockThreshold

`func (o *InStockSubscriptionCreateDataAttributes) SetStockThreshold(v int32)`

SetStockThreshold sets StockThreshold field to given value.

### HasStockThreshold

`func (o *InStockSubscriptionCreateDataAttributes) HasStockThreshold() bool`

HasStockThreshold returns a boolean if a field has been set.

### GetReference

`func (o *InStockSubscriptionCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *InStockSubscriptionCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *InStockSubscriptionCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *InStockSubscriptionCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *InStockSubscriptionCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *InStockSubscriptionCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *InStockSubscriptionCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *InStockSubscriptionCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *InStockSubscriptionCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InStockSubscriptionCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InStockSubscriptionCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InStockSubscriptionCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


