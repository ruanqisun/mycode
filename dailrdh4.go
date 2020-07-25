package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDH4IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH4I
}

type DAILRDH4ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH4O
}

type DAILRDH4RequestForm struct {
	Form []DAILRDH4IDataForm
}

type DAILRDH4ResponseForm struct {
	Form []DAILRDH4ODataForm
}

type DAILRDH4I struct {
	ChkTaskNo                    string  `json:"ChkTaskNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	CtrtNo                       string  `json:"CtrtNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	AfBnkLnChkTaskTypCd          string  `json:"AfBnkLnChkTaskTypCd"`
	AfBnkLnChkManrCd             string  `json:"AfBnkLnChkManrCd"`
	TaskGnrtDt                   string  `json:"TaskGnrtDt"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	AprsEmplyName                string  `json:"AprsEmplyName"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	AprsOrgNm                    string  `json:"AprsOrgNm"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	AlrdyExecFlg                 string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg              string  `json:"AtExecWindowFlg"`
	ExecTms                      int     `json:"ExecTms"`
	ExecRsn                      string  `json:"ExecRsn"`
	ExecConcusComnt              string  `json:"ExecConcusComnt"`
	ChkSponsorEmpnbr             string  `json:"ChkSponsorEmpnbr"`
	ChkSponsorOrgNo              string  `json:"ChkSponsorOrgNo"`
	ChkCorgEmpnbr                string  `json:"ChkCorgEmpnbr"`
	ChkCorgOrgNo                 string  `json:"ChkCorgOrgNo"`
	AprsSugstnCd                 string  `json:"AprsSugstnCd"`
	RctfctnSugstnContent         string  `json:"RctfctnSugstnContent"`
	AprsEmpnbr                   string  `json:"AprsEmpnbr"`
	AprsOrgNo                    string  `json:"AprsOrgNo"`
	AprsDt                       string  `json:"AprsDt"`
	SponsorEmplyName             string  `json:"SponsorEmplyName"`
	CorgEmplyName                string  `json:"CorgEmplyName"`
	SponsorOrgNm                 string  `json:"SponsorOrgNm"`
	CorgOrgNm                    string  `json:"CorgOrgNm"`
	SponsorEmplyTaskNo           string  `json:"SponsorEmplyTaskNo"`
	CorgEmplyTaskNo              string  `json:"CorgEmplyTaskNo"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	AprsEmplyTaskNo              string  `json:"AprsEmplyTaskNo"`
	ReacTotlTms                  int     `json:"ReacTotlTms"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDH4O struct {
	Records                      []DAILRDH4ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDH4ORecords struct {
	ChkTaskNo                    string  `json:"ChkTaskNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	CtrtNo                       string  `json:"CtrtNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	AfBnkLnChkTaskTypCd          string  `json:"AfBnkLnChkTaskTypCd"`
	AfBnkLnChkManrCd             string  `json:"AfBnkLnChkManrCd"`
	TaskGnrtDt                   string  `json:"TaskGnrtDt"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	AprsEmplyName                string  `json:"AprsEmplyName"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	AprsOrgNm                    string  `json:"AprsOrgNm"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	AlrdyExecFlg                 string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg              string  `json:"AtExecWindowFlg"`
	ExecTms                      int     `json:"ExecTms"`
	ExecRsn                      string  `json:"ExecRsn"`
	ExecConcusComnt              string  `json:"ExecConcusComnt"`
	ChkSponsorEmpnbr             string  `json:"ChkSponsorEmpnbr"`
	ChkSponsorOrgNo              string  `json:"ChkSponsorOrgNo"`
	ChkCorgEmpnbr                string  `json:"ChkCorgEmpnbr"`
	ChkCorgOrgNo                 string  `json:"ChkCorgOrgNo"`
	AprsSugstnCd                 string  `json:"AprsSugstnCd"`
	RctfctnSugstnContent         string  `json:"RctfctnSugstnContent"`
	AprsEmpnbr                   string  `json:"AprsEmpnbr"`
	AprsOrgNo                    string  `json:"AprsOrgNo"`
	AprsDt                       string  `json:"AprsDt"`
	SponsorEmplyName             string  `json:"SponsorEmplyName"`
	CorgEmplyName                string  `json:"CorgEmplyName"`
	SponsorOrgNm                 string  `json:"SponsorOrgNm"`
	CorgOrgNm                    string  `json:"CorgOrgNm"`
	SponsorEmplyTaskNo           string  `json:"SponsorEmplyTaskNo"`
	CorgEmplyTaskNo              string  `json:"CorgEmplyTaskNo"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	AprsEmplyTaskNo              string  `json:"AprsEmplyTaskNo"`
	ReacTotlTms                  int     `json:"ReacTotlTms"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDH4RequestForm) PackRequest(dailrdh4I DAILRDH4I) (responseBody []byte, err error) {

	requestForm := DAILRDH4RequestForm{
		Form: []DAILRDH4IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH4I",
				},

				FormData: dailrdh4I,
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
func (o *DAILRDH4RequestForm) UnPackRequest(request []byte) (DAILRDH4I, error) {
	dailrdh4I := DAILRDH4I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh4I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh4I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDH4ResponseForm) PackResponse(dailrdh4O DAILRDH4O) (responseBody []byte, err error) {
	responseForm := DAILRDH4ResponseForm{
		Form: []DAILRDH4ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH4O",
				},
				FormData: dailrdh4O,
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
func (o *DAILRDH4ResponseForm) UnPackResponse(request []byte) (DAILRDH4O, error) {

	dailrdh4O := DAILRDH4O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh4O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh4O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDH4I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
