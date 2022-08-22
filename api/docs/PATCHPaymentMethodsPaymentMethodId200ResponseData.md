# PATCHPaymentMethodsPaymentMethodId200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "payment_methods"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes**](PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**PATCHPaymentMethodsPaymentMethodId200ResponseDataRelationships**](PATCHPaymentMethodsPaymentMethodId200ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHPaymentMethodsPaymentMethodId200ResponseData

`func NewPATCHPaymentMethodsPaymentMethodId200ResponseData() *PATCHPaymentMethodsPaymentMethodId200ResponseData`

NewPATCHPaymentMethodsPaymentMethodId200ResponseData instantiates a new PATCHPaymentMethodsPaymentMethodId200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHPaymentMethodsPaymentMethodId200ResponseDataWithDefaults

`func NewPATCHPaymentMethodsPaymentMethodId200ResponseDataWithDefaults() *PATCHPaymentMethodsPaymentMethodId200ResponseData`

NewPATCHPaymentMethodsPaymentMethodId200ResponseDataWithDefaults instantiates a new PATCHPaymentMethodsPaymentMethodId200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetAttributes() PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetAttributesOk() (*PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) SetAttributes(v PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetRelationships() PATCHPaymentMethodsPaymentMethodId200ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) GetRelationshipsOk() (*PATCHPaymentMethodsPaymentMethodId200ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) SetRelationships(v PATCHPaymentMethodsPaymentMethodId200ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHPaymentMethodsPaymentMethodId200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


