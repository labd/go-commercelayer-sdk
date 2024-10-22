# PriceListSchedulerCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPriceListSchedulers201ResponseDataAttributes**](POSTPriceListSchedulers201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceListSchedulerCreateDataRelationships**](PriceListSchedulerCreateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceListSchedulerCreateData

`func NewPriceListSchedulerCreateData(type_ interface{}, attributes POSTPriceListSchedulers201ResponseDataAttributes, ) *PriceListSchedulerCreateData`

NewPriceListSchedulerCreateData instantiates a new PriceListSchedulerCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceListSchedulerCreateDataWithDefaults

`func NewPriceListSchedulerCreateDataWithDefaults() *PriceListSchedulerCreateData`

NewPriceListSchedulerCreateDataWithDefaults instantiates a new PriceListSchedulerCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceListSchedulerCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceListSchedulerCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceListSchedulerCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PriceListSchedulerCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PriceListSchedulerCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PriceListSchedulerCreateData) GetAttributes() POSTPriceListSchedulers201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceListSchedulerCreateData) GetAttributesOk() (*POSTPriceListSchedulers201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceListSchedulerCreateData) SetAttributes(v POSTPriceListSchedulers201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceListSchedulerCreateData) GetRelationships() PriceListSchedulerCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceListSchedulerCreateData) GetRelationshipsOk() (*PriceListSchedulerCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceListSchedulerCreateData) SetRelationships(v PriceListSchedulerCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceListSchedulerCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


