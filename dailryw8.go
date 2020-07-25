package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW8IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW8I
}

type DAILRYW8ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW8O
}

type DAILRYW8RequestForm struct {
	Form []DAILRYW8IDataForm
}

type DAILRYW8ResponseForm struct {
	Form []DAILRYW8ODataForm
}

type DAILRYW8I struct {
	ComslNo                       string  `json:"ComslNo"`
	DgtCdtlId                     string  `json:"DgtCdtlId"`
	OrgNo                         string  `json:"OrgNo"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	ComslNm                       string  `json:"ComslNm"`
	ComslPictInfo                 string  `json:"ComslPictInfo"`
	CrtTelrNo                     string  `json:"CrtTelrNo"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	CrtDt                         string  `json:"CrtDt"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
	PageNo                        int     `json:"PageNo"`
	PageRecCount                  int     `json:"PageRecCount"`
}

type DAILRYW8O struct {
	Records                      []DAILRYW8ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW8ORecords struct {
	ComslNo                       string  `json:"ComslNo"`
	DgtCdtlId                     string  `json:"DgtCdtlId"`
	OrgNo                         string  `json:"OrgNo"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	ComslNm                       string  `json:"ComslNm"`
	ComslPictInfo                 string  `json:"ComslPictInfo"`
	CrtTelrNo                     string  `json:"CrtTelrNo"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	CrtDt                         string  `json:"CrtDt"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRYW8RequestForm) PackRequest(dailryw8I DAILRYW8I) (responseBody []byte, err error) {

	requestForm := DAILRYW8RequestForm{
		Form: []DAILRYW8IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW8I",
				},

				FormData: dailryw8I,
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
func (o *DAILRYW8RequestForm) UnPackRequest(request []byte) (DAILRYW8I, error) {
	dailryw8I := DAILRYW8I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw8I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw8I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW8ResponseForm) PackResponse(dailryw8O DAILRYW8O) (responseBody []byte, err error) {
	responseForm := DAILRYW8ResponseForm{
		Form: []DAILRYW8ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW8O",
				},
				FormData: dailryw8O,
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
func (o *DAILRYW8ResponseForm) UnPackResponse(request []byte) (DAILRYW8O, error) {

	dailryw8O := DAILRYW8O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw8O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw8O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW8I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
