# POSTSkuListPromotionRules201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AllSkus** | Pointer to **bool** | Indicates if the rule is activated only when all of the SKUs of the list is also part of the order. | [optional] 
**MinQuantity** | Pointer to **int32** | The min quantity of SKUs of the list that must be also part of the order. If positive, overwrites the &#39;all_skus&#39; option. When the SKU list is manual, its items quantities are honoured. | [optional] 

## Methods

### NewPOSTSkuListPromotionRules201ResponseDataAttributes

`func NewPOSTSkuListPromotionRules201ResponseDataAttributes() *POSTSkuListPromotionRules201ResponseDataAttributes`

NewPOSTSkuListPromotionRules201ResponseDataAttributes instantiates a new POSTSkuListPromotionRules201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkuListPromotionRules201ResponseDataAttributesWithDefaults

`func NewPOSTSkuListPromotionRules201ResponseDataAttributesWithDefaults() *POSTSkuListPromotionRules201ResponseDataAttributes`

NewPOSTSkuListPromotionRules201ResponseDataAttributesWithDefaults instantiates a new POSTSkuListPromotionRules201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAllSkus

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetAllSkus() bool`

GetAllSkus returns the AllSkus field if non-nil, zero value otherwise.

### GetAllSkusOk

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetAllSkusOk() (*bool, bool)`

GetAllSkusOk returns a tuple with the AllSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSkus

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetAllSkus(v bool)`

SetAllSkus sets AllSkus field to given value.

### HasAllSkus

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasAllSkus() bool`

HasAllSkus returns a boolean if a field has been set.

### GetMinQuantity

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetMinQuantity() int32`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetMinQuantityOk() (*int32, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetMinQuantity(v int32)`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


