package sdk_booking_service

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/WebXense/ginger/ginger"
	"github.com/WebXense/http"
	"github.com/WebXense/sql"
)

type bookingServiceSdk struct {
	host string
}

func NewBookingServiceSdk(host string) *bookingServiceSdk {
	return &bookingServiceSdk{
		host: host,
	}
}

func (s *bookingServiceSdk) CreateBookObject(name string, allowMultiple, allowOverlap bool) (*BookObjectDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](s.host+url_object, nil, nil, createBookObjectRequest{
		Name:          name,
		AllowMultiple: allowMultiple,
		AllowOverlap:  allowOverlap,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("create fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &BookObjectDTO{}), resp, nil
}

func (s *bookingServiceSdk) ListBookObject(filter *ListBookObjectFilter,
	page *sql.Pagination, sort *sql.Sort) ([]BookObjectDTO, *ginger.Response, error) {

	url := s.host + url_object

	query := make(map[string]string)
	if filter != nil {
		if filter.Name != nil {
			query["name"] = *filter.Name
		}
		if filter.Disabled != nil {
			query["disabled"] = strconv.FormatBool(*filter.Disabled)
		}
		if filter.AllowMultiple != nil {
			query["allow_multiple"] = strconv.FormatBool(*filter.AllowMultiple)
		}
		if filter.AllowOverlap != nil {
			query["allow_overlap"] = strconv.FormatBool(*filter.AllowOverlap)
		}
	}
	if page != nil {
		query["page"] = strconv.Itoa(page.Page)
		query["size"] = strconv.Itoa(page.Size)
	}
	if sort != nil {
		query["asc"] = strconv.FormatBool(sort.Asc)
		query["sortBy"] = sort.SortBy
	}

	status, resp, err := http.Get[ginger.Response](url, nil, query, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("list fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	var data []BookObjectDTO
	for _, d := range resp.Data.([]interface{}) {
		data = append(data, *mapByJson(d, &BookObjectDTO{}))
	}
	return data, resp, nil
}

func (s *bookingServiceSdk) GetBookObject(id uint) (*BookObjectDTO, *ginger.Response, error) {
	status, resp, err := http.Get[ginger.Response](s.host+url_object+"/"+strconv.Itoa(int(id)), nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("get fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &BookObjectDTO{}), resp, nil
}

func (s *bookingServiceSdk) UpdateBookObject(dto *BookObjectDTO) (*BookObjectDTO, *ginger.Response, error) {
	status, resp, err := http.Put[ginger.Response](s.host+url_object, nil, nil, updateBookObjectRequest{
		ID:            dto.ID,
		Name:          dto.Name,
		AllowMultiple: dto.AllowMultiple,
		AllowOverlap:  dto.AllowOverlap,
		Disabled:      dto.Disabled,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("update fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &BookObjectDTO{}), resp, nil
}

func (s *bookingServiceSdk) DeleteBookObject(id uint) (*ginger.Response, error) {
	status, resp, err := http.Delete[ginger.Response](s.host+url_object+"/"+strconv.Itoa(int(id)), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, errors.New("delete fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return resp, nil
	}
	return resp, nil
}

func (s *bookingServiceSdk) Book(bookObjectId, customerId, from, to uint) (*BookingDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](s.host+url_book, nil, nil, bookRequest{
		BookObjectID: bookObjectId,
		CustomerID:   customerId,
		From:         from,
		To:           to,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("book failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &BookingDTO{}), resp, nil
}

func (s *bookingServiceSdk) ListBook(filter *ListBookFilter,
	page *sql.Pagination, sort *sql.Sort) ([]BookingDTO, *ginger.Response, error) {

	url := s.host + url_book

	query := make(map[string]string)
	if filter != nil {
		if filter.CustomerID != nil {
			query["customer_id"] = strconv.Itoa(int(*filter.CustomerID))
		}
		if filter.From != nil {
			query["from"] = strconv.Itoa(int(*filter.From))
		}
		if filter.To != nil {
			query["to"] = strconv.Itoa(int(*filter.To))
		}
	}
	if page != nil {
		query["page"] = strconv.Itoa(page.Page)
		query["size"] = strconv.Itoa(page.Size)
	}
	if sort != nil {
		query["asc"] = strconv.FormatBool(sort.Asc)
		query["sortBy"] = sort.SortBy
	}

	status, resp, err := http.Get[ginger.Response](url, nil, query, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("list fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	var data []BookingDTO
	for _, d := range resp.Data.([]interface{}) {
		data = append(data, *mapByJson(d, &BookingDTO{}))
	}
	return data, resp, nil
}

func (s *bookingServiceSdk) GetBook(id uint) (*BookingDTO, *ginger.Response, error) {
	status, resp, err := http.Get[ginger.Response](s.host+url_book+"/"+strconv.Itoa(int(id)), nil, nil, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("get fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &BookingDTO{}), resp, nil
}

func (s *bookingServiceSdk) UpdateBook(dto *BookingDTO) (*BookingDTO, *ginger.Response, error) {
	status, resp, err := http.Put[ginger.Response](s.host+url_book, nil, nil, updateBookRequest{
		ID:           dto.ID,
		BookObjectID: dto.BookObjectID,
		CustomerID:   dto.CustomerID,
		From:         dto.From,
		To:           dto.To,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("update fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &BookingDTO{}), resp, nil
}

func (s *bookingServiceSdk) DeleteBook(id uint) (*ginger.Response, error) {
	status, resp, err := http.Delete[ginger.Response](s.host+url_book+"/"+strconv.Itoa(int(id)), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if status != 200 {
		return nil, errors.New("delete fragment failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return resp, nil
	}
	return resp, nil
}

func mapByJson[T any](from interface{}, to *T) *T {
	j, err := json.Marshal(from)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(j, to)
	if err != nil {
		return nil
	}
	return to
}
