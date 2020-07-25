package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRXT2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT2I
}

type DAILRXT2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT2O
}

type DAILRXT2RequestForm struct {
	Form []DAILRXT2IDataForm
}

type DAILRXT2ResponseForm struct {
	Form []DAILRXT2ODataForm
}

type DAILRXT2I struct {
	RoleNo                       string  `json:"RoleNo"`
	FuncNo                       string  `json:"FuncNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRXT2O struct {
	Records                      []DAILRXT2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRXT2ORecords struct {
	RoleNo                       string  `json:"RoleNo"`
	FuncNo                       string  `json:"FuncNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRXT2RequestForm) PackRequest(dailrxt2I DAILRXT2I) (responseBody []byte, err error) {

	requestForm := DAILRXT2RequestForm{
		Form: []DAILRXT2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT2I",
				},

				FormData: dailrxt2I,
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
func (o *DAILRXT2RequestForm) UnPackRequest(request []byte) (DAILRXT2I, error) {
	dailrxt2I := DAILRXT2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRXT2ResponseForm) PackResponse(dailrxt2O DAILRXT2O) (responseBody []byte, err error) {
	responseForm := DAILRXT2ResponseForm{
		Form: []DAILRXT2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT2O",
				},
				FormData: dailrxt2O,
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
func (o *DAILRXT2ResponseForm) UnPackResponse(request []byte) (DAILRXT2O, error) {

	dailrxt2O := DAILRXT2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRXT2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
