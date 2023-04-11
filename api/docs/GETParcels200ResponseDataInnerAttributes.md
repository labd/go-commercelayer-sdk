# GETParcels200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the parcel | [optional] 
**Weight** | Pointer to **interface{}** | The parcel weight, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**UnitOfWeight** | Pointer to **interface{}** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | [optional] 
**EelPfc** | Pointer to **interface{}** | When shipping outside the US, you need to provide either an Exemption and Exclusion Legend (EEL) code or a Proof of Filing Citation (PFC). Which you need is based on the value of the goods being shipped. Value can be one of \&quot;EEL\&quot; o \&quot;PFC\&quot;. | [optional] 
**ContentsType** | Pointer to **interface{}** | The type of item you are sending. Can be one of &#39;merchandise&#39;, &#39;gift&#39;, &#39;documents&#39;, &#39;returned_goods&#39;, &#39;sample&#39;, or &#39;other&#39;. | [optional] 
**ContentsExplanation** | Pointer to **interface{}** | If you specify &#39;other&#39; in the &#39;contents_type&#39; attribute, you must supply a brief description in this attribute. | [optional] 
**CustomsCertify** | Pointer to **interface{}** | Indicates if the provided information is accurate | [optional] 
**CustomsSigner** | Pointer to **interface{}** | This is the name of the person who is certifying that the information provided on the customs form is accurate. Use a name of the person in your organization who is responsible for this. | [optional] 
**NonDeliveryOption** | Pointer to **interface{}** | In case the shipment cannot be delivered, this option tells the carrier what you want to happen to the parcel. You can pass either &#39;return&#39;, or &#39;abandon&#39;. The value defaults to &#39;return&#39;. If you pass &#39;abandon&#39;, you will not receive the parcel back if it cannot be delivered. | [optional] 
**RestrictionType** | Pointer to **interface{}** | Describes if your parcel requires any special treatment or quarantine when entering the country. Can be one of &#39;none&#39;, &#39;other&#39;, &#39;quarantine&#39;, or &#39;sanitary_phytosanitary_inspection&#39;. | [optional] 
**RestrictionComments** | Pointer to **interface{}** | If you specify &#39;other&#39; in the restriction type, you must supply a brief description of what is required. | [optional] 
**CustomsInfoRequired** | Pointer to **interface{}** | Indicates if the parcel requires customs info to get the shipping rates. | [optional] 
**ShippingLabelUrl** | Pointer to **interface{}** | The shipping label url, ready to be downloaded and printed. | [optional] 
**ShippingLabelFileType** | Pointer to **interface{}** | The shipping label file type. One of &#39;application/pdf&#39;, &#39;application/zpl&#39;, &#39;application/epl2&#39;, or &#39;image/png&#39;. | [optional] 
**ShippingLabelSize** | Pointer to **interface{}** | The shipping label size. | [optional] 
**ShippingLabelResolution** | Pointer to **interface{}** | The shipping label resolution. | [optional] 
**TrackingNumber** | Pointer to **interface{}** | The tracking number associated to this parcel. | [optional] 
**TrackingStatus** | Pointer to **interface{}** | The tracking status for this parcel, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusDetail** | Pointer to **interface{}** | Additional information about the tracking status, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusUpdatedAt** | Pointer to **interface{}** | Time at which the parcel&#39;s tracking status was last updated. | [optional] 
**TrackingDetails** | Pointer to **interface{}** | The parcel&#39;s full tracking history, automatically updated in real time by the shipping carrier. | [optional] 
**CarrierWeightOz** | Pointer to **interface{}** | The weight of the parcel as measured by the carrier in ounces (if available) | [optional] 
**SignedBy** | Pointer to **interface{}** | The name of the person who signed for the parcel (if available) | [optional] 
**Incoterm** | Pointer to **interface{}** | The type of Incoterm (if available). | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETParcels200ResponseDataInnerAttributes

`func NewGETParcels200ResponseDataInnerAttributes() *GETParcels200ResponseDataInnerAttributes`

NewGETParcels200ResponseDataInnerAttributes instantiates a new GETParcels200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETParcels200ResponseDataInnerAttributesWithDefaults

`func NewGETParcels200ResponseDataInnerAttributesWithDefaults() *GETParcels200ResponseDataInnerAttributes`

NewGETParcels200ResponseDataInnerAttributesWithDefaults instantiates a new GETParcels200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETParcels200ResponseDataInnerAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETParcels200ResponseDataInnerAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETParcels200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetWeight

`func (o *GETParcels200ResponseDataInnerAttributes) GetWeight() interface{}`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetWeightOk() (*interface{}, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GETParcels200ResponseDataInnerAttributes) SetWeight(v interface{})`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GETParcels200ResponseDataInnerAttributes) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetUnitOfWeight

`func (o *GETParcels200ResponseDataInnerAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *GETParcels200ResponseDataInnerAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *GETParcels200ResponseDataInnerAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetEelPfc

`func (o *GETParcels200ResponseDataInnerAttributes) GetEelPfc() interface{}`

GetEelPfc returns the EelPfc field if non-nil, zero value otherwise.

### GetEelPfcOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetEelPfcOk() (*interface{}, bool)`

GetEelPfcOk returns a tuple with the EelPfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEelPfc

`func (o *GETParcels200ResponseDataInnerAttributes) SetEelPfc(v interface{})`

SetEelPfc sets EelPfc field to given value.

### HasEelPfc

`func (o *GETParcels200ResponseDataInnerAttributes) HasEelPfc() bool`

HasEelPfc returns a boolean if a field has been set.

### SetEelPfcNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetEelPfcNil(b bool)`

 SetEelPfcNil sets the value for EelPfc to be an explicit nil

### UnsetEelPfc
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetEelPfc()`

UnsetEelPfc ensures that no value is present for EelPfc, not even an explicit nil
### GetContentsType

`func (o *GETParcels200ResponseDataInnerAttributes) GetContentsType() interface{}`

GetContentsType returns the ContentsType field if non-nil, zero value otherwise.

### GetContentsTypeOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetContentsTypeOk() (*interface{}, bool)`

GetContentsTypeOk returns a tuple with the ContentsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsType

`func (o *GETParcels200ResponseDataInnerAttributes) SetContentsType(v interface{})`

SetContentsType sets ContentsType field to given value.

### HasContentsType

`func (o *GETParcels200ResponseDataInnerAttributes) HasContentsType() bool`

HasContentsType returns a boolean if a field has been set.

### SetContentsTypeNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetContentsTypeNil(b bool)`

 SetContentsTypeNil sets the value for ContentsType to be an explicit nil

### UnsetContentsType
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetContentsType()`

UnsetContentsType ensures that no value is present for ContentsType, not even an explicit nil
### GetContentsExplanation

`func (o *GETParcels200ResponseDataInnerAttributes) GetContentsExplanation() interface{}`

GetContentsExplanation returns the ContentsExplanation field if non-nil, zero value otherwise.

### GetContentsExplanationOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetContentsExplanationOk() (*interface{}, bool)`

GetContentsExplanationOk returns a tuple with the ContentsExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsExplanation

`func (o *GETParcels200ResponseDataInnerAttributes) SetContentsExplanation(v interface{})`

SetContentsExplanation sets ContentsExplanation field to given value.

### HasContentsExplanation

`func (o *GETParcels200ResponseDataInnerAttributes) HasContentsExplanation() bool`

HasContentsExplanation returns a boolean if a field has been set.

### SetContentsExplanationNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetContentsExplanationNil(b bool)`

 SetContentsExplanationNil sets the value for ContentsExplanation to be an explicit nil

### UnsetContentsExplanation
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetContentsExplanation()`

UnsetContentsExplanation ensures that no value is present for ContentsExplanation, not even an explicit nil
### GetCustomsCertify

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsCertify() interface{}`

GetCustomsCertify returns the CustomsCertify field if non-nil, zero value otherwise.

### GetCustomsCertifyOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsCertifyOk() (*interface{}, bool)`

GetCustomsCertifyOk returns a tuple with the CustomsCertify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsCertify

`func (o *GETParcels200ResponseDataInnerAttributes) SetCustomsCertify(v interface{})`

SetCustomsCertify sets CustomsCertify field to given value.

### HasCustomsCertify

`func (o *GETParcels200ResponseDataInnerAttributes) HasCustomsCertify() bool`

HasCustomsCertify returns a boolean if a field has been set.

### SetCustomsCertifyNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetCustomsCertifyNil(b bool)`

 SetCustomsCertifyNil sets the value for CustomsCertify to be an explicit nil

### UnsetCustomsCertify
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetCustomsCertify()`

UnsetCustomsCertify ensures that no value is present for CustomsCertify, not even an explicit nil
### GetCustomsSigner

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsSigner() interface{}`

GetCustomsSigner returns the CustomsSigner field if non-nil, zero value otherwise.

### GetCustomsSignerOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsSignerOk() (*interface{}, bool)`

GetCustomsSignerOk returns a tuple with the CustomsSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsSigner

`func (o *GETParcels200ResponseDataInnerAttributes) SetCustomsSigner(v interface{})`

SetCustomsSigner sets CustomsSigner field to given value.

### HasCustomsSigner

`func (o *GETParcels200ResponseDataInnerAttributes) HasCustomsSigner() bool`

HasCustomsSigner returns a boolean if a field has been set.

### SetCustomsSignerNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetCustomsSignerNil(b bool)`

 SetCustomsSignerNil sets the value for CustomsSigner to be an explicit nil

### UnsetCustomsSigner
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetCustomsSigner()`

UnsetCustomsSigner ensures that no value is present for CustomsSigner, not even an explicit nil
### GetNonDeliveryOption

`func (o *GETParcels200ResponseDataInnerAttributes) GetNonDeliveryOption() interface{}`

GetNonDeliveryOption returns the NonDeliveryOption field if non-nil, zero value otherwise.

### GetNonDeliveryOptionOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetNonDeliveryOptionOk() (*interface{}, bool)`

GetNonDeliveryOptionOk returns a tuple with the NonDeliveryOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDeliveryOption

`func (o *GETParcels200ResponseDataInnerAttributes) SetNonDeliveryOption(v interface{})`

SetNonDeliveryOption sets NonDeliveryOption field to given value.

### HasNonDeliveryOption

`func (o *GETParcels200ResponseDataInnerAttributes) HasNonDeliveryOption() bool`

HasNonDeliveryOption returns a boolean if a field has been set.

### SetNonDeliveryOptionNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetNonDeliveryOptionNil(b bool)`

 SetNonDeliveryOptionNil sets the value for NonDeliveryOption to be an explicit nil

### UnsetNonDeliveryOption
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetNonDeliveryOption()`

UnsetNonDeliveryOption ensures that no value is present for NonDeliveryOption, not even an explicit nil
### GetRestrictionType

`func (o *GETParcels200ResponseDataInnerAttributes) GetRestrictionType() interface{}`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetRestrictionTypeOk() (*interface{}, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *GETParcels200ResponseDataInnerAttributes) SetRestrictionType(v interface{})`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *GETParcels200ResponseDataInnerAttributes) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.

### SetRestrictionTypeNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetRestrictionTypeNil(b bool)`

 SetRestrictionTypeNil sets the value for RestrictionType to be an explicit nil

### UnsetRestrictionType
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetRestrictionType()`

UnsetRestrictionType ensures that no value is present for RestrictionType, not even an explicit nil
### GetRestrictionComments

`func (o *GETParcels200ResponseDataInnerAttributes) GetRestrictionComments() interface{}`

GetRestrictionComments returns the RestrictionComments field if non-nil, zero value otherwise.

### GetRestrictionCommentsOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetRestrictionCommentsOk() (*interface{}, bool)`

GetRestrictionCommentsOk returns a tuple with the RestrictionComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionComments

`func (o *GETParcels200ResponseDataInnerAttributes) SetRestrictionComments(v interface{})`

SetRestrictionComments sets RestrictionComments field to given value.

### HasRestrictionComments

`func (o *GETParcels200ResponseDataInnerAttributes) HasRestrictionComments() bool`

HasRestrictionComments returns a boolean if a field has been set.

### SetRestrictionCommentsNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetRestrictionCommentsNil(b bool)`

 SetRestrictionCommentsNil sets the value for RestrictionComments to be an explicit nil

### UnsetRestrictionComments
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetRestrictionComments()`

UnsetRestrictionComments ensures that no value is present for RestrictionComments, not even an explicit nil
### GetCustomsInfoRequired

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsInfoRequired() interface{}`

GetCustomsInfoRequired returns the CustomsInfoRequired field if non-nil, zero value otherwise.

### GetCustomsInfoRequiredOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCustomsInfoRequiredOk() (*interface{}, bool)`

GetCustomsInfoRequiredOk returns a tuple with the CustomsInfoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsInfoRequired

`func (o *GETParcels200ResponseDataInnerAttributes) SetCustomsInfoRequired(v interface{})`

SetCustomsInfoRequired sets CustomsInfoRequired field to given value.

### HasCustomsInfoRequired

`func (o *GETParcels200ResponseDataInnerAttributes) HasCustomsInfoRequired() bool`

HasCustomsInfoRequired returns a boolean if a field has been set.

### SetCustomsInfoRequiredNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetCustomsInfoRequiredNil(b bool)`

 SetCustomsInfoRequiredNil sets the value for CustomsInfoRequired to be an explicit nil

### UnsetCustomsInfoRequired
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetCustomsInfoRequired()`

UnsetCustomsInfoRequired ensures that no value is present for CustomsInfoRequired, not even an explicit nil
### GetShippingLabelUrl

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelUrl() interface{}`

GetShippingLabelUrl returns the ShippingLabelUrl field if non-nil, zero value otherwise.

### GetShippingLabelUrlOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelUrlOk() (*interface{}, bool)`

GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelUrl

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelUrl(v interface{})`

SetShippingLabelUrl sets ShippingLabelUrl field to given value.

### HasShippingLabelUrl

`func (o *GETParcels200ResponseDataInnerAttributes) HasShippingLabelUrl() bool`

HasShippingLabelUrl returns a boolean if a field has been set.

### SetShippingLabelUrlNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelUrlNil(b bool)`

 SetShippingLabelUrlNil sets the value for ShippingLabelUrl to be an explicit nil

### UnsetShippingLabelUrl
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetShippingLabelUrl()`

UnsetShippingLabelUrl ensures that no value is present for ShippingLabelUrl, not even an explicit nil
### GetShippingLabelFileType

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelFileType() interface{}`

GetShippingLabelFileType returns the ShippingLabelFileType field if non-nil, zero value otherwise.

### GetShippingLabelFileTypeOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelFileTypeOk() (*interface{}, bool)`

GetShippingLabelFileTypeOk returns a tuple with the ShippingLabelFileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelFileType

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelFileType(v interface{})`

SetShippingLabelFileType sets ShippingLabelFileType field to given value.

### HasShippingLabelFileType

`func (o *GETParcels200ResponseDataInnerAttributes) HasShippingLabelFileType() bool`

HasShippingLabelFileType returns a boolean if a field has been set.

### SetShippingLabelFileTypeNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelFileTypeNil(b bool)`

 SetShippingLabelFileTypeNil sets the value for ShippingLabelFileType to be an explicit nil

### UnsetShippingLabelFileType
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetShippingLabelFileType()`

UnsetShippingLabelFileType ensures that no value is present for ShippingLabelFileType, not even an explicit nil
### GetShippingLabelSize

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelSize() interface{}`

GetShippingLabelSize returns the ShippingLabelSize field if non-nil, zero value otherwise.

### GetShippingLabelSizeOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelSizeOk() (*interface{}, bool)`

GetShippingLabelSizeOk returns a tuple with the ShippingLabelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelSize

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelSize(v interface{})`

SetShippingLabelSize sets ShippingLabelSize field to given value.

### HasShippingLabelSize

`func (o *GETParcels200ResponseDataInnerAttributes) HasShippingLabelSize() bool`

HasShippingLabelSize returns a boolean if a field has been set.

### SetShippingLabelSizeNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelSizeNil(b bool)`

 SetShippingLabelSizeNil sets the value for ShippingLabelSize to be an explicit nil

### UnsetShippingLabelSize
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetShippingLabelSize()`

UnsetShippingLabelSize ensures that no value is present for ShippingLabelSize, not even an explicit nil
### GetShippingLabelResolution

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelResolution() interface{}`

GetShippingLabelResolution returns the ShippingLabelResolution field if non-nil, zero value otherwise.

### GetShippingLabelResolutionOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetShippingLabelResolutionOk() (*interface{}, bool)`

GetShippingLabelResolutionOk returns a tuple with the ShippingLabelResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelResolution

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelResolution(v interface{})`

SetShippingLabelResolution sets ShippingLabelResolution field to given value.

### HasShippingLabelResolution

`func (o *GETParcels200ResponseDataInnerAttributes) HasShippingLabelResolution() bool`

HasShippingLabelResolution returns a boolean if a field has been set.

### SetShippingLabelResolutionNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetShippingLabelResolutionNil(b bool)`

 SetShippingLabelResolutionNil sets the value for ShippingLabelResolution to be an explicit nil

### UnsetShippingLabelResolution
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetShippingLabelResolution()`

UnsetShippingLabelResolution ensures that no value is present for ShippingLabelResolution, not even an explicit nil
### GetTrackingNumber

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingNumber() interface{}`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingNumberOk() (*interface{}, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingNumber(v interface{})`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetTrackingStatus

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatus() interface{}`

GetTrackingStatus returns the TrackingStatus field if non-nil, zero value otherwise.

### GetTrackingStatusOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusOk() (*interface{}, bool)`

GetTrackingStatusOk returns a tuple with the TrackingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatus

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingStatus(v interface{})`

SetTrackingStatus sets TrackingStatus field to given value.

### HasTrackingStatus

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingStatus() bool`

HasTrackingStatus returns a boolean if a field has been set.

### SetTrackingStatusNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingStatusNil(b bool)`

 SetTrackingStatusNil sets the value for TrackingStatus to be an explicit nil

### UnsetTrackingStatus
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetTrackingStatus()`

UnsetTrackingStatus ensures that no value is present for TrackingStatus, not even an explicit nil
### GetTrackingStatusDetail

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusDetail() interface{}`

GetTrackingStatusDetail returns the TrackingStatusDetail field if non-nil, zero value otherwise.

### GetTrackingStatusDetailOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusDetailOk() (*interface{}, bool)`

GetTrackingStatusDetailOk returns a tuple with the TrackingStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusDetail

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingStatusDetail(v interface{})`

SetTrackingStatusDetail sets TrackingStatusDetail field to given value.

### HasTrackingStatusDetail

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingStatusDetail() bool`

HasTrackingStatusDetail returns a boolean if a field has been set.

### SetTrackingStatusDetailNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingStatusDetailNil(b bool)`

 SetTrackingStatusDetailNil sets the value for TrackingStatusDetail to be an explicit nil

### UnsetTrackingStatusDetail
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetTrackingStatusDetail()`

UnsetTrackingStatusDetail ensures that no value is present for TrackingStatusDetail, not even an explicit nil
### GetTrackingStatusUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusUpdatedAt() interface{}`

GetTrackingStatusUpdatedAt returns the TrackingStatusUpdatedAt field if non-nil, zero value otherwise.

### GetTrackingStatusUpdatedAtOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingStatusUpdatedAtOk() (*interface{}, bool)`

GetTrackingStatusUpdatedAtOk returns a tuple with the TrackingStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingStatusUpdatedAt(v interface{})`

SetTrackingStatusUpdatedAt sets TrackingStatusUpdatedAt field to given value.

### HasTrackingStatusUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingStatusUpdatedAt() bool`

HasTrackingStatusUpdatedAt returns a boolean if a field has been set.

### SetTrackingStatusUpdatedAtNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingStatusUpdatedAtNil(b bool)`

 SetTrackingStatusUpdatedAtNil sets the value for TrackingStatusUpdatedAt to be an explicit nil

### UnsetTrackingStatusUpdatedAt
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetTrackingStatusUpdatedAt()`

UnsetTrackingStatusUpdatedAt ensures that no value is present for TrackingStatusUpdatedAt, not even an explicit nil
### GetTrackingDetails

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingDetails() interface{}`

GetTrackingDetails returns the TrackingDetails field if non-nil, zero value otherwise.

### GetTrackingDetailsOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetTrackingDetailsOk() (*interface{}, bool)`

GetTrackingDetailsOk returns a tuple with the TrackingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingDetails

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingDetails(v interface{})`

SetTrackingDetails sets TrackingDetails field to given value.

### HasTrackingDetails

`func (o *GETParcels200ResponseDataInnerAttributes) HasTrackingDetails() bool`

HasTrackingDetails returns a boolean if a field has been set.

### SetTrackingDetailsNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetTrackingDetailsNil(b bool)`

 SetTrackingDetailsNil sets the value for TrackingDetails to be an explicit nil

### UnsetTrackingDetails
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetTrackingDetails()`

UnsetTrackingDetails ensures that no value is present for TrackingDetails, not even an explicit nil
### GetCarrierWeightOz

`func (o *GETParcels200ResponseDataInnerAttributes) GetCarrierWeightOz() interface{}`

GetCarrierWeightOz returns the CarrierWeightOz field if non-nil, zero value otherwise.

### GetCarrierWeightOzOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCarrierWeightOzOk() (*interface{}, bool)`

GetCarrierWeightOzOk returns a tuple with the CarrierWeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierWeightOz

`func (o *GETParcels200ResponseDataInnerAttributes) SetCarrierWeightOz(v interface{})`

SetCarrierWeightOz sets CarrierWeightOz field to given value.

### HasCarrierWeightOz

`func (o *GETParcels200ResponseDataInnerAttributes) HasCarrierWeightOz() bool`

HasCarrierWeightOz returns a boolean if a field has been set.

### SetCarrierWeightOzNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetCarrierWeightOzNil(b bool)`

 SetCarrierWeightOzNil sets the value for CarrierWeightOz to be an explicit nil

### UnsetCarrierWeightOz
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetCarrierWeightOz()`

UnsetCarrierWeightOz ensures that no value is present for CarrierWeightOz, not even an explicit nil
### GetSignedBy

`func (o *GETParcels200ResponseDataInnerAttributes) GetSignedBy() interface{}`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetSignedByOk() (*interface{}, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *GETParcels200ResponseDataInnerAttributes) SetSignedBy(v interface{})`

SetSignedBy sets SignedBy field to given value.

### HasSignedBy

`func (o *GETParcels200ResponseDataInnerAttributes) HasSignedBy() bool`

HasSignedBy returns a boolean if a field has been set.

### SetSignedByNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetSignedByNil(b bool)`

 SetSignedByNil sets the value for SignedBy to be an explicit nil

### UnsetSignedBy
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetSignedBy()`

UnsetSignedBy ensures that no value is present for SignedBy, not even an explicit nil
### GetIncoterm

`func (o *GETParcels200ResponseDataInnerAttributes) GetIncoterm() interface{}`

GetIncoterm returns the Incoterm field if non-nil, zero value otherwise.

### GetIncotermOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetIncotermOk() (*interface{}, bool)`

GetIncotermOk returns a tuple with the Incoterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoterm

`func (o *GETParcels200ResponseDataInnerAttributes) SetIncoterm(v interface{})`

SetIncoterm sets Incoterm field to given value.

### HasIncoterm

`func (o *GETParcels200ResponseDataInnerAttributes) HasIncoterm() bool`

HasIncoterm returns a boolean if a field has been set.

### SetIncotermNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetIncotermNil(b bool)`

 SetIncotermNil sets the value for Incoterm to be an explicit nil

### UnsetIncoterm
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetIncoterm()`

UnsetIncoterm ensures that no value is present for Incoterm, not even an explicit nil
### GetCreatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETParcels200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETParcels200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETParcels200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETParcels200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETParcels200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETParcels200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETParcels200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETParcels200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETParcels200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETParcels200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETParcels200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETParcels200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETParcels200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


