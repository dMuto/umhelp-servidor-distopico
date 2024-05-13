package validation

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/dMuto/umhelp-servidor-distopico/presenter/req"
)

func VerifyNewUserRequest(rc io.ReadCloser) (r *req.User, err error){
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("Invalid read closer")
	}

	if err := json.Unmarshal(body, &r); err != nil{
		return nil, errors.New("Invalid JSON payload")
	}

	return r, nil
}