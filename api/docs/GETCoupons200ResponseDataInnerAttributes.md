# GETCoupons200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **interface{}** | The coupon code, that uniquely identifies the coupon within the promotion rule. | [optional] 
**CustomerSingleUse** | Pointer to **interface{}** | Indicates if the coupon can be used just once per customer. | [optional] 
**UsageLimit** | Pointer to **interface{}** | The total number of times this coupon can be used. | [optional] 
**UsageCount** | Pointer to **interface{}** | The number of times this coupon has been used. | [optional] 
**RecipientEmail** | Pointer to **interface{}** | The email address of the associated recipient. When creating or updating a coupon, this is a shortcut to find or create the associated recipient by email. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETCoupons200ResponseDataInnerAttributes

`func NewGETCoupons200ResponseDataInnerAttributes() *GETCoupons200ResponseDataInnerAttributes`

NewGETCoupons200ResponseDataInnerAttributes instantiates a new GETCoupons200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCoupons200ResponseDataInnerAttributesWithDefaults

`func NewGETCoupons200ResponseDataInnerAttributesWithDefaults() *GETCoupons200ResponseDataInnerAttributes`

NewGETCoupons200ResponseDataInnerAttributesWithDefaults instantiates a new GETCoupons200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GETCoupons200ResponseDataInnerAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETCoupons200ResponseDataInnerAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *GETCoupons200ResponseDataInnerAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCustomerSingleUse

`func (o *GETCoupons200ResponseDataInnerAttributes) GetCustomerSingleUse() interface{}`

GetCustomerSingleUse returns the CustomerSingleUse field if non-nil, zero value otherwise.

### GetCustomerSingleUseOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetCustomerSingleUseOk() (*interface{}, bool)`

GetCustomerSingleUseOk returns a tuple with the CustomerSingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSingleUse

`func (o *GETCoupons200ResponseDataInnerAttributes) SetCustomerSingleUse(v interface{})`

SetCustomerSingleUse sets CustomerSingleUse field to given value.

### HasCustomerSingleUse

`func (o *GETCoupons200ResponseDataInnerAttributes) HasCustomerSingleUse() bool`

HasCustomerSingleUse returns a boolean if a field has been set.

### SetCustomerSingleUseNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetCustomerSingleUseNil(b bool)`

 SetCustomerSingleUseNil sets the value for CustomerSingleUse to be an explicit nil

### UnsetCustomerSingleUse
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetCustomerSingleUse()`

UnsetCustomerSingleUse ensures that no value is present for CustomerSingleUse, not even an explicit nil
### GetUsageLimit

`func (o *GETCoupons200ResponseDataInnerAttributes) GetUsageLimit() interface{}`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetUsageLimitOk() (*interface{}, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *GETCoupons200ResponseDataInnerAttributes) SetUsageLimit(v interface{})`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *GETCoupons200ResponseDataInnerAttributes) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimitNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetUsageLimitNil(b bool)`

 SetUsageLimitNil sets the value for UsageLimit to be an explicit nil

### UnsetUsageLimit
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetUsageLimit()`

UnsetUsageLimit ensures that no value is present for UsageLimit, not even an explicit nil
### GetUsageCount

`func (o *GETCoupons200ResponseDataInnerAttributes) GetUsageCount() interface{}`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetUsageCountOk() (*interface{}, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *GETCoupons200ResponseDataInnerAttributes) SetUsageCount(v interface{})`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *GETCoupons200ResponseDataInnerAttributes) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.

### SetUsageCountNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetUsageCountNil(b bool)`

 SetUsageCountNil sets the value for UsageCount to be an explicit nil

### UnsetUsageCount
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetUsageCount()`

UnsetUsageCount ensures that no value is present for UsageCount, not even an explicit nil
### GetRecipientEmail

`func (o *GETCoupons200ResponseDataInnerAttributes) GetRecipientEmail() interface{}`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetRecipientEmailOk() (*interface{}, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *GETCoupons200ResponseDataInnerAttributes) SetRecipientEmail(v interface{})`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *GETCoupons200ResponseDataInnerAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetCreatedAt

`func (o *GETCoupons200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETCoupons200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETCoupons200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETCoupons200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETCoupons200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETCoupons200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETCoupons200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETCoupons200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETCoupons200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETCoupons200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETCoupons200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETCoupons200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETCoupons200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETCoupons200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETCoupons200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETCoupons200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETCoupons200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETCoupons200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


