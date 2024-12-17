package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"` // ID tự động sinh bởi MongoDB
	FirstName        string             `bson:"firstname"`     // Tên người dùng
	LastName         string             `bson:"lastname"`      // Họ người dùng
	Email            string             `bson:"email"`         // Email
	Password         string             `bson:"password"`      // Mật khẩu (cần mã hóa)
	Dob              string             `bson:"dob"`
	Phone            string             `bson:"phone"`
	City             string             `bson:"city"`       // Thành phố
	Postal           string             `bson:"postal"`     // Mã bưu điện
	State            string             `bson:"state"`      // Tỉnh/Thành phố
	Address          string             `bson:"address"`    // Địa chỉ
	CreatedAt        time.Time          `bson:"created_at"` // Thời gian tạo
	UpdatedAt        time.Time          `bson:"updated_at"` // Thời gian cập nhật
	Role             primitive.ObjectID `bson:"role"`       // Vai trò của người dùng
	VerificationCode string             `bson:"verificationcode"`
	Verified         bool               `bson:"verified"`
}

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // ID tự động sinh bởi MongoDB
	Name        string             `bson:"name"`          // Tên vai trò
	Description string             `bson:"description"`   // Mô tả vai trò (tuỳ chọn)
}

type CreateUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // ID tự động sinh bởi MongoDB
	FirstName string             `bson:"firstname"`     // Tên người dùng
	LastName  string             `bson:"lastname"`      // Họ người dùng
	Email     string             `bson:"email"`         // Email
	Password  string             `bson:"password"`      // Mật khẩu (cần mã hóa)
	Dob       string             `bson:"dob"`
	CreatedAt time.Time          `bson:"created_at"` // Thời gian tạo
	Role      primitive.ObjectID `bson:"role"`       // Vai trò của người dùng
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // ID tự động sinh bởi MongoDB
	Name        string             `bson:"name"`          // Tên sản phẩm
	Price       int32              `bson:"price"`         // Giá sản phẩm
	Description string             `bson:"description"`   // Mô tả sản phẩm
	Quantity    int32              `bson:"quantity"`      // Số lượng sản phẩm
	Image       string             `bson:"image"`         // Hình ảnh sản phẩm
	Rating      float32            `bson:"rating"`        // Thêm trường rating
	CreatedAt   time.Time          `bson:"created_at"`    // Thời gian tạo
	UpdatedAt   time.Time          `bson:"updated_at"`    // Thời gian cập nhật
	Tags        []ListTags         `bson:"tags"`          // Danh mục sản phẩm
}

type ListTags struct {
	ID   primitive.ObjectID `bson:"id"`
	Name string             `bson:"name"` // Tên danh mục
}

type Item struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProductID primitive.ObjectID `bson:"product_id"` // ID của sản phẩm
	Name      string             `bson:"name"`       // Tên sản phẩm
	Price     int32              `bson:"price"`      // Giá sản phẩm
	Quantity  int32              `bson:"quantity"`   // Số lượng đặt mua
	Image     string             `bson:"image"`      // Hình ảnh sản phẩm
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type Cart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	User      primitive.ObjectID `bson:"user"`
	Items     []Item             `bson:"items"` // Đổi tên từ Product sang Items
	Totals    int32              `bson:"totals"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type AccessToken struct {
	AccessToken string `json:"accesstoken"`
}

type Order struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UserID          primitive.ObjectID `bson:"user_id"`
	Items           []OrderItem        `bson:"items"`
	TotalAmount     int32              `bson:"total_amount"`
	Status          string             `bson:"status"` // pending, confirmed, shipped, delivered, cancelled
	ShippingDetails ShippingDetails    `bson:"shipping_details"`
	PaymentDetails  PaymentDetails     `bson:"payment_details"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

type OrderItem struct {
	ProductID primitive.ObjectID `bson:"product_id"`
	Name      string             `bson:"name"`
	Price     int32              `bson:"price"`
	Quantity  int32              `bson:"quantity"`
	Image     string             `bson:"image"`
}

type ShippingDetails struct {
	FirstName   string `bson:"first_name"`
	LastName    string `bson:"last_name"`
	Email       string `bson:"email"`
	PhoneNumber string `bson:"phone_number"`
	Address     string `bson:"address"`
	City        string `bson:"city"`
	State       string `bson:"state"`
	ZipCode     string `bson:"zip_code"`
}

type PaymentDetails struct {
	Method        string    `bson:"method"` // COD, Card, etc
	Status        string    `bson:"status"` // pending, completed, failed
	TransactionID string    `bson:"transaction_id,omitempty"`
	PaidAt        time.Time `bson:"paid_at,omitempty"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	ProductID primitive.ObjectID `bson:"product_id"`
	Content   string             `bson:"content"`
	Rating    float32            `bson:"rating"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
