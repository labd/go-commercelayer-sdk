# KlarnaPaymentResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AdyenPaymentResponseDataRelationships**](AdyenPaymentResponseDataRelationships.md) |  | [optional] 

## Methods

### NewKlarnaPaymentResponseData

`func NewKlarnaPaymentResponseData() *KlarnaPaymentResponseData`

NewKlarnaPaymentResponseData instantiates a new KlarnaPaymentResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlarnaPaymentResponseDataWithDefaults

`func NewKlarnaPaymentResponseDataWithDefaults() *KlarnaPaymentResponseData`

NewKlarnaPaymentResponseDataWithDefaults instantiates a new KlarnaPaymentResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KlarnaPaymentResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KlarnaPaymentResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KlarnaPaymentResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KlarnaPaymentResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KlarnaPaymentResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KlarnaPaymentResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KlarnaPaymentResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KlarnaPaymentResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *KlarnaPaymentResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *KlarnaPaymentResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *KlarnaPaymentResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *KlarnaPaymentResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *KlarnaPaymentResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *KlarnaPaymentResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *KlarnaPaymentResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *KlarnaPaymentResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *KlarnaPaymentResponseData) GetRelationships() AdyenPaymentResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *KlarnaPaymentResponseData) GetRelationshipsOk() (*AdyenPaymentResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *KlarnaPaymentResponseData) SetRelationships(v AdyenPaymentResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *KlarnaPaymentResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


