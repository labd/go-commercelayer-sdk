# AdyenPaymentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETAdyenPayments200ResponseDataInnerAttributes**](GETAdyenPayments200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentDataRelationships**](AdyenPaymentDataRelationships.md) |  | [optional] 

## Methods

### NewAdyenPaymentData

`func NewAdyenPaymentData(type_ interface{}, attributes GETAdyenPayments200ResponseDataInnerAttributes, ) *AdyenPaymentData`

NewAdyenPaymentData instantiates a new AdyenPaymentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenPaymentDataWithDefaults

`func NewAdyenPaymentDataWithDefaults() *AdyenPaymentData`

NewAdyenPaymentDataWithDefaults instantiates a new AdyenPaymentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdyenPaymentData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdyenPaymentData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdyenPaymentData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AdyenPaymentData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AdyenPaymentData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *AdyenPaymentData) GetAttributes() GETAdyenPayments200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AdyenPaymentData) GetAttributesOk() (*GETAdyenPayments200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AdyenPaymentData) SetAttributes(v GETAdyenPayments200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AdyenPaymentData) GetRelationships() AdyenPaymentDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AdyenPaymentData) GetRelationshipsOk() (*AdyenPaymentDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AdyenPaymentData) SetRelationships(v AdyenPaymentDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AdyenPaymentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


