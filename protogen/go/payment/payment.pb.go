// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/payment/payment.proto

package payment

import (
	transaction "github.com/danielgom/grpc-proto/protogen/go/transaction"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *transaction.Cart      `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	Currency      string                 `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	TotaAmount    uint32                 `protobuf:"varint,3,opt,name=tota_amount,json=total_amount,proto3" json:"tota_amount,omitempty"`
	Tax           uint32                 `protobuf:"varint,4,opt,name=tax,proto3" json:"tax,omitempty"`
	PromoCode     string                 `protobuf:"bytes,5,opt,name=promo_code,proto3" json:"promo_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	mi := &file_proto_payment_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentRequest) GetCart() *transaction.Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *PaymentRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PaymentRequest) GetTotaAmount() uint32 {
	if x != nil {
		return x.TotaAmount
	}
	return 0
}

func (x *PaymentRequest) GetTax() uint32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *PaymentRequest) GetPromoCode() string {
	if x != nil {
		return x.PromoCode
	}
	return ""
}

type PaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentUuid   string                 `protobuf:"bytes,1,opt,name=payment_uuid,proto3" json:"payment_uuid,omitempty"`
	Confirmed     bool                   `protobuf:"varint,2,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	mi := &file_proto_payment_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentResponse) GetPaymentUuid() string {
	if x != nil {
		return x.PaymentUuid
	}
	return ""
}

func (x *PaymentResponse) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

var File_proto_payment_payment_proto protoreflect.FileDescriptor

const file_proto_payment_payment_proto_rawDesc = "" +
	"\n" +
	"\x1bproto/payment/payment.proto\x12\apayment\x1a\x1cproto/transaction/cart.proto\"\xa8\x01\n" +
	"\x0ePaymentRequest\x12%\n" +
	"\x04cart\x18\x01 \x01(\v2\x11.transaction.CartR\x04cart\x12\x1a\n" +
	"\bcurrency\x18\x02 \x01(\tR\bcurrency\x12!\n" +
	"\vtota_amount\x18\x03 \x01(\rR\ftotal_amount\x12\x10\n" +
	"\x03tax\x18\x04 \x01(\rR\x03tax\x12\x1e\n" +
	"\n" +
	"promo_code\x18\x05 \x01(\tR\n" +
	"promo_code\"S\n" +
	"\x0fPaymentResponse\x12\"\n" +
	"\fpayment_uuid\x18\x01 \x01(\tR\fpayment_uuid\x12\x1c\n" +
	"\tconfirmed\x18\x02 \x01(\bR\tconfirmedB5Z3github.com/danielgom/grpc-proto/protogen/go/paymentb\x06proto3"

var (
	file_proto_payment_payment_proto_rawDescOnce sync.Once
	file_proto_payment_payment_proto_rawDescData []byte
)

func file_proto_payment_payment_proto_rawDescGZIP() []byte {
	file_proto_payment_payment_proto_rawDescOnce.Do(func() {
		file_proto_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_payment_payment_proto_rawDesc), len(file_proto_payment_payment_proto_rawDesc)))
	})
	return file_proto_payment_payment_proto_rawDescData
}

var file_proto_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_payment_payment_proto_goTypes = []any{
	(*PaymentRequest)(nil),   // 0: payment.PaymentRequest
	(*PaymentResponse)(nil),  // 1: payment.PaymentResponse
	(*transaction.Cart)(nil), // 2: transaction.Cart
}
var file_proto_payment_payment_proto_depIdxs = []int32{
	2, // 0: payment.PaymentRequest.cart:type_name -> transaction.Cart
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_payment_payment_proto_init() }
func file_proto_payment_payment_proto_init() {
	if File_proto_payment_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_payment_payment_proto_rawDesc), len(file_proto_payment_payment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_payment_payment_proto_goTypes,
		DependencyIndexes: file_proto_payment_payment_proto_depIdxs,
		MessageInfos:      file_proto_payment_payment_proto_msgTypes,
	}.Build()
	File_proto_payment_payment_proto = out.File
	file_proto_payment_payment_proto_goTypes = nil
	file_proto_payment_payment_proto_depIdxs = nil
}
