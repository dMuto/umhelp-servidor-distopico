package validation

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"

	"github.com/dMuto/umhelp-servidor-distopico/presenter/req"
)

func VerifyNewUserRequest(rc io.ReadCloser) (r *req.User, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("Invalid read closer")
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, errors.New("Invalid JSON payload")
	}

	if len(r.Name) <= 2 {
		return nil, errors.New("First name too short, please type at least 3 characters")
	}

	if len(r.Name) > 40 {
		return nil, errors.New("First name too long, please keep it at under 40 characters")
	}

	if len(r.LastName) <= 2 {
		return nil, errors.New("Last name too short, please type at least 3 characters")
	}

	if len(r.LastName) > 40 {
		return nil, errors.New("Last name too long, please keep it at under 40 characters")
	}

	rgx := regexp.MustCompile(`^[0-9]{3}\.[0-9]{3}\.[0-9]{3}-[0-9]{2}$`)
	if !rgx.MatchString(r.Document){
		return nil, errors.New("Document number invalid, please type 14 characters in the format 123.123.456-50")
	}

	rgsx := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !rgsx.MatchString(r.Email){
		return nil, errors.New("Email is not valid, please type email in the format example@example.com")
	}

	if len(r.Password) <= 7 {
		return nil, errors.New("Password too short, please type at least 8 characters")
	}

	if len(r.Password) > 16 {
		return nil, errors.New("Password Last name too long, please keep it at under 16 characters")
	}

	return r, nil
}
