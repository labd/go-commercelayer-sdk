# ShippingWeightTierCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**POSTShippingWeightTiers201ResponseDataAttributes**](POSTShippingWeightTiers201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ShippingWeightTierCreateDataRelationships**](ShippingWeightTierCreateDataRelationships.md) |  | [optional] 

## Methods

### NewShippingWeightTierCreateData

`func NewShippingWeightTierCreateData(type_ string, attributes POSTShippingWeightTiers201ResponseDataAttributes, ) *ShippingWeightTierCreateData`

NewShippingWeightTierCreateData instantiates a new ShippingWeightTierCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingWeightTierCreateDataWithDefaults

`func NewShippingWeightTierCreateDataWithDefaults() *ShippingWeightTierCreateData`

NewShippingWeightTierCreateDataWithDefaults instantiates a new ShippingWeightTierCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShippingWeightTierCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingWeightTierCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingWeightTierCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ShippingWeightTierCreateData) GetAttributes() POSTShippingWeightTiers201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingWeightTierCreateData) GetAttributesOk() (*POSTShippingWeightTiers201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingWeightTierCreateData) SetAttributes(v POSTShippingWeightTiers201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShippingWeightTierCreateData) GetRelationships() ShippingWeightTierCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingWeightTierCreateData) GetRelationshipsOk() (*ShippingWeightTierCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingWeightTierCreateData) SetRelationships(v ShippingWeightTierCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingWeightTierCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


