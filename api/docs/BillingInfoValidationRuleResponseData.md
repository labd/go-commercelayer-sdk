# BillingInfoValidationRuleResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**BillingInfoValidationRuleResponseDataRelationships**](BillingInfoValidationRuleResponseDataRelationships.md) |  | [optional] 

## Methods

### NewBillingInfoValidationRuleResponseData

`func NewBillingInfoValidationRuleResponseData() *BillingInfoValidationRuleResponseData`

NewBillingInfoValidationRuleResponseData instantiates a new BillingInfoValidationRuleResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInfoValidationRuleResponseDataWithDefaults

`func NewBillingInfoValidationRuleResponseDataWithDefaults() *BillingInfoValidationRuleResponseData`

NewBillingInfoValidationRuleResponseDataWithDefaults instantiates a new BillingInfoValidationRuleResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingInfoValidationRuleResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInfoValidationRuleResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInfoValidationRuleResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingInfoValidationRuleResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BillingInfoValidationRuleResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingInfoValidationRuleResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingInfoValidationRuleResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BillingInfoValidationRuleResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *BillingInfoValidationRuleResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BillingInfoValidationRuleResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BillingInfoValidationRuleResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BillingInfoValidationRuleResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *BillingInfoValidationRuleResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BillingInfoValidationRuleResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BillingInfoValidationRuleResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BillingInfoValidationRuleResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *BillingInfoValidationRuleResponseData) GetRelationships() BillingInfoValidationRuleResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BillingInfoValidationRuleResponseData) GetRelationshipsOk() (*BillingInfoValidationRuleResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BillingInfoValidationRuleResponseData) SetRelationships(v BillingInfoValidationRuleResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BillingInfoValidationRuleResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


