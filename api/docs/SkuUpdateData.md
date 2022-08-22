# SkuUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "skus"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHSkusSkuId200ResponseDataAttributes**](PATCHSkusSkuId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHSkusSkuId200ResponseDataRelationships**](PATCHSkusSkuId200ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewSkuUpdateData

`func NewSkuUpdateData(type_ string, id string, attributes PATCHSkusSkuId200ResponseDataAttributes, ) *SkuUpdateData`

NewSkuUpdateData instantiates a new SkuUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuUpdateDataWithDefaults

`func NewSkuUpdateDataWithDefaults() *SkuUpdateData`

NewSkuUpdateDataWithDefaults instantiates a new SkuUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *SkuUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkuUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkuUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *SkuUpdateData) GetAttributes() PATCHSkusSkuId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuUpdateData) GetAttributesOk() (*PATCHSkusSkuId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuUpdateData) SetAttributes(v PATCHSkusSkuId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuUpdateData) GetRelationships() PATCHSkusSkuId200ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuUpdateData) GetRelationshipsOk() (*PATCHSkusSkuId200ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuUpdateData) SetRelationships(v PATCHSkusSkuId200ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


