package printavo

import (
	"fmt"
	"net/http"
	"time"
)

// OrdersService handles communication with the orders related methods of the
// Printavo API.
//
// Printavo API docs: https://printavo.docs.apiary.io/#reference/orders
type OrdersService struct {
	client *Client
}

type Order struct {
	Id                           int                   `json:"id,omitempty"`
	SalesTax                     float32               `json:"sales_tax,omitempty"`
	TotalUntaxed                 float32               `json:"total_untaxed,omitempty"`
	DiscountAsPercentage         bool                  `json:"discount_as_percentage,omitempty"`
	Discount                     float32               `json:"discount,omitempty"`
	CustomerId                   int                   `json:"customer_id,omitempty"`
	UserId                       int                   `json:"user_id,omitempty"`
	OrderstatusId                int                   `json:"orderstatus_id,omitempty"`
	PublicHash                   string                `json:"public_hash,omitempty"`
	ProductionNotes              interface{}           `json:"production_notes,omitempty"`
	OrderNickname                interface{}           `json:"order_nickname,omitempty"`
	Approved                     bool                  `json:"approved,omitempty"`
	OrderSubtotal                float32               `json:"order_subtotal,omitempty"`
	AmountPaid                   float32               `json:"amount_paid,omitempty"`
	AmountOutstanding            float32               `json:"amount_outstanding,omitempty"`
	ApprovedName                 interface{}           `json:"approved_name,omitempty"`
	VisualId                     int                   `json:"visual_id,omitempty"`
	Stats                        Stats                 `json:"stats,omitempty"`
	Notes                        string                `json:"notes,omitempty"`
	CreatedAt                    time.Time             `json:"created_at,omitempty"`
	UpdatedAt                    time.Time             `json:"updated_at,omitempty"`
	DueDate                      time.Time             `json:"due_date,omitempty"`
	OrderTotal                   float32               `json:"order_total,omitempty"`
	CustomerDueDate              time.Time             `json:"customer_due_date,omitempty"`
	InvoiceDate                  time.Time             `json:"invoice_date,omitempty"`
	PaymentDueDate               time.Time             `json:"payment_due_date,omitempty"`
	CustomCreatedAt              time.Time             `json:"custom_created_at,omitempty"`
	FormattedCustomCreatedAtDate string                `json:"formatted_custom_created_at_date,omitempty"`
	FormattedInvoiceDate         string                `json:"formatted_invoice_date,omitempty"`
	FormattedCustomerDueDate     string                `json:"formatted_customer_due_date,omitempty"`
	FormattedPaymentDueDate      string                `json:"formatted_payment_due_date,omitempty"`
	DeliveryMethodId             int                   `json:"delivery_method_id,omitempty"`
	PaymentTermId                int                   `json:"payment_term_id,omitempty"`
	Expenses                     []Expenses            `json:"expenses,omitempty"`
	Customer                     Customer              `json:"customer,omitempty"`
	Messages                     []Messages            `json:"messages,omitempty"`
	Tasks                        []Tasks               `json:"tasks,omitempty"`
	User                         User                  `json:"user,omitempty"`
	Orderstatus                  Orderstatus           `json:"orderstatus,omitempty"`
	LineitemsAttributes          []LineitemsAttributes `json:"lineitems_attributes,omitempty"`
	OrderFees                    []interface{}         `json:"order_fees,omitempty"`
	Url                          string                `json:"url,omitempty"`
	PublicURL                    string                `json:"public_url,omitempty"`
	Pdf                          string                `json:"pdf,omitempty"`
	Workorder                    string                `json:"workorder,omitempty"`
	PackagingSlip                string                `json:"packaging_slip,omitempty"`
}

type Stats struct {
	Paid bool `json:"paid,omitempty"`
}

type Expenses struct {
	Id              int       `json:"id,omitempty"`
	TransactionDate time.Time `json:"transaction_date,omitempty"`
	Name            string    `json:"name,omitempty"`
	Amount          float32   `json:"amount,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	UserGenerated   bool      `json:"user_generated,omitempty"`
}

type Customer struct {
	FullName  string `json:"full_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Email     string `json:"email,omitempty"`
}

type Messages struct {
	From                    string `json:"from,omitempty"`
	To                      string `json:"to,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	Text                    string `json:"text,omitempty"`
	FormattedDeliveryStatus string `json:"formatted_delivery_status,omitempty"`
}

type Tasks struct {
	Id                int         `json:"id,omitempty"`
	OrderId           int         `json:"order_id,omitempty"`
	TaskName          string      `json:"task_name,omitempty"`
	DueDate           time.Time   `json:"due_date,omitempty"`
	Completed         bool        `json:"completed,omitempty"`
	CompletedDate     interface{} `json:"completed_date,omitempty"`
	CreatedAt         time.Time   `json:"created_at,omitempty"`
	UpdatedAt         time.Time   `json:"updated_at,omitempty"`
	PresetTaskCreated bool        `json:"preset_task_created,omitempty"`
	AssignedTo        string      `json:"assigned_to,omitempty"`
	CompletedBy       interface{} `json:"completed_by,omitempty"`
}

type User struct {
	Name string `json:"name,omitempty"`
}

type Orderstatus struct {
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

type ImagesAttributes struct {
	FilepickerURL        string `json:"filepicker_url,omitempty"`
	MimeType             string `json:"mime_type,omitempty"`
	FullImageURL         string `json:"full_image_url,omitempty"`
	Thumbnail100By100URL string `json:"thumbnail_100_by_100_url,omitempty"`
	DisplayThumbnail     string `json:"display_thumbnail,omitempty"`
}

type LineitemsAttributes struct {
	Id                       int                `json:"id,omitempty"`
	StyleDescription         string             `json:"style_description,omitempty"`
	Taxable                  bool               `json:"taxable,omitempty"`
	StyleNumber              string             `json:"style_number,omitempty"`
	Color                    string             `json:"color,omitempty"`
	SizeOther                int                `json:"size_other,omitempty"`
	SizeYxs                  int                `json:"size_yxs,omitempty"`
	SizeYs                   int                `json:"size_ys,omitempty"`
	SizeYm                   int                `json:"size_ym,omitempty"`
	SizeYl                   int                `json:"size_yl,omitempty"`
	SizeYxl                  int                `json:"size_yxl,omitempty"`
	Size6M                   int                `json:"size_6m,omitempty"`
	Size12M                  int                `json:"size_12m,omitempty"`
	Size18M                  int                `json:"size_18m,omitempty"`
	Size24M                  int                `json:"size_24m,omitempty"`
	Size2T                   int                `json:"size_2t,omitempty"`
	Size3T                   int                `json:"size_3t,omitempty"`
	Size4T                   int                `json:"size_4t,omitempty"`
	Size5T                   int                `json:"size_5t,omitempty"`
	SizeS                    int                `json:"size_s,omitempty"`
	SizeM                    int                `json:"size_m,omitempty"`
	SizeL                    int                `json:"size_l,omitempty"`
	SizeXl                   int                `json:"size_xl,omitempty"`
	Size2Xl                  int                `json:"size_2xl,omitempty"`
	Size3Xl                  int                `json:"size_3xl,omitempty"`
	Size4Xl                  int                `json:"size_4xl,omitempty"`
	Size5Xl                  int                `json:"size_5xl,omitempty"`
	Size6Xl                  int                `json:"size_6xl,omitempty"`
	TotalQuantities          int                `json:"total_quantities,omitempty"`
	GoodsStatus              string             `json:"goods_status,omitempty"`
	PrintLocationsAttributes []interface{}      `json:"print_locations_attributes,omitempty"`
	ImagesAttributes         []ImagesAttributes `json:"images_attributes,omitempty"`
	UnitCost                 float32            `json:"unit_cost,omitempty"`
	Category                 interface{}        `json:"category,omitempty"`
}

func (s *OrdersService) List() (*ListResponse[*Order], *http.Response, error) {
	u := "orders"

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	o := new(ListResponse[*Order])
	resp, err := s.client.Do(req, &o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, err
}

func (s *OrdersService) Show(id int) (*Order, *http.Response, error) {
	u := fmt.Sprintf("orders/%d", id)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	o := new(Order)
	resp, err := s.client.Do(req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, err
}

type OrderSearchOptions struct {
	ListOptions
	Query string `url:"query,omitempty"`
}

func (s *OrdersService) Search(searchOptions *OrderSearchOptions) (*ListResponse[*Order], *http.Response, error) {
	u := "orders/search"

	req, err := s.client.NewRequest(http.MethodGet, u, searchOptions)
	if err != nil {
		return nil, nil, err
	}

	o := new(ListResponse[*Order])
	resp, err := s.client.Do(req, &o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, err
}
