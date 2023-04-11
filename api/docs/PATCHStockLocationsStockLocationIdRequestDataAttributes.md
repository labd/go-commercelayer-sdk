# PATCHStockLocationsStockLocationIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The stock location&#39;s internal name. | [optional] 
**LabelFormat** | Pointer to **interface{}** | The shipping label format for this stock location. Can be one of &#39;PDF&#39;, &#39;ZPL&#39;, &#39;EPL2&#39;, or &#39;PNG&#39; | [optional] 
**SuppressEtd** | Pointer to **interface{}** | Flag it if you want to skip the electronic invoice creation when generating the customs info for this stock location shipments. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHStockLocationsStockLocationIdRequestDataAttributes

`func NewPATCHStockLocationsStockLocationIdRequestDataAttributes() *PATCHStockLocationsStockLocationIdRequestDataAttributes`

NewPATCHStockLocationsStockLocationIdRequestDataAttributes instantiates a new PATCHStockLocationsStockLocationIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStockLocationsStockLocationIdRequestDataAttributesWithDefaults

`func NewPATCHStockLocationsStockLocationIdRequestDataAttributesWithDefaults() *PATCHStockLocationsStockLocationIdRequestDataAttributes`

NewPATCHStockLocationsStockLocationIdRequestDataAttributesWithDefaults instantiates a new PATCHStockLocationsStockLocationIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLabelFormat

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetLabelFormat() interface{}`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetLabelFormatOk() (*interface{}, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetLabelFormat(v interface{})`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### SetLabelFormatNil

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetLabelFormatNil(b bool)`

 SetLabelFormatNil sets the value for LabelFormat to be an explicit nil

### UnsetLabelFormat
`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) UnsetLabelFormat()`

UnsetLabelFormat ensures that no value is present for LabelFormat, not even an explicit nil
### GetSuppressEtd

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetSuppressEtd() interface{}`

GetSuppressEtd returns the SuppressEtd field if non-nil, zero value otherwise.

### GetSuppressEtdOk

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetSuppressEtdOk() (*interface{}, bool)`

GetSuppressEtdOk returns a tuple with the SuppressEtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressEtd

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetSuppressEtd(v interface{})`

SetSuppressEtd sets SuppressEtd field to given value.

### HasSuppressEtd

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) HasSuppressEtd() bool`

HasSuppressEtd returns a boolean if a field has been set.

### SetSuppressEtdNil

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetSuppressEtdNil(b bool)`

 SetSuppressEtdNil sets the value for SuppressEtd to be an explicit nil

### UnsetSuppressEtd
`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) UnsetSuppressEtd()`

UnsetSuppressEtd ensures that no value is present for SuppressEtd, not even an explicit nil
### GetReference

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHStockLocationsStockLocationIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


