# ShippingZoneResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**ShippingZoneResponseDataRelationships**](ShippingZoneResponseDataRelationships.md) |  | [optional] 

## Methods

### NewShippingZoneResponseData

`func NewShippingZoneResponseData() *ShippingZoneResponseData`

NewShippingZoneResponseData instantiates a new ShippingZoneResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingZoneResponseDataWithDefaults

`func NewShippingZoneResponseDataWithDefaults() *ShippingZoneResponseData`

NewShippingZoneResponseDataWithDefaults instantiates a new ShippingZoneResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingZoneResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingZoneResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingZoneResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingZoneResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ShippingZoneResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingZoneResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingZoneResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShippingZoneResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ShippingZoneResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ShippingZoneResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ShippingZoneResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ShippingZoneResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *ShippingZoneResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingZoneResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingZoneResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ShippingZoneResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ShippingZoneResponseData) GetRelationships() ShippingZoneResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingZoneResponseData) GetRelationshipsOk() (*ShippingZoneResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingZoneResponseData) SetRelationships(v ShippingZoneResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingZoneResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


