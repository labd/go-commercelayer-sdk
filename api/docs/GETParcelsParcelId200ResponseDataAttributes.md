# GETParcelsParcelId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the parcel. | [optional] 
**Weight** | Pointer to **interface{}** | The parcel weight, used to automatically calculate the tax rates from the available carrier accounts. | [optional] 
**UnitOfWeight** | Pointer to **interface{}** | The unit of weight. One of &#39;gr&#39;, &#39;oz&#39;, or &#39;lb&#39;. | [optional] 
**EelPfc** | Pointer to **interface{}** | When shipping outside the US, you need to provide either an Exemption and Exclusion Legend (EEL) code or a Proof of Filing Citation (PFC). Which you need is based on the value of the goods being shipped. Value can be one of \&quot;EEL\&quot; o \&quot;PFC\&quot;. | [optional] 
**ContentsType** | Pointer to **interface{}** | The type of item you are sending. | [optional] 
**ContentsExplanation** | Pointer to **interface{}** | If you specify &#39;other&#39; in the &#39;contents_type&#39; attribute, you must supply a brief description in this attribute. | [optional] 
**CustomsCertify** | Pointer to **interface{}** | Indicates if the provided information is accurate. | [optional] 
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
**CarrierWeightOz** | Pointer to **interface{}** | The weight of the parcel as measured by the carrier in ounces (if available). | [optional] 
**SignedBy** | Pointer to **interface{}** | The name of the person who signed for the parcel (if available). | [optional] 
**Incoterm** | Pointer to **interface{}** | The type of Incoterm (if available). | [optional] 
**DeliveryConfirmation** | Pointer to **interface{}** | The type of delivery confirmation option upon delivery. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETParcelsParcelId200ResponseDataAttributes

`func NewGETParcelsParcelId200ResponseDataAttributes() *GETParcelsParcelId200ResponseDataAttributes`

NewGETParcelsParcelId200ResponseDataAttributes instantiates a new GETParcelsParcelId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETParcelsParcelId200ResponseDataAttributesWithDefaults

`func NewGETParcelsParcelId200ResponseDataAttributesWithDefaults() *GETParcelsParcelId200ResponseDataAttributes`

NewGETParcelsParcelId200ResponseDataAttributesWithDefaults instantiates a new GETParcelsParcelId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetWeight

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetWeight() interface{}`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetWeightOk() (*interface{}, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetWeight(v interface{})`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetUnitOfWeight

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetUnitOfWeight() interface{}`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetUnitOfWeightOk() (*interface{}, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetUnitOfWeight(v interface{})`

SetUnitOfWeight sets UnitOfWeight field to given value.

### HasUnitOfWeight

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasUnitOfWeight() bool`

HasUnitOfWeight returns a boolean if a field has been set.

### SetUnitOfWeightNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetUnitOfWeightNil(b bool)`

 SetUnitOfWeightNil sets the value for UnitOfWeight to be an explicit nil

### UnsetUnitOfWeight
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetUnitOfWeight()`

UnsetUnitOfWeight ensures that no value is present for UnitOfWeight, not even an explicit nil
### GetEelPfc

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetEelPfc() interface{}`

GetEelPfc returns the EelPfc field if non-nil, zero value otherwise.

### GetEelPfcOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetEelPfcOk() (*interface{}, bool)`

GetEelPfcOk returns a tuple with the EelPfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEelPfc

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetEelPfc(v interface{})`

SetEelPfc sets EelPfc field to given value.

### HasEelPfc

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasEelPfc() bool`

HasEelPfc returns a boolean if a field has been set.

### SetEelPfcNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetEelPfcNil(b bool)`

 SetEelPfcNil sets the value for EelPfc to be an explicit nil

### UnsetEelPfc
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetEelPfc()`

UnsetEelPfc ensures that no value is present for EelPfc, not even an explicit nil
### GetContentsType

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetContentsType() interface{}`

GetContentsType returns the ContentsType field if non-nil, zero value otherwise.

### GetContentsTypeOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetContentsTypeOk() (*interface{}, bool)`

GetContentsTypeOk returns a tuple with the ContentsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsType

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetContentsType(v interface{})`

SetContentsType sets ContentsType field to given value.

### HasContentsType

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasContentsType() bool`

HasContentsType returns a boolean if a field has been set.

### SetContentsTypeNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetContentsTypeNil(b bool)`

 SetContentsTypeNil sets the value for ContentsType to be an explicit nil

### UnsetContentsType
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetContentsType()`

UnsetContentsType ensures that no value is present for ContentsType, not even an explicit nil
### GetContentsExplanation

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetContentsExplanation() interface{}`

GetContentsExplanation returns the ContentsExplanation field if non-nil, zero value otherwise.

### GetContentsExplanationOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetContentsExplanationOk() (*interface{}, bool)`

GetContentsExplanationOk returns a tuple with the ContentsExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsExplanation

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetContentsExplanation(v interface{})`

SetContentsExplanation sets ContentsExplanation field to given value.

### HasContentsExplanation

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasContentsExplanation() bool`

HasContentsExplanation returns a boolean if a field has been set.

### SetContentsExplanationNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetContentsExplanationNil(b bool)`

 SetContentsExplanationNil sets the value for ContentsExplanation to be an explicit nil

### UnsetContentsExplanation
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetContentsExplanation()`

UnsetContentsExplanation ensures that no value is present for ContentsExplanation, not even an explicit nil
### GetCustomsCertify

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCustomsCertify() interface{}`

GetCustomsCertify returns the CustomsCertify field if non-nil, zero value otherwise.

### GetCustomsCertifyOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCustomsCertifyOk() (*interface{}, bool)`

GetCustomsCertifyOk returns a tuple with the CustomsCertify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsCertify

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCustomsCertify(v interface{})`

SetCustomsCertify sets CustomsCertify field to given value.

### HasCustomsCertify

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasCustomsCertify() bool`

HasCustomsCertify returns a boolean if a field has been set.

### SetCustomsCertifyNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCustomsCertifyNil(b bool)`

 SetCustomsCertifyNil sets the value for CustomsCertify to be an explicit nil

### UnsetCustomsCertify
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetCustomsCertify()`

UnsetCustomsCertify ensures that no value is present for CustomsCertify, not even an explicit nil
### GetCustomsSigner

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCustomsSigner() interface{}`

GetCustomsSigner returns the CustomsSigner field if non-nil, zero value otherwise.

### GetCustomsSignerOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCustomsSignerOk() (*interface{}, bool)`

GetCustomsSignerOk returns a tuple with the CustomsSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsSigner

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCustomsSigner(v interface{})`

SetCustomsSigner sets CustomsSigner field to given value.

### HasCustomsSigner

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasCustomsSigner() bool`

HasCustomsSigner returns a boolean if a field has been set.

### SetCustomsSignerNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCustomsSignerNil(b bool)`

 SetCustomsSignerNil sets the value for CustomsSigner to be an explicit nil

### UnsetCustomsSigner
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetCustomsSigner()`

UnsetCustomsSigner ensures that no value is present for CustomsSigner, not even an explicit nil
### GetNonDeliveryOption

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetNonDeliveryOption() interface{}`

GetNonDeliveryOption returns the NonDeliveryOption field if non-nil, zero value otherwise.

### GetNonDeliveryOptionOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetNonDeliveryOptionOk() (*interface{}, bool)`

GetNonDeliveryOptionOk returns a tuple with the NonDeliveryOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDeliveryOption

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetNonDeliveryOption(v interface{})`

SetNonDeliveryOption sets NonDeliveryOption field to given value.

### HasNonDeliveryOption

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasNonDeliveryOption() bool`

HasNonDeliveryOption returns a boolean if a field has been set.

### SetNonDeliveryOptionNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetNonDeliveryOptionNil(b bool)`

 SetNonDeliveryOptionNil sets the value for NonDeliveryOption to be an explicit nil

### UnsetNonDeliveryOption
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetNonDeliveryOption()`

UnsetNonDeliveryOption ensures that no value is present for NonDeliveryOption, not even an explicit nil
### GetRestrictionType

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetRestrictionType() interface{}`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetRestrictionTypeOk() (*interface{}, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetRestrictionType(v interface{})`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.

### SetRestrictionTypeNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetRestrictionTypeNil(b bool)`

 SetRestrictionTypeNil sets the value for RestrictionType to be an explicit nil

### UnsetRestrictionType
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetRestrictionType()`

UnsetRestrictionType ensures that no value is present for RestrictionType, not even an explicit nil
### GetRestrictionComments

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetRestrictionComments() interface{}`

GetRestrictionComments returns the RestrictionComments field if non-nil, zero value otherwise.

### GetRestrictionCommentsOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetRestrictionCommentsOk() (*interface{}, bool)`

GetRestrictionCommentsOk returns a tuple with the RestrictionComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionComments

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetRestrictionComments(v interface{})`

SetRestrictionComments sets RestrictionComments field to given value.

### HasRestrictionComments

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasRestrictionComments() bool`

HasRestrictionComments returns a boolean if a field has been set.

### SetRestrictionCommentsNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetRestrictionCommentsNil(b bool)`

 SetRestrictionCommentsNil sets the value for RestrictionComments to be an explicit nil

### UnsetRestrictionComments
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetRestrictionComments()`

UnsetRestrictionComments ensures that no value is present for RestrictionComments, not even an explicit nil
### GetCustomsInfoRequired

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCustomsInfoRequired() interface{}`

GetCustomsInfoRequired returns the CustomsInfoRequired field if non-nil, zero value otherwise.

### GetCustomsInfoRequiredOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCustomsInfoRequiredOk() (*interface{}, bool)`

GetCustomsInfoRequiredOk returns a tuple with the CustomsInfoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsInfoRequired

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCustomsInfoRequired(v interface{})`

SetCustomsInfoRequired sets CustomsInfoRequired field to given value.

### HasCustomsInfoRequired

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasCustomsInfoRequired() bool`

HasCustomsInfoRequired returns a boolean if a field has been set.

### SetCustomsInfoRequiredNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCustomsInfoRequiredNil(b bool)`

 SetCustomsInfoRequiredNil sets the value for CustomsInfoRequired to be an explicit nil

### UnsetCustomsInfoRequired
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetCustomsInfoRequired()`

UnsetCustomsInfoRequired ensures that no value is present for CustomsInfoRequired, not even an explicit nil
### GetShippingLabelUrl

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetShippingLabelUrl() interface{}`

GetShippingLabelUrl returns the ShippingLabelUrl field if non-nil, zero value otherwise.

### GetShippingLabelUrlOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetShippingLabelUrlOk() (*interface{}, bool)`

GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelUrl

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetShippingLabelUrl(v interface{})`

SetShippingLabelUrl sets ShippingLabelUrl field to given value.

### HasShippingLabelUrl

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasShippingLabelUrl() bool`

HasShippingLabelUrl returns a boolean if a field has been set.

### SetShippingLabelUrlNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetShippingLabelUrlNil(b bool)`

 SetShippingLabelUrlNil sets the value for ShippingLabelUrl to be an explicit nil

### UnsetShippingLabelUrl
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetShippingLabelUrl()`

UnsetShippingLabelUrl ensures that no value is present for ShippingLabelUrl, not even an explicit nil
### GetShippingLabelFileType

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetShippingLabelFileType() interface{}`

GetShippingLabelFileType returns the ShippingLabelFileType field if non-nil, zero value otherwise.

### GetShippingLabelFileTypeOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetShippingLabelFileTypeOk() (*interface{}, bool)`

GetShippingLabelFileTypeOk returns a tuple with the ShippingLabelFileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelFileType

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetShippingLabelFileType(v interface{})`

SetShippingLabelFileType sets ShippingLabelFileType field to given value.

### HasShippingLabelFileType

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasShippingLabelFileType() bool`

HasShippingLabelFileType returns a boolean if a field has been set.

### SetShippingLabelFileTypeNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetShippingLabelFileTypeNil(b bool)`

 SetShippingLabelFileTypeNil sets the value for ShippingLabelFileType to be an explicit nil

### UnsetShippingLabelFileType
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetShippingLabelFileType()`

UnsetShippingLabelFileType ensures that no value is present for ShippingLabelFileType, not even an explicit nil
### GetShippingLabelSize

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetShippingLabelSize() interface{}`

GetShippingLabelSize returns the ShippingLabelSize field if non-nil, zero value otherwise.

### GetShippingLabelSizeOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetShippingLabelSizeOk() (*interface{}, bool)`

GetShippingLabelSizeOk returns a tuple with the ShippingLabelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelSize

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetShippingLabelSize(v interface{})`

SetShippingLabelSize sets ShippingLabelSize field to given value.

### HasShippingLabelSize

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasShippingLabelSize() bool`

HasShippingLabelSize returns a boolean if a field has been set.

### SetShippingLabelSizeNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetShippingLabelSizeNil(b bool)`

 SetShippingLabelSizeNil sets the value for ShippingLabelSize to be an explicit nil

### UnsetShippingLabelSize
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetShippingLabelSize()`

UnsetShippingLabelSize ensures that no value is present for ShippingLabelSize, not even an explicit nil
### GetShippingLabelResolution

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetShippingLabelResolution() interface{}`

GetShippingLabelResolution returns the ShippingLabelResolution field if non-nil, zero value otherwise.

### GetShippingLabelResolutionOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetShippingLabelResolutionOk() (*interface{}, bool)`

GetShippingLabelResolutionOk returns a tuple with the ShippingLabelResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelResolution

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetShippingLabelResolution(v interface{})`

SetShippingLabelResolution sets ShippingLabelResolution field to given value.

### HasShippingLabelResolution

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasShippingLabelResolution() bool`

HasShippingLabelResolution returns a boolean if a field has been set.

### SetShippingLabelResolutionNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetShippingLabelResolutionNil(b bool)`

 SetShippingLabelResolutionNil sets the value for ShippingLabelResolution to be an explicit nil

### UnsetShippingLabelResolution
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetShippingLabelResolution()`

UnsetShippingLabelResolution ensures that no value is present for ShippingLabelResolution, not even an explicit nil
### GetTrackingNumber

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingNumber() interface{}`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingNumberOk() (*interface{}, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingNumber(v interface{})`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetTrackingStatus

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingStatus() interface{}`

GetTrackingStatus returns the TrackingStatus field if non-nil, zero value otherwise.

### GetTrackingStatusOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingStatusOk() (*interface{}, bool)`

GetTrackingStatusOk returns a tuple with the TrackingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatus

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingStatus(v interface{})`

SetTrackingStatus sets TrackingStatus field to given value.

### HasTrackingStatus

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasTrackingStatus() bool`

HasTrackingStatus returns a boolean if a field has been set.

### SetTrackingStatusNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingStatusNil(b bool)`

 SetTrackingStatusNil sets the value for TrackingStatus to be an explicit nil

### UnsetTrackingStatus
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetTrackingStatus()`

UnsetTrackingStatus ensures that no value is present for TrackingStatus, not even an explicit nil
### GetTrackingStatusDetail

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingStatusDetail() interface{}`

GetTrackingStatusDetail returns the TrackingStatusDetail field if non-nil, zero value otherwise.

### GetTrackingStatusDetailOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingStatusDetailOk() (*interface{}, bool)`

GetTrackingStatusDetailOk returns a tuple with the TrackingStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusDetail

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingStatusDetail(v interface{})`

SetTrackingStatusDetail sets TrackingStatusDetail field to given value.

### HasTrackingStatusDetail

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasTrackingStatusDetail() bool`

HasTrackingStatusDetail returns a boolean if a field has been set.

### SetTrackingStatusDetailNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingStatusDetailNil(b bool)`

 SetTrackingStatusDetailNil sets the value for TrackingStatusDetail to be an explicit nil

### UnsetTrackingStatusDetail
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetTrackingStatusDetail()`

UnsetTrackingStatusDetail ensures that no value is present for TrackingStatusDetail, not even an explicit nil
### GetTrackingStatusUpdatedAt

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingStatusUpdatedAt() interface{}`

GetTrackingStatusUpdatedAt returns the TrackingStatusUpdatedAt field if non-nil, zero value otherwise.

### GetTrackingStatusUpdatedAtOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingStatusUpdatedAtOk() (*interface{}, bool)`

GetTrackingStatusUpdatedAtOk returns a tuple with the TrackingStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusUpdatedAt

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingStatusUpdatedAt(v interface{})`

SetTrackingStatusUpdatedAt sets TrackingStatusUpdatedAt field to given value.

### HasTrackingStatusUpdatedAt

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasTrackingStatusUpdatedAt() bool`

HasTrackingStatusUpdatedAt returns a boolean if a field has been set.

### SetTrackingStatusUpdatedAtNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingStatusUpdatedAtNil(b bool)`

 SetTrackingStatusUpdatedAtNil sets the value for TrackingStatusUpdatedAt to be an explicit nil

### UnsetTrackingStatusUpdatedAt
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetTrackingStatusUpdatedAt()`

UnsetTrackingStatusUpdatedAt ensures that no value is present for TrackingStatusUpdatedAt, not even an explicit nil
### GetTrackingDetails

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingDetails() interface{}`

GetTrackingDetails returns the TrackingDetails field if non-nil, zero value otherwise.

### GetTrackingDetailsOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetTrackingDetailsOk() (*interface{}, bool)`

GetTrackingDetailsOk returns a tuple with the TrackingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingDetails

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingDetails(v interface{})`

SetTrackingDetails sets TrackingDetails field to given value.

### HasTrackingDetails

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasTrackingDetails() bool`

HasTrackingDetails returns a boolean if a field has been set.

### SetTrackingDetailsNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetTrackingDetailsNil(b bool)`

 SetTrackingDetailsNil sets the value for TrackingDetails to be an explicit nil

### UnsetTrackingDetails
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetTrackingDetails()`

UnsetTrackingDetails ensures that no value is present for TrackingDetails, not even an explicit nil
### GetCarrierWeightOz

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCarrierWeightOz() interface{}`

GetCarrierWeightOz returns the CarrierWeightOz field if non-nil, zero value otherwise.

### GetCarrierWeightOzOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCarrierWeightOzOk() (*interface{}, bool)`

GetCarrierWeightOzOk returns a tuple with the CarrierWeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierWeightOz

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCarrierWeightOz(v interface{})`

SetCarrierWeightOz sets CarrierWeightOz field to given value.

### HasCarrierWeightOz

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasCarrierWeightOz() bool`

HasCarrierWeightOz returns a boolean if a field has been set.

### SetCarrierWeightOzNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCarrierWeightOzNil(b bool)`

 SetCarrierWeightOzNil sets the value for CarrierWeightOz to be an explicit nil

### UnsetCarrierWeightOz
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetCarrierWeightOz()`

UnsetCarrierWeightOz ensures that no value is present for CarrierWeightOz, not even an explicit nil
### GetSignedBy

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetSignedBy() interface{}`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetSignedByOk() (*interface{}, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetSignedBy(v interface{})`

SetSignedBy sets SignedBy field to given value.

### HasSignedBy

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasSignedBy() bool`

HasSignedBy returns a boolean if a field has been set.

### SetSignedByNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetSignedByNil(b bool)`

 SetSignedByNil sets the value for SignedBy to be an explicit nil

### UnsetSignedBy
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetSignedBy()`

UnsetSignedBy ensures that no value is present for SignedBy, not even an explicit nil
### GetIncoterm

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetIncoterm() interface{}`

GetIncoterm returns the Incoterm field if non-nil, zero value otherwise.

### GetIncotermOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetIncotermOk() (*interface{}, bool)`

GetIncotermOk returns a tuple with the Incoterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoterm

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetIncoterm(v interface{})`

SetIncoterm sets Incoterm field to given value.

### HasIncoterm

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasIncoterm() bool`

HasIncoterm returns a boolean if a field has been set.

### SetIncotermNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetIncotermNil(b bool)`

 SetIncotermNil sets the value for Incoterm to be an explicit nil

### UnsetIncoterm
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetIncoterm()`

UnsetIncoterm ensures that no value is present for Incoterm, not even an explicit nil
### GetDeliveryConfirmation

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetDeliveryConfirmation() interface{}`

GetDeliveryConfirmation returns the DeliveryConfirmation field if non-nil, zero value otherwise.

### GetDeliveryConfirmationOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetDeliveryConfirmationOk() (*interface{}, bool)`

GetDeliveryConfirmationOk returns a tuple with the DeliveryConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryConfirmation

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetDeliveryConfirmation(v interface{})`

SetDeliveryConfirmation sets DeliveryConfirmation field to given value.

### HasDeliveryConfirmation

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasDeliveryConfirmation() bool`

HasDeliveryConfirmation returns a boolean if a field has been set.

### SetDeliveryConfirmationNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetDeliveryConfirmationNil(b bool)`

 SetDeliveryConfirmationNil sets the value for DeliveryConfirmation to be an explicit nil

### UnsetDeliveryConfirmation
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetDeliveryConfirmation()`

UnsetDeliveryConfirmation ensures that no value is present for DeliveryConfirmation, not even an explicit nil
### GetCreatedAt

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETParcelsParcelId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETParcelsParcelId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETParcelsParcelId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETParcelsParcelId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


