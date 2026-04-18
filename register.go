package gin_auto_router

import (
	"reflect"
)

// Register registers all exported methods of a controller as routes.
//
// Parameters:
//
//	controller - Controller instance (must be a pointer)
//
// Returns:
//
//	bool - Returns true if registration succeeds, false if no methods can be registered
//
// Example:
//
//	import (
//	ginAutoRouter "gitee.com/dengshulei/gin-auto-router"
//	"github.com/gin-gonic/gin"
//	"net/http"
//
//	)
//
//	func init() {
//		ginAutoRouter.Register(&Article{})
//	}
//
//	type Article struct{}
//	func (api *Article) ListGet(c *gin.Context) {
//	}
func Register(controller interface{}) bool {
	// Get the reflect.Value of the controller
	value := reflect.ValueOf(controller)
	// Check if the controller has exported methods, return false if none
	if value.NumMethod() == 0 {
		return false
	}
	valueType := value.Type() // Cache type to avoid repeated lookups
	// Iterate through all methods of the controller
	for i := 0; i < value.NumMethod(); i++ {
		// Get the i-th method (method object, not method name)
		method := value.Method(i)
		// Get method type information
		methodType := valueType.Method(i)
		// value.Type().Method(i).Name: Get the i-th method name
		Routes = append(Routes, Route{
			ControllerName: GetControllerName(controller),
			ActionName:     methodType.Name,
			Action:         method,
			Ext:            ExtModel{},
		})
	}
	return true
}
