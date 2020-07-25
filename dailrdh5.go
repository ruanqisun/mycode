package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDH5IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH5I
}

type DAILRDH5ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH5O
}

type DAILRDH5RequestForm struct {
	Form []DAILRDH5IDataForm
}

type DAILRDH5ResponseForm struct {
	Form []DAILRDH5ODataForm
}

type DAILRDH5I struct {
	TaskNo                       string  `json:"TaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	AfBnkLnBizTypCd              string  `json:"AfBnkLnBizTypCd"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperTypCd                    string  `json:"OperTypCd"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	OperDt                       string  `json:"OperDt"`
	OperBefAfBnkLnJobStusCd      string  `json:"OperBefAfBnkLnJobStusCd"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	AfBnkLnJobDlwthResultCd      string  `json:"AfBnkLnJobDlwthResultCd"`
	OperResultComnt              string  `json:"OperResultComnt"`
	TxSn                         string  `json:"TxSn"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDH5O struct {
	Records                      []DAILRDH5ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDH5ORecords struct {
	TaskNo                       string  `json:"TaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	AfBnkLnBizTypCd              string  `json:"AfBnkLnBizTypCd"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperTypCd                    string  `json:"OperTypCd"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	OperDt                       string  `json:"OperDt"`
	OperBefAfBnkLnJobStusCd      string  `json:"OperBefAfBnkLnJobStusCd"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	AfBnkLnJobDlwthResultCd      string  `json:"AfBnkLnJobDlwthResultCd"`
	OperResultComnt              string  `json:"OperResultComnt"`
	TxSn                         string  `json:"TxSn"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDH5RequestForm) PackRequest(dailrdh5I DAILRDH5I) (responseBody []byte, err error) {

	requestForm := DAILRDH5RequestForm{
		Form: []DAILRDH5IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH5I",
				},

				FormData: dailrdh5I,
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
func (o *DAILRDH5RequestForm) UnPackRequest(request []byte) (DAILRDH5I, error) {
	dailrdh5I := DAILRDH5I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh5I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh5I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDH5ResponseForm) PackResponse(dailrdh5O DAILRDH5O) (responseBody []byte, err error) {
	responseForm := DAILRDH5ResponseForm{
		Form: []DAILRDH5ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH5O",
				},
				FormData: dailrdh5O,
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
func (o *DAILRDH5ResponseForm) UnPackResponse(request []byte) (DAILRDH5O, error) {

	dailrdh5O := DAILRDH5O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh5O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh5O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDH5I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
