package grafana

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"grafana-api/infrastructure/http/common"
	"grafana-api/infrastructure/secretmanager"
)

type IClient interface {
	GetOrg(ctx context.Context, name string) (*OrganizationDTO, error)
	GetUser(ctx context.Context, name string) (*UserDTO, error)
	CreateUser(ctx context.Context, cmd CreateUserCommand) error
}

type client struct {
	conf secretmanager.ClientSecret
}

func NewClient(c secretmanager.ClientSecret) IClient {
	return &client{
		conf: c,
	}
}

func (c *client) GetOrg(ctx context.Context, name string) (*OrganizationDTO, error) {
	r := common.AcquireResource()
	defer r.Release()

	r.Request.SetRequestURI(fmt.Sprintf("%s/api/orgs/name/%s", c.conf.Host, name))
	r.Request.Header.SetMethod(fasthttp.MethodGet)
	r.Request.Header.Set("Authorization", c.conf.APIKey)

	if err := fasthttp.Do(r.Request, r.Response); err != nil {
		return nil, err
	}

	if common.IsFailure(r.Response) {
		return nil, common.NewErrorResponse(r.Response)
	}

	orgDTO := OrganizationDTO{}
	if err := json.Unmarshal(r.Response.Body(), &orgDTO); err != nil {
		return nil, err
	}

	return &orgDTO, nil
}

func (c *client) GetUser(ctx context.Context, name string) (*UserDTO, error) {
	r := common.AcquireResource()
	defer r.Release()

	r.Request.SetRequestURI(fmt.Sprintf("%s/api/users/lookup", c.conf.Host))
	r.Request.Header.SetMethod(fasthttp.MethodGet)
	r.Request.Header.Set("Authorization", c.conf.APIKey)
	r.Request.URI().QueryArgs().Set("loginOrEmail", name)

	if err := fasthttp.Do(r.Request, r.Response); err != nil {
		return nil, err
	}

	if common.IsFailure(r.Response) {
		return nil, common.NewErrorResponse(r.Response)
	}

	userDTO := UserDTO{}
	if err := json.Unmarshal(r.Response.Body(), &userDTO); err != nil {
		return nil, err
	}

	return &userDTO, nil
}

func (c *client) CreateUser(ctx context.Context, cmd CreateUserCommand) error {
	r := common.AcquireResource()
	defer r.Release()

	r.Request.SetRequestURI(fmt.Sprintf("%s/api/admin/users", c.conf.Host))
	r.Request.Header.SetMethod(fasthttp.MethodPost)
	r.Request.Header.Set("Authorization", c.conf.APIKey)
	r.Request.Header.Set("Content-Type", "application/json")

	rawBody, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	r.Request.SetBodyRaw(rawBody)

	if err = fasthttp.Do(r.Request, r.Response); err != nil {
		return err
	}

	if common.IsFailure(r.Response) {
		return common.NewErrorResponse(r.Response)
	}

	return nil
}
