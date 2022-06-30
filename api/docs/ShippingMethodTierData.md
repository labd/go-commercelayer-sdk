# ShippingMethodTierData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "shipping_method_tiers"]
**Attributes** | [**ShippingMethodTierDataAttributes**](ShippingMethodTierDataAttributes.md) |  | 
**Relationships** | Pointer to [**ShippingMethodTierDataRelationships**](ShippingMethodTierDataRelationships.md) |  | [optional] 

## Methods

### NewShippingMethodTierData

`func NewShippingMethodTierData(type_ string, attributes ShippingMethodTierDataAttributes, ) *ShippingMethodTierData`

NewShippingMethodTierData instantiates a new ShippingMethodTierData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodTierDataWithDefaults

`func NewShippingMethodTierDataWithDefaults() *ShippingMethodTierData`

NewShippingMethodTierDataWithDefaults instantiates a new ShippingMethodTierData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShippingMethodTierData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingMethodTierData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingMethodTierData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ShippingMethodTierData) GetAttributes() ShippingMethodTierDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingMethodTierData) GetAttributesOk() (*ShippingMethodTierDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingMethodTierData) SetAttributes(v ShippingMethodTierDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShippingMethodTierData) GetRelationships() ShippingMethodTierDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingMethodTierData) GetRelationshipsOk() (*ShippingMethodTierDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingMethodTierData) SetRelationships(v ShippingMethodTierDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingMethodTierData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


