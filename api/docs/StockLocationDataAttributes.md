# StockLocationDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Unique identifier for the stock location (numeric) | [optional] 
**Name** | Pointer to **string** | The stock location&#39;s internal name. | [optional] 
**LabelFormat** | Pointer to **string** | The shipping label format for this stock location. Can be one of &#39;PDF&#39;, &#39;ZPL&#39;, &#39;EPL2&#39;, or &#39;PNG&#39; | [optional] 
**SuppressEtd** | Pointer to **bool** | Flag it if you want to skip the electronic invoice creation when generating the customs info for this stock location shipments. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewStockLocationDataAttributes

`func NewStockLocationDataAttributes() *StockLocationDataAttributes`

NewStockLocationDataAttributes instantiates a new StockLocationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLocationDataAttributesWithDefaults

`func NewStockLocationDataAttributesWithDefaults() *StockLocationDataAttributes`

NewStockLocationDataAttributesWithDefaults instantiates a new StockLocationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *StockLocationDataAttributes) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *StockLocationDataAttributes) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *StockLocationDataAttributes) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *StockLocationDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetName

`func (o *StockLocationDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StockLocationDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StockLocationDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StockLocationDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabelFormat

`func (o *StockLocationDataAttributes) GetLabelFormat() string`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *StockLocationDataAttributes) GetLabelFormatOk() (*string, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *StockLocationDataAttributes) SetLabelFormat(v string)`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *StockLocationDataAttributes) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### GetSuppressEtd

`func (o *StockLocationDataAttributes) GetSuppressEtd() bool`

GetSuppressEtd returns the SuppressEtd field if non-nil, zero value otherwise.

### GetSuppressEtdOk

`func (o *StockLocationDataAttributes) GetSuppressEtdOk() (*bool, bool)`

GetSuppressEtdOk returns a tuple with the SuppressEtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressEtd

`func (o *StockLocationDataAttributes) SetSuppressEtd(v bool)`

SetSuppressEtd sets SuppressEtd field to given value.

### HasSuppressEtd

`func (o *StockLocationDataAttributes) HasSuppressEtd() bool`

HasSuppressEtd returns a boolean if a field has been set.

### GetId

`func (o *StockLocationDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StockLocationDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StockLocationDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StockLocationDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StockLocationDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StockLocationDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StockLocationDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StockLocationDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StockLocationDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StockLocationDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StockLocationDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StockLocationDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *StockLocationDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StockLocationDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StockLocationDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *StockLocationDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *StockLocationDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *StockLocationDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *StockLocationDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *StockLocationDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *StockLocationDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StockLocationDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StockLocationDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StockLocationDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


