package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX9IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX9I
}

type DAILRSX9ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX9O
}

type DAILRSX9RequestForm struct {
	Form []DAILRSX9IDataForm
}

type DAILRSX9ResponseForm struct {
	Form []DAILRSX9ODataForm
}

type DAILRSX9I struct {
	CrdQurySn                    string  `json:"CrdQurySn"`
	CustNo                       string  `json:"CustNo"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	CrdQuryDt                    string  `json:"CrdQuryDt"`
	CrdQurySuccFlg               string  `json:"CrdQurySuccFlg"`
	CrdQuryResultId              string  `json:"CrdQuryResultId"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRSX9O struct {
	Records                      []DAILRSX9ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX9ORecords struct {
	CrdQurySn                    string  `json:"CrdQurySn"`
	CustNo                       string  `json:"CustNo"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	CrdQuryDt                    string  `json:"CrdQuryDt"`
	CrdQurySuccFlg               string  `json:"CrdQurySuccFlg"`
	CrdQuryResultId              string  `json:"CrdQuryResultId"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX9RequestForm) PackRequest(dailrsx9I DAILRSX9I) (responseBody []byte, err error) {

	requestForm := DAILRSX9RequestForm{
		Form: []DAILRSX9IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX9I",
				},

				FormData: dailrsx9I,
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
func (o *DAILRSX9RequestForm) UnPackRequest(request []byte) (DAILRSX9I, error) {
	dailrsx9I := DAILRSX9I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx9I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx9I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX9ResponseForm) PackResponse(dailrsx9O DAILRSX9O) (responseBody []byte, err error) {
	responseForm := DAILRSX9ResponseForm{
		Form: []DAILRSX9ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX9O",
				},
				FormData: dailrsx9O,
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
func (o *DAILRSX9ResponseForm) UnPackResponse(request []byte) (DAILRSX9O, error) {

	dailrsx9O := DAILRSX9O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx9O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx9O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX9I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
