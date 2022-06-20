# MerchantData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "merchants"]
**Attributes** | [**MerchantDataAttributes**](MerchantDataAttributes.md) |  | 
**Relationships** | Pointer to [**MerchantDataRelationships**](MerchantDataRelationships.md) |  | [optional] 

## Methods

### NewMerchantData

`func NewMerchantData(type_ string, attributes MerchantDataAttributes, ) *MerchantData`

NewMerchantData instantiates a new MerchantData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantDataWithDefaults

`func NewMerchantDataWithDefaults() *MerchantData`

NewMerchantDataWithDefaults instantiates a new MerchantData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MerchantData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *MerchantData) GetAttributes() MerchantDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MerchantData) GetAttributesOk() (*MerchantDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MerchantData) SetAttributes(v MerchantDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *MerchantData) GetRelationships() MerchantDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MerchantData) GetRelationshipsOk() (*MerchantDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MerchantData) SetRelationships(v MerchantDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MerchantData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


