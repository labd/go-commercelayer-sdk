# PriceVolumeTierResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**PriceTierResponseDataRelationships**](PriceTierResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPriceVolumeTierResponseData

`func NewPriceVolumeTierResponseData() *PriceVolumeTierResponseData`

NewPriceVolumeTierResponseData instantiates a new PriceVolumeTierResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceVolumeTierResponseDataWithDefaults

`func NewPriceVolumeTierResponseDataWithDefaults() *PriceVolumeTierResponseData`

NewPriceVolumeTierResponseDataWithDefaults instantiates a new PriceVolumeTierResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriceVolumeTierResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceVolumeTierResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceVolumeTierResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PriceVolumeTierResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PriceVolumeTierResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceVolumeTierResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceVolumeTierResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PriceVolumeTierResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *PriceVolumeTierResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PriceVolumeTierResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PriceVolumeTierResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PriceVolumeTierResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *PriceVolumeTierResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceVolumeTierResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceVolumeTierResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PriceVolumeTierResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PriceVolumeTierResponseData) GetRelationships() PriceTierResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceVolumeTierResponseData) GetRelationshipsOk() (*PriceTierResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceVolumeTierResponseData) SetRelationships(v PriceTierResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceVolumeTierResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


