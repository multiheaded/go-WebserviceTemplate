package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"multiheaded/webservice_template/pkg/crudapi"
	"multiheaded/webservice_template/pkg/datamodel"
	gorm2 "multiheaded/webservice_template/pkg/storage/gorm"
)

// A WebserviceTemplateApp serves as a storage for the instances to access database, object storage, ...
// A new instance of the app will be created by the app's entry point
// The template includes
// * database access via the gorm ORM
type WebserviceTemplateApp struct {
	DB        *gorm.DB
	WebEngine *gin.Engine
}

func (app WebserviceTemplateApp) Run() {
	// and start the backend engine/go-gin on a specific socket
	app.WebEngine.Run("0.0.0.0:8080")
}

func routeEndpoints[T any](grp *gin.RouterGroup, database *gorm.DB) error {
	// generic controller mapping CRUD functions  directly to database operations
	repo := gorm2.NewGormRepository[T](database)

	// generic API handler calling CRUD functions of controller given HTTP requests
	handler := crudapi.NewGinCRUDHandler[T](repo)

	// CRUD API for T
	crudapi.GinRouteCRUD[T](grp, handler)

	return nil
}

// NewInstance sets up a new instance of the WebserviceTemplateApp
func NewInstance() (*WebserviceTemplateApp, error) {
	// connect and login to the database used to persist objects
	db, err := openDatabase()

	if err != nil {
		return nil, err
	}

	// initialize the underlying web framework to rout endpoints eventaully
	web, err := initWebFramework()

	if err != nil {
		return nil, err
	}

	// all apiEndpoint requests are going to be below the /api uri
	apiGroup := web.Group("/api")

	err = routeEndpoints[datamodel.Dummy](apiGroup, db)

	if err != nil {
		return nil, err
	}

	lh := &WebserviceTemplateApp{
		DB:        db,
		WebEngine: web,
	}

	return lh, nil
}
