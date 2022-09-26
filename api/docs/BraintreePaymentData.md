# BraintreePaymentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**BraintreePaymentDataAttributes**](BraintreePaymentDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentDataRelationships**](AdyenPaymentDataRelationships.md) |  | [optional] 

## Methods

### NewBraintreePaymentData

`func NewBraintreePaymentData(type_ string, attributes BraintreePaymentDataAttributes, ) *BraintreePaymentData`

NewBraintreePaymentData instantiates a new BraintreePaymentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBraintreePaymentDataWithDefaults

`func NewBraintreePaymentDataWithDefaults() *BraintreePaymentData`

NewBraintreePaymentDataWithDefaults instantiates a new BraintreePaymentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BraintreePaymentData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BraintreePaymentData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BraintreePaymentData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BraintreePaymentData) GetAttributes() BraintreePaymentDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BraintreePaymentData) GetAttributesOk() (*BraintreePaymentDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BraintreePaymentData) SetAttributes(v BraintreePaymentDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BraintreePaymentData) GetRelationships() AdyenPaymentDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BraintreePaymentData) GetRelationshipsOk() (*AdyenPaymentDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BraintreePaymentData) SetRelationships(v AdyenPaymentDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BraintreePaymentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


