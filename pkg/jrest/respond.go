package jrest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"github.com/Insulince/jlib/pkg/jlog"
)

func Respond(ctx context.Context, w http.ResponseWriter, status int, payload []byte) {
	ctx, logger := jlog.FromContextSafe(ctx)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	if _, err := w.Write(payload); err != nil {
		logger.Error(errors.Wrap(err, "writing response"))
	}
}

func RespondString(ctx context.Context, w http.ResponseWriter, status int, payload string) {
	Respond(ctx, w, status, []byte(payload))
}

func RespondError(ctx context.Context, w http.ResponseWriter, status int, err error) {
	RespondString(ctx, w, status, err.Error())
}

func RespondErrorWithLog(ctx context.Context, w http.ResponseWriter, status int, err error) {
	ctx, logger := jlog.FromContextSafe(ctx)

	logger.Error(err)
	RespondError(ctx, w, status, err)
}

func RespondDefault(ctx context.Context, w http.ResponseWriter, status int) {
	RespondString(ctx, w, status, http.StatusText(status))
}

func RespondJson(ctx context.Context, w http.ResponseWriter, status int, v interface{}) {
	ctx, logger := jlog.FromContextSafe(ctx)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		logger.Error(errors.Wrap(err, "writing response"))
	}
}

func RespondJsonProto(ctx context.Context, w http.ResponseWriter, status int, msg proto.Message) {
	ctx, logger := jlog.FromContextSafe(ctx)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	m := jsonpb.Marshaler{}
	if err := m.Marshal(w, msg); err != nil {
		logger.Error(errors.Wrap(err, "writing response"))
	}
}
