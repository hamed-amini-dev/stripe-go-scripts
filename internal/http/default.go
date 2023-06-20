package http

/*
default http handler use for all end point request for :
 - add origin to response
 - validate request (validate token ...)
 - create custom http request for adding some item we need use to all handler for example session user
 - run handler functions
*/

import (
	"net/http"

	"github.com/hamed-amini-dev/stripe-go-scripts/internal/request"
	"github.com/hamed-amini-dev/stripe-go-scripts/internal/response"
	"github.com/hamed-amini-dev/stripe-go-scripts/pkg/xerrors"
	"google.golang.org/grpc/status"
)

type customHandlerFunc func(r *request.GenericRequest) (interface{}, error)

// DefaultHTTPHandler - Handles every request
func DefaultHTTPHandler(handler customHandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		var err error
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Methods", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, referrer, Accept-Encoding, User-Agent, Host, Connection")
		if req.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}
		var res response.GenericResponse

		// ─── REQUEST VALIDATION ──────────────────────────────────────────
		err = validate(req)
		if res.HandleError(err) {
			res.ResponseJson(writer, code2Http[status.Code(err)])
			return
		}

		// ─── Create Generic Request  ──────────────────────────────────────────
		genReq, err := request.NewGenericHTTPRequestFromHTTPRequest(req)
		if err != nil {
			respondError(writer, code2Http[status.Code(err)], err.Error())
			return
		}

		// ─── HANDLING AND RUN THE REQUEST ────────────────────────────────────────
		res.Data, err = handler(genReq)
		if err != nil {
			res.HandleError(err)
			res.Success = false
			res.ResponseJson(writer, code2Http[status.Code(err)])
			return
		}

		// ─── FORMATTING THE RESPONSE ──────────────────────────────────────
		res.Success = true
		res.ResponseJson(writer, http.StatusOK)
	}
}

// ────────────────────────────────────────────────────────────────────────────────
// validate returns error if the content type of the request is not valid
// If content type is not `application/json` for the POST and PUT methods,
// then the request is invalid
func validate(request *http.Request) error {
	contentType := request.Header.Get("Content-Type")
	switch request.Method {
	case http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodGet:
		_ = contentType
		return nil
	default:
		return xerrors.ErrMethodNotSupported
	}
}
