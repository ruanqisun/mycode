package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW4IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW4I
}

type DAILRYW4ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW4O
}

type DAILRYW4RequestForm struct {
	Form []DAILRYW4IDataForm
}

type DAILRYW4ResponseForm struct {
	Form []DAILRYW4ODataForm
}

type DAILRYW4I struct {
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

type DAILRYW4O struct {
	Records                      []DAILRYW4ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW4ORecords struct {
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
func (o *DAILRYW4RequestForm) PackRequest(dailryw4I DAILRYW4I) (responseBody []byte, err error) {

	requestForm := DAILRYW4RequestForm{
		Form: []DAILRYW4IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW4I",
				},

				FormData: dailryw4I,
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
func (o *DAILRYW4RequestForm) UnPackRequest(request []byte) (DAILRYW4I, error) {
	dailryw4I := DAILRYW4I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw4I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw4I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW4ResponseForm) PackResponse(dailryw4O DAILRYW4O) (responseBody []byte, err error) {
	responseForm := DAILRYW4ResponseForm{
		Form: []DAILRYW4ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW4O",
				},
				FormData: dailryw4O,
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
func (o *DAILRYW4ResponseForm) UnPackResponse(request []byte) (DAILRYW4O, error) {

	dailryw4O := DAILRYW4O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw4O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw4O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW4I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
