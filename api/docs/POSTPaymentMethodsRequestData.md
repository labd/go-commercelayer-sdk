# POSTPaymentMethodsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPaymentMethodsRequestDataAttributes**](POSTPaymentMethodsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTPaymentMethodsRequestDataRelationships**](POSTPaymentMethodsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTPaymentMethodsRequestData

`func NewPOSTPaymentMethodsRequestData(type_ interface{}, attributes POSTPaymentMethodsRequestDataAttributes, ) *POSTPaymentMethodsRequestData`

NewPOSTPaymentMethodsRequestData instantiates a new POSTPaymentMethodsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaymentMethodsRequestDataWithDefaults

`func NewPOSTPaymentMethodsRequestDataWithDefaults() *POSTPaymentMethodsRequestData`

NewPOSTPaymentMethodsRequestDataWithDefaults instantiates a new POSTPaymentMethodsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTPaymentMethodsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTPaymentMethodsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTPaymentMethodsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTPaymentMethodsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTPaymentMethodsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTPaymentMethodsRequestData) GetAttributes() POSTPaymentMethodsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTPaymentMethodsRequestData) GetAttributesOk() (*POSTPaymentMethodsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTPaymentMethodsRequestData) SetAttributes(v POSTPaymentMethodsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTPaymentMethodsRequestData) GetRelationships() POSTPaymentMethodsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTPaymentMethodsRequestData) GetRelationshipsOk() (*POSTPaymentMethodsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTPaymentMethodsRequestData) SetRelationships(v POSTPaymentMethodsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTPaymentMethodsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


