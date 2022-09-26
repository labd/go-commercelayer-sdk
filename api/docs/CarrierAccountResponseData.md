# CarrierAccountResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**CarrierAccountResponseDataRelationships**](CarrierAccountResponseDataRelationships.md) |  | [optional] 

## Methods

### NewCarrierAccountResponseData

`func NewCarrierAccountResponseData() *CarrierAccountResponseData`

NewCarrierAccountResponseData instantiates a new CarrierAccountResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierAccountResponseDataWithDefaults

`func NewCarrierAccountResponseDataWithDefaults() *CarrierAccountResponseData`

NewCarrierAccountResponseDataWithDefaults instantiates a new CarrierAccountResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CarrierAccountResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CarrierAccountResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CarrierAccountResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CarrierAccountResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CarrierAccountResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CarrierAccountResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CarrierAccountResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CarrierAccountResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *CarrierAccountResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CarrierAccountResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CarrierAccountResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CarrierAccountResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *CarrierAccountResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CarrierAccountResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CarrierAccountResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CarrierAccountResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *CarrierAccountResponseData) GetRelationships() CarrierAccountResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CarrierAccountResponseData) GetRelationshipsOk() (*CarrierAccountResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CarrierAccountResponseData) SetRelationships(v CarrierAccountResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CarrierAccountResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


