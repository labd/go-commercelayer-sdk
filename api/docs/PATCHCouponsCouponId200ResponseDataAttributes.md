# PATCHCouponsCouponId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **interface{}** | The coupon code, that uniquely identifies the coupon within the promotion rule. | [optional] 
**CustomerSingleUse** | Pointer to **interface{}** | Indicates if the coupon can be used just once per customer. | [optional] 
**UsageLimit** | Pointer to **interface{}** | The total number of times this coupon can be used. | [optional] 
**RecipientEmail** | Pointer to **interface{}** | The email address of the associated recipient. When creating or updating a coupon, this is a shortcut to find or create the associated recipient by email. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHCouponsCouponId200ResponseDataAttributes

`func NewPATCHCouponsCouponId200ResponseDataAttributes() *PATCHCouponsCouponId200ResponseDataAttributes`

NewPATCHCouponsCouponId200ResponseDataAttributes instantiates a new PATCHCouponsCouponId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHCouponsCouponId200ResponseDataAttributesWithDefaults

`func NewPATCHCouponsCouponId200ResponseDataAttributesWithDefaults() *PATCHCouponsCouponId200ResponseDataAttributes`

NewPATCHCouponsCouponId200ResponseDataAttributesWithDefaults instantiates a new PATCHCouponsCouponId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PATCHCouponsCouponId200ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCustomerSingleUse

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetCustomerSingleUse() interface{}`

GetCustomerSingleUse returns the CustomerSingleUse field if non-nil, zero value otherwise.

### GetCustomerSingleUseOk

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetCustomerSingleUseOk() (*interface{}, bool)`

GetCustomerSingleUseOk returns a tuple with the CustomerSingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSingleUse

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetCustomerSingleUse(v interface{})`

SetCustomerSingleUse sets CustomerSingleUse field to given value.

### HasCustomerSingleUse

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasCustomerSingleUse() bool`

HasCustomerSingleUse returns a boolean if a field has been set.

### SetCustomerSingleUseNil

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetCustomerSingleUseNil(b bool)`

 SetCustomerSingleUseNil sets the value for CustomerSingleUse to be an explicit nil

### UnsetCustomerSingleUse
`func (o *PATCHCouponsCouponId200ResponseDataAttributes) UnsetCustomerSingleUse()`

UnsetCustomerSingleUse ensures that no value is present for CustomerSingleUse, not even an explicit nil
### GetUsageLimit

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetUsageLimit() interface{}`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetUsageLimitOk() (*interface{}, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetUsageLimit(v interface{})`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimitNil

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetUsageLimitNil(b bool)`

 SetUsageLimitNil sets the value for UsageLimit to be an explicit nil

### UnsetUsageLimit
`func (o *PATCHCouponsCouponId200ResponseDataAttributes) UnsetUsageLimit()`

UnsetUsageLimit ensures that no value is present for UsageLimit, not even an explicit nil
### GetRecipientEmail

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetRecipientEmail() interface{}`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetRecipientEmailOk() (*interface{}, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetRecipientEmail(v interface{})`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *PATCHCouponsCouponId200ResponseDataAttributes) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetReference

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHCouponsCouponId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHCouponsCouponId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHCouponsCouponId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHCouponsCouponId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


