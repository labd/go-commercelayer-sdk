# PATCHStockLocationsStockLocationId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The stock location&#39;s internal name. | [optional] 
**Code** | Pointer to **interface{}** | A string that you can use to identify the stock location (must be unique within the environment). | [optional] 
**LabelFormat** | Pointer to **interface{}** | The shipping label format for this stock location. Can be one of &#39;PDF&#39;, &#39;ZPL&#39;, &#39;EPL2&#39;, or &#39;PNG&#39;. | [optional] 
**SuppressEtd** | Pointer to **interface{}** | Flag it if you want to skip the electronic invoice creation when generating the customs info for this stock location shipments. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHStockLocationsStockLocationId200ResponseDataAttributes

`func NewPATCHStockLocationsStockLocationId200ResponseDataAttributes() *PATCHStockLocationsStockLocationId200ResponseDataAttributes`

NewPATCHStockLocationsStockLocationId200ResponseDataAttributes instantiates a new PATCHStockLocationsStockLocationId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStockLocationsStockLocationId200ResponseDataAttributesWithDefaults

`func NewPATCHStockLocationsStockLocationId200ResponseDataAttributesWithDefaults() *PATCHStockLocationsStockLocationId200ResponseDataAttributes`

NewPATCHStockLocationsStockLocationId200ResponseDataAttributesWithDefaults instantiates a new PATCHStockLocationsStockLocationId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLabelFormat

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetLabelFormat() interface{}`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetLabelFormatOk() (*interface{}, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetLabelFormat(v interface{})`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### SetLabelFormatNil

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetLabelFormatNil(b bool)`

 SetLabelFormatNil sets the value for LabelFormat to be an explicit nil

### UnsetLabelFormat
`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) UnsetLabelFormat()`

UnsetLabelFormat ensures that no value is present for LabelFormat, not even an explicit nil
### GetSuppressEtd

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetSuppressEtd() interface{}`

GetSuppressEtd returns the SuppressEtd field if non-nil, zero value otherwise.

### GetSuppressEtdOk

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetSuppressEtdOk() (*interface{}, bool)`

GetSuppressEtdOk returns a tuple with the SuppressEtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressEtd

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetSuppressEtd(v interface{})`

SetSuppressEtd sets SuppressEtd field to given value.

### HasSuppressEtd

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) HasSuppressEtd() bool`

HasSuppressEtd returns a boolean if a field has been set.

### SetSuppressEtdNil

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetSuppressEtdNil(b bool)`

 SetSuppressEtdNil sets the value for SuppressEtd to be an explicit nil

### UnsetSuppressEtd
`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) UnsetSuppressEtd()`

UnsetSuppressEtd ensures that no value is present for SuppressEtd, not even an explicit nil
### GetReference

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHStockLocationsStockLocationId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


