# PaymentMethodCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**POSTPaymentMethods201ResponseDataAttributes**](POSTPaymentMethods201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PaymentMethodCreateDataRelationships**](PaymentMethodCreateDataRelationships.md) |  | [optional] 

## Methods

### NewPaymentMethodCreateData

`func NewPaymentMethodCreateData(type_ string, attributes POSTPaymentMethods201ResponseDataAttributes, ) *PaymentMethodCreateData`

NewPaymentMethodCreateData instantiates a new PaymentMethodCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCreateDataWithDefaults

`func NewPaymentMethodCreateDataWithDefaults() *PaymentMethodCreateData`

NewPaymentMethodCreateDataWithDefaults instantiates a new PaymentMethodCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PaymentMethodCreateData) GetAttributes() POSTPaymentMethods201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaymentMethodCreateData) GetAttributesOk() (*POSTPaymentMethods201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaymentMethodCreateData) SetAttributes(v POSTPaymentMethods201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaymentMethodCreateData) GetRelationships() PaymentMethodCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaymentMethodCreateData) GetRelationshipsOk() (*PaymentMethodCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaymentMethodCreateData) SetRelationships(v PaymentMethodCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaymentMethodCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


