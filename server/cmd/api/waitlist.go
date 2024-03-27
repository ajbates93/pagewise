package main

import (
	"errors"
	"net/http"
	"pagewise/internal/data"
)

func (app *application) registerWaitlistUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	waitlist := &data.Waitlist{
		Email: input.Email,
	}

	err = app.models.Waitlist.Insert(waitlist)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			// Add some validation stuff here
			app.failedValidationResponse(w, r, map[string]string{"email": "a user with this email address already exists"})
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.mailer.Send(waitlist.Email, "waitlist_confirmation.tmpl", waitlist)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"success": true, "message": "email added to waitlist"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
