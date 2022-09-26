# StripePaymentCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**StripePaymentCreateDataAttributes**](StripePaymentCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentCreateDataRelationships**](AdyenPaymentCreateDataRelationships.md) |  | [optional] 

## Methods

### NewStripePaymentCreateData

`func NewStripePaymentCreateData(type_ string, attributes StripePaymentCreateDataAttributes, ) *StripePaymentCreateData`

NewStripePaymentCreateData instantiates a new StripePaymentCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripePaymentCreateDataWithDefaults

`func NewStripePaymentCreateDataWithDefaults() *StripePaymentCreateData`

NewStripePaymentCreateDataWithDefaults instantiates a new StripePaymentCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StripePaymentCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripePaymentCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripePaymentCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *StripePaymentCreateData) GetAttributes() StripePaymentCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StripePaymentCreateData) GetAttributesOk() (*StripePaymentCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StripePaymentCreateData) SetAttributes(v StripePaymentCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StripePaymentCreateData) GetRelationships() AdyenPaymentCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StripePaymentCreateData) GetRelationshipsOk() (*AdyenPaymentCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StripePaymentCreateData) SetRelationships(v AdyenPaymentCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StripePaymentCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


