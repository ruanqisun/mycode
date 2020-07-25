package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRXT3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT3I
}

type DAILRXT3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT3O
}

type DAILRXT3RequestForm struct {
	Form []DAILRXT3IDataForm
}

type DAILRXT3ResponseForm struct {
	Form []DAILRXT3ODataForm
}

type DAILRXT3I struct {
	RoleNo                       string  `json:"RoleNo"`
	SensTypCd                    string  `json:"SensTypCd"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTm                        string  `json:"CrtTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRXT3O struct {
	Records                      []DAILRXT3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRXT3ORecords struct {
	RoleNo                       string  `json:"RoleNo"`
	SensTypCd                    string  `json:"SensTypCd"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTm                        string  `json:"CrtTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRXT3RequestForm) PackRequest(dailrxt3I DAILRXT3I) (responseBody []byte, err error) {

	requestForm := DAILRXT3RequestForm{
		Form: []DAILRXT3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT3I",
				},

				FormData: dailrxt3I,
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
func (o *DAILRXT3RequestForm) UnPackRequest(request []byte) (DAILRXT3I, error) {
	dailrxt3I := DAILRXT3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRXT3ResponseForm) PackResponse(dailrxt3O DAILRXT3O) (responseBody []byte, err error) {
	responseForm := DAILRXT3ResponseForm{
		Form: []DAILRXT3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT3O",
				},
				FormData: dailrxt3O,
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
func (o *DAILRXT3ResponseForm) UnPackResponse(request []byte) (DAILRXT3O, error) {

	dailrxt3O := DAILRXT3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRXT3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
