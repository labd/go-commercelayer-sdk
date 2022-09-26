# TaxRuleUpdateDataAttributes

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
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewTaxRuleUpdateDataAttributes

`func NewTaxRuleUpdateDataAttributes() *TaxRuleUpdateDataAttributes`

NewTaxRuleUpdateDataAttributes instantiates a new TaxRuleUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRuleUpdateDataAttributesWithDefaults

`func NewTaxRuleUpdateDataAttributesWithDefaults() *TaxRuleUpdateDataAttributes`

NewTaxRuleUpdateDataAttributesWithDefaults instantiates a new TaxRuleUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaxRuleUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxRuleUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxRuleUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxRuleUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxRate

`func (o *TaxRuleUpdateDataAttributes) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *TaxRuleUpdateDataAttributes) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *TaxRuleUpdateDataAttributes) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *TaxRuleUpdateDataAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetCountryCodeRegex

`func (o *TaxRuleUpdateDataAttributes) GetCountryCodeRegex() string`

GetCountryCodeRegex returns the CountryCodeRegex field if non-nil, zero value otherwise.

### GetCountryCodeRegexOk

`func (o *TaxRuleUpdateDataAttributes) GetCountryCodeRegexOk() (*string, bool)`

GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeRegex

`func (o *TaxRuleUpdateDataAttributes) SetCountryCodeRegex(v string)`

SetCountryCodeRegex sets CountryCodeRegex field to given value.

### HasCountryCodeRegex

`func (o *TaxRuleUpdateDataAttributes) HasCountryCodeRegex() bool`

HasCountryCodeRegex returns a boolean if a field has been set.

### GetNotCountryCodeRegex

`func (o *TaxRuleUpdateDataAttributes) GetNotCountryCodeRegex() string`

GetNotCountryCodeRegex returns the NotCountryCodeRegex field if non-nil, zero value otherwise.

### GetNotCountryCodeRegexOk

`func (o *TaxRuleUpdateDataAttributes) GetNotCountryCodeRegexOk() (*string, bool)`

GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCountryCodeRegex

`func (o *TaxRuleUpdateDataAttributes) SetNotCountryCodeRegex(v string)`

SetNotCountryCodeRegex sets NotCountryCodeRegex field to given value.

### HasNotCountryCodeRegex

`func (o *TaxRuleUpdateDataAttributes) HasNotCountryCodeRegex() bool`

HasNotCountryCodeRegex returns a boolean if a field has been set.

### GetStateCodeRegex

`func (o *TaxRuleUpdateDataAttributes) GetStateCodeRegex() string`

GetStateCodeRegex returns the StateCodeRegex field if non-nil, zero value otherwise.

### GetStateCodeRegexOk

`func (o *TaxRuleUpdateDataAttributes) GetStateCodeRegexOk() (*string, bool)`

GetStateCodeRegexOk returns a tuple with the StateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCodeRegex

`func (o *TaxRuleUpdateDataAttributes) SetStateCodeRegex(v string)`

SetStateCodeRegex sets StateCodeRegex field to given value.

### HasStateCodeRegex

`func (o *TaxRuleUpdateDataAttributes) HasStateCodeRegex() bool`

HasStateCodeRegex returns a boolean if a field has been set.

### GetNotStateCodeRegex

`func (o *TaxRuleUpdateDataAttributes) GetNotStateCodeRegex() string`

GetNotStateCodeRegex returns the NotStateCodeRegex field if non-nil, zero value otherwise.

### GetNotStateCodeRegexOk

`func (o *TaxRuleUpdateDataAttributes) GetNotStateCodeRegexOk() (*string, bool)`

GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStateCodeRegex

`func (o *TaxRuleUpdateDataAttributes) SetNotStateCodeRegex(v string)`

SetNotStateCodeRegex sets NotStateCodeRegex field to given value.

### HasNotStateCodeRegex

`func (o *TaxRuleUpdateDataAttributes) HasNotStateCodeRegex() bool`

HasNotStateCodeRegex returns a boolean if a field has been set.

### GetZipCodeRegex

`func (o *TaxRuleUpdateDataAttributes) GetZipCodeRegex() string`

GetZipCodeRegex returns the ZipCodeRegex field if non-nil, zero value otherwise.

### GetZipCodeRegexOk

`func (o *TaxRuleUpdateDataAttributes) GetZipCodeRegexOk() (*string, bool)`

GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeRegex

`func (o *TaxRuleUpdateDataAttributes) SetZipCodeRegex(v string)`

SetZipCodeRegex sets ZipCodeRegex field to given value.

### HasZipCodeRegex

`func (o *TaxRuleUpdateDataAttributes) HasZipCodeRegex() bool`

HasZipCodeRegex returns a boolean if a field has been set.

### GetNotZipCodeRegex

`func (o *TaxRuleUpdateDataAttributes) GetNotZipCodeRegex() string`

GetNotZipCodeRegex returns the NotZipCodeRegex field if non-nil, zero value otherwise.

### GetNotZipCodeRegexOk

`func (o *TaxRuleUpdateDataAttributes) GetNotZipCodeRegexOk() (*string, bool)`

GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotZipCodeRegex

`func (o *TaxRuleUpdateDataAttributes) SetNotZipCodeRegex(v string)`

SetNotZipCodeRegex sets NotZipCodeRegex field to given value.

### HasNotZipCodeRegex

`func (o *TaxRuleUpdateDataAttributes) HasNotZipCodeRegex() bool`

HasNotZipCodeRegex returns a boolean if a field has been set.

### GetFreightTaxable

`func (o *TaxRuleUpdateDataAttributes) GetFreightTaxable() bool`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *TaxRuleUpdateDataAttributes) GetFreightTaxableOk() (*bool, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *TaxRuleUpdateDataAttributes) SetFreightTaxable(v bool)`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *TaxRuleUpdateDataAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### GetPaymentMethodTaxable

`func (o *TaxRuleUpdateDataAttributes) GetPaymentMethodTaxable() bool`

GetPaymentMethodTaxable returns the PaymentMethodTaxable field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableOk

`func (o *TaxRuleUpdateDataAttributes) GetPaymentMethodTaxableOk() (*bool, bool)`

GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxable

`func (o *TaxRuleUpdateDataAttributes) SetPaymentMethodTaxable(v bool)`

SetPaymentMethodTaxable sets PaymentMethodTaxable field to given value.

### HasPaymentMethodTaxable

`func (o *TaxRuleUpdateDataAttributes) HasPaymentMethodTaxable() bool`

HasPaymentMethodTaxable returns a boolean if a field has been set.

### GetGiftCardTaxable

`func (o *TaxRuleUpdateDataAttributes) GetGiftCardTaxable() bool`

GetGiftCardTaxable returns the GiftCardTaxable field if non-nil, zero value otherwise.

### GetGiftCardTaxableOk

`func (o *TaxRuleUpdateDataAttributes) GetGiftCardTaxableOk() (*bool, bool)`

GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardTaxable

`func (o *TaxRuleUpdateDataAttributes) SetGiftCardTaxable(v bool)`

SetGiftCardTaxable sets GiftCardTaxable field to given value.

### HasGiftCardTaxable

`func (o *TaxRuleUpdateDataAttributes) HasGiftCardTaxable() bool`

HasGiftCardTaxable returns a boolean if a field has been set.

### GetAdjustmentTaxable

`func (o *TaxRuleUpdateDataAttributes) GetAdjustmentTaxable() bool`

GetAdjustmentTaxable returns the AdjustmentTaxable field if non-nil, zero value otherwise.

### GetAdjustmentTaxableOk

`func (o *TaxRuleUpdateDataAttributes) GetAdjustmentTaxableOk() (*bool, bool)`

GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxable

`func (o *TaxRuleUpdateDataAttributes) SetAdjustmentTaxable(v bool)`

SetAdjustmentTaxable sets AdjustmentTaxable field to given value.

### HasAdjustmentTaxable

`func (o *TaxRuleUpdateDataAttributes) HasAdjustmentTaxable() bool`

HasAdjustmentTaxable returns a boolean if a field has been set.

### GetReference

`func (o *TaxRuleUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TaxRuleUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TaxRuleUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TaxRuleUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *TaxRuleUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *TaxRuleUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *TaxRuleUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *TaxRuleUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *TaxRuleUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TaxRuleUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TaxRuleUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TaxRuleUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


