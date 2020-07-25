package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDK7IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK7I
}

type DAILRDK7ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK7O
}

type DAILRDK7RequestForm struct {
	Form []DAILRDK7IDataForm
}

type DAILRDK7ResponseForm struct {
	Form []DAILRDK7ODataForm
}

type DAILRDK7I struct {
	LoanDubilNo                      string  `json:"LoanDubilNo"`
	AdjDt                            string  `json:"AdjDt"`
	SeqNo                            int  `json:"SeqNo"`
	CustNo                           string  `json:"CustNo"`
	AdjManrCd                        string  `json:"AdjManrCd"`
	AdjPersonEmpnbr                  string  `json:"AdjPersonEmpnbr"`
	AdjPersonOrgNo                   string  `json:"AdjPersonOrgNo"`
	AdjBefLoanRiskClsfCd             string  `json:"AdjBefLoanRiskClsfCd"`
	AdjAftLoanRiskClsfCd             string  `json:"AdjAftLoanRiskClsfCd"`
	CurtprdCompEtimLnRiskClsfDt      string  `json:"CurtprdCompEtimLnRiskClsfDt"`
	InendIdtfyLoanRiskClsfDt         string  `json:"InendIdtfyLoanRiskClsfDt"`
	InendIdtfyLoanRiskClsCd          string  `json:"InendIdtfyLoanRiskClsCd"`
	CurtprdMchLoanRiskClsfDt         string  `json:"CurtprdMchLoanRiskClsfDt"`
	CurtprdArtgclIdtfyLoanRiskClsfCd string  `json:"CurtprdArtgclIdtfyLoanRiskClsfCd"`
	CurtprdArtgclIdtfyLoanRiskClsfDt string  `json:"CurtprdArtgclIdtfyLoanRiskClsfDt"`
	LoanRiskClsfStusCd  			 string  `json:"LoanRiskClsfStusCd"`
	SugstnDescr		       	  		 string  `json:"SugstnDescr"`
	FinlModfyDt                      string  `json:"FinlModfyDt"`
	FinlModfyTm                      string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                   string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                  string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK7O struct {
	Records                      []DAILRDK7ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK7ORecords struct {
	LoanDubilNo                      string  `json:"LoanDubilNo"`
	AdjDt                            string  `json:"AdjDt"`
	SeqNo                            int  `json:"SeqNo"`
	CustNo                           string  `json:"CustNo"`
	AdjManrCd                        string  `json:"AdjManrCd"`
	AdjPersonEmpnbr                  string  `json:"AdjPersonEmpnbr"`
	AdjPersonOrgNo                   string  `json:"AdjPersonOrgNo"`
	AdjBefLoanRiskClsfCd             string  `json:"AdjBefLoanRiskClsfCd"`
	AdjAftLoanRiskClsfCd             string  `json:"AdjAftLoanRiskClsfCd"`
	CurtprdCompEtimLnRiskClsfDt      string  `json:"CurtprdCompEtimLnRiskClsfDt"`
	InendIdtfyLoanRiskClsfDt         string  `json:"InendIdtfyLoanRiskClsfDt"`
	InendIdtfyLoanRiskClsCd          string  `json:"InendIdtfyLoanRiskClsCd"`
	CurtprdMchLoanRiskClsfDt         string  `json:"CurtprdMchLoanRiskClsfDt"`
	CurtprdArtgclIdtfyLoanRiskClsfCd string  `json:"CurtprdArtgclIdtfyLoanRiskClsfCd"`
	CurtprdArtgclIdtfyLoanRiskClsfDt string  `json:"CurtprdArtgclIdtfyLoanRiskClsfDt"`
	LoanRiskClsfStusCd  			 string  `json:"LoanRiskClsfStusCd"`
	SugstnDescr		       	  		 string  `json:"SugstnDescr"`
	FinlModfyDt                      string  `json:"FinlModfyDt"`
	FinlModfyTm                      string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                   string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                  string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDK7RequestForm) PackRequest(dailrdk7I DAILRDK7I) (responseBody []byte, err error) {

	requestForm := DAILRDK7RequestForm{
		Form: []DAILRDK7IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK7I",
				},
				FormData: dailrdk7I,
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
func (o *DAILRDK7RequestForm) UnPackRequest(request []byte) (DAILRDK7I, error) {
	dailrdk7I := DAILRDK7I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk7I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk7I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDK7ResponseForm) PackResponse(dailrdk7O DAILRDK7O) (responseBody []byte, err error) {
	responseForm := DAILRDK7ResponseForm{
		Form: []DAILRDK7ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK7O",
				},
				FormData: dailrdk7O,
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
func (o *DAILRDK7ResponseForm) UnPackResponse(request []byte) (DAILRDK7O, error) {

	dailrdk7O := DAILRDK7O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk7O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk7O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDK7I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
