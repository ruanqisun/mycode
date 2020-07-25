package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRED3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRED3I
}

type DAILRED3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRED3O
}

type DAILRED3RequestForm struct {
	Form []DAILRED3IDataForm
}

type DAILRED3ResponseForm struct {
	Form []DAILRED3ODataForm
}

type DAILRED3I struct {
	ChgAplyNo                     string  `json:"ChgAplyNo"`
	CustNo                        string  `json:"CustNo"`
	CrdtAplySn                    string  `json:"CrdtAplySn"`
	AplyAdjRsnCd                  string  `json:"AplyAdjRsnCd"`
	AplyAdjTypCd                  string  `json:"AplyAdjTypCd"`
	AdjDrctCd                     string  `json:"AdjDrctCd"`
	CurCd                         string  `json:"CurCd"`
	ChgBefAvalQta                 float64  `json:"ChgBefAvalQta"`
	ChgAftQta                     float64  `json:"ChgAftQta"`
	ChgAftAvalQta                 float64  `json:"ChgAftAvalQta"`
	ChgRsn                        string  `json:"ChgRsn"`
	ExmnvrfyStusCd                string  `json:"ExmnvrfyStusCd"`
	ExmnvrfySugstnComnt           string  `json:"ExmnvrfySugstnComnt"`
	TxSn                          string  `json:"TxSn"`
	CrtDt                         string  `json:"CrtDt"`
	CrtEmpnbr                     string  `json:"CrtEmpnbr"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRED3O struct {
	Records                      []DAILRED3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRED3ORecords struct {
	ChgAplyNo                     string  `json:"ChgAplyNo"`
	CustNo                        string  `json:"CustNo"`
	CrdtAplySn                    string  `json:"CrdtAplySn"`
	AplyAdjRsnCd                  string  `json:"AplyAdjRsnCd"`
	AplyAdjTypCd                  string  `json:"AplyAdjTypCd"`
	AdjDrctCd                     string  `json:"AdjDrctCd"`
	CurCd                         string  `json:"CurCd"`
	ChgBefAvalQta                 float64  `json:"ChgBefAvalQta"`
	ChgAftQta                     float64  `json:"ChgAftQta"`
	ChgAftAvalQta                 float64  `json:"ChgAftAvalQta"`
	ChgRsn                        string  `json:"ChgRsn"`
	ExmnvrfyStusCd                string  `json:"ExmnvrfyStusCd"`
	ExmnvrfySugstnComnt           string  `json:"ExmnvrfySugstnComnt"`
	TxSn                          string  `json:"TxSn"`
	CrtDt                         string  `json:"CrtDt"`
	CrtEmpnbr                     string  `json:"CrtEmpnbr"`
	CrtOrgNo                      string  `json:"CrtOrgNo"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRED3RequestForm) PackRequest(dailred3I DAILRED3I) (responseBody []byte, err error) {

	requestForm := DAILRED3RequestForm{
		Form: []DAILRED3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRED3I",
				},

				FormData: dailred3I,
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
func (o *DAILRED3RequestForm) UnPackRequest(request []byte) (DAILRED3I, error) {
	dailred3I := DAILRED3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailred3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailred3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRED3ResponseForm) PackResponse(dailred3O DAILRED3O) (responseBody []byte, err error) {
	responseForm := DAILRED3ResponseForm{
		Form: []DAILRED3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRED3O",
				},
				FormData: dailred3O,
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
func (o *DAILRED3ResponseForm) UnPackResponse(request []byte) (DAILRED3O, error) {

	dailred3O := DAILRED3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailred3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailred3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRED3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
