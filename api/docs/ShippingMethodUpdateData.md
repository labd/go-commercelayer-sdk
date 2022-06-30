# ShippingMethodUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "shipping_methods"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**ShippingMethodUpdateDataAttributes**](ShippingMethodUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**ShippingMethodCreateDataRelationships**](ShippingMethodCreateDataRelationships.md) |  | [optional] 

## Methods

### NewShippingMethodUpdateData

`func NewShippingMethodUpdateData(type_ string, id string, attributes ShippingMethodUpdateDataAttributes, ) *ShippingMethodUpdateData`

NewShippingMethodUpdateData instantiates a new ShippingMethodUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodUpdateDataWithDefaults

`func NewShippingMethodUpdateDataWithDefaults() *ShippingMethodUpdateData`

NewShippingMethodUpdateDataWithDefaults instantiates a new ShippingMethodUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShippingMethodUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingMethodUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingMethodUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ShippingMethodUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingMethodUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingMethodUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ShippingMethodUpdateData) GetAttributes() ShippingMethodUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingMethodUpdateData) GetAttributesOk() (*ShippingMethodUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingMethodUpdateData) SetAttributes(v ShippingMethodUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShippingMethodUpdateData) GetRelationships() ShippingMethodCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingMethodUpdateData) GetRelationshipsOk() (*ShippingMethodCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingMethodUpdateData) SetRelationships(v ShippingMethodCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingMethodUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


