package controller

import (
	"abduselam-arabianmejlis/delivery/controller"
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/mocks"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type productControllerSuite struct {
	// we need this to use the suite functionalities from testify
	suite.Suite
	// the mocked version of the usecase
	usecase *mocks.ProductUseCase
	// the functionalities we need to test
	controller *controller.ProductController
	// testing server to be used the handler
	testingServer   *httptest.Server
}


func (suite *productControllerSuite) SetupSuite() {
	// create a mocked version of usecase
	usecase := new(mocks.ProductUseCase)
	// inject the usecase to be used by handler
	controller := controller.NewProductController(usecase, nil)

	// create default server using gin, then register all endpoints
	router := gin.Default()

	router.POST("/products", controller.CreateProduct)
	router.GET("/products/:id", controller.GetProductByID)
	router.GET("/products", controller.GetProducts)
	router.PUT("/products/:id", controller.UpdateProduct)
	router.DELETE("/products/:id", controller.DeleteProduct)
	router.GET("/products/search", controller.SearchProducts)
	router.POST("/products/:id/like", controller.LikeProduct)
	router.POST("/products/:id/unlike", controller.UnlikeProduct)

	// create and run the testing server
	testingServer := httptest.NewServer(router)

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.testingServer = testingServer
	suite.usecase = usecase
	suite.controller = controller
}

func (suite *productControllerSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

func (suite *productControllerSuite) TestCreateProduct_Positive() {
	// an example tweet for the test
	product := domain.Product{
		ID: primitive.NewObjectID(),
		Title: "Product Name",
		Images: []string{"image1", "image2"},
		Description: "Product Description",
		Category: "Product Category",
		Features: []string{"feature1", "feature2"},
		Tags: []string{"tag1", "tag2"},
		ColorOptions: map[string]string{"blue":"image3", "red": "image4",},
		IsAvailable: true,
		Views: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.usecase.On("CreateProduct",mock.Anything, &product).Return(&product, nil)

	// marshalling and some assertion
	requestBody, err := json.Marshal(&product)
	suite.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/products", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// unmarshalling the response
	responseBody := domain.SuccessResponse{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// running assertions to make sure that our method does the correct thing
	suite.Equal(http.StatusCreated, response.StatusCode) 
	suite.Equal(responseBody.Message, "Product Created Successfully")
	suite.usecase.AssertExpectations(suite.T())
}


func TestProductController(t *testing.T) {
	suite.Run(t, new(productControllerSuite))
}
