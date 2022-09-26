# ManualTaxCalculatorResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**ManualTaxCalculatorResponseDataRelationships**](ManualTaxCalculatorResponseDataRelationships.md) |  | [optional] 

## Methods

### NewManualTaxCalculatorResponseData

`func NewManualTaxCalculatorResponseData() *ManualTaxCalculatorResponseData`

NewManualTaxCalculatorResponseData instantiates a new ManualTaxCalculatorResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualTaxCalculatorResponseDataWithDefaults

`func NewManualTaxCalculatorResponseDataWithDefaults() *ManualTaxCalculatorResponseData`

NewManualTaxCalculatorResponseDataWithDefaults instantiates a new ManualTaxCalculatorResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualTaxCalculatorResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualTaxCalculatorResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualTaxCalculatorResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManualTaxCalculatorResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ManualTaxCalculatorResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManualTaxCalculatorResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManualTaxCalculatorResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManualTaxCalculatorResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ManualTaxCalculatorResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ManualTaxCalculatorResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ManualTaxCalculatorResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ManualTaxCalculatorResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *ManualTaxCalculatorResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManualTaxCalculatorResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManualTaxCalculatorResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ManualTaxCalculatorResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ManualTaxCalculatorResponseData) GetRelationships() ManualTaxCalculatorResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ManualTaxCalculatorResponseData) GetRelationshipsOk() (*ManualTaxCalculatorResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ManualTaxCalculatorResponseData) SetRelationships(v ManualTaxCalculatorResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ManualTaxCalculatorResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


