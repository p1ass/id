package zerologgcloud

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog/hlog"
)

func NewCloudLoggingRequestHandler() func(http.Handler) http.Handler {
	return hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		log.Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Str("userAgent", r.Header.Get("User-Agent")).
			Str("referer", r.Header.Get("Referer")).
			Str("proto", r.Proto).
			Msg("")
	},
	)
}
