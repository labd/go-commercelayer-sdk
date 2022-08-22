# PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "sku_list_promotion_rules"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTSkuListPromotionRules201ResponseDataAttributes**](POSTSkuListPromotionRules201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships**](PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData

`func NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData() *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData`

NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData instantiates a new PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataWithDefaults

`func NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataWithDefaults() *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData`

NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataWithDefaults instantiates a new PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetAttributes() POSTSkuListPromotionRules201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetAttributesOk() (*POSTSkuListPromotionRules201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) SetAttributes(v POSTSkuListPromotionRules201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetRelationships() PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) GetRelationshipsOk() (*PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) SetRelationships(v PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


