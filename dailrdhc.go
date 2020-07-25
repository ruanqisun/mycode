package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDHCIDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDHCI
}

type DAILRDHCODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDHCO
}

type DAILRDHCRequestForm struct {
	Form []DAILRDHCIDataForm
}

type DAILRDHCResponseForm struct {
	Form []DAILRDHCODataForm
}

type DAILRDHCI struct {
	ChkTaskNo                    string  `json:"ChkTaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	HndlrEmpnbr                  string  `json:"HndlrEmpnbr"`
	ResultCnfrmRemrkInfo         string  `json:"ResultCnfrmRemrkInfo"`
	ReacCustManrCd               string  `json:"ReacCustManrCd"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CorgDlwthResultCd            string  `json:"CorgDlwthResultCd"`
	CollTms                      int     `json:"CollTms"`
	EmbFlgCd                     string  `json:"EmbFlgCd"`
	AcrdgWhenRprincPayIntFlgCd   string  `json:"AcrdgWhenRprincPayIntFlgCd"`
	FamilyWorkMangNormlFlgCd     string  `json:"FamilyWorkMangNormlFlgCd"`
	BrwerAndFamlMebrCprtChkFlgCd string  `json:"BrwerAndFamlMebrCprtChkFlgCd"`
	ChkDlwthRemrk                string  `json:"ChkDlwthRemrk"`
	CotactTelNo                  string  `json:"CotactTelNo"`
	CurrAddr                     string  `json:"CurrAddr"`
	Lgtd                         float64 `json:"Lgtd"`
	LgtdDrctCd                   string  `json:"LgtdDrctCd"`
	Lttd                         float64 `json:"Lttd"`
	LttdDrctCd                   string  `json:"LttdDrctCd"`
	UpdAftRsdnceScAdcmCd         string  `json:"UpdAftRsdnceScAdcmCd"`
	UpdAftRsdnceDtlAddr          string  `json:"UpdAftRsdnceDtlAddr"`
	UpdAftRsdnceAddr             string  `json:"UpdAftRsdnceAddr"`
	UpdAftCmunicAddrScAdcmCd     string  `json:"UpdAftCmunicAddrScAdcmCd"`
	UpdAftCmunicDtlAddr          string  `json:"UpdAftCmunicDtlAddr"`
	UpdAftCmunicAddr             string  `json:"UpdAftCmunicAddr"`
	UpdGrndPostCd                string  `json:"UpdGrndPostCd"`
	MrgncyConterName             string  `json:"MrgncyConterName"`
	AndBrwerRelCd                string  `json:"AndBrwerRelCd"`
	MrgncyConterTelNo            string  `json:"MrgncyConterTelNo"`
	SponsorEmpnbr                string  `json:"SponsorEmpnbr"`
	SponsorOrgNo                 string  `json:"SponsorOrgNo"`
	CorgEmpnbr                   string  `json:"CorgEmpnbr"`
	CorgOrgNo                    string  `json:"CorgOrgNo"`
	RiskAnlyzInfo                string  `json:"RiskAnlyzInfo"`
	OneslfCmtdInfo               string  `json:"OneslfCmtdInfo"`
	DlwthSugstnCd                string  `json:"DlwthSugstnCd"`
	DlwthSugstnComnt             string  `json:"DlwthSugstnComnt"`
	DlwthEmpnbr                  string  `json:"DlwthEmpnbr"`
	DlwthOrgNo                   string  `json:"DlwthOrgNo"`
	DlwthDt                      string  `json:"DlwthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDHCO struct {
	Records                      []DAILRDHCORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDHCORecords struct {
	ChkTaskNo                    string  `json:"ChkTaskNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	TaskDlwthResultNo            string  `json:"TaskDlwthResultNo"`
	HndlrEmpnbr                  string  `json:"HndlrEmpnbr"`
	ResultCnfrmRemrkInfo         string  `json:"ResultCnfrmRemrkInfo"`
	ReacCustManrCd               string  `json:"ReacCustManrCd"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CorgDlwthResultCd            string  `json:"CorgDlwthResultCd"`
	CollTms                      int     `json:"CollTms"`
	EmbFlgCd                     string  `json:"EmbFlgCd"`
	AcrdgWhenRprincPayIntFlgCd   string  `json:"AcrdgWhenRprincPayIntFlgCd"`
	FamilyWorkMangNormlFlgCd     string  `json:"FamilyWorkMangNormlFlgCd"`
	BrwerAndFamlMebrCprtChkFlgCd string  `json:"BrwerAndFamlMebrCprtChkFlgCd"`
	ChkDlwthRemrk                string  `json:"ChkDlwthRemrk"`
	CotactTelNo                  string  `json:"CotactTelNo"`
	CurrAddr                     string  `json:"CurrAddr"`
	Lgtd                         float64 `json:"Lgtd"`
	LgtdDrctCd                   string  `json:"LgtdDrctCd"`
	Lttd                         float64 `json:"Lttd"`
	LttdDrctCd                   string  `json:"LttdDrctCd"`
	UpdAftRsdnceScAdcmCd         string  `json:"UpdAftRsdnceScAdcmCd"`
	UpdAftRsdnceDtlAddr          string  `json:"UpdAftRsdnceDtlAddr"`
	UpdAftRsdnceAddr             string  `json:"UpdAftRsdnceAddr"`
	UpdAftCmunicAddrScAdcmCd     string  `json:"UpdAftCmunicAddrScAdcmCd"`
	UpdAftCmunicDtlAddr          string  `json:"UpdAftCmunicDtlAddr"`
	UpdAftCmunicAddr             string  `json:"UpdAftCmunicAddr"`
	UpdGrndPostCd                string  `json:"UpdGrndPostCd"`
	MrgncyConterName             string  `json:"MrgncyConterName"`
	AndBrwerRelCd                string  `json:"AndBrwerRelCd"`
	MrgncyConterTelNo            string  `json:"MrgncyConterTelNo"`
	SponsorEmpnbr                string  `json:"SponsorEmpnbr"`
	SponsorOrgNo                 string  `json:"SponsorOrgNo"`
	CorgEmpnbr                   string  `json:"CorgEmpnbr"`
	CorgOrgNo                    string  `json:"CorgOrgNo"`
	RiskAnlyzInfo                string  `json:"RiskAnlyzInfo"`
	OneslfCmtdInfo               string  `json:"OneslfCmtdInfo"`
	DlwthSugstnCd                string  `json:"DlwthSugstnCd"`
	DlwthSugstnComnt             string  `json:"DlwthSugstnComnt"`
	DlwthEmpnbr                  string  `json:"DlwthEmpnbr"`
	DlwthOrgNo                   string  `json:"DlwthOrgNo"`
	DlwthDt                      string  `json:"DlwthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDHCRequestForm) PackRequest(dailrdhcI DAILRDHCI) (responseBody []byte, err error) {

	requestForm := DAILRDHCRequestForm{
		Form: []DAILRDHCIDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDHCI",
				},

				FormData: dailrdhcI,
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
func (o *DAILRDHCRequestForm) UnPackRequest(request []byte) (DAILRDHCI, error) {
	dailrdhcI := DAILRDHCI{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdhcI, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdhcI, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDHCResponseForm) PackResponse(dailrdhcO DAILRDHCO) (responseBody []byte, err error) {
	responseForm := DAILRDHCResponseForm{
		Form: []DAILRDHCODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDHCO",
				},
				FormData: dailrdhcO,
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
func (o *DAILRDHCResponseForm) UnPackResponse(request []byte) (DAILRDHCO, error) {

	dailrdhcO := DAILRDHCO{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdhcO, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdhcO, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDHCI) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
