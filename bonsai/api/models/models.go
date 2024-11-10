package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // ID tự động sinh bởi MongoDB
	Username  string             `bson:"username"`      // Tên người dùng
	Email     string             `bson:"email"`         // Email
	Password  string             `bson:"password"`      // Mật khẩu (cần mã hóa)
	Dob       string             `bson:"dob"`
	CreatedAt time.Time          `bson:"created_at"` // Thời gian tạo
	UpdatedAt time.Time          `bson:"updated_at"` // Thời gian cập nhật
	Role      primitive.ObjectID `bson:"role"`       // Vai trò của người dùng
}

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // ID tự động sinh bởi MongoDB
	Name        string             `bson:"name"`          // Tên vai trò
	Description string             `bson:"description"`   // Mô tả vai trò (tuỳ chọn)
}

type CreateUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // ID tự động sinh bởi MongoDB
	Username  string             `bson:"username"`      // Tên người dùng
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
	Image       []string           `bson:"image"`         // Hình ảnh sản phẩm
	CreatedAt   time.Time          `bson:"created_at"`    // Thời gian tạo
	UpdatedAt   time.Time          `bson:"updated_at"`    // Thời gian cập nhật
}

type Categories struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // ID tự động sinh bởi MongoDB
	Name        string             `bson:"name"`          // Tên danh mục
	Description string             `bson:"description"`   // Mô tả danh mục
}

type Item struct {
	ProductID primitive.ObjectID `bson:"product"`    // Sản phẩm
	Price     int32              `bson:"price"`      // Giá sản phẩm
	Quantity  int32              `bson:"quantity"`   // Số lượng sản phẩm
	CreatedAt time.Time          `bson:"created_at"` // Thời gian tạo
	UpdatedAt time.Time          `bson:"updated_at"` // Thời gian cập nhật
}

type Cart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // ID tự động sinh bởi MongoDB
	User      primitive.ObjectID `bson:"user"`          // Người dùng
	Product   []Item             `bson:"items"`         // Sản phẩm
	Totals    int32              `bson:"totals"`        // Số lượng sản phẩm
	CreatedAt time.Time          `bson:"created_at"`    // Thời gian tạo
	UpdatedAt time.Time          `bson:"updated_at"`    // Thời gian cập nhật
}

type AccessToken struct {
	AccessToken string `json:"accesstoken"`
}
