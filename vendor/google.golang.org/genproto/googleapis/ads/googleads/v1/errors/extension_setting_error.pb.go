// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/extension_setting_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v1/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible extension setting errors.
type ExtensionSettingErrorEnum_ExtensionSettingError int32

const (
	// Enum unspecified.
	ExtensionSettingErrorEnum_UNSPECIFIED ExtensionSettingErrorEnum_ExtensionSettingError = 0
	// The received error code is not known in this version.
	ExtensionSettingErrorEnum_UNKNOWN ExtensionSettingErrorEnum_ExtensionSettingError = 1
	// A platform restriction was provided without input extensions or existing
	// extensions.
	ExtensionSettingErrorEnum_EXTENSIONS_REQUIRED ExtensionSettingErrorEnum_ExtensionSettingError = 2
	// The provided feed type does not correspond to the provided extensions.
	ExtensionSettingErrorEnum_FEED_TYPE_EXTENSION_TYPE_MISMATCH ExtensionSettingErrorEnum_ExtensionSettingError = 3
	// The provided feed type cannot be used.
	ExtensionSettingErrorEnum_INVALID_FEED_TYPE ExtensionSettingErrorEnum_ExtensionSettingError = 4
	// The provided feed type cannot be used at the customer level.
	ExtensionSettingErrorEnum_INVALID_FEED_TYPE_FOR_CUSTOMER_EXTENSION_SETTING ExtensionSettingErrorEnum_ExtensionSettingError = 5
	// Cannot change a feed item field on a CREATE operation.
	ExtensionSettingErrorEnum_CANNOT_CHANGE_FEED_ITEM_ON_CREATE ExtensionSettingErrorEnum_ExtensionSettingError = 6
	// Cannot update an extension that is not already in this setting.
	ExtensionSettingErrorEnum_CANNOT_UPDATE_NEWLY_CREATED_EXTENSION ExtensionSettingErrorEnum_ExtensionSettingError = 7
	// There is no existing AdGroupExtensionSetting for this type.
	ExtensionSettingErrorEnum_NO_EXISTING_AD_GROUP_EXTENSION_SETTING_FOR_TYPE ExtensionSettingErrorEnum_ExtensionSettingError = 8
	// There is no existing CampaignExtensionSetting for this type.
	ExtensionSettingErrorEnum_NO_EXISTING_CAMPAIGN_EXTENSION_SETTING_FOR_TYPE ExtensionSettingErrorEnum_ExtensionSettingError = 9
	// There is no existing CustomerExtensionSetting for this type.
	ExtensionSettingErrorEnum_NO_EXISTING_CUSTOMER_EXTENSION_SETTING_FOR_TYPE ExtensionSettingErrorEnum_ExtensionSettingError = 10
	// The AdGroupExtensionSetting already exists. UPDATE should be used to
	// modify the existing AdGroupExtensionSetting.
	ExtensionSettingErrorEnum_AD_GROUP_EXTENSION_SETTING_ALREADY_EXISTS ExtensionSettingErrorEnum_ExtensionSettingError = 11
	// The CampaignExtensionSetting already exists. UPDATE should be used to
	// modify the existing CampaignExtensionSetting.
	ExtensionSettingErrorEnum_CAMPAIGN_EXTENSION_SETTING_ALREADY_EXISTS ExtensionSettingErrorEnum_ExtensionSettingError = 12
	// The CustomerExtensionSetting already exists. UPDATE should be used to
	// modify the existing CustomerExtensionSetting.
	ExtensionSettingErrorEnum_CUSTOMER_EXTENSION_SETTING_ALREADY_EXISTS ExtensionSettingErrorEnum_ExtensionSettingError = 13
	// An active ad group feed already exists for this place holder type.
	ExtensionSettingErrorEnum_AD_GROUP_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE ExtensionSettingErrorEnum_ExtensionSettingError = 14
	// An active campaign feed already exists for this place holder type.
	ExtensionSettingErrorEnum_CAMPAIGN_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE ExtensionSettingErrorEnum_ExtensionSettingError = 15
	// An active customer feed already exists for this place holder type.
	ExtensionSettingErrorEnum_CUSTOMER_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE ExtensionSettingErrorEnum_ExtensionSettingError = 16
	// Value is not within the accepted range.
	ExtensionSettingErrorEnum_VALUE_OUT_OF_RANGE ExtensionSettingErrorEnum_ExtensionSettingError = 17
	// Cannot simultaneously set specified field with final urls.
	ExtensionSettingErrorEnum_CANNOT_SET_FIELD_WITH_FINAL_URLS ExtensionSettingErrorEnum_ExtensionSettingError = 18
	// Must set field with final urls.
	ExtensionSettingErrorEnum_FINAL_URLS_NOT_SET ExtensionSettingErrorEnum_ExtensionSettingError = 19
	// Phone number for a call extension is invalid.
	ExtensionSettingErrorEnum_INVALID_PHONE_NUMBER ExtensionSettingErrorEnum_ExtensionSettingError = 20
	// Phone number for a call extension is not supported for the given country
	// code.
	ExtensionSettingErrorEnum_PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY ExtensionSettingErrorEnum_ExtensionSettingError = 21
	// A carrier specific number in short format is not allowed for call
	// extensions.
	ExtensionSettingErrorEnum_CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED ExtensionSettingErrorEnum_ExtensionSettingError = 22
	// Premium rate numbers are not allowed for call extensions.
	ExtensionSettingErrorEnum_PREMIUM_RATE_NUMBER_NOT_ALLOWED ExtensionSettingErrorEnum_ExtensionSettingError = 23
	// Phone number type for a call extension is not allowed.
	ExtensionSettingErrorEnum_DISALLOWED_NUMBER_TYPE ExtensionSettingErrorEnum_ExtensionSettingError = 24
	// Phone number for a call extension does not meet domestic format
	// requirements.
	ExtensionSettingErrorEnum_INVALID_DOMESTIC_PHONE_NUMBER_FORMAT ExtensionSettingErrorEnum_ExtensionSettingError = 25
	// Vanity phone numbers (i.e. those including letters) are not allowed for
	// call extensions.
	ExtensionSettingErrorEnum_VANITY_PHONE_NUMBER_NOT_ALLOWED ExtensionSettingErrorEnum_ExtensionSettingError = 26
	// Country code provided for a call extension is invalid.
	ExtensionSettingErrorEnum_INVALID_COUNTRY_CODE ExtensionSettingErrorEnum_ExtensionSettingError = 27
	// Call conversion type id provided for a call extension is invalid.
	ExtensionSettingErrorEnum_INVALID_CALL_CONVERSION_TYPE_ID ExtensionSettingErrorEnum_ExtensionSettingError = 28
	// For a call extension, the customer is not whitelisted for call tracking.
	ExtensionSettingErrorEnum_CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING ExtensionSettingErrorEnum_ExtensionSettingError = 29
	// Call tracking is not supported for the given country for a call
	// extension.
	ExtensionSettingErrorEnum_CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY ExtensionSettingErrorEnum_ExtensionSettingError = 30
	// App id provided for an app extension is invalid.
	ExtensionSettingErrorEnum_INVALID_APP_ID ExtensionSettingErrorEnum_ExtensionSettingError = 31
	// Quotation marks present in the review text for a review extension.
	ExtensionSettingErrorEnum_QUOTES_IN_REVIEW_EXTENSION_SNIPPET ExtensionSettingErrorEnum_ExtensionSettingError = 32
	// Hyphen character present in the review text for a review extension.
	ExtensionSettingErrorEnum_HYPHENS_IN_REVIEW_EXTENSION_SNIPPET ExtensionSettingErrorEnum_ExtensionSettingError = 33
	// A blacklisted review source name or url was provided for a review
	// extension.
	ExtensionSettingErrorEnum_REVIEW_EXTENSION_SOURCE_NOT_ELIGIBLE ExtensionSettingErrorEnum_ExtensionSettingError = 34
	// Review source name should not be found in the review text.
	ExtensionSettingErrorEnum_SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT ExtensionSettingErrorEnum_ExtensionSettingError = 35
	// Field must be set.
	ExtensionSettingErrorEnum_MISSING_FIELD ExtensionSettingErrorEnum_ExtensionSettingError = 36
	// Inconsistent currency codes.
	ExtensionSettingErrorEnum_INCONSISTENT_CURRENCY_CODES ExtensionSettingErrorEnum_ExtensionSettingError = 37
	// Price extension cannot have duplicated headers.
	ExtensionSettingErrorEnum_PRICE_EXTENSION_HAS_DUPLICATED_HEADERS ExtensionSettingErrorEnum_ExtensionSettingError = 38
	// Price item cannot have duplicated header and description.
	ExtensionSettingErrorEnum_PRICE_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION ExtensionSettingErrorEnum_ExtensionSettingError = 39
	// Price extension has too few items
	ExtensionSettingErrorEnum_PRICE_EXTENSION_HAS_TOO_FEW_ITEMS ExtensionSettingErrorEnum_ExtensionSettingError = 40
	// Price extension has too many items
	ExtensionSettingErrorEnum_PRICE_EXTENSION_HAS_TOO_MANY_ITEMS ExtensionSettingErrorEnum_ExtensionSettingError = 41
	// The input value is not currently supported.
	ExtensionSettingErrorEnum_UNSUPPORTED_VALUE ExtensionSettingErrorEnum_ExtensionSettingError = 42
	// Unknown or unsupported device preference.
	ExtensionSettingErrorEnum_INVALID_DEVICE_PREFERENCE ExtensionSettingErrorEnum_ExtensionSettingError = 43
	// Invalid feed item schedule end time (i.e., endHour = 24 and
	// endMinute != 0).
	ExtensionSettingErrorEnum_INVALID_SCHEDULE_END ExtensionSettingErrorEnum_ExtensionSettingError = 45
	// Date time zone does not match the account's time zone.
	ExtensionSettingErrorEnum_DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE ExtensionSettingErrorEnum_ExtensionSettingError = 47
	// Overlapping feed item schedule times (e.g., 7-10AM and 8-11AM) are not
	// allowed.
	ExtensionSettingErrorEnum_OVERLAPPING_SCHEDULES_NOT_ALLOWED ExtensionSettingErrorEnum_ExtensionSettingError = 48
	// Feed item schedule end time must be after start time.
	ExtensionSettingErrorEnum_SCHEDULE_END_NOT_AFTER_START ExtensionSettingErrorEnum_ExtensionSettingError = 49
	// There are too many feed item schedules per day.
	ExtensionSettingErrorEnum_TOO_MANY_SCHEDULES_PER_DAY ExtensionSettingErrorEnum_ExtensionSettingError = 50
	// Cannot edit the same extension feed item more than once in the same
	// request.
	ExtensionSettingErrorEnum_DUPLICATE_EXTENSION_FEED_ITEM_EDIT ExtensionSettingErrorEnum_ExtensionSettingError = 51
	// Invalid structured snippet header.
	ExtensionSettingErrorEnum_INVALID_SNIPPETS_HEADER ExtensionSettingErrorEnum_ExtensionSettingError = 52
	// Phone number with call tracking enabled is not supported for the
	// specified country.
	ExtensionSettingErrorEnum_PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY ExtensionSettingErrorEnum_ExtensionSettingError = 53
	// The targeted adgroup must belong to the targeted campaign.
	ExtensionSettingErrorEnum_CAMPAIGN_TARGETING_MISMATCH ExtensionSettingErrorEnum_ExtensionSettingError = 54
	// The feed used by the ExtensionSetting is removed and cannot be operated
	// on. Remove the ExtensionSetting to allow a new one to be created using
	// an active feed.
	ExtensionSettingErrorEnum_CANNOT_OPERATE_ON_REMOVED_FEED ExtensionSettingErrorEnum_ExtensionSettingError = 55
	// The ExtensionFeedItem type is required for this operation.
	ExtensionSettingErrorEnum_EXTENSION_TYPE_REQUIRED ExtensionSettingErrorEnum_ExtensionSettingError = 56
	// The matching function that links the extension feed to the customer,
	// campaign, or ad group is not compatible with the ExtensionSetting
	// services.
	ExtensionSettingErrorEnum_INCOMPATIBLE_UNDERLYING_MATCHING_FUNCTION ExtensionSettingErrorEnum_ExtensionSettingError = 57
	// Start date must be before end date.
	ExtensionSettingErrorEnum_START_DATE_AFTER_END_DATE ExtensionSettingErrorEnum_ExtensionSettingError = 58
	// Input price is not in a valid format.
	ExtensionSettingErrorEnum_INVALID_PRICE_FORMAT ExtensionSettingErrorEnum_ExtensionSettingError = 59
	// The promotion time is invalid.
	ExtensionSettingErrorEnum_PROMOTION_INVALID_TIME ExtensionSettingErrorEnum_ExtensionSettingError = 60
	// Cannot set both percent discount and money discount fields.
	ExtensionSettingErrorEnum_PROMOTION_CANNOT_SET_PERCENT_DISCOUNT_AND_MONEY_DISCOUNT ExtensionSettingErrorEnum_ExtensionSettingError = 61
	// Cannot set both promotion code and orders over amount fields.
	ExtensionSettingErrorEnum_PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT ExtensionSettingErrorEnum_ExtensionSettingError = 62
	// This field has too many decimal places specified.
	ExtensionSettingErrorEnum_TOO_MANY_DECIMAL_PLACES_SPECIFIED ExtensionSettingErrorEnum_ExtensionSettingError = 63
	// The language code is not valid.
	ExtensionSettingErrorEnum_INVALID_LANGUAGE_CODE ExtensionSettingErrorEnum_ExtensionSettingError = 64
	// The language is not supported.
	ExtensionSettingErrorEnum_UNSUPPORTED_LANGUAGE ExtensionSettingErrorEnum_ExtensionSettingError = 65
	// Customer hasn't consented for call recording, which is required for
	// adding/updating call extensions. Please see
	// https://support.google.com/google-ads/answer/7412639.
	ExtensionSettingErrorEnum_CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED ExtensionSettingErrorEnum_ExtensionSettingError = 66
	// The UPDATE operation does not specify any fields other than the resource
	// name in the update mask.
	ExtensionSettingErrorEnum_EXTENSION_SETTING_UPDATE_IS_A_NOOP ExtensionSettingErrorEnum_ExtensionSettingError = 67
)

var ExtensionSettingErrorEnum_ExtensionSettingError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "EXTENSIONS_REQUIRED",
	3:  "FEED_TYPE_EXTENSION_TYPE_MISMATCH",
	4:  "INVALID_FEED_TYPE",
	5:  "INVALID_FEED_TYPE_FOR_CUSTOMER_EXTENSION_SETTING",
	6:  "CANNOT_CHANGE_FEED_ITEM_ON_CREATE",
	7:  "CANNOT_UPDATE_NEWLY_CREATED_EXTENSION",
	8:  "NO_EXISTING_AD_GROUP_EXTENSION_SETTING_FOR_TYPE",
	9:  "NO_EXISTING_CAMPAIGN_EXTENSION_SETTING_FOR_TYPE",
	10: "NO_EXISTING_CUSTOMER_EXTENSION_SETTING_FOR_TYPE",
	11: "AD_GROUP_EXTENSION_SETTING_ALREADY_EXISTS",
	12: "CAMPAIGN_EXTENSION_SETTING_ALREADY_EXISTS",
	13: "CUSTOMER_EXTENSION_SETTING_ALREADY_EXISTS",
	14: "AD_GROUP_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	15: "CAMPAIGN_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	16: "CUSTOMER_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	17: "VALUE_OUT_OF_RANGE",
	18: "CANNOT_SET_FIELD_WITH_FINAL_URLS",
	19: "FINAL_URLS_NOT_SET",
	20: "INVALID_PHONE_NUMBER",
	21: "PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY",
	22: "CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED",
	23: "PREMIUM_RATE_NUMBER_NOT_ALLOWED",
	24: "DISALLOWED_NUMBER_TYPE",
	25: "INVALID_DOMESTIC_PHONE_NUMBER_FORMAT",
	26: "VANITY_PHONE_NUMBER_NOT_ALLOWED",
	27: "INVALID_COUNTRY_CODE",
	28: "INVALID_CALL_CONVERSION_TYPE_ID",
	29: "CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING",
	30: "CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY",
	31: "INVALID_APP_ID",
	32: "QUOTES_IN_REVIEW_EXTENSION_SNIPPET",
	33: "HYPHENS_IN_REVIEW_EXTENSION_SNIPPET",
	34: "REVIEW_EXTENSION_SOURCE_NOT_ELIGIBLE",
	35: "SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT",
	36: "MISSING_FIELD",
	37: "INCONSISTENT_CURRENCY_CODES",
	38: "PRICE_EXTENSION_HAS_DUPLICATED_HEADERS",
	39: "PRICE_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION",
	40: "PRICE_EXTENSION_HAS_TOO_FEW_ITEMS",
	41: "PRICE_EXTENSION_HAS_TOO_MANY_ITEMS",
	42: "UNSUPPORTED_VALUE",
	43: "INVALID_DEVICE_PREFERENCE",
	45: "INVALID_SCHEDULE_END",
	47: "DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE",
	48: "OVERLAPPING_SCHEDULES_NOT_ALLOWED",
	49: "SCHEDULE_END_NOT_AFTER_START",
	50: "TOO_MANY_SCHEDULES_PER_DAY",
	51: "DUPLICATE_EXTENSION_FEED_ITEM_EDIT",
	52: "INVALID_SNIPPETS_HEADER",
	53: "PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY",
	54: "CAMPAIGN_TARGETING_MISMATCH",
	55: "CANNOT_OPERATE_ON_REMOVED_FEED",
	56: "EXTENSION_TYPE_REQUIRED",
	57: "INCOMPATIBLE_UNDERLYING_MATCHING_FUNCTION",
	58: "START_DATE_AFTER_END_DATE",
	59: "INVALID_PRICE_FORMAT",
	60: "PROMOTION_INVALID_TIME",
	61: "PROMOTION_CANNOT_SET_PERCENT_DISCOUNT_AND_MONEY_DISCOUNT",
	62: "PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT",
	63: "TOO_MANY_DECIMAL_PLACES_SPECIFIED",
	64: "INVALID_LANGUAGE_CODE",
	65: "UNSUPPORTED_LANGUAGE",
	66: "CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED",
	67: "EXTENSION_SETTING_UPDATE_IS_A_NOOP",
}
var ExtensionSettingErrorEnum_ExtensionSettingError_value = map[string]int32{
	"UNSPECIFIED":                       0,
	"UNKNOWN":                           1,
	"EXTENSIONS_REQUIRED":               2,
	"FEED_TYPE_EXTENSION_TYPE_MISMATCH": 3,
	"INVALID_FEED_TYPE":                 4,
	"INVALID_FEED_TYPE_FOR_CUSTOMER_EXTENSION_SETTING":           5,
	"CANNOT_CHANGE_FEED_ITEM_ON_CREATE":                          6,
	"CANNOT_UPDATE_NEWLY_CREATED_EXTENSION":                      7,
	"NO_EXISTING_AD_GROUP_EXTENSION_SETTING_FOR_TYPE":            8,
	"NO_EXISTING_CAMPAIGN_EXTENSION_SETTING_FOR_TYPE":            9,
	"NO_EXISTING_CUSTOMER_EXTENSION_SETTING_FOR_TYPE":            10,
	"AD_GROUP_EXTENSION_SETTING_ALREADY_EXISTS":                  11,
	"CAMPAIGN_EXTENSION_SETTING_ALREADY_EXISTS":                  12,
	"CUSTOMER_EXTENSION_SETTING_ALREADY_EXISTS":                  13,
	"AD_GROUP_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":          14,
	"CAMPAIGN_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":          15,
	"CUSTOMER_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":          16,
	"VALUE_OUT_OF_RANGE":                                         17,
	"CANNOT_SET_FIELD_WITH_FINAL_URLS":                           18,
	"FINAL_URLS_NOT_SET":                                         19,
	"INVALID_PHONE_NUMBER":                                       20,
	"PHONE_NUMBER_NOT_SUPPORTED_FOR_COUNTRY":                     21,
	"CARRIER_SPECIFIC_SHORT_NUMBER_NOT_ALLOWED":                  22,
	"PREMIUM_RATE_NUMBER_NOT_ALLOWED":                            23,
	"DISALLOWED_NUMBER_TYPE":                                     24,
	"INVALID_DOMESTIC_PHONE_NUMBER_FORMAT":                       25,
	"VANITY_PHONE_NUMBER_NOT_ALLOWED":                            26,
	"INVALID_COUNTRY_CODE":                                       27,
	"INVALID_CALL_CONVERSION_TYPE_ID":                            28,
	"CUSTOMER_NOT_WHITELISTED_FOR_CALLTRACKING":                  29,
	"CALLTRACKING_NOT_SUPPORTED_FOR_COUNTRY":                     30,
	"INVALID_APP_ID":                                             31,
	"QUOTES_IN_REVIEW_EXTENSION_SNIPPET":                         32,
	"HYPHENS_IN_REVIEW_EXTENSION_SNIPPET":                        33,
	"REVIEW_EXTENSION_SOURCE_NOT_ELIGIBLE":                       34,
	"SOURCE_NAME_IN_REVIEW_EXTENSION_TEXT":                       35,
	"MISSING_FIELD":                                              36,
	"INCONSISTENT_CURRENCY_CODES":                                37,
	"PRICE_EXTENSION_HAS_DUPLICATED_HEADERS":                     38,
	"PRICE_ITEM_HAS_DUPLICATED_HEADER_AND_DESCRIPTION":           39,
	"PRICE_EXTENSION_HAS_TOO_FEW_ITEMS":                          40,
	"PRICE_EXTENSION_HAS_TOO_MANY_ITEMS":                         41,
	"UNSUPPORTED_VALUE":                                          42,
	"INVALID_DEVICE_PREFERENCE":                                  43,
	"INVALID_SCHEDULE_END":                                       45,
	"DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE":                     47,
	"OVERLAPPING_SCHEDULES_NOT_ALLOWED":                          48,
	"SCHEDULE_END_NOT_AFTER_START":                               49,
	"TOO_MANY_SCHEDULES_PER_DAY":                                 50,
	"DUPLICATE_EXTENSION_FEED_ITEM_EDIT":                         51,
	"INVALID_SNIPPETS_HEADER":                                    52,
	"PHONE_NUMBER_NOT_SUPPORTED_WITH_CALLTRACKING_FOR_COUNTRY":   53,
	"CAMPAIGN_TARGETING_MISMATCH":                                54,
	"CANNOT_OPERATE_ON_REMOVED_FEED":                             55,
	"EXTENSION_TYPE_REQUIRED":                                    56,
	"INCOMPATIBLE_UNDERLYING_MATCHING_FUNCTION":                  57,
	"START_DATE_AFTER_END_DATE":                                  58,
	"INVALID_PRICE_FORMAT":                                       59,
	"PROMOTION_INVALID_TIME":                                     60,
	"PROMOTION_CANNOT_SET_PERCENT_DISCOUNT_AND_MONEY_DISCOUNT":   61,
	"PROMOTION_CANNOT_SET_PROMOTION_CODE_AND_ORDERS_OVER_AMOUNT": 62,
	"TOO_MANY_DECIMAL_PLACES_SPECIFIED":                          63,
	"INVALID_LANGUAGE_CODE":                                      64,
	"UNSUPPORTED_LANGUAGE":                                       65,
	"CUSTOMER_CONSENT_FOR_CALL_RECORDING_REQUIRED":               66,
	"EXTENSION_SETTING_UPDATE_IS_A_NOOP":                         67,
}

func (x ExtensionSettingErrorEnum_ExtensionSettingError) String() string {
	return proto.EnumName(ExtensionSettingErrorEnum_ExtensionSettingError_name, int32(x))
}
func (ExtensionSettingErrorEnum_ExtensionSettingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_extension_setting_error_36d74918ab80bb6d, []int{0, 0}
}

// Container for enum describing validation errors of extension settings.
type ExtensionSettingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtensionSettingErrorEnum) Reset()         { *m = ExtensionSettingErrorEnum{} }
func (m *ExtensionSettingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ExtensionSettingErrorEnum) ProtoMessage()    {}
func (*ExtensionSettingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_extension_setting_error_36d74918ab80bb6d, []int{0}
}
func (m *ExtensionSettingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionSettingErrorEnum.Unmarshal(m, b)
}
func (m *ExtensionSettingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionSettingErrorEnum.Marshal(b, m, deterministic)
}
func (dst *ExtensionSettingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionSettingErrorEnum.Merge(dst, src)
}
func (m *ExtensionSettingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ExtensionSettingErrorEnum.Size(m)
}
func (m *ExtensionSettingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionSettingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionSettingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtensionSettingErrorEnum)(nil), "google.ads.googleads.v1.errors.ExtensionSettingErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v1.errors.ExtensionSettingErrorEnum_ExtensionSettingError", ExtensionSettingErrorEnum_ExtensionSettingError_name, ExtensionSettingErrorEnum_ExtensionSettingError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/extension_setting_error.proto", fileDescriptor_extension_setting_error_36d74918ab80bb6d)
}

var fileDescriptor_extension_setting_error_36d74918ab80bb6d = []byte{
	// 1313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x5d, 0x73, 0x53, 0xb7,
	0x16, 0xbd, 0x84, 0x7b, 0xe1, 0x5e, 0x71, 0x01, 0x21, 0x08, 0x90, 0x00, 0x01, 0xc2, 0x37, 0x2d,
	0x36, 0x29, 0xd0, 0x52, 0x43, 0x69, 0x95, 0xa3, 0x6d, 0x5b, 0xc3, 0x39, 0x92, 0x90, 0x74, 0x1c,
	0xcc, 0x64, 0x46, 0x93, 0x36, 0x99, 0x4c, 0x66, 0xc0, 0x66, 0xe2, 0x94, 0xe9, 0x4f, 0xe9, 0x73,
	0x1f, 0xfb, 0x53, 0xfa, 0x53, 0xfa, 0xde, 0x97, 0x3e, 0x75, 0xb6, 0xce, 0x87, 0x9d, 0x38, 0xf1,
	0xf0, 0x64, 0x1f, 0x69, 0xad, 0xbd, 0xb7, 0xb6, 0x96, 0x96, 0x44, 0x5e, 0x6e, 0x0f, 0x87, 0xdb,
	0xef, 0xb7, 0x9a, 0x1b, 0x9b, 0xa3, 0x66, 0xf1, 0x17, 0xff, 0x7d, 0x5a, 0x69, 0x6e, 0xed, 0xee,
	0x0e, 0x77, 0x47, 0xcd, 0xad, 0x5f, 0xf6, 0xb6, 0x06, 0xa3, 0x9d, 0xe1, 0x20, 0x8c, 0xb6, 0xf6,
	0xf6, 0x76, 0x06, 0xdb, 0x21, 0x4e, 0x34, 0x3e, 0xee, 0x0e, 0xf7, 0x86, 0x6c, 0xa9, 0xa0, 0x34,
	0x36, 0x36, 0x47, 0x8d, 0x9a, 0xdd, 0xf8, 0xb4, 0xd2, 0x28, 0xd8, 0x8b, 0x57, 0xab, 0xe8, 0x1f,
	0x77, 0x9a, 0x1b, 0x83, 0xc1, 0x70, 0x6f, 0x63, 0x6f, 0x67, 0x38, 0x18, 0x15, 0xec, 0xe5, 0x5f,
	0x2f, 0x90, 0x05, 0xa8, 0xe2, 0xbb, 0x22, 0x3c, 0x20, 0x11, 0x06, 0x3f, 0x7f, 0x58, 0xfe, 0xfb,
	0x3c, 0x99, 0x3f, 0x74, 0x96, 0x9d, 0x25, 0xa7, 0x72, 0xe5, 0x0c, 0x24, 0xb2, 0x2d, 0x41, 0xd0,
	0x7f, 0xb1, 0x53, 0xe4, 0x64, 0xae, 0x5e, 0x2b, 0xbd, 0xa6, 0xe8, 0x31, 0x76, 0x89, 0x9c, 0x87,
	0xb7, 0x1e, 0x94, 0x93, 0x5a, 0xb9, 0x60, 0xe1, 0x4d, 0x2e, 0x2d, 0x08, 0x3a, 0xc7, 0xee, 0x90,
	0x9b, 0x6d, 0x00, 0x11, 0x7c, 0xdf, 0x40, 0xa8, 0x21, 0xc5, 0x67, 0x26, 0x5d, 0xc6, 0x7d, 0xd2,
	0xa5, 0xc7, 0xd9, 0x3c, 0x39, 0x27, 0x55, 0x8f, 0xa7, 0x52, 0x84, 0x1a, 0x4e, 0xff, 0xcd, 0x9e,
	0x92, 0xc7, 0x53, 0xc3, 0xa1, 0xad, 0x6d, 0x48, 0x72, 0xe7, 0x75, 0x06, 0x76, 0x22, 0xa4, 0x03,
	0xef, 0xa5, 0xea, 0xd0, 0xff, 0x60, 0xce, 0x84, 0x2b, 0xa5, 0x7d, 0x48, 0xba, 0x5c, 0x75, 0xa0,
	0xe0, 0x4a, 0x0f, 0x59, 0xd0, 0x2a, 0x24, 0x16, 0xb8, 0x07, 0x7a, 0x82, 0x3d, 0x20, 0x77, 0x4a,
	0x58, 0x6e, 0x04, 0xf7, 0x10, 0x14, 0xac, 0xa5, 0xfd, 0x72, 0x5e, 0x8c, 0x23, 0xd3, 0x93, 0xec,
	0x09, 0x69, 0x2a, 0x1d, 0xe0, 0xad, 0x74, 0x98, 0x22, 0x70, 0x11, 0x3a, 0x56, 0xe7, 0x66, 0x3a,
	0x7b, 0x2c, 0x2e, 0x16, 0xff, 0xdf, 0x83, 0xa4, 0x84, 0x67, 0x86, 0xcb, 0x8e, 0x9a, 0x45, 0xfa,
	0xdf, 0x14, 0xe9, 0xc8, 0x75, 0x8e, 0x49, 0x84, 0x3d, 0x22, 0x0f, 0x66, 0x94, 0xc4, 0x53, 0x0b,
	0x5c, 0xf4, 0x8b, 0xa0, 0x8e, 0x9e, 0x42, 0xf8, 0x8c, 0x62, 0x0e, 0xc0, 0xff, 0x1f, 0xe1, 0x47,
	0x97, 0x71, 0x00, 0x7e, 0x9a, 0x3d, 0x23, 0x2b, 0x75, 0x31, 0xb1, 0xf1, 0xfb, 0x11, 0xb1, 0x6e,
	0x93, 0xf2, 0x04, 0xba, 0x3a, 0x15, 0x50, 0xae, 0xe1, 0x0c, 0xd2, 0xea, 0xa2, 0x3e, 0x9b, 0x76,
	0x36, 0xd2, 0xaa, 0xe2, 0x3e, 0x9b, 0x46, 0xd9, 0x45, 0xc2, 0x7a, 0x3c, 0xcd, 0x21, 0xe8, 0xdc,
	0x07, 0xdd, 0x0e, 0x16, 0x75, 0x42, 0xcf, 0xb1, 0xdb, 0xe4, 0x46, 0xa9, 0x09, 0x07, 0x3e, 0xb4,
	0x25, 0xa4, 0x22, 0xac, 0x49, 0xdf, 0x0d, 0x6d, 0xa9, 0x78, 0x1a, 0x72, 0x9b, 0x3a, 0xca, 0x90,
	0x3d, 0xfe, 0x0e, 0x25, 0x9a, 0x9e, 0x67, 0x97, 0xc9, 0x85, 0x4a, 0xae, 0xa6, 0xab, 0x15, 0x04,
	0x95, 0x67, 0xab, 0x60, 0xe9, 0x05, 0xf6, 0x90, 0xdc, 0x9d, 0x1c, 0x29, 0x38, 0xb9, 0x31, 0xda,
	0xa2, 0xd4, 0xa2, 0xa2, 0x75, 0xae, 0xbc, 0xed, 0xd3, 0xf9, 0x62, 0x7b, 0xac, 0x95, 0x60, 0x43,
	0x79, 0xde, 0x92, 0xe0, 0xba, 0xda, 0xfa, 0x49, 0x32, 0x4f, 0x53, 0xbd, 0x06, 0x82, 0x5e, 0x64,
	0xb7, 0xc8, 0x75, 0x63, 0x21, 0x93, 0x79, 0x16, 0x6c, 0x54, 0xf1, 0x34, 0xe8, 0x12, 0x5b, 0x24,
	0x17, 0x85, 0x74, 0xe5, 0x77, 0x05, 0x89, 0xbd, 0xb8, 0xcc, 0xee, 0x93, 0xdb, 0x55, 0xd5, 0x42,
	0x67, 0xe0, 0xbc, 0x4c, 0xf6, 0x95, 0x8f, 0xe5, 0x65, 0xdc, 0xd3, 0x05, 0x4c, 0xd5, 0xe3, 0x4a,
	0xfa, 0x7e, 0x98, 0x5a, 0x4c, 0x95, 0x6a, 0x71, 0xb2, 0x09, 0xe5, 0x9a, 0x42, 0xa2, 0x05, 0xd0,
	0x2b, 0x48, 0xaf, 0x67, 0x78, 0x9a, 0x86, 0x44, 0xab, 0x1e, 0xd8, 0xb1, 0x1f, 0x48, 0x41, 0xaf,
	0xee, 0x53, 0x1b, 0x06, 0x5e, 0xeb, 0x4a, 0x0f, 0xa9, 0x74, 0x75, 0x9f, 0x78, 0x9a, 0x7a, 0xcb,
	0x93, 0xd7, 0x78, 0xd6, 0xaf, 0x61, 0x63, 0x27, 0x47, 0x66, 0x34, 0x76, 0x89, 0x31, 0x72, 0xa6,
	0xca, 0xcf, 0x8d, 0xc1, 0x74, 0xd7, 0xd9, 0x5d, 0xb2, 0xfc, 0x26, 0xd7, 0x1e, 0x5c, 0x90, 0x2a,
	0x58, 0xe8, 0x49, 0x58, 0x9b, 0x14, 0xb9, 0x92, 0xc6, 0x80, 0xa7, 0x37, 0xd8, 0x3d, 0x72, 0xab,
	0xdb, 0x37, 0x5d, 0x50, 0xb3, 0x81, 0x37, 0xb1, 0x9b, 0xd3, 0xb3, 0x3a, 0xb7, 0x09, 0xc4, 0xda,
	0x20, 0x95, 0x1d, 0xb9, 0x9a, 0x02, 0x5d, 0x46, 0x64, 0x35, 0xc1, 0x33, 0x38, 0x34, 0xac, 0x87,
	0xb7, 0x9e, 0xde, 0x62, 0xe7, 0xc8, 0xe9, 0x4c, 0x3a, 0x17, 0x4f, 0x3d, 0x4a, 0x92, 0xde, 0x66,
	0xd7, 0xc9, 0x15, 0xa9, 0x12, 0xad, 0x1c, 0x76, 0x46, 0xf9, 0x90, 0xe4, 0xd6, 0x82, 0x4a, 0x8a,
	0x5e, 0x3b, 0x7a, 0x27, 0x2a, 0xce, 0xca, 0x64, 0xd2, 0x74, 0xbb, 0xdc, 0x05, 0x91, 0x9b, 0x54,
	0x26, 0xd1, 0xe0, 0xba, 0xc0, 0x05, 0x58, 0x47, 0xef, 0xa2, 0xcd, 0x16, 0xd8, 0xe8, 0x91, 0x87,
	0xc2, 0x02, 0x57, 0x22, 0x08, 0x70, 0x89, 0x95, 0xc6, 0xa3, 0x29, 0xde, 0x43, 0x9b, 0x3d, 0x2c,
	0x83, 0xd7, 0x3a, 0xb4, 0x61, 0x2d, 0xc6, 0x72, 0xf4, 0x3e, 0x76, 0xf8, 0x28, 0x58, 0xc6, 0x55,
	0xbf, 0xc4, 0x3d, 0xc0, 0x2b, 0x20, 0x57, 0xe3, 0xad, 0x8b, 0xc7, 0x93, 0x3e, 0x64, 0xd7, 0xc8,
	0x42, 0xad, 0x4e, 0xe8, 0x61, 0x1c, 0x63, 0xa1, 0x0d, 0xb8, 0x58, 0xa0, 0x5f, 0x4c, 0xaa, 0xcd,
	0x25, 0x5d, 0x10, 0x79, 0x0a, 0x01, 0x94, 0xa0, 0x8f, 0xb0, 0x01, 0xd1, 0xd5, 0xbd, 0xcc, 0x20,
	0x64, 0xb9, 0xf3, 0x61, 0x35, 0x36, 0x99, 0x27, 0x51, 0x11, 0xc5, 0xf8, 0x3b, 0xad, 0x80, 0x36,
	0x71, 0x29, 0xba, 0x07, 0x36, 0xe5, 0xc6, 0x60, 0x93, 0xab, 0x48, 0x6e, 0x9f, 0xb4, 0x1f, 0xb3,
	0x1b, 0xe4, 0xea, 0x64, 0x92, 0x62, 0xb6, 0xed, 0xf1, 0xa0, 0x7a, 0x6e, 0x3d, 0x5d, 0x61, 0x4b,
	0x64, 0xb1, 0x5e, 0xd8, 0x38, 0x8a, 0x01, 0x1b, 0x04, 0xef, 0xd3, 0xaf, 0xb0, 0x19, 0x75, 0x6b,
	0x27, 0x1a, 0x32, 0xbe, 0xa0, 0x40, 0x48, 0x4f, 0x9f, 0xb0, 0x2b, 0xe4, 0x52, 0xbd, 0xac, 0x42,
	0x5a, 0xae, 0xdc, 0x08, 0xfa, 0x94, 0xbd, 0x24, 0xcf, 0x67, 0x98, 0x49, 0xb4, 0xab, 0x7d, 0x67,
	0x62, 0xf2, 0x14, 0x3c, 0x43, 0xe5, 0xd4, 0x46, 0xeb, 0xb9, 0xed, 0x40, 0xb4, 0xf1, 0xfa, 0x2e,
	0xfe, 0x9a, 0x2d, 0x93, 0xa5, 0xd2, 0x03, 0xb5, 0x81, 0x68, 0x29, 0x1a, 0xa5, 0x99, 0xe9, 0x1e,
	0x14, 0xd7, 0x30, 0xfd, 0x06, 0xeb, 0x3b, 0x70, 0x99, 0xd7, 0x77, 0xfe, 0x73, 0x3c, 0xc2, 0xa8,
	0xcd, 0xcc, 0x70, 0x8f, 0x52, 0x0f, 0xb9, 0x12, 0x60, 0xd3, 0x7e, 0x4c, 0x83, 0x39, 0x62, 0x4d,
	0xb9, 0x4a, 0xa2, 0x8e, 0xbe, 0xc5, 0x1d, 0x8e, 0xed, 0x0b, 0x71, 0xbb, 0x8a, 0x7e, 0x62, 0x77,
	0xf1, 0x93, 0xb6, 0xf6, 0x99, 0x6a, 0xd4, 0x51, 0x69, 0x47, 0x2f, 0xd0, 0xd4, 0x8c, 0xd5, 0x99,
	0xc6, 0x38, 0xa1, 0xc2, 0xe0, 0xb6, 0xd2, 0x97, 0xb1, 0x47, 0xf5, 0xdc, 0x84, 0xa5, 0x1b, 0xb0,
	0x09, 0x9e, 0x17, 0x21, 0x5d, 0x21, 0x03, 0xd4, 0x75, 0xa6, 0x15, 0xf4, 0xeb, 0x21, 0xfa, 0x1d,
	0x7b, 0x45, 0x5a, 0x87, 0xb3, 0xc7, 0x83, 0x5a, 0x40, 0xe4, 0x6a, 0x8b, 0x87, 0x28, 0xa0, 0x7c,
	0x02, 0xcf, 0x22, 0xff, 0x15, 0xea, 0xa9, 0x96, 0x81, 0x80, 0x44, 0x66, 0x3c, 0x2d, 0x6e, 0x21,
	0x17, 0xc6, 0x4f, 0xa8, 0xef, 0xd9, 0x02, 0x99, 0xaf, 0xca, 0x4e, 0xb9, 0xea, 0xe4, 0xbc, 0x03,
	0x85, 0x57, 0xfe, 0x80, 0xab, 0x9e, 0x3c, 0x0d, 0xd5, 0x34, 0xe5, 0xec, 0x31, 0xf9, 0xb2, 0x36,
	0x48, 0x34, 0x00, 0x5c, 0x4d, 0x65, 0x8c, 0xc1, 0x42, 0xa2, 0xad, 0xc0, 0x06, 0xd7, 0xfb, 0xb1,
	0x8a, 0xa2, 0x9b, 0xbe, 0xb7, 0xcb, 0x37, 0x8f, 0x74, 0x81, 0x07, 0xa5, 0xb5, 0xa1, 0xc9, 0xea,
	0x5f, 0xc7, 0xc8, 0xf2, 0x4f, 0xc3, 0x0f, 0x8d, 0xd9, 0xef, 0xcb, 0xd5, 0xc5, 0x43, 0x1f, 0x88,
	0x06, 0x5f, 0x97, 0xe6, 0xd8, 0x3b, 0x51, 0xb2, 0xb7, 0x87, 0xef, 0x37, 0x06, 0xdb, 0x8d, 0xe1,
	0xee, 0x76, 0x73, 0x7b, 0x6b, 0x10, 0xdf, 0x9e, 0xd5, 0x5b, 0xf7, 0xe3, 0xce, 0xe8, 0xa8, 0xa7,
	0xef, 0x8b, 0xe2, 0xe7, 0xb7, 0xb9, 0xe3, 0x1d, 0xce, 0x7f, 0x9f, 0x5b, 0xea, 0x14, 0xc1, 0xf8,
	0xe6, 0xa8, 0x51, 0xfc, 0xc5, 0x7f, 0xbd, 0x95, 0x46, 0x4c, 0x39, 0xfa, 0xa3, 0x02, 0xac, 0xf3,
	0xcd, 0xd1, 0x7a, 0x0d, 0x58, 0xef, 0xad, 0xac, 0x17, 0x80, 0x3f, 0xe7, 0x96, 0x8b, 0xd1, 0x56,
	0x8b, 0x6f, 0x8e, 0x5a, 0xad, 0x1a, 0xd2, 0x6a, 0xf5, 0x56, 0x5a, 0xad, 0x02, 0xf4, 0xe3, 0x89,
	0x58, 0xdd, 0x93, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x85, 0x42, 0xdd, 0x97, 0x0b, 0x00,
	0x00,
}
