# MerchantUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "merchants"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**MerchantUpdateDataAttributes**](MerchantUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**MerchantUpdateDataRelationships**](MerchantUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewMerchantUpdateData

`func NewMerchantUpdateData(type_ string, id string, attributes MerchantUpdateDataAttributes, ) *MerchantUpdateData`

NewMerchantUpdateData instantiates a new MerchantUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantUpdateDataWithDefaults

`func NewMerchantUpdateDataWithDefaults() *MerchantUpdateData`

NewMerchantUpdateDataWithDefaults instantiates a new MerchantUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MerchantUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *MerchantUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *MerchantUpdateData) GetAttributes() MerchantUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MerchantUpdateData) GetAttributesOk() (*MerchantUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MerchantUpdateData) SetAttributes(v MerchantUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *MerchantUpdateData) GetRelationships() MerchantUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MerchantUpdateData) GetRelationshipsOk() (*MerchantUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MerchantUpdateData) SetRelationships(v MerchantUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MerchantUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


