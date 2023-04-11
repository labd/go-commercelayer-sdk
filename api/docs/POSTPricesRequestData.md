# POSTPricesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPricesRequestDataAttributes**](POSTPricesRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTPricesRequestDataRelationships**](POSTPricesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTPricesRequestData

`func NewPOSTPricesRequestData(type_ interface{}, attributes POSTPricesRequestDataAttributes, ) *POSTPricesRequestData`

NewPOSTPricesRequestData instantiates a new POSTPricesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPricesRequestDataWithDefaults

`func NewPOSTPricesRequestDataWithDefaults() *POSTPricesRequestData`

NewPOSTPricesRequestDataWithDefaults instantiates a new POSTPricesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTPricesRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTPricesRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTPricesRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTPricesRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTPricesRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTPricesRequestData) GetAttributes() POSTPricesRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTPricesRequestData) GetAttributesOk() (*POSTPricesRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTPricesRequestData) SetAttributes(v POSTPricesRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTPricesRequestData) GetRelationships() POSTPricesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTPricesRequestData) GetRelationshipsOk() (*POSTPricesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTPricesRequestData) SetRelationships(v POSTPricesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTPricesRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


