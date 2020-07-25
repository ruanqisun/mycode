package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRXT4IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT4I
}

type DAILRXT4ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT4O
}

type DAILRXT4RequestForm struct {
	Form []DAILRXT4IDataForm
}

type DAILRXT4ResponseForm struct {
	Form []DAILRXT4ODataForm
}

type DAILRXT4I struct {
	OptnNm                       string  `json:"OptnNm"`
	AvalFlg                      string  `json:"AvalFlg"`
	PareClsId                    string  `json:"PareClsId"`
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

type DAILRXT4O struct {
	Records                      []DAILRXT4ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRXT4ORecords struct {
	OptnNm                       string  `json:"OptnNm"`
	AvalFlg                      string  `json:"AvalFlg"`
	PareClsId                    string  `json:"PareClsId"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTm                        string  `json:"CrtTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRXT4RequestForm) PackRequest(dailrxt4I DAILRXT4I) (responseBody []byte, err error) {

	requestForm := DAILRXT4RequestForm{
		Form: []DAILRXT4IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT4I",
				},

				FormData: dailrxt4I,
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
func (o *DAILRXT4RequestForm) UnPackRequest(request []byte) (DAILRXT4I, error) {
	dailrxt4I := DAILRXT4I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt4I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt4I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRXT4ResponseForm) PackResponse(dailrxt4O DAILRXT4O) (responseBody []byte, err error) {
	responseForm := DAILRXT4ResponseForm{
		Form: []DAILRXT4ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT4O",
				},
				FormData: dailrxt4O,
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
func (o *DAILRXT4ResponseForm) UnPackResponse(request []byte) (DAILRXT4O, error) {

	dailrxt4O := DAILRXT4O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt4O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt4O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRXT4I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
