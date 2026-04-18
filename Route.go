package gin_auto_router

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"reflect"
	"slices"
	"strings"
)

// Route represents route registration and dispatch information
type Route struct {
	ControllerName string        // Controller struct name (e.g., Article, extracted from *controller.Article)
	ActionName     string        // Controller method name (e.g., ListGet)
	Action         reflect.Value // Controller method reflect value (can be called directly)
	Ext            ExtModel      // Route extension configuration parameters
}

// ExtModel is the route extension configuration model
// NamingConvention: naming method, supported: kebab-case (default), snake_case, camelCase, PascalCase
// HTTPMethod: HTTP request method, supported: {"post", "get", "put", "patch", "head", "options", "delete", "any"}
// URLClass: URL path class part, corresponds to controller name, e.g., /path-to-class/*****
// URLAction: URL path action part, corresponds to method name, e.g., /****/action-name
type ExtModel struct {
	NamingConvention string // URL naming convention (e.g., kebab-case, default)
	HTTPMethod       string // HTTP request method (support: get/post/put/delete/any, etc.)
	URLClass         string // URL path class (corresponds to controller name), e.g., article
	URLAction        string // URL path action (corresponds to method name), e.g., list
}

// Routes stores all registered methods
var Routes = make([]Route, 0)

// Bind performs route binding
func (r *Route) Bind(routerGroup *gin.RouterGroup) {
	r.ParseMetadata()
	switch r.Ext.HTTPMethod {
	case "get":
		routerGroup.GET(r.FullPath(), r.Handler())
	case "put":
		routerGroup.PUT(r.FullPath(), r.Handler())
	case "patch":
		routerGroup.PATCH(r.FullPath(), r.Handler())
	case "head":
		routerGroup.HEAD(r.FullPath(), r.Handler())
	case "options":
		routerGroup.OPTIONS(r.FullPath(), r.Handler())
	case "delete":
		routerGroup.DELETE(r.FullPath(), r.Handler())
	case "any":
		routerGroup.Any(r.FullPath(), r.Handler())
	// Any registers a route that matches all the HTTP methods.
	// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
	// case "post":    // The DEFAULT VALUE is "post".
	default:
		routerGroup.POST(r.FullPath(), r.Handler())
	}
}

// ParseMetadata parses metadata
func (r *Route) ParseMetadata() {
	// Default naming rule is: kebab-case
	urlAction := ToKebabCase(r.ActionName)

	// Split urlAction string and take the last segment to match request type
	// e.g., for list_get, info_push, extract: get and push
	fields := strings.Split(urlAction, "-")
	httpMethod := fields[len(fields)-1]

	if slices.Contains([]string{"post", "get", "put", "patch", "head", "options", "delete", "any"}, httpMethod) {
		// Pattern matched, remove the last segment from urlAction
		// e.g., "user_get" becomes "user"
		urlAction = urlAction[:len(urlAction)-len(httpMethod)-1]
	} else {
		// Pattern not matched, default method is: POST, urlAction remains unchanged
		httpMethod = "post"
	}
	r.Ext.HTTPMethod = httpMethod
	r.Ext.URLClass = r.ConvertByNamingConvention(ToKebabCase(r.ControllerName))
	r.Ext.URLAction = r.ConvertByNamingConvention(urlAction)
}

// FullPath returns the path to be bound
//
// Combines r.Ext.URLClass and r.Ext.URLAction to form the complete path, e.g., /path-to-class/action-name.
//
// Parameters:
//
//	none
//
// Returns:
//
//	string: Complete path string
func (r *Route) FullPath() (path string) {
	// Join path, e.g., /path-to-class/action-name
	path, _ = url.JoinPath("/", r.Ext.URLClass, r.Ext.URLAction)
	return
}

// Handler converts the controller method to gin.HandlerFunc.
// This passes gin.Context as a parameter to the method, limited to one parameter: gin.Context.
// Method definition example: func (api *Article) ListGet(c *gin.Context)
//
// Converts the Action method in the Route struct to gin.HandlerFunc for use in routing.
//
// Parameters:
//
//	none
//
// Returns:
//
//	gin.HandlerFunc: A function type that can be used for Gin route handling
func (r *Route) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		arguments := make([]reflect.Value, 1)
		arguments[0] = reflect.ValueOf(c)
		r.Action.Call(arguments)
	}
}

// ConvertByNamingConvention converts string to the specified naming format
//
// Converts a kebab-case formatted string to the specified format based on r.Ext.NamingConvention.
//
// Supported naming conventions:
//   - "kebab-case": kebab-case (default, no conversion)
//   - "snake_case": snake_case (e.g., article-list -> article_list)
//   - "camelCase": camelCase (e.g., article-list -> articleList)
//   - "PascalCase": PascalCase (e.g., article-list -> ArticleList)
//
// Parameters:
//
//	str: Input kebab-case formatted string
//
// Returns:
//
//	String converted according to the naming convention
func (r *Route) ConvertByNamingConvention(str string) string {
	// Return empty string directly to avoid panic
	if str == "" {
		return ""
	}
	switch r.Ext.NamingConvention {
	case "camelCase":
		return KebabCaseToCamelCase(str)
	case "PascalCase":
		return KebabCaseToPascalCase(str)
	case "snake_case":
		return KebabCaseToSnakeCase(str)
	case "kebab-case":
		fallthrough //
	default:
		return str
	}
}
