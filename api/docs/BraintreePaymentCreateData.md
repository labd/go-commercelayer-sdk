# BraintreePaymentCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "braintree_payments"]
**Attributes** | [**POSTBraintreePayments201ResponseDataAttributes**](POSTBraintreePayments201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTAdyenPayments201ResponseDataRelationships**](POSTAdyenPayments201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewBraintreePaymentCreateData

`func NewBraintreePaymentCreateData(type_ string, attributes POSTBraintreePayments201ResponseDataAttributes, ) *BraintreePaymentCreateData`

NewBraintreePaymentCreateData instantiates a new BraintreePaymentCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBraintreePaymentCreateDataWithDefaults

`func NewBraintreePaymentCreateDataWithDefaults() *BraintreePaymentCreateData`

NewBraintreePaymentCreateDataWithDefaults instantiates a new BraintreePaymentCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BraintreePaymentCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BraintreePaymentCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BraintreePaymentCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BraintreePaymentCreateData) GetAttributes() POSTBraintreePayments201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BraintreePaymentCreateData) GetAttributesOk() (*POSTBraintreePayments201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BraintreePaymentCreateData) SetAttributes(v POSTBraintreePayments201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BraintreePaymentCreateData) GetRelationships() POSTAdyenPayments201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BraintreePaymentCreateData) GetRelationshipsOk() (*POSTAdyenPayments201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BraintreePaymentCreateData) SetRelationships(v POSTAdyenPayments201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BraintreePaymentCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


