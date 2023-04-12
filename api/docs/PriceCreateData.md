# PriceCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPrices201ResponseDataAttributes**](POSTPrices201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceCreateDataRelationships**](PriceCreateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceCreateData

`func NewPriceCreateData(type_ interface{}, attributes POSTPrices201ResponseDataAttributes, ) *PriceCreateData`

NewPriceCreateData instantiates a new PriceCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceCreateDataWithDefaults

`func NewPriceCreateDataWithDefaults() *PriceCreateData`

NewPriceCreateDataWithDefaults instantiates a new PriceCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PriceCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PriceCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PriceCreateData) GetAttributes() POSTPrices201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceCreateData) GetAttributesOk() (*POSTPrices201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceCreateData) SetAttributes(v POSTPrices201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceCreateData) GetRelationships() PriceCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceCreateData) GetRelationshipsOk() (*PriceCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceCreateData) SetRelationships(v PriceCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


