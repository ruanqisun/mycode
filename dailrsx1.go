package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX1I
}

type DAILRSX1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX1O
}

type DAILRSX1RequestForm struct {
	Form []DAILRSX1IDataForm
}

type DAILRSX1ResponseForm struct {
	Form []DAILRSX1ODataForm
}

type DAILRSX1I struct {
	CrdtAplySn                    string  `json:"CrdtAplySn"`
	CustNo                        string  `json:"CustNo"`
	StrQtyCustFlg                 string  `json:"StrQtyCustFlg"`
	AplyTypCd                     string  `json:"AplyTypCd"`
	RecomCustMgrEmpnbr            string  `json:"RecomCustMgrEmpnbr"`
	ApyfLprOrgNo                  string  `json:"ApyfLprOrgNo"`
	CrdtAplyDt                    string  `json:"CrdtAplyDt"`
	InitCredQta                   string  `json:"InitCredQta"`
	AplyChnlCd                    string  `json:"AplyChnlCd"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	LoanProdtVersNo               string  `json:"LoanProdtVersNo"`
	LoanProdtNm                   string  `json:"LoanProdtNm"`
	QtaMltpLendFlg                string  `json:"QtaMltpLendFlg"`
	WhtlPreCrdtQta                float64 `json:"WhtlPreCrdtQta"`
	AplyCrdtQta                   float64 `json:"AplyCrdtQta"`
	LoanCurCd                     string  `json:"LoanCurCd"`
	LoanUsageCd                   string  `json:"LoanUsageCd"`
	LoanUsageRplshDescr           string  `json:"LoanUsageRplshDescr"`
	LoanInvtIndusCd               string  `json:"LoanInvtIndusCd"`
	CrdtAplyStusCd                string  `json:"CrdtAplyStusCd"`
	CrdtAplyTm                    string  `json:"CrdtAplyTm"`
	SponsorCustMgrEmpnbr          string  `json:"SponsorCustMgrEmpnbr"`
	CorgCustMgrEmpnbr             string  `json:"CorgCustMgrEmpnbr"`
	ActngMgmtOrgNo                string  `json:"ActngMgmtOrgNo"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX1O struct {
	Records                      []DAILRSX1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX1ORecords struct {
	CrdtAplySn                    string  `json:"CrdtAplySn"`
	CustNo                        string  `json:"CustNo"`
	StrQtyCustFlg                 string  `json:"StrQtyCustFlg"`
	AplyTypCd                     string  `json:"AplyTypCd"`
	RecomCustMgrEmpnbr            string  `json:"RecomCustMgrEmpnbr"`
	ApyfLprOrgNo                  string  `json:"ApyfLprOrgNo"`
	CrdtAplyDt                    string  `json:"CrdtAplyDt"`
	InitCredQta                   string  `json:"InitCredQta"`
	AplyChnlCd                    string  `json:"AplyChnlCd"`
	LoanProdtNo                   string  `json:"LoanProdtNo"`
	LoanProdtVersNo               string  `json:"LoanProdtVersNo"`
	LoanProdtNm                   string  `json:"LoanProdtNm"`
	QtaMltpLendFlg                string  `json:"QtaMltpLendFlg"`
	WhtlPreCrdtQta                float64 `json:"WhtlPreCrdtQta"`
	AplyCrdtQta                   float64 `json:"AplyCrdtQta"`
	LoanCurCd                     string  `json:"LoanCurCd"`
	LoanUsageCd                   string  `json:"LoanUsageCd"`
	LoanUsageRplshDescr           string  `json:"LoanUsageRplshDescr"`
	LoanInvtIndusCd               string  `json:"LoanInvtIndusCd"`
	CrdtAplyStusCd                string  `json:"CrdtAplyStusCd"`
	CrdtAplyTm                    string  `json:"CrdtAplyTm"`
	SponsorCustMgrEmpnbr          string  `json:"SponsorCustMgrEmpnbr"`
	CorgCustMgrEmpnbr             string  `json:"CorgCustMgrEmpnbr"`
	ActngMgmtOrgNo                string  `json:"ActngMgmtOrgNo"`
	FinlModfyDt                   string  `json:"FinlModfyDt"`
	FinlModfyTm                   string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo               string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX1RequestForm) PackRequest(dailrsx1I DAILRSX1I) (responseBody []byte, err error) {

	requestForm := DAILRSX1RequestForm{
		Form: []DAILRSX1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX1I",
				},

				FormData: dailrsx1I,
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
func (o *DAILRSX1RequestForm) UnPackRequest(request []byte) (DAILRSX1I, error) {
	dailrsx1I := DAILRSX1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX1ResponseForm) PackResponse(dailrsx1O DAILRSX1O) (responseBody []byte, err error) {
	responseForm := DAILRSX1ResponseForm{
		Form: []DAILRSX1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX1O",
				},
				FormData: dailrsx1O,
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
func (o *DAILRSX1ResponseForm) UnPackResponse(request []byte) (DAILRSX1O, error) {

	dailrsx1O := DAILRSX1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
