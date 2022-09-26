# CheckoutComPaymentResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AdyenPaymentResponseDataRelationships**](AdyenPaymentResponseDataRelationships.md) |  | [optional] 

## Methods

### NewCheckoutComPaymentResponseData

`func NewCheckoutComPaymentResponseData() *CheckoutComPaymentResponseData`

NewCheckoutComPaymentResponseData instantiates a new CheckoutComPaymentResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutComPaymentResponseDataWithDefaults

`func NewCheckoutComPaymentResponseDataWithDefaults() *CheckoutComPaymentResponseData`

NewCheckoutComPaymentResponseDataWithDefaults instantiates a new CheckoutComPaymentResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckoutComPaymentResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutComPaymentResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutComPaymentResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckoutComPaymentResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CheckoutComPaymentResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutComPaymentResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutComPaymentResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckoutComPaymentResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *CheckoutComPaymentResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CheckoutComPaymentResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CheckoutComPaymentResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CheckoutComPaymentResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *CheckoutComPaymentResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CheckoutComPaymentResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CheckoutComPaymentResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CheckoutComPaymentResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *CheckoutComPaymentResponseData) GetRelationships() AdyenPaymentResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CheckoutComPaymentResponseData) GetRelationshipsOk() (*AdyenPaymentResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CheckoutComPaymentResponseData) SetRelationships(v AdyenPaymentResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CheckoutComPaymentResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


