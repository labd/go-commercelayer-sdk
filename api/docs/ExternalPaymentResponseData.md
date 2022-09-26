# ExternalPaymentResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**ExternalPaymentResponseDataRelationships**](ExternalPaymentResponseDataRelationships.md) |  | [optional] 

## Methods

### NewExternalPaymentResponseData

`func NewExternalPaymentResponseData() *ExternalPaymentResponseData`

NewExternalPaymentResponseData instantiates a new ExternalPaymentResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentResponseDataWithDefaults

`func NewExternalPaymentResponseDataWithDefaults() *ExternalPaymentResponseData`

NewExternalPaymentResponseDataWithDefaults instantiates a new ExternalPaymentResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalPaymentResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalPaymentResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalPaymentResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalPaymentResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ExternalPaymentResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalPaymentResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalPaymentResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalPaymentResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ExternalPaymentResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalPaymentResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalPaymentResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExternalPaymentResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *ExternalPaymentResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalPaymentResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalPaymentResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ExternalPaymentResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ExternalPaymentResponseData) GetRelationships() ExternalPaymentResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalPaymentResponseData) GetRelationshipsOk() (*ExternalPaymentResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalPaymentResponseData) SetRelationships(v ExternalPaymentResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalPaymentResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


