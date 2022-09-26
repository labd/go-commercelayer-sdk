# GETTaxRules200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The tax rule internal name. | [optional] 
**TaxRate** | Pointer to **float32** | The tax rate for this ruke. | [optional] 
**CountryCodeRegex** | Pointer to **string** | The regex that will be evaluated to match the shipping address country code. | [optional] 
**NotCountryCodeRegex** | Pointer to **string** | The regex that will be evaluated as negative match for the shipping address country code. | [optional] 
**StateCodeRegex** | Pointer to **string** | The regex that will be evaluated to match the shipping address state code. | [optional] 
**NotStateCodeRegex** | Pointer to **string** | The regex that will be evaluated as negative match for the shipping address state code. | [optional] 
**ZipCodeRegex** | Pointer to **string** | The regex that will be evaluated to match the shipping address zip code. | [optional] 
**NotZipCodeRegex** | Pointer to **string** | The regex that will be evaluated as negative match for the shipping zip country code. | [optional] 
**FreightTaxable** | Pointer to **bool** | Indicates if the freight is taxable. | [optional] 
**PaymentMethodTaxable** | Pointer to **bool** | Indicates if the payment method is taxable. | [optional] 
**GiftCardTaxable** | Pointer to **bool** | Indicates if gift cards are taxable. | [optional] 
**AdjustmentTaxable** | Pointer to **bool** | Indicates if adjustemnts are taxable. | [optional] 
**Breakdown** | Pointer to **map[string]interface{}** | The breakdown for this tax rule (if calculated). | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETTaxRules200ResponseDataInnerAttributes

`func NewGETTaxRules200ResponseDataInnerAttributes() *GETTaxRules200ResponseDataInnerAttributes`

NewGETTaxRules200ResponseDataInnerAttributes instantiates a new GETTaxRules200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETTaxRules200ResponseDataInnerAttributesWithDefaults

`func NewGETTaxRules200ResponseDataInnerAttributesWithDefaults() *GETTaxRules200ResponseDataInnerAttributes`

NewGETTaxRules200ResponseDataInnerAttributesWithDefaults instantiates a new GETTaxRules200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxRate

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetCountryCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetCountryCodeRegex() string`

GetCountryCodeRegex returns the CountryCodeRegex field if non-nil, zero value otherwise.

### GetCountryCodeRegexOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetCountryCodeRegexOk() (*string, bool)`

GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetCountryCodeRegex(v string)`

SetCountryCodeRegex sets CountryCodeRegex field to given value.

### HasCountryCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasCountryCodeRegex() bool`

HasCountryCodeRegex returns a boolean if a field has been set.

### GetNotCountryCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotCountryCodeRegex() string`

GetNotCountryCodeRegex returns the NotCountryCodeRegex field if non-nil, zero value otherwise.

### GetNotCountryCodeRegexOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotCountryCodeRegexOk() (*string, bool)`

GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCountryCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetNotCountryCodeRegex(v string)`

SetNotCountryCodeRegex sets NotCountryCodeRegex field to given value.

### HasNotCountryCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasNotCountryCodeRegex() bool`

HasNotCountryCodeRegex returns a boolean if a field has been set.

### GetStateCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetStateCodeRegex() string`

GetStateCodeRegex returns the StateCodeRegex field if non-nil, zero value otherwise.

### GetStateCodeRegexOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetStateCodeRegexOk() (*string, bool)`

GetStateCodeRegexOk returns a tuple with the StateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetStateCodeRegex(v string)`

SetStateCodeRegex sets StateCodeRegex field to given value.

### HasStateCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasStateCodeRegex() bool`

HasStateCodeRegex returns a boolean if a field has been set.

### GetNotStateCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotStateCodeRegex() string`

GetNotStateCodeRegex returns the NotStateCodeRegex field if non-nil, zero value otherwise.

### GetNotStateCodeRegexOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotStateCodeRegexOk() (*string, bool)`

GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStateCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetNotStateCodeRegex(v string)`

SetNotStateCodeRegex sets NotStateCodeRegex field to given value.

### HasNotStateCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasNotStateCodeRegex() bool`

HasNotStateCodeRegex returns a boolean if a field has been set.

### GetZipCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetZipCodeRegex() string`

GetZipCodeRegex returns the ZipCodeRegex field if non-nil, zero value otherwise.

### GetZipCodeRegexOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetZipCodeRegexOk() (*string, bool)`

GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetZipCodeRegex(v string)`

SetZipCodeRegex sets ZipCodeRegex field to given value.

### HasZipCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasZipCodeRegex() bool`

HasZipCodeRegex returns a boolean if a field has been set.

### GetNotZipCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotZipCodeRegex() string`

GetNotZipCodeRegex returns the NotZipCodeRegex field if non-nil, zero value otherwise.

### GetNotZipCodeRegexOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetNotZipCodeRegexOk() (*string, bool)`

GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotZipCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetNotZipCodeRegex(v string)`

SetNotZipCodeRegex sets NotZipCodeRegex field to given value.

### HasNotZipCodeRegex

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasNotZipCodeRegex() bool`

HasNotZipCodeRegex returns a boolean if a field has been set.

### GetFreightTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetFreightTaxable() bool`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetFreightTaxableOk() (*bool, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetFreightTaxable(v bool)`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### GetPaymentMethodTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetPaymentMethodTaxable() bool`

GetPaymentMethodTaxable returns the PaymentMethodTaxable field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetPaymentMethodTaxableOk() (*bool, bool)`

GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetPaymentMethodTaxable(v bool)`

SetPaymentMethodTaxable sets PaymentMethodTaxable field to given value.

### HasPaymentMethodTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasPaymentMethodTaxable() bool`

HasPaymentMethodTaxable returns a boolean if a field has been set.

### GetGiftCardTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetGiftCardTaxable() bool`

GetGiftCardTaxable returns the GiftCardTaxable field if non-nil, zero value otherwise.

### GetGiftCardTaxableOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetGiftCardTaxableOk() (*bool, bool)`

GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetGiftCardTaxable(v bool)`

SetGiftCardTaxable sets GiftCardTaxable field to given value.

### HasGiftCardTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasGiftCardTaxable() bool`

HasGiftCardTaxable returns a boolean if a field has been set.

### GetAdjustmentTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetAdjustmentTaxable() bool`

GetAdjustmentTaxable returns the AdjustmentTaxable field if non-nil, zero value otherwise.

### GetAdjustmentTaxableOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetAdjustmentTaxableOk() (*bool, bool)`

GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetAdjustmentTaxable(v bool)`

SetAdjustmentTaxable sets AdjustmentTaxable field to given value.

### HasAdjustmentTaxable

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasAdjustmentTaxable() bool`

HasAdjustmentTaxable returns a boolean if a field has been set.

### GetBreakdown

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetBreakdown() map[string]interface{}`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetBreakdownOk() (*map[string]interface{}, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetBreakdown(v map[string]interface{})`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetId

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETTaxRules200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETTaxRules200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETTaxRules200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


