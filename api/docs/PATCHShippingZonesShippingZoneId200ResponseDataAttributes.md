# PATCHShippingZonesShippingZoneId200ResponseDataAttributes

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

### NewPATCHShippingZonesShippingZoneId200ResponseDataAttributes

`func NewPATCHShippingZonesShippingZoneId200ResponseDataAttributes() *PATCHShippingZonesShippingZoneId200ResponseDataAttributes`

NewPATCHShippingZonesShippingZoneId200ResponseDataAttributes instantiates a new PATCHShippingZonesShippingZoneId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHShippingZonesShippingZoneId200ResponseDataAttributesWithDefaults

`func NewPATCHShippingZonesShippingZoneId200ResponseDataAttributesWithDefaults() *PATCHShippingZonesShippingZoneId200ResponseDataAttributes`

NewPATCHShippingZonesShippingZoneId200ResponseDataAttributesWithDefaults instantiates a new PATCHShippingZonesShippingZoneId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountryCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetCountryCodeRegex() string`

GetCountryCodeRegex returns the CountryCodeRegex field if non-nil, zero value otherwise.

### GetCountryCodeRegexOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetCountryCodeRegexOk() (*string, bool)`

GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetCountryCodeRegex(v string)`

SetCountryCodeRegex sets CountryCodeRegex field to given value.

### HasCountryCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasCountryCodeRegex() bool`

HasCountryCodeRegex returns a boolean if a field has been set.

### GetNotCountryCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotCountryCodeRegex() string`

GetNotCountryCodeRegex returns the NotCountryCodeRegex field if non-nil, zero value otherwise.

### GetNotCountryCodeRegexOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotCountryCodeRegexOk() (*string, bool)`

GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCountryCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetNotCountryCodeRegex(v string)`

SetNotCountryCodeRegex sets NotCountryCodeRegex field to given value.

### HasNotCountryCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasNotCountryCodeRegex() bool`

HasNotCountryCodeRegex returns a boolean if a field has been set.

### GetStateCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetStateCodeRegex() string`

GetStateCodeRegex returns the StateCodeRegex field if non-nil, zero value otherwise.

### GetStateCodeRegexOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetStateCodeRegexOk() (*string, bool)`

GetStateCodeRegexOk returns a tuple with the StateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetStateCodeRegex(v string)`

SetStateCodeRegex sets StateCodeRegex field to given value.

### HasStateCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasStateCodeRegex() bool`

HasStateCodeRegex returns a boolean if a field has been set.

### GetNotStateCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotStateCodeRegex() string`

GetNotStateCodeRegex returns the NotStateCodeRegex field if non-nil, zero value otherwise.

### GetNotStateCodeRegexOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotStateCodeRegexOk() (*string, bool)`

GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStateCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetNotStateCodeRegex(v string)`

SetNotStateCodeRegex sets NotStateCodeRegex field to given value.

### HasNotStateCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasNotStateCodeRegex() bool`

HasNotStateCodeRegex returns a boolean if a field has been set.

### GetZipCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetZipCodeRegex() string`

GetZipCodeRegex returns the ZipCodeRegex field if non-nil, zero value otherwise.

### GetZipCodeRegexOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetZipCodeRegexOk() (*string, bool)`

GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetZipCodeRegex(v string)`

SetZipCodeRegex sets ZipCodeRegex field to given value.

### HasZipCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasZipCodeRegex() bool`

HasZipCodeRegex returns a boolean if a field has been set.

### GetNotZipCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotZipCodeRegex() string`

GetNotZipCodeRegex returns the NotZipCodeRegex field if non-nil, zero value otherwise.

### GetNotZipCodeRegexOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotZipCodeRegexOk() (*string, bool)`

GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotZipCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetNotZipCodeRegex(v string)`

SetNotZipCodeRegex sets NotZipCodeRegex field to given value.

### HasNotZipCodeRegex

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasNotZipCodeRegex() bool`

HasNotZipCodeRegex returns a boolean if a field has been set.

### GetReference

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


