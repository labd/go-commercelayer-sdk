# FreeShippingPromotionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "free_shipping_promotions"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes**](PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalPromotionCreateDataRelationships**](ExternalPromotionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewFreeShippingPromotionUpdateData

`func NewFreeShippingPromotionUpdateData(type_ string, id string, attributes PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes, ) *FreeShippingPromotionUpdateData`

NewFreeShippingPromotionUpdateData instantiates a new FreeShippingPromotionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeShippingPromotionUpdateDataWithDefaults

`func NewFreeShippingPromotionUpdateDataWithDefaults() *FreeShippingPromotionUpdateData`

NewFreeShippingPromotionUpdateDataWithDefaults instantiates a new FreeShippingPromotionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FreeShippingPromotionUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FreeShippingPromotionUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FreeShippingPromotionUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FreeShippingPromotionUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FreeShippingPromotionUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FreeShippingPromotionUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *FreeShippingPromotionUpdateData) GetAttributes() PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FreeShippingPromotionUpdateData) GetAttributesOk() (*PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FreeShippingPromotionUpdateData) SetAttributes(v PATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FreeShippingPromotionUpdateData) GetRelationships() ExternalPromotionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FreeShippingPromotionUpdateData) GetRelationshipsOk() (*ExternalPromotionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FreeShippingPromotionUpdateData) SetRelationships(v ExternalPromotionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FreeShippingPromotionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


