package repository

import (
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepository struct {
	database   mongo.Database
	collection string
}

func NewProductRepository(db mongo.Database, collection string) domain.ProductRepository {
	return &productRepository{
		database:   db,
		collection: collection,
	}
}

func (r *productRepository) CreateProduct(c context.Context, product *domain.Product) (domain.Product, error) {
	collection := r.database.Collection(r.collection)
	_, err := collection.InsertOne(c, product)
	return *product, err
}

func (r *productRepository) GetProduct(c context.Context, id string) (*domain.Product, error) {
	collection := r.database.Collection(r.collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product domain.Product
	err = collection.FindOne(c, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // product not found
		}
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetProducts(c context.Context, pagination *domain.Pagination) ([]*domain.Product, error) {
	collection := r.database.Collection(r.collection)

	var products []*domain.Product
	filter := bson.M{}
	skip := int64((pagination.Page - 1) * pagination.PageSize)
	limit := int64(pagination.PageSize)
	opts := &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := collection.Find(c, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (r *productRepository) UpdateProduct(c context.Context, product *domain.Product) error {
	collection := r.database.Collection(r.collection)

	_, err := collection.UpdateOne(
		c,
		bson.M{"_id": product.ID},
		bson.M{"$set": product},
	)
	return err
}

func (r *productRepository) DeleteProduct(c context.Context, id string) error {
	collection := r.database.Collection(r.collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": objectID})
	return err
}

func (r *productRepository) LikeProduct(c context.Context, productID string, userID string) error {
	collection := r.database.Collection(r.collection)

	productObjectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": productObjectID}
	update := bson.M{"$addToSet": bson.M{"likes": userObjectID}}
	_, err = collection.UpdateOne(c, filter, update)
	return err
}

func (r *productRepository) UnlikeProduct(c context.Context, productID string, userID string) error {
	collection := r.database.Collection(r.collection)

	productObjectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": productObjectID}
	update := bson.M{"$pull": bson.M{"likes": userObjectID}}
	_, err = collection.UpdateOne(c, filter, update)
	return err
}

func (r *productRepository) SearchProducts(c context.Context, keyword string) ([]*domain.Product, error) {
	collection := r.database.Collection(r.collection)

	var products []*domain.Product
	filter := bson.M{
		"$text": bson.M{
			"$search": keyword,
		},
	}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}