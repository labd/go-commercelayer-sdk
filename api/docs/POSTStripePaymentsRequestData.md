# POSTStripePaymentsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTStripePaymentsRequestDataAttributes**](POSTStripePaymentsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTAdyenPaymentsRequestDataRelationships**](POSTAdyenPaymentsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTStripePaymentsRequestData

`func NewPOSTStripePaymentsRequestData(type_ interface{}, attributes POSTStripePaymentsRequestDataAttributes, ) *POSTStripePaymentsRequestData`

NewPOSTStripePaymentsRequestData instantiates a new POSTStripePaymentsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStripePaymentsRequestDataWithDefaults

`func NewPOSTStripePaymentsRequestDataWithDefaults() *POSTStripePaymentsRequestData`

NewPOSTStripePaymentsRequestDataWithDefaults instantiates a new POSTStripePaymentsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTStripePaymentsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTStripePaymentsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTStripePaymentsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTStripePaymentsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTStripePaymentsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTStripePaymentsRequestData) GetAttributes() POSTStripePaymentsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTStripePaymentsRequestData) GetAttributesOk() (*POSTStripePaymentsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTStripePaymentsRequestData) SetAttributes(v POSTStripePaymentsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTStripePaymentsRequestData) GetRelationships() POSTAdyenPaymentsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTStripePaymentsRequestData) GetRelationshipsOk() (*POSTAdyenPaymentsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTStripePaymentsRequestData) SetRelationships(v POSTAdyenPaymentsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTStripePaymentsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


