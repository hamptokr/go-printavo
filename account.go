package printavo

import (
	"net/http"
)

// AccountService handles communication with the account related methods of the
// Printavo API.
//
// Printavo API docs: https://printavo.docs.apiary.io/#reference/account
type AccountService struct {
	client *Client
}

type Account struct {
	Address1                     string                       `json:"address1"`
	Address2                     string                       `json:"address2"`
	AllowMobileAppAccess         bool                         `json:"allow_mobile_app_access?"`
	City                         string                       `json:"city"`
	CompanyName                  string                       `json:"company_name"`
	Country                      string                       `json:"country"`
	CurrentVisualID              int                          `json:"current_visual_id"`
	ID                           int                          `json:"id"`
	InvoiceTemplate              interface{}                  `json:"invoice_template"`
	InvoiceTemplateSubjectLine   interface{}                  `json:"invoice_template_subject_line"`
	InvoiceTitle                 string                       `json:"invoice_title"`
	MasterEmail                  string                       `json:"master_email"`
	PaymentProcessorPresent      bool                         `json:"payment_processor_present?"`
	PaymentRequestTemplate       interface{}                  `json:"payment_request_template"`
	Phone                        string                       `json:"phone"`
	Plan                         string                       `json:"plan"`
	PlanAllowsPricingCalculation bool                         `json:"plan_allows_pricing_calculation?"`
	PlanAllowsQrCoding           bool                         `json:"plan_allows_qr_coding?"`
	SalesTax                     float32                      `json:"sales_tax"`
	ShowInvoicePoNumber          bool                         `json:"show_invoice_po_number"`
	SocialFacebookLink           string                       `json:"social_facebook_link"`
	SocialTwitterLink            string                       `json:"social_twitter_link"`
	State                        string                       `json:"state"`
	TermsAndConditions           interface{}                  `json:"terms_and_conditions"`
	UseBroderGlobalProducts      bool                         `json:"use_broder_global_products"`
	UseSanmarGlobalProducts      bool                         `json:"use_sanmar_global_products"`
	UseTscApparelGlobalProducts  bool                         `json:"use_tsc_apparel_global_products"`
	Website                      string                       `json:"website"`
	Zip                          string                       `json:"zip"`
	PricingMatrices              []PricingMatrices            `json:"pricing_matrices"`
	PricingMatricesCells         []PricingMatricesCells       `json:"pricing_matrices_cells"`
	EmailTemplates               []EmailTemplates             `json:"email_templates"`
	LineItemSizeOptionAttributes LineItemSizeOptionAttributes `json:"line_item_size_option_attributes"`
}

type PricingMatrices struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	ProductPricingOption string `json:"product_pricing_option"`
}

type PricingMatricesCells struct {
	MatrixID       int    `json:"matrix_id"`
	ColHeaderID    int    `json:"col_header_id"`
	ColHeaderLabel string `json:"col_header_label"`
}

type EmailTemplates struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SubjectLine string `json:"subject_line"`
	Message     string `json:"message"`
}

type LineItemSizeOptionAttributes struct {
	StyleNumber bool   `json:"style_number"`
	Category    bool   `json:"category"`
	Color       bool   `json:"color"`
	SizeYxs     bool   `json:"size_yxs"`
	SizeYs      bool   `json:"size_ys"`
	SizeYm      bool   `json:"size_ym"`
	SizeYl      bool   `json:"size_yl"`
	SizeYxl     bool   `json:"size_yxl"`
	SizeXs      bool   `json:"size_xs"`
	SizeS       bool   `json:"size_s"`
	SizeM       bool   `json:"size_m"`
	SizeL       bool   `json:"size_l"`
	SizeXl      bool   `json:"size_xl"`
	Size2Xl     bool   `json:"size_2xl"`
	Size3Xl     bool   `json:"size_3xl"`
	Size4Xl     bool   `json:"size_4xl"`
	Size5Xl     bool   `json:"size_5xl"`
	Size6Xl     bool   `json:"size_6xl"`
	SizeOther   bool   `json:"size_other"`
	Size6M      bool   `json:"size_6m"`
	Size12M     bool   `json:"size_12m"`
	Size18M     bool   `json:"size_18m"`
	Size24M     bool   `json:"size_24m"`
	Size2T      bool   `json:"size_2t"`
	Size3T      bool   `json:"size_3t"`
	Size4T      bool   `json:"size_4t"`
	Size5T      bool   `json:"size_5t"`
	UpdatedAt   string `json:"updated_at"`
}

func (s *AccountService) Show() (*Account, *http.Response, error) {
	u := "accounts"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(Account)
	resp, err := s.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, err
}
