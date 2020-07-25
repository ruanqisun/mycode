package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDH2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH2I
}

type DAILRDH2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH2O
}

type DAILRDH2RequestForm struct {
	Form []DAILRDH2IDataForm
}

type DAILRDH2ResponseForm struct {
	Form []DAILRDH2ODataForm
}

type DAILRDH2I struct {
	WarnTaskNo                   string  `json:"WarnTaskNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	OrgNo                        string  `json:"OrgNo"`
	OrgNm                        string  `json:"OrgNm"`
	CrtDt                        string  `json:"CrtDt"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	WarnSponsorEmpnbr            string  `json:"WarnSponsorEmpnbr"`
	WarnSponsorOrgNo             string  `json:"WarnSponsorOrgNo"`
	WarnCorgEmpnbr               string  `json:"WarnCorgEmpnbr"`
	WarnCorgOrgNo                string  `json:"WarnCorgOrgNo"`
	AlrdyExecFlg                 string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg              string  `json:"AtExecWindowFlg"`
	ExecTms                      int     `json:"ExecTms"`
	ExecConcusCd                 string  `json:"ExecConcusCd"`
	ExecRsn                      string  `json:"ExecRsn"`
	CtrtNo                       string  `json:"CtrtNo"`
	SponsorEmplyName             string  `json:"SponsorEmplyName"`
	CorgEmplyName                string  `json:"CorgEmplyName"`
	SponsorOrgNm                 string  `json:"SponsorOrgNm"`
	CorgOrgNm                    string  `json:"CorgOrgNm"`
	SponsorEmplyTaskNo           string  `json:"SponsorEmplyTaskNo"`
	CorgEmplyTaskNo              string  `json:"CorgEmplyTaskNo"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	WarnSignalTypCd              string  `json:"WarnSignalTypCd"`
	AdjQtaDlwthResultCd          string  `json:"AdjQtaDlwthResultCd"`
	DlwthDt                      string  `json:"DlwthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDH2O struct {
	Records                      []DAILRDH2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDH2ORecords struct {
	WarnTaskNo                   string  `json:"WarnTaskNo"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	OrgNo                        string  `json:"OrgNo"`
	OrgNm                        string  `json:"OrgNm"`
	CrtDt                        string  `json:"CrtDt"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	WarnSponsorEmpnbr            string  `json:"WarnSponsorEmpnbr"`
	WarnSponsorOrgNo             string  `json:"WarnSponsorOrgNo"`
	WarnCorgEmpnbr               string  `json:"WarnCorgEmpnbr"`
	WarnCorgOrgNo                string  `json:"WarnCorgOrgNo"`
	AlrdyExecFlg                 string  `json:"AlrdyExecFlg"`
	AtExecWindowFlg              string  `json:"AtExecWindowFlg"`
	ExecTms                      int     `json:"ExecTms"`
	ExecConcusCd                 string  `json:"ExecConcusCd"`
	ExecRsn                      string  `json:"ExecRsn"`
	CtrtNo                       string  `json:"CtrtNo"`
	SponsorEmplyName             string  `json:"SponsorEmplyName"`
	CorgEmplyName                string  `json:"CorgEmplyName"`
	SponsorOrgNm                 string  `json:"SponsorOrgNm"`
	CorgOrgNm                    string  `json:"CorgOrgNm"`
	SponsorEmplyTaskNo           string  `json:"SponsorEmplyTaskNo"`
	CorgEmplyTaskNo              string  `json:"CorgEmplyTaskNo"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	WarnSignalTypCd              string  `json:"WarnSignalTypCd"`
	AdjQtaDlwthResultCd          string  `json:"AdjQtaDlwthResultCd"`
	DlwthDt                      string  `json:"DlwthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDH2RequestForm) PackRequest(dailrdh2I DAILRDH2I) (responseBody []byte, err error) {

	requestForm := DAILRDH2RequestForm{
		Form: []DAILRDH2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH2I",
				},

				FormData: dailrdh2I,
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
func (o *DAILRDH2RequestForm) UnPackRequest(request []byte) (DAILRDH2I, error) {
	dailrdh2I := DAILRDH2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDH2ResponseForm) PackResponse(dailrdh2O DAILRDH2O) (responseBody []byte, err error) {
	responseForm := DAILRDH2ResponseForm{
		Form: []DAILRDH2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH2O",
				},
				FormData: dailrdh2O,
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
func (o *DAILRDH2ResponseForm) UnPackResponse(request []byte) (DAILRDH2O, error) {

	dailrdh2O := DAILRDH2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDH2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
