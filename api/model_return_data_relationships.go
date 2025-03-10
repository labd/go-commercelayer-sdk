/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReturnDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnDataRelationships{}

// ReturnDataRelationships struct for ReturnDataRelationships
type ReturnDataRelationships struct {
	Order              *AdyenPaymentDataRelationshipsOrder             `json:"order,omitempty"`
	Customer           *CouponRecipientDataRelationshipsCustomer       `json:"customer,omitempty"`
	StockLocation      *DeliveryLeadTimeDataRelationshipsStockLocation `json:"stock_location,omitempty"`
	OriginAddress      *BingGeocoderDataRelationshipsAddresses         `json:"origin_address,omitempty"`
	DestinationAddress *BingGeocoderDataRelationshipsAddresses         `json:"destination_address,omitempty"`
	ReferenceCapture   *AuthorizationDataRelationshipsCaptures         `json:"reference_capture,omitempty"`
	ReferenceRefund    *CaptureDataRelationshipsRefunds                `json:"reference_refund,omitempty"`
	ReturnLineItems    *LineItemDataRelationshipsReturnLineItems       `json:"return_line_items,omitempty"`
	Attachments        *AuthorizationDataRelationshipsAttachments      `json:"attachments,omitempty"`
	ResourceErrors     *OrderDataRelationshipsResourceErrors           `json:"resource_errors,omitempty"`
	Events             *AddressDataRelationshipsEvents                 `json:"events,omitempty"`
	Tags               *AddressDataRelationshipsTags                   `json:"tags,omitempty"`
	Versions           *AddressDataRelationshipsVersions               `json:"versions,omitempty"`
}

// NewReturnDataRelationships instantiates a new ReturnDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnDataRelationships() *ReturnDataRelationships {
	this := ReturnDataRelationships{}
	return &this
}

// NewReturnDataRelationshipsWithDefaults instantiates a new ReturnDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnDataRelationshipsWithDefaults() *ReturnDataRelationships {
	this := ReturnDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *ReturnDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *ReturnDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *ReturnDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetOriginAddress returns the OriginAddress field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetOriginAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || IsNil(o.OriginAddress) {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.OriginAddress
}

// GetOriginAddressOk returns a tuple with the OriginAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetOriginAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || IsNil(o.OriginAddress) {
		return nil, false
	}
	return o.OriginAddress, true
}

// HasOriginAddress returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasOriginAddress() bool {
	if o != nil && !IsNil(o.OriginAddress) {
		return true
	}

	return false
}

// SetOriginAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the OriginAddress field.
func (o *ReturnDataRelationships) SetOriginAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.OriginAddress = &v
}

// GetDestinationAddress returns the DestinationAddress field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetDestinationAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || IsNil(o.DestinationAddress) {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetDestinationAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || IsNil(o.DestinationAddress) {
		return nil, false
	}
	return o.DestinationAddress, true
}

// HasDestinationAddress returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasDestinationAddress() bool {
	if o != nil && !IsNil(o.DestinationAddress) {
		return true
	}

	return false
}

// SetDestinationAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the DestinationAddress field.
func (o *ReturnDataRelationships) SetDestinationAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.DestinationAddress = &v
}

// GetReferenceCapture returns the ReferenceCapture field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetReferenceCapture() AuthorizationDataRelationshipsCaptures {
	if o == nil || IsNil(o.ReferenceCapture) {
		var ret AuthorizationDataRelationshipsCaptures
		return ret
	}
	return *o.ReferenceCapture
}

// GetReferenceCaptureOk returns a tuple with the ReferenceCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetReferenceCaptureOk() (*AuthorizationDataRelationshipsCaptures, bool) {
	if o == nil || IsNil(o.ReferenceCapture) {
		return nil, false
	}
	return o.ReferenceCapture, true
}

// HasReferenceCapture returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasReferenceCapture() bool {
	if o != nil && !IsNil(o.ReferenceCapture) {
		return true
	}

	return false
}

// SetReferenceCapture gets a reference to the given AuthorizationDataRelationshipsCaptures and assigns it to the ReferenceCapture field.
func (o *ReturnDataRelationships) SetReferenceCapture(v AuthorizationDataRelationshipsCaptures) {
	o.ReferenceCapture = &v
}

// GetReferenceRefund returns the ReferenceRefund field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetReferenceRefund() CaptureDataRelationshipsRefunds {
	if o == nil || IsNil(o.ReferenceRefund) {
		var ret CaptureDataRelationshipsRefunds
		return ret
	}
	return *o.ReferenceRefund
}

// GetReferenceRefundOk returns a tuple with the ReferenceRefund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetReferenceRefundOk() (*CaptureDataRelationshipsRefunds, bool) {
	if o == nil || IsNil(o.ReferenceRefund) {
		return nil, false
	}
	return o.ReferenceRefund, true
}

// HasReferenceRefund returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasReferenceRefund() bool {
	if o != nil && !IsNil(o.ReferenceRefund) {
		return true
	}

	return false
}

// SetReferenceRefund gets a reference to the given CaptureDataRelationshipsRefunds and assigns it to the ReferenceRefund field.
func (o *ReturnDataRelationships) SetReferenceRefund(v CaptureDataRelationshipsRefunds) {
	o.ReferenceRefund = &v
}

// GetReturnLineItems returns the ReturnLineItems field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetReturnLineItems() LineItemDataRelationshipsReturnLineItems {
	if o == nil || IsNil(o.ReturnLineItems) {
		var ret LineItemDataRelationshipsReturnLineItems
		return ret
	}
	return *o.ReturnLineItems
}

// GetReturnLineItemsOk returns a tuple with the ReturnLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetReturnLineItemsOk() (*LineItemDataRelationshipsReturnLineItems, bool) {
	if o == nil || IsNil(o.ReturnLineItems) {
		return nil, false
	}
	return o.ReturnLineItems, true
}

// HasReturnLineItems returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasReturnLineItems() bool {
	if o != nil && !IsNil(o.ReturnLineItems) {
		return true
	}

	return false
}

// SetReturnLineItems gets a reference to the given LineItemDataRelationshipsReturnLineItems and assigns it to the ReturnLineItems field.
func (o *ReturnDataRelationships) SetReturnLineItems(v LineItemDataRelationshipsReturnLineItems) {
	o.ReturnLineItems = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ReturnDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetResourceErrors returns the ResourceErrors field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetResourceErrors() OrderDataRelationshipsResourceErrors {
	if o == nil || IsNil(o.ResourceErrors) {
		var ret OrderDataRelationshipsResourceErrors
		return ret
	}
	return *o.ResourceErrors
}

// GetResourceErrorsOk returns a tuple with the ResourceErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetResourceErrorsOk() (*OrderDataRelationshipsResourceErrors, bool) {
	if o == nil || IsNil(o.ResourceErrors) {
		return nil, false
	}
	return o.ResourceErrors, true
}

// HasResourceErrors returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasResourceErrors() bool {
	if o != nil && !IsNil(o.ResourceErrors) {
		return true
	}

	return false
}

// SetResourceErrors gets a reference to the given OrderDataRelationshipsResourceErrors and assigns it to the ResourceErrors field.
func (o *ReturnDataRelationships) SetResourceErrors(v OrderDataRelationshipsResourceErrors) {
	o.ResourceErrors = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *ReturnDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetTags() AddressDataRelationshipsTags {
	if o == nil || IsNil(o.Tags) {
		var ret AddressDataRelationshipsTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given AddressDataRelationshipsTags and assigns it to the Tags field.
func (o *ReturnDataRelationships) SetTags(v AddressDataRelationshipsTags) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ReturnDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ReturnDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *ReturnDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o ReturnDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	if !IsNil(o.OriginAddress) {
		toSerialize["origin_address"] = o.OriginAddress
	}
	if !IsNil(o.DestinationAddress) {
		toSerialize["destination_address"] = o.DestinationAddress
	}
	if !IsNil(o.ReferenceCapture) {
		toSerialize["reference_capture"] = o.ReferenceCapture
	}
	if !IsNil(o.ReferenceRefund) {
		toSerialize["reference_refund"] = o.ReferenceRefund
	}
	if !IsNil(o.ReturnLineItems) {
		toSerialize["return_line_items"] = o.ReturnLineItems
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.ResourceErrors) {
		toSerialize["resource_errors"] = o.ResourceErrors
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableReturnDataRelationships struct {
	value *ReturnDataRelationships
	isSet bool
}

func (v NullableReturnDataRelationships) Get() *ReturnDataRelationships {
	return v.value
}

func (v *NullableReturnDataRelationships) Set(val *ReturnDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnDataRelationships(val *ReturnDataRelationships) *NullableReturnDataRelationships {
	return &NullableReturnDataRelationships{value: val, isSet: true}
}

func (v NullableReturnDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
