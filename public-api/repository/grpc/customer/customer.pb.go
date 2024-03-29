// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0--rc1
// source: public-api/repository/grpc/proto/customer.proto

package customer

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestDataCustomer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *RequestDataCustomer) Reset() {
	*x = RequestDataCustomer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDataCustomer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDataCustomer) ProtoMessage() {}

func (x *RequestDataCustomer) ProtoReflect() protoreflect.Message {
	mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDataCustomer.ProtoReflect.Descriptor instead.
func (*RequestDataCustomer) Descriptor() ([]byte, []int) {
	return file_public_api_repository_grpc_proto_customer_proto_rawDescGZIP(), []int{0}
}

func (x *RequestDataCustomer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RequestDataCustomer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestDataCustomer) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type RequestDataCheckout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email            string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber      string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	ProductSalesId   int64  `protobuf:"varint,4,opt,name=product_sales_id,json=productSalesId,proto3" json:"product_sales_id,omitempty"`
	Qty              int64  `protobuf:"varint,5,opt,name=qty,proto3" json:"qty,omitempty"`
	TotalPricing     int64  `protobuf:"varint,6,opt,name=total_pricing,json=totalPricing,proto3" json:"total_pricing,omitempty"`
	PaymentReference string `protobuf:"bytes,7,opt,name=payment_reference,json=paymentReference,proto3" json:"payment_reference,omitempty"`
	PaymentDomain    string `protobuf:"bytes,8,opt,name=payment_domain,json=paymentDomain,proto3" json:"payment_domain,omitempty"`
	CustomerId       int64  `protobuf:"varint,9,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ListKey          string `protobuf:"bytes,10,opt,name=list_key,json=listKey,proto3" json:"list_key,omitempty"`
	Invoice          string `protobuf:"bytes,11,opt,name=invoice,proto3" json:"invoice,omitempty"`
	TypeDuration     string `protobuf:"bytes,12,opt,name=type_duration,json=typeDuration,proto3" json:"type_duration,omitempty"`
	Pricing          string `protobuf:"bytes,13,opt,name=pricing,proto3" json:"pricing,omitempty"`
	Discount         string `protobuf:"bytes,14,opt,name=discount,proto3" json:"discount,omitempty"`
	Tax              string `protobuf:"bytes,15,opt,name=tax,proto3" json:"tax,omitempty"`
	Status           string `protobuf:"bytes,16,opt,name=status,proto3" json:"status,omitempty"`
	References       string `protobuf:"bytes,17,opt,name=references,proto3" json:"references,omitempty"`
}

func (x *RequestDataCheckout) Reset() {
	*x = RequestDataCheckout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDataCheckout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDataCheckout) ProtoMessage() {}

func (x *RequestDataCheckout) ProtoReflect() protoreflect.Message {
	mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDataCheckout.ProtoReflect.Descriptor instead.
func (*RequestDataCheckout) Descriptor() ([]byte, []int) {
	return file_public_api_repository_grpc_proto_customer_proto_rawDescGZIP(), []int{1}
}

func (x *RequestDataCheckout) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RequestDataCheckout) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestDataCheckout) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RequestDataCheckout) GetProductSalesId() int64 {
	if x != nil {
		return x.ProductSalesId
	}
	return 0
}

func (x *RequestDataCheckout) GetQty() int64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *RequestDataCheckout) GetTotalPricing() int64 {
	if x != nil {
		return x.TotalPricing
	}
	return 0
}

func (x *RequestDataCheckout) GetPaymentReference() string {
	if x != nil {
		return x.PaymentReference
	}
	return ""
}

func (x *RequestDataCheckout) GetPaymentDomain() string {
	if x != nil {
		return x.PaymentDomain
	}
	return ""
}

func (x *RequestDataCheckout) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *RequestDataCheckout) GetListKey() string {
	if x != nil {
		return x.ListKey
	}
	return ""
}

func (x *RequestDataCheckout) GetInvoice() string {
	if x != nil {
		return x.Invoice
	}
	return ""
}

func (x *RequestDataCheckout) GetTypeDuration() string {
	if x != nil {
		return x.TypeDuration
	}
	return ""
}

func (x *RequestDataCheckout) GetPricing() string {
	if x != nil {
		return x.Pricing
	}
	return ""
}

func (x *RequestDataCheckout) GetDiscount() string {
	if x != nil {
		return x.Discount
	}
	return ""
}

func (x *RequestDataCheckout) GetTax() string {
	if x != nil {
		return x.Tax
	}
	return ""
}

func (x *RequestDataCheckout) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RequestDataCheckout) GetReferences() string {
	if x != nil {
		return x.References
	}
	return ""
}

type RequestProductIDAndLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Limit     string `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *RequestProductIDAndLimit) Reset() {
	*x = RequestProductIDAndLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestProductIDAndLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestProductIDAndLimit) ProtoMessage() {}

func (x *RequestProductIDAndLimit) ProtoReflect() protoreflect.Message {
	mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestProductIDAndLimit.ProtoReflect.Descriptor instead.
func (*RequestProductIDAndLimit) Descriptor() ([]byte, []int) {
	return file_public_api_repository_grpc_proto_customer_proto_rawDescGZIP(), []int{2}
}

func (x *RequestProductIDAndLimit) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *RequestProductIDAndLimit) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

type RequestCheckStok struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int32 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *RequestCheckStok) Reset() {
	*x = RequestCheckStok{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCheckStok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCheckStok) ProtoMessage() {}

func (x *RequestCheckStok) ProtoReflect() protoreflect.Message {
	mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCheckStok.ProtoReflect.Descriptor instead.
func (*RequestCheckStok) Descriptor() ([]byte, []int) {
	return file_public_api_repository_grpc_proto_customer_proto_rawDescGZIP(), []int{3}
}

func (x *RequestCheckStok) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type RequestProductSerialNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber string `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
}

func (x *RequestProductSerialNumber) Reset() {
	*x = RequestProductSerialNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestProductSerialNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestProductSerialNumber) ProtoMessage() {}

func (x *RequestProductSerialNumber) ProtoReflect() protoreflect.Message {
	mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestProductSerialNumber.ProtoReflect.Descriptor instead.
func (*RequestProductSerialNumber) Descriptor() ([]byte, []int) {
	return file_public_api_repository_grpc_proto_customer_proto_rawDescGZIP(), []int{4}
}

func (x *RequestProductSerialNumber) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

type ResponseDataCheckout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email            string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber      string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	ProductSalesId   int64  `protobuf:"varint,4,opt,name=product_sales_id,json=productSalesId,proto3" json:"product_sales_id,omitempty"`
	Qty              int64  `protobuf:"varint,5,opt,name=qty,proto3" json:"qty,omitempty"`
	TotalPricing     int64  `protobuf:"varint,6,opt,name=total_pricing,json=totalPricing,proto3" json:"total_pricing,omitempty"`
	PaymentReference string `protobuf:"bytes,7,opt,name=payment_reference,json=paymentReference,proto3" json:"payment_reference,omitempty"`
	PaymentDomain    string `protobuf:"bytes,8,opt,name=payment_domain,json=paymentDomain,proto3" json:"payment_domain,omitempty"`
	CustomerId       int64  `protobuf:"varint,9,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ListKey          string `protobuf:"bytes,10,opt,name=list_key,json=listKey,proto3" json:"list_key,omitempty"`
	Invoice          string `protobuf:"bytes,11,opt,name=invoice,proto3" json:"invoice,omitempty"`
	TypeDuration     string `protobuf:"bytes,12,opt,name=type_duration,json=typeDuration,proto3" json:"type_duration,omitempty"`
	Pricing          string `protobuf:"bytes,13,opt,name=pricing,proto3" json:"pricing,omitempty"`
	Discount         string `protobuf:"bytes,14,opt,name=discount,proto3" json:"discount,omitempty"`
	Tax              string `protobuf:"bytes,15,opt,name=tax,proto3" json:"tax,omitempty"`
	Status           string `protobuf:"bytes,16,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseDataCheckout) Reset() {
	*x = ResponseDataCheckout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDataCheckout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDataCheckout) ProtoMessage() {}

func (x *ResponseDataCheckout) ProtoReflect() protoreflect.Message {
	mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDataCheckout.ProtoReflect.Descriptor instead.
func (*ResponseDataCheckout) Descriptor() ([]byte, []int) {
	return file_public_api_repository_grpc_proto_customer_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseDataCheckout) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResponseDataCheckout) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseDataCheckout) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *ResponseDataCheckout) GetProductSalesId() int64 {
	if x != nil {
		return x.ProductSalesId
	}
	return 0
}

func (x *ResponseDataCheckout) GetQty() int64 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *ResponseDataCheckout) GetTotalPricing() int64 {
	if x != nil {
		return x.TotalPricing
	}
	return 0
}

func (x *ResponseDataCheckout) GetPaymentReference() string {
	if x != nil {
		return x.PaymentReference
	}
	return ""
}

func (x *ResponseDataCheckout) GetPaymentDomain() string {
	if x != nil {
		return x.PaymentDomain
	}
	return ""
}

func (x *ResponseDataCheckout) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ResponseDataCheckout) GetListKey() string {
	if x != nil {
		return x.ListKey
	}
	return ""
}

func (x *ResponseDataCheckout) GetInvoice() string {
	if x != nil {
		return x.Invoice
	}
	return ""
}

func (x *ResponseDataCheckout) GetTypeDuration() string {
	if x != nil {
		return x.TypeDuration
	}
	return ""
}

func (x *ResponseDataCheckout) GetPricing() string {
	if x != nil {
		return x.Pricing
	}
	return ""
}

func (x *ResponseDataCheckout) GetDiscount() string {
	if x != nil {
		return x.Discount
	}
	return ""
}

func (x *ResponseDataCheckout) GetTax() string {
	if x != nil {
		return x.Tax
	}
	return ""
}

func (x *ResponseDataCheckout) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_api_repository_grpc_proto_customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_public_api_repository_grpc_proto_customer_proto_rawDescGZIP(), []int{6}
}

var File_public_api_repository_grpc_proto_customer_proto protoreflect.FileDescriptor

var file_public_api_repository_grpc_proto_customer_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x13, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x92, 0x04, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x71, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x78, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x22, 0x4f, 0x0a, 0x18, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x31, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x1a, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xf3, 0x03, 0x0a, 0x14,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x61, 0x6c,
	0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x4b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x78, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x95, 0x03, 0x0a, 0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x0c, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0f, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x12, 0x1d, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x1a,
	0x17, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x24, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x40, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x74, 0x6f, 0x6b, 0x12, 0x1a, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f,
	0x6b, 0x1a, 0x17, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_api_repository_grpc_proto_customer_proto_rawDescOnce sync.Once
	file_public_api_repository_grpc_proto_customer_proto_rawDescData = file_public_api_repository_grpc_proto_customer_proto_rawDesc
)

func file_public_api_repository_grpc_proto_customer_proto_rawDescGZIP() []byte {
	file_public_api_repository_grpc_proto_customer_proto_rawDescOnce.Do(func() {
		file_public_api_repository_grpc_proto_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_api_repository_grpc_proto_customer_proto_rawDescData)
	})
	return file_public_api_repository_grpc_proto_customer_proto_rawDescData
}

var file_public_api_repository_grpc_proto_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_public_api_repository_grpc_proto_customer_proto_goTypes = []interface{}{
	(*RequestDataCustomer)(nil),        // 0: customer.RequestDataCustomer
	(*RequestDataCheckout)(nil),        // 1: customer.RequestDataCheckout
	(*RequestProductIDAndLimit)(nil),   // 2: customer.RequestProductIDAndLimit
	(*RequestCheckStok)(nil),           // 3: customer.RequestCheckStok
	(*RequestProductSerialNumber)(nil), // 4: customer.RequestProductSerialNumber
	(*ResponseDataCheckout)(nil),       // 5: customer.ResponseDataCheckout
	(*EmptyResponse)(nil),              // 6: customer.EmptyResponse
}
var file_public_api_repository_grpc_proto_customer_proto_depIdxs = []int32{
	0, // 0: customer.CustomerUseCaseService.PostCustomer:input_type -> customer.RequestDataCustomer
	1, // 1: customer.CustomerUseCaseService.PostCheckout:input_type -> customer.RequestDataCheckout
	1, // 2: customer.CustomerUseCaseService.PartnerCheckout:input_type -> customer.RequestDataCheckout
	4, // 3: customer.CustomerUseCaseService.GetCheckoutByKeyNumber:input_type -> customer.RequestProductSerialNumber
	3, // 4: customer.CustomerUseCaseService.CheckStok:input_type -> customer.RequestCheckStok
	6, // 5: customer.CustomerUseCaseService.PostCustomer:output_type -> customer.EmptyResponse
	6, // 6: customer.CustomerUseCaseService.PostCheckout:output_type -> customer.EmptyResponse
	6, // 7: customer.CustomerUseCaseService.PartnerCheckout:output_type -> customer.EmptyResponse
	5, // 8: customer.CustomerUseCaseService.GetCheckoutByKeyNumber:output_type -> customer.ResponseDataCheckout
	6, // 9: customer.CustomerUseCaseService.CheckStok:output_type -> customer.EmptyResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_public_api_repository_grpc_proto_customer_proto_init() }
func file_public_api_repository_grpc_proto_customer_proto_init() {
	if File_public_api_repository_grpc_proto_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_api_repository_grpc_proto_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDataCustomer); i {
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
		file_public_api_repository_grpc_proto_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDataCheckout); i {
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
		file_public_api_repository_grpc_proto_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestProductIDAndLimit); i {
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
		file_public_api_repository_grpc_proto_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCheckStok); i {
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
		file_public_api_repository_grpc_proto_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestProductSerialNumber); i {
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
		file_public_api_repository_grpc_proto_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDataCheckout); i {
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
		file_public_api_repository_grpc_proto_customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_public_api_repository_grpc_proto_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_api_repository_grpc_proto_customer_proto_goTypes,
		DependencyIndexes: file_public_api_repository_grpc_proto_customer_proto_depIdxs,
		MessageInfos:      file_public_api_repository_grpc_proto_customer_proto_msgTypes,
	}.Build()
	File_public_api_repository_grpc_proto_customer_proto = out.File
	file_public_api_repository_grpc_proto_customer_proto_rawDesc = nil
	file_public_api_repository_grpc_proto_customer_proto_goTypes = nil
	file_public_api_repository_grpc_proto_customer_proto_depIdxs = nil
}
