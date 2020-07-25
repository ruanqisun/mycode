package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0015IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0015I
}

type IL8R0015ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0015O
}

type IL8R0015RequestForm struct {
	Form []IL8R0015IDataForm
}

type IL8R0015ResponseForm struct {
	Form []IL8R0015ODataForm
}

type IL8R0015I struct {
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	CurrAcctnDt                    string  `json:"CurrAcctnDt"`
	LstoneAcctnDt                  string  `json:"LstoneAcctnDt"`
	NxtoneAcctnDt                  string  `json:"NxtoneAcctnDt"`
	AcctnDayFlgCd                  string  `json:"AcctnDayFlgCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0015O struct {
	Records                      []IL8R0015ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0015ORecords struct {
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	CurrAcctnDt                    string  `json:"CurrAcctnDt"`
	LstoneAcctnDt                  string  `json:"LstoneAcctnDt"`
	NxtoneAcctnDt                  string  `json:"NxtoneAcctnDt"`
	AcctnDayFlgCd                  string  `json:"AcctnDayFlgCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0015RequestForm) PackRequest(il8r0015I IL8R0015I) (responseBody []byte, err error) {

	requestForm := IL8R0015RequestForm{
		Form: []IL8R0015IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0015I",
				},

				FormData: il8r0015I,
			},
		},
	}

	responseBody, err = json.Marshal(requestForm)
	if err != nil {
		return nil, errors.Wrap(err, 0, constant.REQPACKERR)
	}

	return responseBody, nil
}

// @Desc Parsing request message
func (o *IL8R0015RequestForm) UnPackRequest(request []byte) (IL8R0015I, error) {
	il8r0015I := IL8R0015I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0015I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0015I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0015ResponseForm) PackResponse(il8r0015O IL8R0015O) (responseBody []byte, err error) {
	responseForm := IL8R0015ResponseForm{
		Form: []IL8R0015ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0015O",
				},
				FormData: il8r0015O,
			},
		},
	}

	responseBody, err = json.Marshal(responseForm)
	if err != nil {
		return nil, errors.Wrap(err, 0, constant.RSPPACKERR)
	}

	return responseBody, nil
}

// @Desc Parsing response message
func (o *IL8R0015ResponseForm) UnPackResponse(request []byte) (IL8R0015O, error) {

	il8r0015O := IL8R0015O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0015O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0015O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0015I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
