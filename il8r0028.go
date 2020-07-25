package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0028IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0028I
}

type IL8R0028ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0028O
}

type IL8R0028RequestForm struct {
	Form []IL8R0028IDataForm
}

type IL8R0028ResponseForm struct {
	Form []IL8R0028ODataForm
}

type IL8R0028I struct {
	KeprcdNo                       string  `json:"KeprcdNo"`
	CustNo                         string  `json:"CustNo"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	CtrtNo                         string  `json:"CtrtNo"`
	ForsprtDt                      string  `json:"ForsprtDt"`
	ForsprtEquipNo                 string  `json:"ForsprtEquipNo"`
	KeprcdStusCd                   string  `json:"KeprcdStusCd"`
	CrtDt                          string  `json:"CrtDt"`
	ForsprtAmt                     float64  `json:"ForsprtAmt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0028O struct {
	Records                      []IL8R0028ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0028ORecords struct {
	KeprcdNo                       string  `json:"KeprcdNo"`
	CustNo                         string  `json:"CustNo"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	CtrtNo                         string  `json:"CtrtNo"`
	ForsprtDt                      string  `json:"ForsprtDt"`
	ForsprtEquipNo                 string  `json:"ForsprtEquipNo"`
	KeprcdStusCd                   string  `json:"KeprcdStusCd"`
	CrtDt                          string  `json:"CrtDt"`
	ForsprtAmt                     float64  `json:"ForsprtAmt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0028RequestForm) PackRequest(il8r0028I IL8R0028I) (responseBody []byte, err error) {

	requestForm := IL8R0028RequestForm{
		Form: []IL8R0028IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0028I",
				},

				FormData: il8r0028I,
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
func (o *IL8R0028RequestForm) UnPackRequest(request []byte) (IL8R0028I, error) {
	il8r0028I := IL8R0028I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0028I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0028I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0028ResponseForm) PackResponse(il8r0028O IL8R0028O) (responseBody []byte, err error) {
	responseForm := IL8R0028ResponseForm{
		Form: []IL8R0028ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0028O",
				},
				FormData: il8r0028O,
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
func (o *IL8R0028ResponseForm) UnPackResponse(request []byte) (IL8R0028O, error) {

	il8r0028O := IL8R0028O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0028O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0028O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0028I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
