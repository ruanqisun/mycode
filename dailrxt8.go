package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRXT8IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT8I
}

type DAILRXT8ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT8O
}

type DAILRXT8RequestForm struct {
	Form []DAILRXT8IDataForm
}

type DAILRXT8ResponseForm struct {
	Form []DAILRXT8ODataForm
}

type DAILRXT8I struct {
	ProcesModelId                string  `json:"ProcesModelId"`
	ProcesModelVersNo            int     `json:"ProcesModelVersNo"`
	ProcesModelNm                string  `json:"ProcesModelNm"`
	ProcesTmplId                 string  `json:"ProcesTmplId"`
	ProcesGraphBase64Cd          string  `json:"ProcesGraphBase64Cd"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTm                        string  `json:"CrtTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRXT8O struct {
	Records                      []DAILRXT8ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRXT8ORecords struct {
	ProcesModelId                string  `json:"ProcesModelId"`
	ProcesModelVersNo            int     `json:"ProcesModelVersNo"`
	ProcesModelNm                string  `json:"ProcesModelNm"`
	ProcesTmplId                 string  `json:"ProcesTmplId"`
	ProcesGraphBase64Cd          string  `json:"ProcesGraphBase64Cd"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTm                        string  `json:"CrtTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRXT8RequestForm) PackRequest(dailrxt8I DAILRXT8I) (responseBody []byte, err error) {

	requestForm := DAILRXT8RequestForm{
		Form: []DAILRXT8IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT8I",
				},

				FormData: dailrxt8I,
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
func (o *DAILRXT8RequestForm) UnPackRequest(request []byte) (DAILRXT8I, error) {
	dailrxt8I := DAILRXT8I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt8I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt8I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRXT8ResponseForm) PackResponse(dailrxt8O DAILRXT8O) (responseBody []byte, err error) {
	responseForm := DAILRXT8ResponseForm{
		Form: []DAILRXT8ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT8O",
				},
				FormData: dailrxt8O,
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
func (o *DAILRXT8ResponseForm) UnPackResponse(request []byte) (DAILRXT8O, error) {

	dailrxt8O := DAILRXT8O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt8O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt8O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRXT8I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
