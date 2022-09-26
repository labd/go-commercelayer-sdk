# CouponCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The coupon code, that uniquely identifies the coupon within the promotion rule. | 
**CustomerSingleUse** | Pointer to **bool** | Indicates if the coupon can be used just once per customer. | [optional] 
**UsageLimit** | **int32** | The total number of times this coupon can be used. | 
**RecipientEmail** | Pointer to **string** | The email address of the associated recipient. When creating or updating a coupon, this is a shortcut to find or create the associated recipient by email. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewCouponCreateDataAttributes

`func NewCouponCreateDataAttributes(code string, usageLimit int32, ) *CouponCreateDataAttributes`

NewCouponCreateDataAttributes instantiates a new CouponCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCreateDataAttributesWithDefaults

`func NewCouponCreateDataAttributesWithDefaults() *CouponCreateDataAttributes`

NewCouponCreateDataAttributesWithDefaults instantiates a new CouponCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CouponCreateDataAttributes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CouponCreateDataAttributes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CouponCreateDataAttributes) SetCode(v string)`

SetCode sets Code field to given value.


### GetCustomerSingleUse

`func (o *CouponCreateDataAttributes) GetCustomerSingleUse() bool`

GetCustomerSingleUse returns the CustomerSingleUse field if non-nil, zero value otherwise.

### GetCustomerSingleUseOk

`func (o *CouponCreateDataAttributes) GetCustomerSingleUseOk() (*bool, bool)`

GetCustomerSingleUseOk returns a tuple with the CustomerSingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSingleUse

`func (o *CouponCreateDataAttributes) SetCustomerSingleUse(v bool)`

SetCustomerSingleUse sets CustomerSingleUse field to given value.

### HasCustomerSingleUse

`func (o *CouponCreateDataAttributes) HasCustomerSingleUse() bool`

HasCustomerSingleUse returns a boolean if a field has been set.

### GetUsageLimit

`func (o *CouponCreateDataAttributes) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *CouponCreateDataAttributes) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *CouponCreateDataAttributes) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.


### GetRecipientEmail

`func (o *CouponCreateDataAttributes) GetRecipientEmail() string`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *CouponCreateDataAttributes) GetRecipientEmailOk() (*string, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *CouponCreateDataAttributes) SetRecipientEmail(v string)`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *CouponCreateDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### GetReference

`func (o *CouponCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CouponCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CouponCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CouponCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *CouponCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *CouponCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *CouponCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *CouponCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *CouponCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CouponCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CouponCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CouponCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


