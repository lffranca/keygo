package oauth2

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func New(options Options) (*Client, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	manager := manage.NewDefaultManager()

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore := store.NewClientStore()
	if err := clientStore.Set(*options.ID, &models.Client{
		ID:     *options.ID,
		Secret: *options.Secret,
		Domain: *options.Domain,
	}); err != nil {
		return nil, err
	}

	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)

	client := new(Client)
	client.common.Client = client
	client.manager = manager
	client.server = srv

	return client, nil
}

type Client struct {
	common  service
	manager oauth2.Manager
	server  *server.Server
}

type service struct {
	Client *Client
}
