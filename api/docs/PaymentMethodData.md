# PaymentMethodData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**GETPaymentMethods200ResponseDataInnerAttributes**](GETPaymentMethods200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**PaymentMethodDataRelationships**](PaymentMethodDataRelationships.md) |  | [optional] 

## Methods

### NewPaymentMethodData

`func NewPaymentMethodData(type_ string, attributes GETPaymentMethods200ResponseDataInnerAttributes, ) *PaymentMethodData`

NewPaymentMethodData instantiates a new PaymentMethodData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodDataWithDefaults

`func NewPaymentMethodDataWithDefaults() *PaymentMethodData`

NewPaymentMethodDataWithDefaults instantiates a new PaymentMethodData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PaymentMethodData) GetAttributes() GETPaymentMethods200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaymentMethodData) GetAttributesOk() (*GETPaymentMethods200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaymentMethodData) SetAttributes(v GETPaymentMethods200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaymentMethodData) GetRelationships() PaymentMethodDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaymentMethodData) GetRelationshipsOk() (*PaymentMethodDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaymentMethodData) SetRelationships(v PaymentMethodDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaymentMethodData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


