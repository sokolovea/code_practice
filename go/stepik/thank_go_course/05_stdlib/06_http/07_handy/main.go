package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// начало решения

// Handy предоставляет удобный интерфейс
// для выполнения HTTP-запросов
type Handy struct {
	client *http.Client
	url    string
	header *http.Header
	param  *url.Values
	body   io.Reader
	err    error
}

// NewHandy создает новый экземпляр Handy
func NewHandy() *Handy {
	return &Handy{client: http.DefaultClient,
		header: &http.Header{},
		param:  &url.Values{},
	}
}

// URL устанавливает URL, на который пойдет запрос
func (h *Handy) URL(uri string) *Handy {
	h.url = uri
	return h
}

// Client устанавливает HTTP-клиента
// вместо умолчательного http.DefaultClient
func (h *Handy) Client(client *http.Client) *Handy {
	h.client = client
	return h
}

// Header устанавливает значение заголовка
func (h *Handy) Header(key, value string) *Handy {
	h.header.Set(key, value)
	return h
}

// Param устанавливает значение URL-параметра
func (h *Handy) Param(key, value string) *Handy {
	h.param.Set(key, value)
	return h
}

// Form устанавливает данные, которые будут закодированы
// как application/x-www-form-urlencoded и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) Form(form map[string]string) *Handy {
	h.header.Set("Content-Type", "application/x-www-form-urlencoded")
	sb := strings.Builder{}
	pairNumber := 1
	for key, value := range form {
		sb.WriteString(key)
		sb.WriteByte('=')
		sb.WriteString(value)
		if pairNumber != len(form) {
			sb.WriteByte('&')
		}
		pairNumber++
	}
	h.body = strings.NewReader(sb.String())
	return h
}

// JSON устанавливает данные, которые будут закодированы
// как application/json и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) JSON(v any) *Handy {
	h.header.Set("Content-Type", "application/json")
	b, err := json.Marshal(v)
	h.err = err
	if h.err != nil {
		h.body = bytes.NewReader(b)
	}
	return h
}

// Get выполняет GET-запрос с настроенными ранее параметрами
func (h *Handy) Get() *HandyResponse {
	if h.err != nil {
		return &HandyResponse{err: h.err}
	}
	req, err := http.NewRequest(http.MethodGet, h.url, nil)
	if err != nil {
		return &HandyResponse{err: err}
	}
	req.Header = *h.header
	req.URL.RawQuery = h.param.Encode()
	resp, err := h.client.Do(req)
	if err != nil {
		return &HandyResponse{err: err}
	}
	return &HandyResponse{StatusCode: resp.StatusCode, body: resp.Body, err: nil}
}

// Post выполняет POST-запрос с настроенными ранее параметрами
func (h *Handy) Post() *HandyResponse {
	if h.err != nil {
		return &HandyResponse{err: h.err}
	}
	req, err := http.NewRequest(http.MethodPost, h.url, h.body)
	if err != nil {
		return &HandyResponse{err: err}
	}
	req.Header = *h.header
	req.URL.RawQuery = h.param.Encode()
	resp, err := h.client.Do(req)
	if err != nil {
		return &HandyResponse{err: err}
	}
	return &HandyResponse{StatusCode: resp.StatusCode, body: resp.Body, err: nil}
}

// HandyResponse представляет ответ на HTTP-запрос
type HandyResponse struct {
	StatusCode int
	body       io.ReadCloser
	err        error
}

// OK возвращает true, если во время выполнения запроса
// не произошло ошибок, а код HTTP-статуса ответа равен 200
func (r *HandyResponse) OK() bool {
	return r.err == nil && r.StatusCode == 200
}

// Bytes возвращает тело ответа как срез байт
func (r *HandyResponse) Bytes() []byte {
	if r.body != nil {
		defer r.body.Close()
		res, err := io.ReadAll(r.body)
		r.err = err
		if err != nil {
			return []byte{}
		}
		return res
	}
	return []byte{}
}

// String возвращает тело ответа как строку
func (r *HandyResponse) String() string {
	return string(r.Bytes())
}

// JSON декодирует тело ответа из JSON и сохраняет
// результат по адресу, на который указывает v
func (r *HandyResponse) JSON(v any) {
	rawData := r.Bytes()
	r.err = json.Unmarshal(rawData, v)
	// работает аналогично json.Unmarshal()
	// если при декодировании произошла ошибка,
	// она должна быть доступна через r.Err()
}

// Err возвращает ошибку, которая возникла при выполнении запроса
// или обработке ответа
func (r *HandyResponse) Err() error {
	return r.err
}

// конец решения

func main() {
	{
		// примеры запросов

		// GET-запрос с параметрами
		NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()

		// HTTP-заголовки
		NewHandy().
			URL("https://httpbingo.org/get").
			Header("Accept", "text/html").
			Header("Authorization", "Bearer 1234567890").
			Get()

		// POST формы
		params := map[string]string{
			"brand":    "lg",
			"category": "tv",
		}
		NewHandy().URL("https://httpbingo.org/post").Form(params).Post()

		// POST JSON-документа
		NewHandy().URL("https://httpbingo.org/post").JSON(params).Post()
	}

	{
		// пример обработки ответа

		// отправляем GET-запрос с параметрами
		resp := NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()
		if !resp.OK() {
			panic(resp.String())
		}

		// декодируем ответ в JSON
		var data map[string]any
		resp.JSON(&data)

		fmt.Println(data["url"])
		// "https://httpbingo.org/get"
		fmt.Println(data["args"])
		// map[id:[42]]
	}
}
