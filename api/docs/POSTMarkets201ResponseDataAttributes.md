# POSTMarkets201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The market&#39;s internal name. | 
**Code** | Pointer to **interface{}** | A string that you can use to identify the market (must be unique within the environment). | [optional] 
**FacebookPixelId** | Pointer to **interface{}** | The Facebook Pixed ID. | [optional] 
**CheckoutUrl** | Pointer to **interface{}** | The checkout URL for this market. | [optional] 
**ExternalPricesUrl** | Pointer to **interface{}** | The URL used to overwrite prices by an external source. | [optional] 
**ExternalOrderValidationUrl** | Pointer to **interface{}** | The URL used to validate orders by an external source. | [optional] 
**ShippingCostCutoff** | Pointer to **interface{}** | When specified indicates the maximum number of shipping line items with cost that will be added to an order. | [optional] 
**Disable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as disabled. | [optional] 
**Enable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as enabled. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTMarkets201ResponseDataAttributes

`func NewPOSTMarkets201ResponseDataAttributes(name interface{}, ) *POSTMarkets201ResponseDataAttributes`

NewPOSTMarkets201ResponseDataAttributes instantiates a new POSTMarkets201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTMarkets201ResponseDataAttributesWithDefaults

`func NewPOSTMarkets201ResponseDataAttributesWithDefaults() *POSTMarkets201ResponseDataAttributes`

NewPOSTMarkets201ResponseDataAttributesWithDefaults instantiates a new POSTMarkets201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTMarkets201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTMarkets201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTMarkets201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTMarkets201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTMarkets201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *POSTMarkets201ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *POSTMarkets201ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *POSTMarkets201ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *POSTMarkets201ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *POSTMarkets201ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *POSTMarkets201ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetFacebookPixelId

`func (o *POSTMarkets201ResponseDataAttributes) GetFacebookPixelId() interface{}`

GetFacebookPixelId returns the FacebookPixelId field if non-nil, zero value otherwise.

### GetFacebookPixelIdOk

`func (o *POSTMarkets201ResponseDataAttributes) GetFacebookPixelIdOk() (*interface{}, bool)`

GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookPixelId

`func (o *POSTMarkets201ResponseDataAttributes) SetFacebookPixelId(v interface{})`

SetFacebookPixelId sets FacebookPixelId field to given value.

### HasFacebookPixelId

`func (o *POSTMarkets201ResponseDataAttributes) HasFacebookPixelId() bool`

HasFacebookPixelId returns a boolean if a field has been set.

### SetFacebookPixelIdNil

`func (o *POSTMarkets201ResponseDataAttributes) SetFacebookPixelIdNil(b bool)`

 SetFacebookPixelIdNil sets the value for FacebookPixelId to be an explicit nil

### UnsetFacebookPixelId
`func (o *POSTMarkets201ResponseDataAttributes) UnsetFacebookPixelId()`

UnsetFacebookPixelId ensures that no value is present for FacebookPixelId, not even an explicit nil
### GetCheckoutUrl

`func (o *POSTMarkets201ResponseDataAttributes) GetCheckoutUrl() interface{}`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *POSTMarkets201ResponseDataAttributes) GetCheckoutUrlOk() (*interface{}, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *POSTMarkets201ResponseDataAttributes) SetCheckoutUrl(v interface{})`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *POSTMarkets201ResponseDataAttributes) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### SetCheckoutUrlNil

`func (o *POSTMarkets201ResponseDataAttributes) SetCheckoutUrlNil(b bool)`

 SetCheckoutUrlNil sets the value for CheckoutUrl to be an explicit nil

### UnsetCheckoutUrl
`func (o *POSTMarkets201ResponseDataAttributes) UnsetCheckoutUrl()`

UnsetCheckoutUrl ensures that no value is present for CheckoutUrl, not even an explicit nil
### GetExternalPricesUrl

`func (o *POSTMarkets201ResponseDataAttributes) GetExternalPricesUrl() interface{}`

GetExternalPricesUrl returns the ExternalPricesUrl field if non-nil, zero value otherwise.

### GetExternalPricesUrlOk

`func (o *POSTMarkets201ResponseDataAttributes) GetExternalPricesUrlOk() (*interface{}, bool)`

GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPricesUrl

`func (o *POSTMarkets201ResponseDataAttributes) SetExternalPricesUrl(v interface{})`

SetExternalPricesUrl sets ExternalPricesUrl field to given value.

### HasExternalPricesUrl

`func (o *POSTMarkets201ResponseDataAttributes) HasExternalPricesUrl() bool`

HasExternalPricesUrl returns a boolean if a field has been set.

### SetExternalPricesUrlNil

`func (o *POSTMarkets201ResponseDataAttributes) SetExternalPricesUrlNil(b bool)`

 SetExternalPricesUrlNil sets the value for ExternalPricesUrl to be an explicit nil

### UnsetExternalPricesUrl
`func (o *POSTMarkets201ResponseDataAttributes) UnsetExternalPricesUrl()`

UnsetExternalPricesUrl ensures that no value is present for ExternalPricesUrl, not even an explicit nil
### GetExternalOrderValidationUrl

`func (o *POSTMarkets201ResponseDataAttributes) GetExternalOrderValidationUrl() interface{}`

GetExternalOrderValidationUrl returns the ExternalOrderValidationUrl field if non-nil, zero value otherwise.

### GetExternalOrderValidationUrlOk

`func (o *POSTMarkets201ResponseDataAttributes) GetExternalOrderValidationUrlOk() (*interface{}, bool)`

GetExternalOrderValidationUrlOk returns a tuple with the ExternalOrderValidationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderValidationUrl

`func (o *POSTMarkets201ResponseDataAttributes) SetExternalOrderValidationUrl(v interface{})`

SetExternalOrderValidationUrl sets ExternalOrderValidationUrl field to given value.

### HasExternalOrderValidationUrl

`func (o *POSTMarkets201ResponseDataAttributes) HasExternalOrderValidationUrl() bool`

HasExternalOrderValidationUrl returns a boolean if a field has been set.

### SetExternalOrderValidationUrlNil

`func (o *POSTMarkets201ResponseDataAttributes) SetExternalOrderValidationUrlNil(b bool)`

 SetExternalOrderValidationUrlNil sets the value for ExternalOrderValidationUrl to be an explicit nil

### UnsetExternalOrderValidationUrl
`func (o *POSTMarkets201ResponseDataAttributes) UnsetExternalOrderValidationUrl()`

UnsetExternalOrderValidationUrl ensures that no value is present for ExternalOrderValidationUrl, not even an explicit nil
### GetShippingCostCutoff

`func (o *POSTMarkets201ResponseDataAttributes) GetShippingCostCutoff() interface{}`

GetShippingCostCutoff returns the ShippingCostCutoff field if non-nil, zero value otherwise.

### GetShippingCostCutoffOk

`func (o *POSTMarkets201ResponseDataAttributes) GetShippingCostCutoffOk() (*interface{}, bool)`

GetShippingCostCutoffOk returns a tuple with the ShippingCostCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostCutoff

`func (o *POSTMarkets201ResponseDataAttributes) SetShippingCostCutoff(v interface{})`

SetShippingCostCutoff sets ShippingCostCutoff field to given value.

### HasShippingCostCutoff

`func (o *POSTMarkets201ResponseDataAttributes) HasShippingCostCutoff() bool`

HasShippingCostCutoff returns a boolean if a field has been set.

### SetShippingCostCutoffNil

`func (o *POSTMarkets201ResponseDataAttributes) SetShippingCostCutoffNil(b bool)`

 SetShippingCostCutoffNil sets the value for ShippingCostCutoff to be an explicit nil

### UnsetShippingCostCutoff
`func (o *POSTMarkets201ResponseDataAttributes) UnsetShippingCostCutoff()`

UnsetShippingCostCutoff ensures that no value is present for ShippingCostCutoff, not even an explicit nil
### GetDisable

`func (o *POSTMarkets201ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *POSTMarkets201ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *POSTMarkets201ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *POSTMarkets201ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *POSTMarkets201ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *POSTMarkets201ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *POSTMarkets201ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *POSTMarkets201ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *POSTMarkets201ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *POSTMarkets201ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *POSTMarkets201ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *POSTMarkets201ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetReference

`func (o *POSTMarkets201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTMarkets201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTMarkets201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTMarkets201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTMarkets201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTMarkets201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTMarkets201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTMarkets201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTMarkets201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTMarkets201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTMarkets201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTMarkets201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTMarkets201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTMarkets201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTMarkets201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTMarkets201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTMarkets201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTMarkets201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


