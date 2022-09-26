# BraintreeGatewayResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**BraintreeGatewayResponseDataRelationships**](BraintreeGatewayResponseDataRelationships.md) |  | [optional] 

## Methods

### NewBraintreeGatewayResponseData

`func NewBraintreeGatewayResponseData() *BraintreeGatewayResponseData`

NewBraintreeGatewayResponseData instantiates a new BraintreeGatewayResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBraintreeGatewayResponseDataWithDefaults

`func NewBraintreeGatewayResponseDataWithDefaults() *BraintreeGatewayResponseData`

NewBraintreeGatewayResponseDataWithDefaults instantiates a new BraintreeGatewayResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BraintreeGatewayResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BraintreeGatewayResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BraintreeGatewayResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BraintreeGatewayResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BraintreeGatewayResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BraintreeGatewayResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BraintreeGatewayResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BraintreeGatewayResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *BraintreeGatewayResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BraintreeGatewayResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BraintreeGatewayResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BraintreeGatewayResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *BraintreeGatewayResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BraintreeGatewayResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BraintreeGatewayResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BraintreeGatewayResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BraintreeGatewayResponseData) GetRelationships() BraintreeGatewayResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BraintreeGatewayResponseData) GetRelationshipsOk() (*BraintreeGatewayResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BraintreeGatewayResponseData) SetRelationships(v BraintreeGatewayResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BraintreeGatewayResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


