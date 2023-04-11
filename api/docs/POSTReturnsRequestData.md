# POSTReturnsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTAdyenPaymentsRequestDataAttributes**](POSTAdyenPaymentsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTReturnsRequestDataRelationships**](POSTReturnsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTReturnsRequestData

`func NewPOSTReturnsRequestData(type_ interface{}, attributes POSTAdyenPaymentsRequestDataAttributes, ) *POSTReturnsRequestData`

NewPOSTReturnsRequestData instantiates a new POSTReturnsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTReturnsRequestDataWithDefaults

`func NewPOSTReturnsRequestDataWithDefaults() *POSTReturnsRequestData`

NewPOSTReturnsRequestDataWithDefaults instantiates a new POSTReturnsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTReturnsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTReturnsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTReturnsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTReturnsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTReturnsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTReturnsRequestData) GetAttributes() POSTAdyenPaymentsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTReturnsRequestData) GetAttributesOk() (*POSTAdyenPaymentsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTReturnsRequestData) SetAttributes(v POSTAdyenPaymentsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTReturnsRequestData) GetRelationships() POSTReturnsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTReturnsRequestData) GetRelationshipsOk() (*POSTReturnsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTReturnsRequestData) SetRelationships(v POSTReturnsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTReturnsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


