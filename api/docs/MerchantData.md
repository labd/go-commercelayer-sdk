# MerchantData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETMerchantsMerchantId200ResponseDataAttributes**](GETMerchantsMerchantId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**MerchantDataRelationships**](MerchantDataRelationships.md) |  | [optional] 

## Methods

### NewMerchantData

`func NewMerchantData(type_ interface{}, attributes GETMerchantsMerchantId200ResponseDataAttributes, ) *MerchantData`

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

`func (o *MerchantData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *MerchantData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MerchantData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *MerchantData) GetAttributes() GETMerchantsMerchantId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MerchantData) GetAttributesOk() (*GETMerchantsMerchantId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MerchantData) SetAttributes(v GETMerchantsMerchantId200ResponseDataAttributes)`

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


