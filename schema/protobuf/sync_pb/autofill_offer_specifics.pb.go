// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: autofill_offer_specifics.proto

package sync_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Next tag: 11
type AutofillOfferSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id for this offer data. Will be used as the client tag.
	Id *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The link leading to the offer details page on Gpay app. Will be populated
	// on Android only.
	OfferDetailsUrl *string `protobuf:"bytes,2,opt,name=offer_details_url,json=offerDetailsUrl" json:"offer_details_url,omitempty"`
	// Merchant domain and merchant app package name refers to the merchant this
	// offer is applied to.
	MerchantDomain     []string `protobuf:"bytes,3,rep,name=merchant_domain,json=merchantDomain" json:"merchant_domain,omitempty"`
	MerchantAppPackage []string `protobuf:"bytes,4,rep,name=merchant_app_package,json=merchantAppPackage" json:"merchant_app_package,omitempty"`
	// The expiry of this offer. Will be represented in the form of unix epoch
	// time in seconds. Once the offer is expired it will not be shown in the
	// client.
	OfferExpiryDate *int64 `protobuf:"varint,5,opt,name=offer_expiry_date,json=offerExpiryDate" json:"offer_expiry_date,omitempty"`
	// The unique offer data for different offer types.
	//
	// Types that are assignable to TypeSpecificOfferData:
	//	*AutofillOfferSpecifics_CardLinkedOfferData_
	//	*AutofillOfferSpecifics_PromoCodeOfferData_
	TypeSpecificOfferData isAutofillOfferSpecifics_TypeSpecificOfferData `protobuf_oneof:"type_specific_offer_data"`
	DisplayStrings        *AutofillOfferSpecifics_DisplayStrings         `protobuf:"bytes,10,opt,name=display_strings,json=displayStrings" json:"display_strings,omitempty"`
	// The reward type of the offer. Will be used to generate the display text in
	// the UI. Each type has its own client side text template.
	//
	// Types that are assignable to RewardType:
	//	*AutofillOfferSpecifics_PercentageReward_
	//	*AutofillOfferSpecifics_FixedAmountReward_
	RewardType isAutofillOfferSpecifics_RewardType `protobuf_oneof:"reward_type"`
}

func (x *AutofillOfferSpecifics) Reset() {
	*x = AutofillOfferSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autofill_offer_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutofillOfferSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutofillOfferSpecifics) ProtoMessage() {}

func (x *AutofillOfferSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_autofill_offer_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutofillOfferSpecifics.ProtoReflect.Descriptor instead.
func (*AutofillOfferSpecifics) Descriptor() ([]byte, []int) {
	return file_autofill_offer_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *AutofillOfferSpecifics) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *AutofillOfferSpecifics) GetOfferDetailsUrl() string {
	if x != nil && x.OfferDetailsUrl != nil {
		return *x.OfferDetailsUrl
	}
	return ""
}

func (x *AutofillOfferSpecifics) GetMerchantDomain() []string {
	if x != nil {
		return x.MerchantDomain
	}
	return nil
}

func (x *AutofillOfferSpecifics) GetMerchantAppPackage() []string {
	if x != nil {
		return x.MerchantAppPackage
	}
	return nil
}

func (x *AutofillOfferSpecifics) GetOfferExpiryDate() int64 {
	if x != nil && x.OfferExpiryDate != nil {
		return *x.OfferExpiryDate
	}
	return 0
}

func (m *AutofillOfferSpecifics) GetTypeSpecificOfferData() isAutofillOfferSpecifics_TypeSpecificOfferData {
	if m != nil {
		return m.TypeSpecificOfferData
	}
	return nil
}

func (x *AutofillOfferSpecifics) GetCardLinkedOfferData() *AutofillOfferSpecifics_CardLinkedOfferData {
	if x, ok := x.GetTypeSpecificOfferData().(*AutofillOfferSpecifics_CardLinkedOfferData_); ok {
		return x.CardLinkedOfferData
	}
	return nil
}

func (x *AutofillOfferSpecifics) GetPromoCodeOfferData() *AutofillOfferSpecifics_PromoCodeOfferData {
	if x, ok := x.GetTypeSpecificOfferData().(*AutofillOfferSpecifics_PromoCodeOfferData_); ok {
		return x.PromoCodeOfferData
	}
	return nil
}

func (x *AutofillOfferSpecifics) GetDisplayStrings() *AutofillOfferSpecifics_DisplayStrings {
	if x != nil {
		return x.DisplayStrings
	}
	return nil
}

func (m *AutofillOfferSpecifics) GetRewardType() isAutofillOfferSpecifics_RewardType {
	if m != nil {
		return m.RewardType
	}
	return nil
}

func (x *AutofillOfferSpecifics) GetPercentageReward() *AutofillOfferSpecifics_PercentageReward {
	if x, ok := x.GetRewardType().(*AutofillOfferSpecifics_PercentageReward_); ok {
		return x.PercentageReward
	}
	return nil
}

func (x *AutofillOfferSpecifics) GetFixedAmountReward() *AutofillOfferSpecifics_FixedAmountReward {
	if x, ok := x.GetRewardType().(*AutofillOfferSpecifics_FixedAmountReward_); ok {
		return x.FixedAmountReward
	}
	return nil
}

type isAutofillOfferSpecifics_TypeSpecificOfferData interface {
	isAutofillOfferSpecifics_TypeSpecificOfferData()
}

type AutofillOfferSpecifics_CardLinkedOfferData_ struct {
	CardLinkedOfferData *AutofillOfferSpecifics_CardLinkedOfferData `protobuf:"bytes,6,opt,name=card_linked_offer_data,json=cardLinkedOfferData,oneof"`
}

type AutofillOfferSpecifics_PromoCodeOfferData_ struct {
	PromoCodeOfferData *AutofillOfferSpecifics_PromoCodeOfferData `protobuf:"bytes,9,opt,name=promo_code_offer_data,json=promoCodeOfferData,oneof"`
}

func (*AutofillOfferSpecifics_CardLinkedOfferData_) isAutofillOfferSpecifics_TypeSpecificOfferData() {
}

func (*AutofillOfferSpecifics_PromoCodeOfferData_) isAutofillOfferSpecifics_TypeSpecificOfferData() {}

type isAutofillOfferSpecifics_RewardType interface {
	isAutofillOfferSpecifics_RewardType()
}

type AutofillOfferSpecifics_PercentageReward_ struct {
	PercentageReward *AutofillOfferSpecifics_PercentageReward `protobuf:"bytes,7,opt,name=percentage_reward,json=percentageReward,oneof"`
}

type AutofillOfferSpecifics_FixedAmountReward_ struct {
	FixedAmountReward *AutofillOfferSpecifics_FixedAmountReward `protobuf:"bytes,8,opt,name=fixed_amount_reward,json=fixedAmountReward,oneof"`
}

func (*AutofillOfferSpecifics_PercentageReward_) isAutofillOfferSpecifics_RewardType() {}

func (*AutofillOfferSpecifics_FixedAmountReward_) isAutofillOfferSpecifics_RewardType() {}

// Proto containing data specific to a card-linked offer.
type AutofillOfferSpecifics_CardLinkedOfferData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server id of the card to which the offer is linked.
	InstrumentId []int64 `protobuf:"varint,3,rep,name=instrument_id,json=instrumentId" json:"instrument_id,omitempty"`
}

func (x *AutofillOfferSpecifics_CardLinkedOfferData) Reset() {
	*x = AutofillOfferSpecifics_CardLinkedOfferData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autofill_offer_specifics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutofillOfferSpecifics_CardLinkedOfferData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutofillOfferSpecifics_CardLinkedOfferData) ProtoMessage() {}

func (x *AutofillOfferSpecifics_CardLinkedOfferData) ProtoReflect() protoreflect.Message {
	mi := &file_autofill_offer_specifics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutofillOfferSpecifics_CardLinkedOfferData.ProtoReflect.Descriptor instead.
func (*AutofillOfferSpecifics_CardLinkedOfferData) Descriptor() ([]byte, []int) {
	return file_autofill_offer_specifics_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AutofillOfferSpecifics_CardLinkedOfferData) GetInstrumentId() []int64 {
	if x != nil {
		return x.InstrumentId
	}
	return nil
}

// Proto containing data specific to a promo code offer.
type AutofillOfferSpecifics_PromoCodeOfferData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The actual promo code which can be applied at checkout.
	PromoCode *string `protobuf:"bytes,1,opt,name=promo_code,json=promoCode" json:"promo_code,omitempty"`
}

func (x *AutofillOfferSpecifics_PromoCodeOfferData) Reset() {
	*x = AutofillOfferSpecifics_PromoCodeOfferData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autofill_offer_specifics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutofillOfferSpecifics_PromoCodeOfferData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutofillOfferSpecifics_PromoCodeOfferData) ProtoMessage() {}

func (x *AutofillOfferSpecifics_PromoCodeOfferData) ProtoReflect() protoreflect.Message {
	mi := &file_autofill_offer_specifics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutofillOfferSpecifics_PromoCodeOfferData.ProtoReflect.Descriptor instead.
func (*AutofillOfferSpecifics_PromoCodeOfferData) Descriptor() ([]byte, []int) {
	return file_autofill_offer_specifics_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AutofillOfferSpecifics_PromoCodeOfferData) GetPromoCode() string {
	if x != nil && x.PromoCode != nil {
		return *x.PromoCode
	}
	return ""
}

// Strings to be shown in client UI, based on the offer type and details.
type AutofillOfferSpecifics_DisplayStrings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A message translated in the user's GPay app locale, explaining the value
	// of the offer. For example, a promo code offer might display
	// "$5 off on shoes, up to $50."
	ValuePropText *string `protobuf:"bytes,1,opt,name=value_prop_text,json=valuePropText" json:"value_prop_text,omitempty"`
	// A message translated in the user's GPay app locale and shown on mobile as
	// a link, prompting the user to click it to learn more about the offer.
	// Generally, "See details".
	SeeDetailsTextMobile *string `protobuf:"bytes,2,opt,name=see_details_text_mobile,json=seeDetailsTextMobile" json:"see_details_text_mobile,omitempty"`
	// A message translated in the user's GPay app locale and shown on desktop
	// (not as a link), informing the user that exclusions and restrictions may
	// apply to the value prop text. Generally, "Terms apply."
	SeeDetailsTextDesktop *string `protobuf:"bytes,3,opt,name=see_details_text_desktop,json=seeDetailsTextDesktop" json:"see_details_text_desktop,omitempty"`
	// A message translated in the user's GPay app locale and shown on mobile,
	// instructing them on how to redeem the offer. For example, a promo code
	// offer might display "Tap the promo code field at checkout to autofill
	// it."
	UsageInstructionsTextMobile *string `protobuf:"bytes,4,opt,name=usage_instructions_text_mobile,json=usageInstructionsTextMobile" json:"usage_instructions_text_mobile,omitempty"`
	// A message translated in the user's GPay app locale and shown on desktop,
	// instructing them on how to redeem the offer. For example, a promo code
	// offer might display "Click the promo code field at checkout to autofill
	// it."
	UsageInstructionsTextDesktop *string `protobuf:"bytes,5,opt,name=usage_instructions_text_desktop,json=usageInstructionsTextDesktop" json:"usage_instructions_text_desktop,omitempty"`
}

func (x *AutofillOfferSpecifics_DisplayStrings) Reset() {
	*x = AutofillOfferSpecifics_DisplayStrings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autofill_offer_specifics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutofillOfferSpecifics_DisplayStrings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutofillOfferSpecifics_DisplayStrings) ProtoMessage() {}

func (x *AutofillOfferSpecifics_DisplayStrings) ProtoReflect() protoreflect.Message {
	mi := &file_autofill_offer_specifics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutofillOfferSpecifics_DisplayStrings.ProtoReflect.Descriptor instead.
func (*AutofillOfferSpecifics_DisplayStrings) Descriptor() ([]byte, []int) {
	return file_autofill_offer_specifics_proto_rawDescGZIP(), []int{0, 2}
}

func (x *AutofillOfferSpecifics_DisplayStrings) GetValuePropText() string {
	if x != nil && x.ValuePropText != nil {
		return *x.ValuePropText
	}
	return ""
}

func (x *AutofillOfferSpecifics_DisplayStrings) GetSeeDetailsTextMobile() string {
	if x != nil && x.SeeDetailsTextMobile != nil {
		return *x.SeeDetailsTextMobile
	}
	return ""
}

func (x *AutofillOfferSpecifics_DisplayStrings) GetSeeDetailsTextDesktop() string {
	if x != nil && x.SeeDetailsTextDesktop != nil {
		return *x.SeeDetailsTextDesktop
	}
	return ""
}

func (x *AutofillOfferSpecifics_DisplayStrings) GetUsageInstructionsTextMobile() string {
	if x != nil && x.UsageInstructionsTextMobile != nil {
		return *x.UsageInstructionsTextMobile
	}
	return ""
}

func (x *AutofillOfferSpecifics_DisplayStrings) GetUsageInstructionsTextDesktop() string {
	if x != nil && x.UsageInstructionsTextDesktop != nil {
		return *x.UsageInstructionsTextDesktop
	}
	return ""
}

// This value will be shown in the offer text template as "XXX% cashback".
// Percentage has a range of (0, 100].
type AutofillOfferSpecifics_PercentageReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The string contains a number and a percent sign.
	Percentage *string `protobuf:"bytes,1,opt,name=percentage" json:"percentage,omitempty"`
}

func (x *AutofillOfferSpecifics_PercentageReward) Reset() {
	*x = AutofillOfferSpecifics_PercentageReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autofill_offer_specifics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutofillOfferSpecifics_PercentageReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutofillOfferSpecifics_PercentageReward) ProtoMessage() {}

func (x *AutofillOfferSpecifics_PercentageReward) ProtoReflect() protoreflect.Message {
	mi := &file_autofill_offer_specifics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutofillOfferSpecifics_PercentageReward.ProtoReflect.Descriptor instead.
func (*AutofillOfferSpecifics_PercentageReward) Descriptor() ([]byte, []int) {
	return file_autofill_offer_specifics_proto_rawDescGZIP(), []int{0, 3}
}

func (x *AutofillOfferSpecifics_PercentageReward) GetPercentage() string {
	if x != nil && x.Percentage != nil {
		return *x.Percentage
	}
	return ""
}

// This value will be shown in the offer text template as "XXX$ off".
type AutofillOfferSpecifics_FixedAmountReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The string contains a number and a currency sign.
	Amount *string `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
}

func (x *AutofillOfferSpecifics_FixedAmountReward) Reset() {
	*x = AutofillOfferSpecifics_FixedAmountReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autofill_offer_specifics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutofillOfferSpecifics_FixedAmountReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutofillOfferSpecifics_FixedAmountReward) ProtoMessage() {}

func (x *AutofillOfferSpecifics_FixedAmountReward) ProtoReflect() protoreflect.Message {
	mi := &file_autofill_offer_specifics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutofillOfferSpecifics_FixedAmountReward.ProtoReflect.Descriptor instead.
func (*AutofillOfferSpecifics_FixedAmountReward) Descriptor() ([]byte, []int) {
	return file_autofill_offer_specifics_proto_rawDescGZIP(), []int{0, 4}
}

func (x *AutofillOfferSpecifics_FixedAmountReward) GetAmount() string {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return ""
}

var File_autofill_offer_specifics_proto protoreflect.FileDescriptor

var file_autofill_offer_specifics_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x8f, 0x0a, 0x0a, 0x16, 0x41, 0x75,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x55, 0x72, 0x6c,
	0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x41, 0x70, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x6a, 0x0a, 0x16, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70,
	0x62, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x6e,
	0x6b, 0x65, 0x64, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x13,
	0x63, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x67, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x43,
	0x6f, 0x64, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x57, 0x0a, 0x0f,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e,
	0x41, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5f, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x73, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x48, 0x01, 0x52, 0x10, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x63, 0x0a, 0x13, 0x66, 0x69, 0x78, 0x65, 0x64, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x01, 0x52, 0x11, 0x66, 0x69, 0x78, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x1a, 0x46, 0x0a, 0x13, 0x43,
	0x61, 0x72, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08,
	0x02, 0x10, 0x03, 0x1a, 0x33, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0xb4, 0x02, 0x0a, 0x0e, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x73, 0x65, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x65, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x54, 0x65, 0x78, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x65,
	0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64,
	0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x65,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x54, 0x65, 0x78, 0x74, 0x44, 0x65, 0x73, 0x6b,
	0x74, 0x6f, 0x70, 0x12, 0x43, 0x0a, 0x1e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65,
	0x78, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x45, 0x0a, 0x1f, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1c, 0x75, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x78, 0x74, 0x44, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x1a,
	0x32, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x1a, 0x2b, 0x0a, 0x11, 0x46, 0x69, 0x78, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x1a, 0x0a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0d, 0x0a, 0x0b,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x2b, 0x0a, 0x25, 0x6f,
	0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_autofill_offer_specifics_proto_rawDescOnce sync.Once
	file_autofill_offer_specifics_proto_rawDescData = file_autofill_offer_specifics_proto_rawDesc
)

func file_autofill_offer_specifics_proto_rawDescGZIP() []byte {
	file_autofill_offer_specifics_proto_rawDescOnce.Do(func() {
		file_autofill_offer_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_autofill_offer_specifics_proto_rawDescData)
	})
	return file_autofill_offer_specifics_proto_rawDescData
}

var file_autofill_offer_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_autofill_offer_specifics_proto_goTypes = []interface{}{
	(*AutofillOfferSpecifics)(nil),                     // 0: sync_pb.AutofillOfferSpecifics
	(*AutofillOfferSpecifics_CardLinkedOfferData)(nil), // 1: sync_pb.AutofillOfferSpecifics.CardLinkedOfferData
	(*AutofillOfferSpecifics_PromoCodeOfferData)(nil),  // 2: sync_pb.AutofillOfferSpecifics.PromoCodeOfferData
	(*AutofillOfferSpecifics_DisplayStrings)(nil),      // 3: sync_pb.AutofillOfferSpecifics.DisplayStrings
	(*AutofillOfferSpecifics_PercentageReward)(nil),    // 4: sync_pb.AutofillOfferSpecifics.PercentageReward
	(*AutofillOfferSpecifics_FixedAmountReward)(nil),   // 5: sync_pb.AutofillOfferSpecifics.FixedAmountReward
}
var file_autofill_offer_specifics_proto_depIdxs = []int32{
	1, // 0: sync_pb.AutofillOfferSpecifics.card_linked_offer_data:type_name -> sync_pb.AutofillOfferSpecifics.CardLinkedOfferData
	2, // 1: sync_pb.AutofillOfferSpecifics.promo_code_offer_data:type_name -> sync_pb.AutofillOfferSpecifics.PromoCodeOfferData
	3, // 2: sync_pb.AutofillOfferSpecifics.display_strings:type_name -> sync_pb.AutofillOfferSpecifics.DisplayStrings
	4, // 3: sync_pb.AutofillOfferSpecifics.percentage_reward:type_name -> sync_pb.AutofillOfferSpecifics.PercentageReward
	5, // 4: sync_pb.AutofillOfferSpecifics.fixed_amount_reward:type_name -> sync_pb.AutofillOfferSpecifics.FixedAmountReward
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_autofill_offer_specifics_proto_init() }
func file_autofill_offer_specifics_proto_init() {
	if File_autofill_offer_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autofill_offer_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutofillOfferSpecifics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_autofill_offer_specifics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutofillOfferSpecifics_CardLinkedOfferData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_autofill_offer_specifics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutofillOfferSpecifics_PromoCodeOfferData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_autofill_offer_specifics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutofillOfferSpecifics_DisplayStrings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_autofill_offer_specifics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutofillOfferSpecifics_PercentageReward); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_autofill_offer_specifics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutofillOfferSpecifics_FixedAmountReward); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_autofill_offer_specifics_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AutofillOfferSpecifics_CardLinkedOfferData_)(nil),
		(*AutofillOfferSpecifics_PromoCodeOfferData_)(nil),
		(*AutofillOfferSpecifics_PercentageReward_)(nil),
		(*AutofillOfferSpecifics_FixedAmountReward_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_autofill_offer_specifics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_autofill_offer_specifics_proto_goTypes,
		DependencyIndexes: file_autofill_offer_specifics_proto_depIdxs,
		MessageInfos:      file_autofill_offer_specifics_proto_msgTypes,
	}.Build()
	File_autofill_offer_specifics_proto = out.File
	file_autofill_offer_specifics_proto_rawDesc = nil
	file_autofill_offer_specifics_proto_goTypes = nil
	file_autofill_offer_specifics_proto_depIdxs = nil
}
