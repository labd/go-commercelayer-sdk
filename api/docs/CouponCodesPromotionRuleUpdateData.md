# CouponCodesPromotionRuleUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "coupon_codes_promotion_rules"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**POSTAdyenPayments201ResponseDataAttributes**](POSTAdyenPayments201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**GETCouponCodesPromotionRules200ResponseDataInnerRelationships**](GETCouponCodesPromotionRules200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewCouponCodesPromotionRuleUpdateData

`func NewCouponCodesPromotionRuleUpdateData(type_ string, id string, attributes POSTAdyenPayments201ResponseDataAttributes, ) *CouponCodesPromotionRuleUpdateData`

NewCouponCodesPromotionRuleUpdateData instantiates a new CouponCodesPromotionRuleUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodesPromotionRuleUpdateDataWithDefaults

`func NewCouponCodesPromotionRuleUpdateDataWithDefaults() *CouponCodesPromotionRuleUpdateData`

NewCouponCodesPromotionRuleUpdateDataWithDefaults instantiates a new CouponCodesPromotionRuleUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponCodesPromotionRuleUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponCodesPromotionRuleUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponCodesPromotionRuleUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CouponCodesPromotionRuleUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponCodesPromotionRuleUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponCodesPromotionRuleUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CouponCodesPromotionRuleUpdateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponCodesPromotionRuleUpdateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponCodesPromotionRuleUpdateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponCodesPromotionRuleUpdateData) GetRelationships() GETCouponCodesPromotionRules200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponCodesPromotionRuleUpdateData) GetRelationshipsOk() (*GETCouponCodesPromotionRules200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponCodesPromotionRuleUpdateData) SetRelationships(v GETCouponCodesPromotionRules200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponCodesPromotionRuleUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


