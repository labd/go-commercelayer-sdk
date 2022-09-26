# ShippingMethodTierResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**ShippingMethodTierResponseDataRelationships**](ShippingMethodTierResponseDataRelationships.md) |  | [optional] 

## Methods

### NewShippingMethodTierResponseData

`func NewShippingMethodTierResponseData() *ShippingMethodTierResponseData`

NewShippingMethodTierResponseData instantiates a new ShippingMethodTierResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodTierResponseDataWithDefaults

`func NewShippingMethodTierResponseDataWithDefaults() *ShippingMethodTierResponseData`

NewShippingMethodTierResponseDataWithDefaults instantiates a new ShippingMethodTierResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingMethodTierResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingMethodTierResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingMethodTierResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingMethodTierResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ShippingMethodTierResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingMethodTierResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingMethodTierResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShippingMethodTierResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ShippingMethodTierResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ShippingMethodTierResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ShippingMethodTierResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ShippingMethodTierResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *ShippingMethodTierResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingMethodTierResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingMethodTierResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ShippingMethodTierResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ShippingMethodTierResponseData) GetRelationships() ShippingMethodTierResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingMethodTierResponseData) GetRelationshipsOk() (*ShippingMethodTierResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingMethodTierResponseData) SetRelationships(v ShippingMethodTierResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingMethodTierResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


