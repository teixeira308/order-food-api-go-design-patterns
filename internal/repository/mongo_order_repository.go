package repository

import (
	"context"
	"errors"
	"time"

	"github.com/teixeira308/order-food-api-go-design-patterns/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrderRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}
type mongoOrder struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CustomerID    string             `bson:"customer_id"`
	Items         []domain.OrderItem `bson:"items"` // certifique que OrderItem tem bson tags!
	Status        domain.OrderStatus `bson:"status"`
	CreatedAt     time.Time          `bson:"created_at"`
	DeliveryLimit time.Duration      `bson:"delivery_limit,omitempty"`
	OrderType     string             `bson:"order_type"`
}

func NewMongoOrderRepository(client *mongo.Client, dbName, collectionName string) *MongoOrderRepository {
	return &MongoOrderRepository{
		collection: client.Database(dbName).Collection(collectionName),
		ctx:        context.Background(),
	}
}

func (r *MongoOrderRepository) Save(order domain.Order) (string, error) {
	// Insere o pedido, o campo ID vazio deixa o Mongo criar o ObjectID
	mo := mongoOrder{
		CustomerID: order.GetCustomerID(),
		Items:      order.GetItems(), // isso deve retornar []OrderItem com product_id preenchido
		Status:     order.GetStatus(),
		CreatedAt:  order.GetCreatedAt(),
		OrderType:  order.GetOrderType(),
		// deliveryLimit, se for pedido express
	}
	res, err := r.collection.InsertOne(r.ctx, mo)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (r *MongoOrderRepository) FindByID(id string) (domain.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	// Tenta decodificar como RegularOrder
	var regOrder domain.RegularOrder
	err = r.collection.FindOne(r.ctx, filter).Decode(&regOrder)
	if err == nil && regOrder.OrderType == domain.OrderTypeRegular {
		return &regOrder, nil
	}

	// Se falhar, tenta como ExpressOrder
	var expOrder domain.ExpressOrder
	err = r.collection.FindOne(r.ctx, filter).Decode(&expOrder)
	if err == nil && expOrder.OrderType == domain.OrderTypeExpress {
		return &expOrder, nil
	}

	// Não achou
	return nil, errors.New("order not found")
}

func (r *MongoOrderRepository) List() ([]domain.Order, error) {
	cur, err := r.collection.Find(r.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(r.ctx)

	var orders []domain.Order
	for cur.Next(r.ctx) {
		// Para cada documento tenta detectar o tipo (pode ser otimizado)
		var regOrder domain.RegularOrder
		if err := cur.Decode(&regOrder); err == nil && regOrder.OrderType == domain.OrderTypeRegular {
			orders = append(orders, &regOrder)
			continue
		}

		var expOrder domain.ExpressOrder
		if err := cur.Decode(&expOrder); err == nil && expOrder.OrderType == domain.OrderTypeExpress {
			orders = append(orders, &expOrder)
			continue
		}
		// Ignora documentos inválidos
	}

	return orders, nil
}
