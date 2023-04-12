# POSTCoupons201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **interface{}** | The coupon code, that uniquely identifies the coupon within the promotion rule. | 
**CustomerSingleUse** | Pointer to **interface{}** | Indicates if the coupon can be used just once per customer. | [optional] 
**UsageLimit** | **interface{}** | The total number of times this coupon can be used. | 
**RecipientEmail** | Pointer to **interface{}** | The email address of the associated recipient. When creating or updating a coupon, this is a shortcut to find or create the associated recipient by email. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTCoupons201ResponseDataAttributes

`func NewPOSTCoupons201ResponseDataAttributes(code interface{}, usageLimit interface{}, ) *POSTCoupons201ResponseDataAttributes`

NewPOSTCoupons201ResponseDataAttributes instantiates a new POSTCoupons201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCoupons201ResponseDataAttributesWithDefaults

`func NewPOSTCoupons201ResponseDataAttributesWithDefaults() *POSTCoupons201ResponseDataAttributes`

NewPOSTCoupons201ResponseDataAttributesWithDefaults instantiates a new POSTCoupons201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *POSTCoupons201ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTCoupons201ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTCoupons201ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *POSTCoupons201ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *POSTCoupons201ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCustomerSingleUse

`func (o *POSTCoupons201ResponseDataAttributes) GetCustomerSingleUse() interface{}`

GetCustomerSingleUse returns the CustomerSingleUse field if non-nil, zero value otherwise.

### GetCustomerSingleUseOk

`func (o *POSTCoupons201ResponseDataAttributes) GetCustomerSingleUseOk() (*interface{}, bool)`

GetCustomerSingleUseOk returns a tuple with the CustomerSingleUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSingleUse

`func (o *POSTCoupons201ResponseDataAttributes) SetCustomerSingleUse(v interface{})`

SetCustomerSingleUse sets CustomerSingleUse field to given value.

### HasCustomerSingleUse

`func (o *POSTCoupons201ResponseDataAttributes) HasCustomerSingleUse() bool`

HasCustomerSingleUse returns a boolean if a field has been set.

### SetCustomerSingleUseNil

`func (o *POSTCoupons201ResponseDataAttributes) SetCustomerSingleUseNil(b bool)`

 SetCustomerSingleUseNil sets the value for CustomerSingleUse to be an explicit nil

### UnsetCustomerSingleUse
`func (o *POSTCoupons201ResponseDataAttributes) UnsetCustomerSingleUse()`

UnsetCustomerSingleUse ensures that no value is present for CustomerSingleUse, not even an explicit nil
### GetUsageLimit

`func (o *POSTCoupons201ResponseDataAttributes) GetUsageLimit() interface{}`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *POSTCoupons201ResponseDataAttributes) GetUsageLimitOk() (*interface{}, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *POSTCoupons201ResponseDataAttributes) SetUsageLimit(v interface{})`

SetUsageLimit sets UsageLimit field to given value.


### SetUsageLimitNil

`func (o *POSTCoupons201ResponseDataAttributes) SetUsageLimitNil(b bool)`

 SetUsageLimitNil sets the value for UsageLimit to be an explicit nil

### UnsetUsageLimit
`func (o *POSTCoupons201ResponseDataAttributes) UnsetUsageLimit()`

UnsetUsageLimit ensures that no value is present for UsageLimit, not even an explicit nil
### GetRecipientEmail

`func (o *POSTCoupons201ResponseDataAttributes) GetRecipientEmail() interface{}`

GetRecipientEmail returns the RecipientEmail field if non-nil, zero value otherwise.

### GetRecipientEmailOk

`func (o *POSTCoupons201ResponseDataAttributes) GetRecipientEmailOk() (*interface{}, bool)`

GetRecipientEmailOk returns a tuple with the RecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientEmail

`func (o *POSTCoupons201ResponseDataAttributes) SetRecipientEmail(v interface{})`

SetRecipientEmail sets RecipientEmail field to given value.

### HasRecipientEmail

`func (o *POSTCoupons201ResponseDataAttributes) HasRecipientEmail() bool`

HasRecipientEmail returns a boolean if a field has been set.

### SetRecipientEmailNil

`func (o *POSTCoupons201ResponseDataAttributes) SetRecipientEmailNil(b bool)`

 SetRecipientEmailNil sets the value for RecipientEmail to be an explicit nil

### UnsetRecipientEmail
`func (o *POSTCoupons201ResponseDataAttributes) UnsetRecipientEmail()`

UnsetRecipientEmail ensures that no value is present for RecipientEmail, not even an explicit nil
### GetReference

`func (o *POSTCoupons201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTCoupons201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTCoupons201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTCoupons201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTCoupons201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTCoupons201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTCoupons201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTCoupons201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTCoupons201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTCoupons201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTCoupons201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTCoupons201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTCoupons201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTCoupons201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTCoupons201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTCoupons201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTCoupons201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTCoupons201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


