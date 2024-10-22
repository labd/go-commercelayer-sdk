# PriceListSchedulerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes**](GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceListSchedulerDataRelationships**](PriceListSchedulerDataRelationships.md) |  | [optional] 

## Methods

### NewPriceListSchedulerData

`func NewPriceListSchedulerData(type_ interface{}, attributes GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes, ) *PriceListSchedulerData`

NewPriceListSchedulerData instantiates a new PriceListSchedulerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceListSchedulerDataWithDefaults

`func NewPriceListSchedulerDataWithDefaults() *PriceListSchedulerData`

NewPriceListSchedulerDataWithDefaults instantiates a new PriceListSchedulerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceListSchedulerData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceListSchedulerData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceListSchedulerData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PriceListSchedulerData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PriceListSchedulerData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PriceListSchedulerData) GetAttributes() GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceListSchedulerData) GetAttributesOk() (*GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceListSchedulerData) SetAttributes(v GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceListSchedulerData) GetRelationships() PriceListSchedulerDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceListSchedulerData) GetRelationshipsOk() (*PriceListSchedulerDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceListSchedulerData) SetRelationships(v PriceListSchedulerDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceListSchedulerData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


