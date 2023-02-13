package sdk_booking_service

type BookObjectDTO struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Disabled      bool   `json:"disabled"`
	AllowMultiple bool   `json:"allow_multiple"`
	AllowOverlap  bool   `json:"allow_overlap"`
}

type BookingDTO struct {
	ID           uint `json:"id"`
	CustomerID   uint `json:"customer_id"`
	BookObjectID uint `json:"book_object_id"`
	From         uint `json:"from"`
	To           uint `json:"to"`
}

type ListBookObjectFilter struct {
	Name          *string `form:"name"`
	Disabled      *bool   `form:"disabled"`
	AllowMultiple *bool   `form:"allow_multiple"`
	AllowOverlap  *bool   `form:"allow_overlap"`
}

type ListBookFilter struct {
	CustomerID *uint `form:"customer_id"`
	From       *uint `form:"from"`
	To         *uint `form:"to"`
}

type createBookObjectRequest struct {
	Name          string `json:"name" binding:"required"`
	AllowMultiple bool   `json:"allow_multiple" binding:"required"`
	AllowOverlap  bool   `json:"allow_overlap" binding:"required"`
}

type updateBookObjectRequest struct {
	ID            uint   `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	AllowMultiple bool   `json:"allow_multiple" binding:"required"`
	AllowOverlap  bool   `json:"allow_overlap" binding:"required"`
	Disabled      bool   `json:"disabled" binding:"required"`
}

type bookRequest struct {
	CustomerID   uint `json:"customer_id" binding:"required"`
	BookObjectID uint `json:"book_object_id" binding:"required"`
	From         uint `json:"from" binding:"required"`
	To           uint `json:"to" binding:"required"`
}

type updateBookRequest struct {
	ID           uint `json:"id" binding:"required"`
	BookObjectID uint `json:"book_object" binding:"required"`
	CustomerID   uint `json:"customer_id" binding:"required"`
	From         uint `json:"from" binding:"required"`
	To           uint `json:"to" binding:"required"`
}
