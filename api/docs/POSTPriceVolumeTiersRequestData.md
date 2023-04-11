# POSTPriceVolumeTiersRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPriceVolumeTiersRequestDataAttributes**](POSTPriceVolumeTiersRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTPriceFrequencyTiersRequestDataRelationships**](POSTPriceFrequencyTiersRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTPriceVolumeTiersRequestData

`func NewPOSTPriceVolumeTiersRequestData(type_ interface{}, attributes POSTPriceVolumeTiersRequestDataAttributes, ) *POSTPriceVolumeTiersRequestData`

NewPOSTPriceVolumeTiersRequestData instantiates a new POSTPriceVolumeTiersRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPriceVolumeTiersRequestDataWithDefaults

`func NewPOSTPriceVolumeTiersRequestDataWithDefaults() *POSTPriceVolumeTiersRequestData`

NewPOSTPriceVolumeTiersRequestDataWithDefaults instantiates a new POSTPriceVolumeTiersRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTPriceVolumeTiersRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTPriceVolumeTiersRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTPriceVolumeTiersRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTPriceVolumeTiersRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTPriceVolumeTiersRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTPriceVolumeTiersRequestData) GetAttributes() POSTPriceVolumeTiersRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTPriceVolumeTiersRequestData) GetAttributesOk() (*POSTPriceVolumeTiersRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTPriceVolumeTiersRequestData) SetAttributes(v POSTPriceVolumeTiersRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTPriceVolumeTiersRequestData) GetRelationships() POSTPriceFrequencyTiersRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTPriceVolumeTiersRequestData) GetRelationshipsOk() (*POSTPriceFrequencyTiersRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTPriceVolumeTiersRequestData) SetRelationships(v POSTPriceFrequencyTiersRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTPriceVolumeTiersRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


