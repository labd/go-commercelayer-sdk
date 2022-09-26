# StripeGatewayResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**StripeGatewayResponseDataRelationships**](StripeGatewayResponseDataRelationships.md) |  | [optional] 

## Methods

### NewStripeGatewayResponseData

`func NewStripeGatewayResponseData() *StripeGatewayResponseData`

NewStripeGatewayResponseData instantiates a new StripeGatewayResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeGatewayResponseDataWithDefaults

`func NewStripeGatewayResponseDataWithDefaults() *StripeGatewayResponseData`

NewStripeGatewayResponseDataWithDefaults instantiates a new StripeGatewayResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StripeGatewayResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeGatewayResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeGatewayResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripeGatewayResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *StripeGatewayResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeGatewayResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeGatewayResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StripeGatewayResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *StripeGatewayResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StripeGatewayResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StripeGatewayResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StripeGatewayResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *StripeGatewayResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StripeGatewayResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StripeGatewayResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *StripeGatewayResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *StripeGatewayResponseData) GetRelationships() StripeGatewayResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StripeGatewayResponseData) GetRelationshipsOk() (*StripeGatewayResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StripeGatewayResponseData) SetRelationships(v StripeGatewayResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StripeGatewayResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


