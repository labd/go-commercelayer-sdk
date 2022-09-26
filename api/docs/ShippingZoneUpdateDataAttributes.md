# ShippingZoneUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The shipping zone&#39;s internal name. | [optional] 
**CountryCodeRegex** | Pointer to **string** | The regex that will be evaluated to match the shipping address country code. | [optional] 
**NotCountryCodeRegex** | Pointer to **string** | The regex that will be evaluated as negative match for the shipping address country code. | [optional] 
**StateCodeRegex** | Pointer to **string** | The regex that will be evaluated to match the shipping address state code. | [optional] 
**NotStateCodeRegex** | Pointer to **string** | The regex that will be evaluated as negative match for the shipping address state code. | [optional] 
**ZipCodeRegex** | Pointer to **string** | The regex that will be evaluated to match the shipping address zip code. | [optional] 
**NotZipCodeRegex** | Pointer to **string** | The regex that will be evaluated as negative match for the shipping zip country code. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewShippingZoneUpdateDataAttributes

`func NewShippingZoneUpdateDataAttributes() *ShippingZoneUpdateDataAttributes`

NewShippingZoneUpdateDataAttributes instantiates a new ShippingZoneUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingZoneUpdateDataAttributesWithDefaults

`func NewShippingZoneUpdateDataAttributesWithDefaults() *ShippingZoneUpdateDataAttributes`

NewShippingZoneUpdateDataAttributesWithDefaults instantiates a new ShippingZoneUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShippingZoneUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingZoneUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingZoneUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShippingZoneUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountryCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) GetCountryCodeRegex() string`

GetCountryCodeRegex returns the CountryCodeRegex field if non-nil, zero value otherwise.

### GetCountryCodeRegexOk

`func (o *ShippingZoneUpdateDataAttributes) GetCountryCodeRegexOk() (*string, bool)`

GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) SetCountryCodeRegex(v string)`

SetCountryCodeRegex sets CountryCodeRegex field to given value.

### HasCountryCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) HasCountryCodeRegex() bool`

HasCountryCodeRegex returns a boolean if a field has been set.

### GetNotCountryCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) GetNotCountryCodeRegex() string`

GetNotCountryCodeRegex returns the NotCountryCodeRegex field if non-nil, zero value otherwise.

### GetNotCountryCodeRegexOk

`func (o *ShippingZoneUpdateDataAttributes) GetNotCountryCodeRegexOk() (*string, bool)`

GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCountryCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) SetNotCountryCodeRegex(v string)`

SetNotCountryCodeRegex sets NotCountryCodeRegex field to given value.

### HasNotCountryCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) HasNotCountryCodeRegex() bool`

HasNotCountryCodeRegex returns a boolean if a field has been set.

### GetStateCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) GetStateCodeRegex() string`

GetStateCodeRegex returns the StateCodeRegex field if non-nil, zero value otherwise.

### GetStateCodeRegexOk

`func (o *ShippingZoneUpdateDataAttributes) GetStateCodeRegexOk() (*string, bool)`

GetStateCodeRegexOk returns a tuple with the StateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) SetStateCodeRegex(v string)`

SetStateCodeRegex sets StateCodeRegex field to given value.

### HasStateCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) HasStateCodeRegex() bool`

HasStateCodeRegex returns a boolean if a field has been set.

### GetNotStateCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) GetNotStateCodeRegex() string`

GetNotStateCodeRegex returns the NotStateCodeRegex field if non-nil, zero value otherwise.

### GetNotStateCodeRegexOk

`func (o *ShippingZoneUpdateDataAttributes) GetNotStateCodeRegexOk() (*string, bool)`

GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStateCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) SetNotStateCodeRegex(v string)`

SetNotStateCodeRegex sets NotStateCodeRegex field to given value.

### HasNotStateCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) HasNotStateCodeRegex() bool`

HasNotStateCodeRegex returns a boolean if a field has been set.

### GetZipCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) GetZipCodeRegex() string`

GetZipCodeRegex returns the ZipCodeRegex field if non-nil, zero value otherwise.

### GetZipCodeRegexOk

`func (o *ShippingZoneUpdateDataAttributes) GetZipCodeRegexOk() (*string, bool)`

GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) SetZipCodeRegex(v string)`

SetZipCodeRegex sets ZipCodeRegex field to given value.

### HasZipCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) HasZipCodeRegex() bool`

HasZipCodeRegex returns a boolean if a field has been set.

### GetNotZipCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) GetNotZipCodeRegex() string`

GetNotZipCodeRegex returns the NotZipCodeRegex field if non-nil, zero value otherwise.

### GetNotZipCodeRegexOk

`func (o *ShippingZoneUpdateDataAttributes) GetNotZipCodeRegexOk() (*string, bool)`

GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotZipCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) SetNotZipCodeRegex(v string)`

SetNotZipCodeRegex sets NotZipCodeRegex field to given value.

### HasNotZipCodeRegex

`func (o *ShippingZoneUpdateDataAttributes) HasNotZipCodeRegex() bool`

HasNotZipCodeRegex returns a boolean if a field has been set.

### GetReference

`func (o *ShippingZoneUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ShippingZoneUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ShippingZoneUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ShippingZoneUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ShippingZoneUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ShippingZoneUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ShippingZoneUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ShippingZoneUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ShippingZoneUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShippingZoneUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShippingZoneUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShippingZoneUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


