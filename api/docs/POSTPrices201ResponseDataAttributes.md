# POSTPrices201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **string** | The code of the associated SKU. When creating a price, either a valid sku_code or a SKU relationship must be present. | [optional] 
**AmountCents** | **int32** | The SKU price amount for the associated price list, in cents. | 
**CompareAtAmountCents** | **int32** | The compared price amount, in cents. Useful to display a percentage discount. | 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTPrices201ResponseDataAttributes

`func NewPOSTPrices201ResponseDataAttributes(amountCents int32, compareAtAmountCents int32, ) *POSTPrices201ResponseDataAttributes`

NewPOSTPrices201ResponseDataAttributes instantiates a new POSTPrices201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPrices201ResponseDataAttributesWithDefaults

`func NewPOSTPrices201ResponseDataAttributesWithDefaults() *POSTPrices201ResponseDataAttributes`

NewPOSTPrices201ResponseDataAttributesWithDefaults instantiates a new POSTPrices201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *POSTPrices201ResponseDataAttributes) GetSkuCode() string`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *POSTPrices201ResponseDataAttributes) GetSkuCodeOk() (*string, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *POSTPrices201ResponseDataAttributes) SetSkuCode(v string)`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *POSTPrices201ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### GetAmountCents

`func (o *POSTPrices201ResponseDataAttributes) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *POSTPrices201ResponseDataAttributes) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *POSTPrices201ResponseDataAttributes) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetCompareAtAmountCents

`func (o *POSTPrices201ResponseDataAttributes) GetCompareAtAmountCents() int32`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *POSTPrices201ResponseDataAttributes) GetCompareAtAmountCentsOk() (*int32, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *POSTPrices201ResponseDataAttributes) SetCompareAtAmountCents(v int32)`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.


### GetReference

`func (o *POSTPrices201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTPrices201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTPrices201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTPrices201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTPrices201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTPrices201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTPrices201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTPrices201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTPrices201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTPrices201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTPrices201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTPrices201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


