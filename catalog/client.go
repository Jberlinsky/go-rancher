package catalog

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type RancherBaseClientImpl struct {
	Opts          *ClientOpts
	Schemas       *Schemas
	Types         map[string]Schema
	customHeaders map[string]string
}

type RancherBaseClient interface {
	Websocket(string, map[string][]string) (*websocket.Conn, *http.Response, error)
	List(string, *ListOpts, interface{}) error
	Post(string, interface{}, interface{}) error
	GetLink(Resource, string, interface{}) error
	Create(string, interface{}, interface{}) error
	Update(string, *Resource, interface{}, interface{}) error
	ById(string, string, interface{}) error
	Delete(*Resource) error
	Reload(*Resource, interface{}) error
	Action(string, string, *Resource, interface{}, interface{}) error
	GetOpts() *ClientOpts
	GetSchemas() *Schemas
	GetTypes() map[string]Schema
	SetCustomHeaders(headers map[string]string)

	DoGet(string, *ListOpts, interface{}) error
	doGet(string, *ListOpts, interface{}) error
	doList(string, *ListOpts, interface{}) error
	doNext(string, interface{}) error
	doModify(string, string, interface{}, interface{}) error
	doCreate(string, interface{}, interface{}) error
	doUpdate(string, *Resource, interface{}, interface{}) error
	doById(string, string, interface{}) error
	doResourceDelete(string, *Resource) error
	doAction(string, string, *Resource, interface{}, interface{}) error
}
