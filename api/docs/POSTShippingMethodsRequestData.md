# POSTShippingMethodsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTShippingMethodsRequestDataAttributes**](POSTShippingMethodsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTShippingMethodsRequestDataRelationships**](POSTShippingMethodsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTShippingMethodsRequestData

`func NewPOSTShippingMethodsRequestData(type_ interface{}, attributes POSTShippingMethodsRequestDataAttributes, ) *POSTShippingMethodsRequestData`

NewPOSTShippingMethodsRequestData instantiates a new POSTShippingMethodsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTShippingMethodsRequestDataWithDefaults

`func NewPOSTShippingMethodsRequestDataWithDefaults() *POSTShippingMethodsRequestData`

NewPOSTShippingMethodsRequestDataWithDefaults instantiates a new POSTShippingMethodsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTShippingMethodsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTShippingMethodsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTShippingMethodsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTShippingMethodsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTShippingMethodsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTShippingMethodsRequestData) GetAttributes() POSTShippingMethodsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTShippingMethodsRequestData) GetAttributesOk() (*POSTShippingMethodsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTShippingMethodsRequestData) SetAttributes(v POSTShippingMethodsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTShippingMethodsRequestData) GetRelationships() POSTShippingMethodsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTShippingMethodsRequestData) GetRelationshipsOk() (*POSTShippingMethodsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTShippingMethodsRequestData) SetRelationships(v POSTShippingMethodsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTShippingMethodsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


