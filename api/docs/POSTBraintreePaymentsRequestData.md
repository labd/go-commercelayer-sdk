# POSTBraintreePaymentsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTBraintreePaymentsRequestDataAttributes**](POSTBraintreePaymentsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTAdyenPaymentsRequestDataRelationships**](POSTAdyenPaymentsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTBraintreePaymentsRequestData

`func NewPOSTBraintreePaymentsRequestData(type_ interface{}, attributes POSTBraintreePaymentsRequestDataAttributes, ) *POSTBraintreePaymentsRequestData`

NewPOSTBraintreePaymentsRequestData instantiates a new POSTBraintreePaymentsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTBraintreePaymentsRequestDataWithDefaults

`func NewPOSTBraintreePaymentsRequestDataWithDefaults() *POSTBraintreePaymentsRequestData`

NewPOSTBraintreePaymentsRequestDataWithDefaults instantiates a new POSTBraintreePaymentsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTBraintreePaymentsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTBraintreePaymentsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTBraintreePaymentsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTBraintreePaymentsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTBraintreePaymentsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTBraintreePaymentsRequestData) GetAttributes() POSTBraintreePaymentsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTBraintreePaymentsRequestData) GetAttributesOk() (*POSTBraintreePaymentsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTBraintreePaymentsRequestData) SetAttributes(v POSTBraintreePaymentsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTBraintreePaymentsRequestData) GetRelationships() POSTAdyenPaymentsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTBraintreePaymentsRequestData) GetRelationshipsOk() (*POSTAdyenPaymentsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTBraintreePaymentsRequestData) SetRelationships(v POSTAdyenPaymentsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTBraintreePaymentsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


