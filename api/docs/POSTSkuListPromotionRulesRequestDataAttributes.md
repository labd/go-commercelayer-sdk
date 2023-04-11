# POSTSkuListPromotionRulesRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AllSkus** | Pointer to **interface{}** | Indicates if the rule is activated only when all of the SKUs of the list is also part of the order. | [optional] 
**MinQuantity** | Pointer to **interface{}** | The min quantity of SKUs of the list that must be also part of the order. If positive, overwrites the &#39;all_skus&#39; option. When the SKU list is manual, its items quantities are honoured. | [optional] 

## Methods

### NewPOSTSkuListPromotionRulesRequestDataAttributes

`func NewPOSTSkuListPromotionRulesRequestDataAttributes() *POSTSkuListPromotionRulesRequestDataAttributes`

NewPOSTSkuListPromotionRulesRequestDataAttributes instantiates a new POSTSkuListPromotionRulesRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkuListPromotionRulesRequestDataAttributesWithDefaults

`func NewPOSTSkuListPromotionRulesRequestDataAttributesWithDefaults() *POSTSkuListPromotionRulesRequestDataAttributes`

NewPOSTSkuListPromotionRulesRequestDataAttributesWithDefaults instantiates a new POSTSkuListPromotionRulesRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTSkuListPromotionRulesRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTSkuListPromotionRulesRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTSkuListPromotionRulesRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAllSkus

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetAllSkus() interface{}`

GetAllSkus returns the AllSkus field if non-nil, zero value otherwise.

### GetAllSkusOk

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetAllSkusOk() (*interface{}, bool)`

GetAllSkusOk returns a tuple with the AllSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSkus

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetAllSkus(v interface{})`

SetAllSkus sets AllSkus field to given value.

### HasAllSkus

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) HasAllSkus() bool`

HasAllSkus returns a boolean if a field has been set.

### SetAllSkusNil

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetAllSkusNil(b bool)`

 SetAllSkusNil sets the value for AllSkus to be an explicit nil

### UnsetAllSkus
`func (o *POSTSkuListPromotionRulesRequestDataAttributes) UnsetAllSkus()`

UnsetAllSkus ensures that no value is present for AllSkus, not even an explicit nil
### GetMinQuantity

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetMinQuantity() interface{}`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) GetMinQuantityOk() (*interface{}, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetMinQuantity(v interface{})`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### SetMinQuantityNil

`func (o *POSTSkuListPromotionRulesRequestDataAttributes) SetMinQuantityNil(b bool)`

 SetMinQuantityNil sets the value for MinQuantity to be an explicit nil

### UnsetMinQuantity
`func (o *POSTSkuListPromotionRulesRequestDataAttributes) UnsetMinQuantity()`

UnsetMinQuantity ensures that no value is present for MinQuantity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


