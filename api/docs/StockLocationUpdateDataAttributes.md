# StockLocationUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The stock location&#39;s internal name. | [optional] 
**LabelFormat** | Pointer to **string** | The shipping label format for this stock location. Can be one of &#39;PDF&#39;, &#39;ZPL&#39;, &#39;EPL2&#39;, or &#39;PNG&#39; | [optional] 
**SuppressEtd** | Pointer to **bool** | Flag it if you want to skip the electronic invoice creation when generating the customs info for this stock location shipments. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewStockLocationUpdateDataAttributes

`func NewStockLocationUpdateDataAttributes() *StockLocationUpdateDataAttributes`

NewStockLocationUpdateDataAttributes instantiates a new StockLocationUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLocationUpdateDataAttributesWithDefaults

`func NewStockLocationUpdateDataAttributesWithDefaults() *StockLocationUpdateDataAttributes`

NewStockLocationUpdateDataAttributesWithDefaults instantiates a new StockLocationUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StockLocationUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StockLocationUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StockLocationUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StockLocationUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabelFormat

`func (o *StockLocationUpdateDataAttributes) GetLabelFormat() string`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *StockLocationUpdateDataAttributes) GetLabelFormatOk() (*string, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *StockLocationUpdateDataAttributes) SetLabelFormat(v string)`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *StockLocationUpdateDataAttributes) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### GetSuppressEtd

`func (o *StockLocationUpdateDataAttributes) GetSuppressEtd() bool`

GetSuppressEtd returns the SuppressEtd field if non-nil, zero value otherwise.

### GetSuppressEtdOk

`func (o *StockLocationUpdateDataAttributes) GetSuppressEtdOk() (*bool, bool)`

GetSuppressEtdOk returns a tuple with the SuppressEtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressEtd

`func (o *StockLocationUpdateDataAttributes) SetSuppressEtd(v bool)`

SetSuppressEtd sets SuppressEtd field to given value.

### HasSuppressEtd

`func (o *StockLocationUpdateDataAttributes) HasSuppressEtd() bool`

HasSuppressEtd returns a boolean if a field has been set.

### GetReference

`func (o *StockLocationUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StockLocationUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StockLocationUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *StockLocationUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *StockLocationUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *StockLocationUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *StockLocationUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *StockLocationUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *StockLocationUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StockLocationUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StockLocationUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StockLocationUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


