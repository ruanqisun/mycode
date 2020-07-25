package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRXT0IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT0I
}

type DAILRXT0ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT0O
}

type DAILRXT0RequestForm struct {
	Form []DAILRXT0IDataForm
}

type DAILRXT0ResponseForm struct {
	Form []DAILRXT0ODataForm
}

type DAILRXT0I struct {
	SysPmn                       string  `json:"SysPmn"`
	OrgNo                        string  `json:"OrgNo"`
	SysPmt                       string  `json:"SysPmt"`
	SysPml                       int     `json:"SysPml"`
	SysPmd                       string  `json:"SysPmd"`
	SysPmv                       string  `json:"SysPmv"`
	SysCd6                       string  `json:"SysCd6"`
	SysRds                       string  `json:"SysRds"`
	SysSmk                       string  `json:"SysSmk"`
	SysSmk1                      string  `json:"SysSmk1"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRXT0O struct {
	Records                      []DAILRXT0ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRXT0ORecords struct {
	SysPmn                       string  `json:"SysPmn"`
	OrgNo                        string  `json:"OrgNo"`
	SysPmt                       string  `json:"SysPmt"`
	SysPml                       int     `json:"SysPml"`
	SysPmd                       string  `json:"SysPmd"`
	SysPmv                       string  `json:"SysPmv"`
	SysCd6                       string  `json:"SysCd6"`
	SysRds                       string  `json:"SysRds"`
	SysSmk                       string  `json:"SysSmk"`
	SysSmk1                      string  `json:"SysSmk1"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRXT0RequestForm) PackRequest(dailrxt0I DAILRXT0I) (responseBody []byte, err error) {

	requestForm := DAILRXT0RequestForm{
		Form: []DAILRXT0IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT0I",
				},

				FormData: dailrxt0I,
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
func (o *DAILRXT0RequestForm) UnPackRequest(request []byte) (DAILRXT0I, error) {
	dailrxt0I := DAILRXT0I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt0I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt0I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRXT0ResponseForm) PackResponse(dailrxt0O DAILRXT0O) (responseBody []byte, err error) {
	responseForm := DAILRXT0ResponseForm{
		Form: []DAILRXT0ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT0O",
				},
				FormData: dailrxt0O,
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
func (o *DAILRXT0ResponseForm) UnPackResponse(request []byte) (DAILRXT0O, error) {

	dailrxt0O := DAILRXT0O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt0O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt0O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRXT0I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
