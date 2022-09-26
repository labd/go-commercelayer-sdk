# CustomerPaymentSourceResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**CustomerPaymentSourceResponseDataRelationships**](CustomerPaymentSourceResponseDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerPaymentSourceResponseData

`func NewCustomerPaymentSourceResponseData() *CustomerPaymentSourceResponseData`

NewCustomerPaymentSourceResponseData instantiates a new CustomerPaymentSourceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentSourceResponseDataWithDefaults

`func NewCustomerPaymentSourceResponseDataWithDefaults() *CustomerPaymentSourceResponseData`

NewCustomerPaymentSourceResponseDataWithDefaults instantiates a new CustomerPaymentSourceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerPaymentSourceResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerPaymentSourceResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerPaymentSourceResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerPaymentSourceResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CustomerPaymentSourceResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerPaymentSourceResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerPaymentSourceResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerPaymentSourceResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *CustomerPaymentSourceResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomerPaymentSourceResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomerPaymentSourceResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomerPaymentSourceResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *CustomerPaymentSourceResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerPaymentSourceResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerPaymentSourceResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CustomerPaymentSourceResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *CustomerPaymentSourceResponseData) GetRelationships() CustomerPaymentSourceResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerPaymentSourceResponseData) GetRelationshipsOk() (*CustomerPaymentSourceResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerPaymentSourceResponseData) SetRelationships(v CustomerPaymentSourceResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerPaymentSourceResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


