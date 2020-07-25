package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRYW3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW3I
}

type DAILRYW3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRYW3O
}

type DAILRYW3RequestForm struct {
	Form []DAILRYW3IDataForm
}

type DAILRYW3ResponseForm struct {
	Form []DAILRYW3ODataForm
}

type DAILRYW3I struct {
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	ProdtProcesVersNo            int     `json:"ProdtProcesVersNo"`
	OrgNo                        string  `json:"OrgNo"`
	BizProcesNm                  string  `json:"BizProcesNm"`
	ProcesTypCd                  string  `json:"ProcesTypCd"`
	ProcesModelId                string  `json:"ProcesModelId"`
	ProcesModelVersNo            int     `json:"ProcesModelVersNo"`
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

type DAILRYW3O struct {
	Records                      []DAILRYW3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRYW3ORecords struct {
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	ProdtProcesVersNo            int     `json:"ProdtProcesVersNo"`
	OrgNo                        string  `json:"OrgNo"`
	BizProcesNm                  string  `json:"BizProcesNm"`
	ProcesTypCd                  string  `json:"ProcesTypCd"`
	ProcesModelId                string  `json:"ProcesModelId"`
	ProcesModelVersNo            int     `json:"ProcesModelVersNo"`
	CrtTelrNo                    string  `json:"CrtTelrNo"`
	CrtDt                        string  `json:"CrtDt"`
	CrtTm                        string  `json:"CrtTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRYW3RequestForm) PackRequest(dailryw3I DAILRYW3I) (responseBody []byte, err error) {

	requestForm := DAILRYW3RequestForm{
		Form: []DAILRYW3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW3I",
				},

				FormData: dailryw3I,
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
func (o *DAILRYW3RequestForm) UnPackRequest(request []byte) (DAILRYW3I, error) {
	dailryw3I := DAILRYW3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailryw3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRYW3ResponseForm) PackResponse(dailryw3O DAILRYW3O) (responseBody []byte, err error) {
	responseForm := DAILRYW3ResponseForm{
		Form: []DAILRYW3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRYW3O",
				},
				FormData: dailryw3O,
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
func (o *DAILRYW3ResponseForm) UnPackResponse(request []byte) (DAILRYW3O, error) {

	dailryw3O := DAILRYW3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailryw3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailryw3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRYW3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
