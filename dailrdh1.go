package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDH1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH1I
}

type DAILRDH1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH1O
}

type DAILRDH1RequestForm struct {
	Form []DAILRDH1IDataForm
}

type DAILRDH1ResponseForm struct {
	Form []DAILRDH1ODataForm
}

type DAILRDH1I struct {
	CollTaskNo                   string  `json:"CollTaskNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	OrgNo                        string  `json:"OrgNo"`
	OrgNm                        string  `json:"OrgNm"`
	CrtDt                        string  `json:"CrtDt"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	AlrdyExecFlg                 string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg              string  `json:"AtExecWindowFlg"`
	ExecTms                      int     `json:"ExecTms"`
	ExecRsn                      string  `json:"ExecRsn"`
	ExecConcusComnt              string  `json:"ExecConcusComnt"`
	CollSponsorEmpnbr            string  `json:"CollSponsorEmpnbr"`
	CollSponsorOrgNo             string  `json:"CollSponsorOrgNo"`
	CollCorgEmpnbr               string  `json:"CollCorgEmpnbr"`
	CollCorgOrgNo                string  `json:"CollCorgOrgNo"`
	CtrtNo                       string  `json:"CtrtNo"`
	ReacTotlTms                  int     `json:"ReacTotlTms"`
	DlwthDt                      string  `json:"DlwthDt"`
	SponsorEmplyTaskNo           string  `json:"SponsorEmplyTaskNo"`
	CorgEmplyTaskNo              string  `json:"CorgEmplyTaskNo"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	SponsorEmplyName             string  `json:"SponsorEmplyName"`
	CorgEmplyName                string  `json:"CorgEmplyName"`
	SponsorOrgNm                 string  `json:"SponsorOrgNm"`
	CorgOrgNm                    string  `json:"CorgOrgNm"`
	ReacCustTms                  int     `json:"ReacCustTms"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDH1O struct {
	Records                      []DAILRDH1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDH1ORecords struct {
	CollTaskNo                   string  `json:"CollTaskNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	OrgNo                        string  `json:"OrgNo"`
	OrgNm                        string  `json:"OrgNm"`
	CrtDt                        string  `json:"CrtDt"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	AlrdyExecFlg                 string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg              string  `json:"AtExecWindowFlg"`
	ExecTms                      int     `json:"ExecTms"`
	ExecRsn                      string  `json:"ExecRsn"`
	ExecConcusComnt              string  `json:"ExecConcusComnt"`
	CollSponsorEmpnbr            string  `json:"CollSponsorEmpnbr"`
	CollSponsorOrgNo             string  `json:"CollSponsorOrgNo"`
	CollCorgEmpnbr               string  `json:"CollCorgEmpnbr"`
	CollCorgOrgNo                string  `json:"CollCorgOrgNo"`
	CtrtNo                       string  `json:"CtrtNo"`
	ReacTotlTms                  int     `json:"ReacTotlTms"`
	DlwthDt                      string  `json:"DlwthDt"`
	SponsorEmplyTaskNo           string  `json:"SponsorEmplyTaskNo"`
	CorgEmplyTaskNo              string  `json:"CorgEmplyTaskNo"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	SponsorEmplyName             string  `json:"SponsorEmplyName"`
	CorgEmplyName                string  `json:"CorgEmplyName"`
	SponsorOrgNm                 string  `json:"SponsorOrgNm"`
	CorgOrgNm                    string  `json:"CorgOrgNm"`
	ReacCustTms                  int     `json:"ReacCustTms"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDH1RequestForm) PackRequest(dailrdh1I DAILRDH1I) (responseBody []byte, err error) {

	requestForm := DAILRDH1RequestForm{
		Form: []DAILRDH1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH1I",
				},

				FormData: dailrdh1I,
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
func (o *DAILRDH1RequestForm) UnPackRequest(request []byte) (DAILRDH1I, error) {
	dailrdh1I := DAILRDH1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDH1ResponseForm) PackResponse(dailrdh1O DAILRDH1O) (responseBody []byte, err error) {
	responseForm := DAILRDH1ResponseForm{
		Form: []DAILRDH1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH1O",
				},
				FormData: dailrdh1O,
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
func (o *DAILRDH1ResponseForm) UnPackResponse(request []byte) (DAILRDH1O, error) {

	dailrdh1O := DAILRDH1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDH1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
