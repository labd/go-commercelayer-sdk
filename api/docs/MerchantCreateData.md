# MerchantCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTMerchants201ResponseDataAttributes**](POSTMerchants201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**MerchantCreateDataRelationships**](MerchantCreateDataRelationships.md) |  | [optional] 

## Methods

### NewMerchantCreateData

`func NewMerchantCreateData(type_ interface{}, attributes POSTMerchants201ResponseDataAttributes, ) *MerchantCreateData`

NewMerchantCreateData instantiates a new MerchantCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCreateDataWithDefaults

`func NewMerchantCreateDataWithDefaults() *MerchantCreateData`

NewMerchantCreateDataWithDefaults instantiates a new MerchantCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MerchantCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *MerchantCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MerchantCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *MerchantCreateData) GetAttributes() POSTMerchants201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MerchantCreateData) GetAttributesOk() (*POSTMerchants201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MerchantCreateData) SetAttributes(v POSTMerchants201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *MerchantCreateData) GetRelationships() MerchantCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MerchantCreateData) GetRelationshipsOk() (*MerchantCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MerchantCreateData) SetRelationships(v MerchantCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MerchantCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


