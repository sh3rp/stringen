package server

import (
	"encoding/base64"
	"encoding/json"
	"github.com/sh3rp/stringen/pkg/codec"
	"github.com/sh3rp/stringen/pkg/log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var InputDataField = "input"
var CountDataField = "count"
var RawResponseDataField = "raw"
var CharCountDataField = "num"
var CharTypeDataField = "type"

type Service interface {
	Serve(string) error
}

type service struct {
	router *http.ServeMux
}

func NewService(mux *http.ServeMux) Service {
	return &service{mux}
}

func (s *service) Serve(port string) error {
	log.LOGGER.Info().Msgf("Compiling routes...")
	s.routes()
	log.LOGGER.Info().Msgf("Starting server on %s", port)
	return http.ListenAndServe(port, s.router)
}

func (s *service) routes() {
	s.router.HandleFunc("/v1/sha256", s.requestLogger(s.sha256HashService()))
	s.router.HandleFunc("/v1/uuid", s.requestLogger(s.uuidService()))
	s.router.HandleFunc("/v1/base64Encode", s.requestLogger(s.base64EncodeService()))
	s.router.HandleFunc("/v1/base64Decode", s.requestLogger(s.base64DecodeService()))
	s.router.HandleFunc("/v1/ulid", s.requestLogger(s.ulidService()))
	s.router.HandleFunc("/v1/random", s.requestLogger(s.randomCharsService()))
}

func (s *service) requestLogger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		log.LOGGER.Info().Msgf("[r] src: %s - %s (%s) [%s]", r.RemoteAddr, r.URL.Path, r.Method, r.Form)
		h(w, r)
	}
}

func (s *service) randomCharsService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var numChars int
			if r.URL.Query().Get(CharCountDataField) != "" {
				numChars, _ = strconv.Atoi(r.URL.Query().Get(CharCountDataField))
			} else {
				numChars = 20
			}
			var charType int
			if r.URL.Query().Get(CharTypeDataField) != "" {
				charType, _ = strconv.Atoi(r.URL.Query().Get(CharTypeDataField))
			} else {
				charType = codec.CharTypeAlphaNumericSpecial
			}

			str := codec.GenRandomCharacters(numChars, codec.CharType(charType))

			if isRaw(r) {
				w.Write([]byte(str))
				return
			}
			successResponse(w, str)
		}
	}
}

func (s *service) base64DecodeService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			r.ParseForm()
			input := r.Form.Get(InputDataField)

			encoding, _ := base64.StdEncoding.DecodeString(input)
			if isRaw(r) {
				w.Write(encoding)
				return
			}
			successResponse(w, string(encoding))
		}
	}
}

func (s *service) base64EncodeService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			r.ParseForm()
			input := r.Form.Get(InputDataField)

			encoding := base64.StdEncoding.EncodeToString([]byte(input))
			if isRaw(r) {
				w.Write([]byte(encoding))
				return
			}
			successResponse(w, encoding)
		}
	}
}

func (s *service) ulidService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			count := 1
			if r.URL.Query().Get(CountDataField) != "" {
				countStr := r.URL.Query().Get(CountDataField)
				count, _ = strconv.Atoi(countStr)
			}
			var ids []string
			for i := 0; i < count; i++ {
				ids = append(ids, strings.ToLower(codec.GenUlid()))
			}
			str := strings.Join(ids, ",")
			if isRaw(r) {
				w.Write([]byte(str))
				return
			}
			successResponse(w, ids)
		}
	}
}

func (s *service) uuidService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			count := 1
			if r.URL.Query().Get(CountDataField) != "" {
				countStr := r.URL.Query().Get(CountDataField)
				count, _ = strconv.Atoi(countStr)
			}
			var ids []string
			for i := 0; i < count; i++ {
				id, _ := uuid.NewRandom()
				ids = append(ids, id.String())
			}
			str := strings.Join(ids, ",")
			if isRaw(r) {
				w.Write([]byte(str))
				return
			}
			successResponse(w, ids)
		}
	}
}

func (s *service) sha256HashService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			randChars := codec.GenRandomCharacters(20, codec.CharTypeAlphaNumericSpecial)
			hashStr := codec.GenSha256Hash(randChars)
			if isRaw(r) {
				w.Write([]byte(hashStr))
				return
			}
			successResponse(w, hashStr)
		case "POST":
			r.ParseForm()
			input := r.Form.Get(InputDataField)
			hashStr := codec.GenSha256Hash(input)
			if isRaw(r) {
				w.Write([]byte(hashStr))
				return
			}
			successResponse(w, hashStr)
		default:

		}
	}
}

func successResponse(w http.ResponseWriter, data interface{}) {
	writeResponse(w, 0, "Ok", data)
}

func writeResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	response := &Response{code, message, data}

	bytes, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(bytes)
}

func isRaw(r *http.Request) bool {
	return r.Form.Get(RawResponseDataField) != "" || r.URL.Query().Get(RawResponseDataField) != ""
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
