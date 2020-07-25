package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRXT5IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT5I
}

type DAILRXT5ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT5O
}

type DAILRXT5RequestForm struct {
	Form []DAILRXT5IDataForm
}

type DAILRXT5ResponseForm struct {
	Form []DAILRXT5ODataForm
}

type DAILRXT5I struct {
	KeprcdNo                     string  `json:"KeprcdNo"`
	DocTmplTypCd                 string  `json:"DocTmplTypCd"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	TmplSubTypCd                 string  `json:"TmplSubTypCd"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	OrgNo                        string  `json:"OrgNo"`
	EfftDt                       string  `json:"EfftDt"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtTm                        string  `json:"CrtTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRXT5O struct {
	Records                      []DAILRXT5ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRXT5ORecords struct {
	KeprcdNo                     string  `json:"KeprcdNo"`
	DocTmplTypCd                 string  `json:"DocTmplTypCd"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	TmplSubTypCd                 string  `json:"TmplSubTypCd"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	OrgNo                        string  `json:"OrgNo"`
	EfftDt                       string  `json:"EfftDt"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtTm                        string  `json:"CrtTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRXT5RequestForm) PackRequest(dailrxt5I DAILRXT5I) (responseBody []byte, err error) {

	requestForm := DAILRXT5RequestForm{
		Form: []DAILRXT5IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT5I",
				},

				FormData: dailrxt5I,
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
func (o *DAILRXT5RequestForm) UnPackRequest(request []byte) (DAILRXT5I, error) {
	dailrxt5I := DAILRXT5I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt5I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt5I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRXT5ResponseForm) PackResponse(dailrxt5O DAILRXT5O) (responseBody []byte, err error) {
	responseForm := DAILRXT5ResponseForm{
		Form: []DAILRXT5ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT5O",
				},
				FormData: dailrxt5O,
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
func (o *DAILRXT5ResponseForm) UnPackResponse(request []byte) (DAILRXT5O, error) {

	dailrxt5O := DAILRXT5O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt5O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt5O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRXT5I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
