# PaypalPaymentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "paypal_payments"]
**Attributes** | [**PaypalPaymentDataAttributes**](PaypalPaymentDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentDataRelationships**](AdyenPaymentDataRelationships.md) |  | [optional] 

## Methods

### NewPaypalPaymentData

`func NewPaypalPaymentData(type_ string, attributes PaypalPaymentDataAttributes, ) *PaypalPaymentData`

NewPaypalPaymentData instantiates a new PaypalPaymentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalPaymentDataWithDefaults

`func NewPaypalPaymentDataWithDefaults() *PaypalPaymentData`

NewPaypalPaymentDataWithDefaults instantiates a new PaypalPaymentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaypalPaymentData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaypalPaymentData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaypalPaymentData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PaypalPaymentData) GetAttributes() PaypalPaymentDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaypalPaymentData) GetAttributesOk() (*PaypalPaymentDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaypalPaymentData) SetAttributes(v PaypalPaymentDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaypalPaymentData) GetRelationships() AdyenPaymentDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaypalPaymentData) GetRelationshipsOk() (*AdyenPaymentDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaypalPaymentData) SetRelationships(v AdyenPaymentDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaypalPaymentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


