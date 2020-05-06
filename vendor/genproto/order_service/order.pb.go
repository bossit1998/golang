// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order_service

import (
	fmt "fmt"
	courier_service "genproto/courier_service"
	fare_service "genproto/fare_service"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Location struct {
	Long                 float32  `protobuf:"fixed32,1,opt,name=long,proto3" json:"long,omitempty"`
	Lat                  float32  `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLong() float32 {
	if m != nil {
		return m.Long
	}
	return 0
}

func (m *Location) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

type Order struct {
	Id                   string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerName         string                   `protobuf:"bytes,2,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	ToLocation           *Location                `protobuf:"bytes,3,opt,name=to_location,json=toLocation,proto3" json:"to_location,omitempty"`
	ToAddress            string                   `protobuf:"bytes,4,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	CourierId            *wrappers.StringValue    `protobuf:"bytes,5,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`
	FareId               string                   `protobuf:"bytes,6,opt,name=fare_id,json=fareId,proto3" json:"fare_id,omitempty"`
	StatusId             string                   `protobuf:"bytes,7,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	CreatedAt            string                   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Description          string                   `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	CoId                 string                   `protobuf:"bytes,10,opt,name=co_id,json=coId,proto3" json:"co_id,omitempty"`
	CreatorTypeId        string                   `protobuf:"bytes,11,opt,name=creator_type_id,json=creatorTypeId,proto3" json:"creator_type_id,omitempty"`
	UserId               string                   `protobuf:"bytes,12,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Steps                []*Step                  `protobuf:"bytes,13,rep,name=steps,proto3" json:"steps,omitempty"`
	Fare                 *fare_service.Fare       `protobuf:"bytes,14,opt,name=fare,proto3" json:"fare,omitempty"`
	CoDeliveryPrice      float32                  `protobuf:"fixed32,15,opt,name=co_delivery_price,json=coDeliveryPrice,proto3" json:"co_delivery_price,omitempty"`
	DeliveryPrice        float32                  `protobuf:"fixed32,16,opt,name=delivery_price,json=deliveryPrice,proto3" json:"delivery_price,omitempty"`
	Courier              *courier_service.Courier `protobuf:"bytes,17,opt,name=courier,proto3" json:"courier,omitempty"`
	CustomerPhoneNumber  string                   `protobuf:"bytes,18,opt,name=customer_phone_number,json=customerPhoneNumber,proto3" json:"customer_phone_number,omitempty"`
	FinishedAt           string                   `protobuf:"bytes,19,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	OrderAmount          float32                  `protobuf:"fixed32,20,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
}

func (m *Order) GetToLocation() *Location {
	if m != nil {
		return m.ToLocation
	}
	return nil
}

func (m *Order) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *Order) GetCourierId() *wrappers.StringValue {
	if m != nil {
		return m.CourierId
	}
	return nil
}

func (m *Order) GetFareId() string {
	if m != nil {
		return m.FareId
	}
	return ""
}

func (m *Order) GetStatusId() string {
	if m != nil {
		return m.StatusId
	}
	return ""
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Order) GetCoId() string {
	if m != nil {
		return m.CoId
	}
	return ""
}

func (m *Order) GetCreatorTypeId() string {
	if m != nil {
		return m.CreatorTypeId
	}
	return ""
}

func (m *Order) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Order) GetSteps() []*Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *Order) GetFare() *fare_service.Fare {
	if m != nil {
		return m.Fare
	}
	return nil
}

func (m *Order) GetCoDeliveryPrice() float32 {
	if m != nil {
		return m.CoDeliveryPrice
	}
	return 0
}

func (m *Order) GetDeliveryPrice() float32 {
	if m != nil {
		return m.DeliveryPrice
	}
	return 0
}

func (m *Order) GetCourier() *courier_service.Courier {
	if m != nil {
		return m.Courier
	}
	return nil
}

func (m *Order) GetCustomerPhoneNumber() string {
	if m != nil {
		return m.CustomerPhoneNumber
	}
	return ""
}

func (m *Order) GetFinishedAt() string {
	if m != nil {
		return m.FinishedAt
	}
	return ""
}

func (m *Order) GetOrderAmount() float32 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

type Step struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchName           string     `protobuf:"bytes,2,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	Location             *Location  `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Address              string     `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	DestinationAddress   string     `protobuf:"bytes,5,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	PhoneNumber          string     `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Products             []*Product `protobuf:"bytes,7,rep,name=products,proto3" json:"products,omitempty"`
	Description          string     `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	OrderNo              uint64     `protobuf:"varint,9,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	Status               string     `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	StepAmount           float32    `protobuf:"fixed32,11,opt,name=step_amount,json=stepAmount,proto3" json:"step_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}
func (*Step) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *Step) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Step.Unmarshal(m, b)
}
func (m *Step) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Step.Marshal(b, m, deterministic)
}
func (m *Step) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Step.Merge(m, src)
}
func (m *Step) XXX_Size() int {
	return xxx_messageInfo_Step.Size(m)
}
func (m *Step) XXX_DiscardUnknown() {
	xxx_messageInfo_Step.DiscardUnknown(m)
}

var xxx_messageInfo_Step proto.InternalMessageInfo

func (m *Step) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Step) GetBranchName() string {
	if m != nil {
		return m.BranchName
	}
	return ""
}

func (m *Step) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Step) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Step) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *Step) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Step) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *Step) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Step) GetOrderNo() uint64 {
	if m != nil {
		return m.OrderNo
	}
	return 0
}

func (m *Step) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Step) GetStepAmount() float32 {
	if m != nil {
		return m.StepAmount
	}
	return 0
}

type Product struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity             float32  `protobuf:"fixed32,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	TotalAmount          float32  `protobuf:"fixed32,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetQuantity() float32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetTotalAmount() float32 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Location)(nil), "genproto.Location")
	proto.RegisterType((*Order)(nil), "genproto.Order")
	proto.RegisterType((*Step)(nil), "genproto.Step")
	proto.RegisterType((*Product)(nil), "genproto.Product")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0x55, 0x82, 0x9d, 0x38, 0xd7, 0x49, 0x80, 0x09, 0x0f, 0xe6, 0xf1, 0x1e, 0x25, 0xa4, 0x1f,
	0x42, 0xad, 0x9a, 0x54, 0xb0, 0xec, 0x2a, 0x6d, 0x55, 0x29, 0x52, 0x45, 0x91, 0xa9, 0xba, 0xe8,
	0xc6, 0x9a, 0x78, 0x86, 0x30, 0x52, 0xe2, 0x71, 0xc7, 0x63, 0xaa, 0x6c, 0xfb, 0x3b, 0xfb, 0x4b,
	0xba, 0xaa, 0xe6, 0x8e, 0x1d, 0x0c, 0x6c, 0xba, 0x9b, 0x7b, 0xce, 0x61, 0x7c, 0xee, 0xe1, 0x4c,
	0x20, 0x54, 0x9a, 0x0b, 0x3d, 0xce, 0xb4, 0x32, 0x8a, 0x04, 0x0b, 0x91, 0xe2, 0xe9, 0xf0, 0xc9,
	0x42, 0xa9, 0xc5, 0x52, 0x4c, 0x70, 0x9a, 0x17, 0xd7, 0x93, 0x1f, 0x9a, 0x65, 0x99, 0xd0, 0xb9,
	0x53, 0x1e, 0x1e, 0x5c, 0x33, 0x2d, 0xe2, 0x5c, 0xe8, 0x5b, 0x99, 0x88, 0x89, 0x1d, 0x4a, 0xe2,
	0x28, 0x51, 0x85, 0x96, 0x42, 0x6f, 0xb8, 0x72, 0x76, 0xf4, 0xe8, 0x0d, 0x04, 0x9f, 0x54, 0xc2,
	0x8c, 0x54, 0x29, 0x21, 0xe0, 0x2d, 0x55, 0xba, 0xa0, 0x8d, 0x61, 0xe3, 0xb4, 0x19, 0xe1, 0x99,
	0xec, 0xc0, 0xd6, 0x92, 0x19, 0xda, 0x44, 0xc8, 0x1e, 0x47, 0xbf, 0x7c, 0xf0, 0x3f, 0x5b, 0x8f,
	0xa4, 0x0f, 0x4d, 0xc9, 0x51, 0xdd, 0x89, 0x9a, 0x92, 0x93, 0xa7, 0xd0, 0x4b, 0x8a, 0xdc, 0xa8,
	0x95, 0xd0, 0x71, 0xca, 0x56, 0x02, 0xff, 0xaa, 0x13, 0x75, 0x2b, 0xf0, 0x82, 0xad, 0x04, 0x39,
	0x87, 0xd0, 0xa8, 0x78, 0x59, 0x7e, 0x93, 0x6e, 0x0d, 0x1b, 0xa7, 0xe1, 0x19, 0x19, 0x57, 0x8b,
	0x8e, 0x2b, 0x37, 0x11, 0x18, 0xb5, 0x71, 0x76, 0x04, 0x60, 0x54, 0xcc, 0x38, 0xd7, 0x22, 0xcf,
	0xa9, 0x87, 0xd7, 0x76, 0x8c, 0x9a, 0x3a, 0x80, 0xbc, 0x05, 0xa8, 0xb6, 0x94, 0x9c, 0xfa, 0x78,
	0xe5, 0xff, 0x63, 0x97, 0xd8, 0xb8, 0x4a, 0x6c, 0x7c, 0x65, 0xb4, 0x4c, 0x17, 0x5f, 0xd9, 0xb2,
	0x10, 0x51, 0xa7, 0xd4, 0xcf, 0x38, 0x39, 0x80, 0x36, 0x66, 0x27, 0x39, 0x6d, 0xe1, 0xc5, 0x2d,
	0x3b, 0xce, 0x38, 0xf9, 0x0f, 0x3a, 0xb9, 0x61, 0xa6, 0xc8, 0x2d, 0xd5, 0x46, 0x2a, 0x70, 0xc0,
	0x8c, 0x5b, 0x47, 0x89, 0x16, 0xcc, 0x08, 0x1e, 0x33, 0x43, 0x03, 0xe7, 0xa8, 0x44, 0xa6, 0x86,
	0x0c, 0x21, 0xe4, 0x22, 0x4f, 0xb4, 0xcc, 0x70, 0xcb, 0x0e, 0xf2, 0x75, 0x88, 0x0c, 0xc0, 0x4f,
	0x94, 0xbd, 0x19, 0x90, 0xf3, 0x12, 0x35, 0xe3, 0xe4, 0x05, 0x6c, 0xe3, 0x1d, 0x4a, 0xc7, 0x66,
	0x9d, 0xa1, 0xa7, 0x10, 0xe9, 0x5e, 0x09, 0x7f, 0x59, 0x67, 0xc2, 0x79, 0x2e, 0x72, 0xb7, 0x6d,
	0xd7, 0x79, 0xb6, 0xe3, 0x8c, 0x93, 0x67, 0xe0, 0xe7, 0x46, 0x64, 0x39, 0xed, 0x0d, 0xb7, 0x4e,
	0xc3, 0xb3, 0xfe, 0x5d, 0xae, 0x57, 0x46, 0x64, 0x91, 0x23, 0xc9, 0x08, 0x3c, 0xbb, 0x23, 0xed,
	0x63, 0x52, 0x35, 0xd1, 0x47, 0xa6, 0x45, 0x84, 0x1c, 0x79, 0x09, 0xbb, 0x89, 0x8a, 0xb9, 0x58,
	0xca, 0x5b, 0xa1, 0xd7, 0x71, 0xa6, 0x65, 0x22, 0xe8, 0x36, 0xd6, 0x60, 0x3b, 0x51, 0x1f, 0x4a,
	0xfc, 0xd2, 0xc2, 0xe4, 0x39, 0xf4, 0x1f, 0x08, 0x77, 0x50, 0xd8, 0xe3, 0xf7, 0x64, 0xaf, 0xa0,
	0x5d, 0xc6, 0x4e, 0x77, 0xf1, 0xcb, 0xbb, 0x77, 0x5f, 0x7e, 0xef, 0x88, 0xa8, 0x52, 0x90, 0x33,
	0xf8, 0x67, 0x53, 0xa6, 0xec, 0x46, 0xa5, 0x22, 0x4e, 0x8b, 0xd5, 0x5c, 0x68, 0x4a, 0x70, 0xe1,
	0x41, 0x45, 0x5e, 0x5a, 0xee, 0x02, 0x29, 0x72, 0x0c, 0xe1, 0xb5, 0x4c, 0x65, 0x7e, 0xe3, 0xfe,
	0x2b, 0x03, 0x54, 0x42, 0x05, 0x4d, 0x0d, 0x39, 0x81, 0x2e, 0x3e, 0xaf, 0x98, 0xad, 0x54, 0x91,
	0x1a, 0xba, 0x87, 0x36, 0xdd, 0x93, 0x9b, 0x22, 0x34, 0xfa, 0xdd, 0x04, 0xcf, 0x66, 0xf5, 0xa8,
	0xdd, 0xc7, 0x10, 0xce, 0x35, 0x4b, 0x93, 0x9b, 0x7a, 0xb7, 0xc1, 0x41, 0xd8, 0xec, 0x31, 0x04,
	0x7f, 0x51, 0xeb, 0x8d, 0x86, 0x50, 0x68, 0xdf, 0x6f, 0x74, 0x35, 0x92, 0x09, 0x0c, 0xb8, 0xc8,
	0x8d, 0x4c, 0x51, 0xb8, 0xe9, 0xbd, 0x8f, 0x2a, 0x52, 0xa3, 0xaa, 0x07, 0x70, 0x02, 0xdd, 0x7b,
	0x19, 0xb9, 0x22, 0x87, 0x59, 0x2d, 0x9b, 0xd7, 0x10, 0x64, 0x5a, 0xf1, 0x22, 0x31, 0x39, 0x6d,
	0x63, 0x39, 0x6a, 0xe9, 0x5f, 0x3a, 0x26, 0xda, 0x48, 0x1e, 0x16, 0x38, 0x78, 0x5c, 0xe0, 0x7f,
	0x21, 0x70, 0x59, 0xa6, 0x0a, 0xfb, 0xed, 0x45, 0x6d, 0x9c, 0x2f, 0x14, 0xd9, 0x87, 0x96, 0x7b,
	0x28, 0x65, 0xb9, 0xcb, 0xc9, 0x46, 0x68, 0x0b, 0x58, 0xa5, 0x1f, 0x62, 0xfa, 0x60, 0xa1, 0x32,
	0xfc, 0x9f, 0x0d, 0x68, 0x97, 0x5e, 0x1e, 0xe5, 0x4f, 0xc0, 0xab, 0x05, 0x8f, 0x67, 0x72, 0x08,
	0xc1, 0xf7, 0x82, 0xa5, 0x46, 0x9a, 0x35, 0x46, 0xde, 0x8c, 0x36, 0x33, 0xd9, 0x03, 0xdf, 0x75,
	0xd1, 0x43, 0xc2, 0x0d, 0x36, 0x29, 0xa3, 0x0c, 0x5b, 0x56, 0x1e, 0x7c, 0xd7, 0x00, 0xc4, 0x9c,
	0x89, 0x77, 0xf4, 0xdb, 0x7e, 0x15, 0xcc, 0xc4, 0x6d, 0x58, 0xfe, 0x74, 0xce, 0x5b, 0x08, 0x9e,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xaa, 0xe8, 0x4e, 0xa4, 0x05, 0x00, 0x00,
}
