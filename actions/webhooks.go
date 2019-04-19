package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/ulthuan/glapi/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Webhook)
// DB Table: Plural (webhooks)
// Resource: Plural (Webhooks)
// Path: Plural (/webhooks)
// View Template Folder: Plural (/templates/webhooks/)

// WebhooksResource is the resource for the Webhook model
type WebhooksResource struct {
	buffalo.Resource
}

// List gets all Webhooks. This function is mapped to the path
// GET /webhooks
func (v WebhooksResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	webhooks := &models.Webhooks{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Webhooks from the DB
	if err := q.All(webhooks); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, webhooks))
}

// Show gets the data for one Webhook. This function is mapped to
// the path GET /webhooks/{webhook_id}
func (v WebhooksResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Webhook
	webhook := &models.Webhook{}

	// To find the Webhook the parameter webhook_id is used.
	if err := tx.Find(webhook, c.Param("webhook_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, webhook))
}

// Create adds a Webhook to the DB. This function is mapped to the
// path POST /webhooks
func (v WebhooksResource) Create(c buffalo.Context) error {
	// Allocate an empty Webhook
	webhook := &models.Webhook{}

	webhookcard := &models.WebhookCardEvent{}

	webhookaction := &models.CardWebHook{}

	// Bind webhook to the html form elements
	if err := c.Bind(webhookcard); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	webhook.Action = webhookcard.Action
	webhook.BoardID = webhookcard.Board.ID
	webhook.CardID = webhookcard.Card.ID
	webhook.SenderUsername = webhookcard.Sender.Username

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(webhook)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, webhook))
	}

	webhookaction.ExecuteAction(webhookcard)

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "webhook.created.success"))
	// and redirect to the webhooks index page
	return c.Render(201, r.Auto(c, webhook))
}

// Destroy deletes a Webhook from the DB. This function is mapped
// to the path DELETE /webhooks/{webhook_id}
func (v WebhooksResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Webhook
	webhook := &models.Webhook{}

	// To find the Webhook the parameter webhook_id is used.
	if err := tx.Find(webhook, c.Param("webhook_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(webhook); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "webhook.destroyed.success"))
	// Redirect to the webhooks index page
	return c.Render(200, r.Auto(c, webhook))
}
