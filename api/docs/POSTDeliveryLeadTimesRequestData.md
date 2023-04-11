# POSTDeliveryLeadTimesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTDeliveryLeadTimesRequestDataAttributes**](POSTDeliveryLeadTimesRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTDeliveryLeadTimesRequestDataRelationships**](POSTDeliveryLeadTimesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTDeliveryLeadTimesRequestData

`func NewPOSTDeliveryLeadTimesRequestData(type_ interface{}, attributes POSTDeliveryLeadTimesRequestDataAttributes, ) *POSTDeliveryLeadTimesRequestData`

NewPOSTDeliveryLeadTimesRequestData instantiates a new POSTDeliveryLeadTimesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTDeliveryLeadTimesRequestDataWithDefaults

`func NewPOSTDeliveryLeadTimesRequestDataWithDefaults() *POSTDeliveryLeadTimesRequestData`

NewPOSTDeliveryLeadTimesRequestDataWithDefaults instantiates a new POSTDeliveryLeadTimesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTDeliveryLeadTimesRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTDeliveryLeadTimesRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTDeliveryLeadTimesRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTDeliveryLeadTimesRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTDeliveryLeadTimesRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTDeliveryLeadTimesRequestData) GetAttributes() POSTDeliveryLeadTimesRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTDeliveryLeadTimesRequestData) GetAttributesOk() (*POSTDeliveryLeadTimesRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTDeliveryLeadTimesRequestData) SetAttributes(v POSTDeliveryLeadTimesRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTDeliveryLeadTimesRequestData) GetRelationships() POSTDeliveryLeadTimesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTDeliveryLeadTimesRequestData) GetRelationshipsOk() (*POSTDeliveryLeadTimesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTDeliveryLeadTimesRequestData) SetRelationships(v POSTDeliveryLeadTimesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTDeliveryLeadTimesRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


