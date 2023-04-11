# POSTBraintreeGatewaysRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTBraintreeGatewaysRequestDataAttributes**](POSTBraintreeGatewaysRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTBraintreeGatewaysRequestDataRelationships**](POSTBraintreeGatewaysRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTBraintreeGatewaysRequestData

`func NewPOSTBraintreeGatewaysRequestData(type_ interface{}, attributes POSTBraintreeGatewaysRequestDataAttributes, ) *POSTBraintreeGatewaysRequestData`

NewPOSTBraintreeGatewaysRequestData instantiates a new POSTBraintreeGatewaysRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBraintreeGatewaysRequestDataWithDefaults

`func NewPOSTBraintreeGatewaysRequestDataWithDefaults() *POSTBraintreeGatewaysRequestData`

NewPOSTBraintreeGatewaysRequestDataWithDefaults instantiates a new POSTBraintreeGatewaysRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTBraintreeGatewaysRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTBraintreeGatewaysRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTBraintreeGatewaysRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTBraintreeGatewaysRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTBraintreeGatewaysRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTBraintreeGatewaysRequestData) GetAttributes() POSTBraintreeGatewaysRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTBraintreeGatewaysRequestData) GetAttributesOk() (*POSTBraintreeGatewaysRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTBraintreeGatewaysRequestData) SetAttributes(v POSTBraintreeGatewaysRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTBraintreeGatewaysRequestData) GetRelationships() POSTBraintreeGatewaysRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTBraintreeGatewaysRequestData) GetRelationshipsOk() (*POSTBraintreeGatewaysRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTBraintreeGatewaysRequestData) SetRelationships(v POSTBraintreeGatewaysRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTBraintreeGatewaysRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


