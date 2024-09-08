package domain

import (
	"context"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ProductsCollection = "products"
	LikesCollection = "likes"
	ColorOptionsCollection = "color_options"
	ImageUploadFolder = "uploads"
	ImageQuality = 80  // The quality of the image to be uploaded out of 100

)

type Product struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name    	 string             `json:"author" bson:"author" required:"true"`
	Images       []string           `json:"images" bson:"images"`
	Description  string             `json:"title" bson:"title" required:"true"`
	Category     string    			`json:"category" bson:"category" required:"true"`
	Features     []string           `json:"features" bson:"features"`
	Tags      	 []string           `json:"tags" bson:"tags"`
	ColorOptions []ColorOption      `json:"color_options" bson:"color_options"`
	IsAvailable  bool      			`json:"is_available" bson:"is_available"`
	Views        int       			`json:"views" bson:"views"`
	CreatedAt 	 time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt 	 time.Time          `json:"updated_at" bson:"updated_at"`
}

type Like struct {
	ID      	primitive.ObjectID `json:"id" bson:"_id"`
	UserID  	primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductID   primitive.ObjectID `json:"product_id" bson:"product_id"`
	IsLiked 	bool               `json:"is_liked" bson:"is_liked"`
}

type ColorOption struct {
    Name  string `json:"name" bson:"name"`
    Image string `json:"image" bson:"image"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type ProductUseCase interface {
	CreateProduct(c context.Context, product *Product) (Product, error)
	GetProduct(c context.Context, id string) (*Product, error)
	GetProducts(c context.Context, pagination *Pagination) ([]*Product, error)
	UpdateProduct(c context.Context, product *Product, product_id string) error
	DeleteProduct(c context.Context, id string) error
	UploadProductImages(ctx context.Context, files map[string]io.Reader, serverAdress string) ([]string, error)
	SearchProducts(ctx context.Context, query string) ([]*Product, error)
	LikeProduct(c context.Context, productID string, userID string) error
	UnlikeProduct(c context.Context, productID string, userID string) error
}

type ProductRepository interface {
	CreateProduct(c context.Context, Product *Product) (Product, error)
	GetProduct(c context.Context, id string) (*Product, error)
	GetProducts(c context.Context, pagination *Pagination) ([]*Product, error)
	UpdateProduct(c context.Context, Product *Product) error
	DeleteProduct(c context.Context, id string) error
	SearchProducts(ctx context.Context, query string) ([]*Product, error)
	LikeProduct(c context.Context, ProductID string, userID string) error
	UnlikeProduct(c context.Context, ProductID string, userID string) error
}

type LikeRepository interface {
	LikeProduct(c context.Context, ProductID, userID primitive.ObjectID) error
	UnLikeProduct(c context.Context, ProductID, userID primitive.ObjectID) error
	DeleteLike(c context.Context, ProductID, userID primitive.ObjectID) error
	GetLike(ctx context.Context, userID, ProductID primitive.ObjectID) (*Like, error)
}

type LikeUsecase interface {
	LikeProduct(c context.Context, ProductID, userID primitive.ObjectID) error
	UnLikeProduct(c context.Context, ProductID, userID primitive.ObjectID) error
	DeleteLike(c context.Context, ProductID, userID primitive.ObjectID) error
	GetLike(ctx context.Context, userID, ProductID primitive.ObjectID) (*Like, error)
}