# POSTPriceLists201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The price list&#39;s internal name | 
**CurrencyCode** | **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | 
**TaxIncluded** | Pointer to **bool** | Indicates if the associated prices include taxes. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTPriceLists201ResponseDataAttributes

`func NewPOSTPriceLists201ResponseDataAttributes(name string, currencyCode string, ) *POSTPriceLists201ResponseDataAttributes`

NewPOSTPriceLists201ResponseDataAttributes instantiates a new POSTPriceLists201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPriceLists201ResponseDataAttributesWithDefaults

`func NewPOSTPriceLists201ResponseDataAttributesWithDefaults() *POSTPriceLists201ResponseDataAttributes`

NewPOSTPriceLists201ResponseDataAttributesWithDefaults instantiates a new POSTPriceLists201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTPriceLists201ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTPriceLists201ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTPriceLists201ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCurrencyCode

`func (o *POSTPriceLists201ResponseDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTPriceLists201ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTPriceLists201ResponseDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetTaxIncluded

`func (o *POSTPriceLists201ResponseDataAttributes) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *POSTPriceLists201ResponseDataAttributes) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *POSTPriceLists201ResponseDataAttributes) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *POSTPriceLists201ResponseDataAttributes) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### GetReference

`func (o *POSTPriceLists201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTPriceLists201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTPriceLists201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTPriceLists201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTPriceLists201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTPriceLists201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTPriceLists201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTPriceLists201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTPriceLists201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTPriceLists201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTPriceLists201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTPriceLists201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


