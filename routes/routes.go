package routes

import (
	"net/http"

	"github.com/drinks/handlers"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"AddDrink",
		"POST",
		"/",
		handlers.AddDrink,
	},
	Route{
		"UpdateDrink",
		"UPDATE",
		"/",
		handlers.UpdateDrink,
	},
	Route{
		"DeleteDrink",
		"DELETE",
		"/",
		handlers.DeleteDrink,
	},
}
