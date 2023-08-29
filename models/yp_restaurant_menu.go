package models

import "github.com/astaxie/beego/orm"

type RestroMenu struct {
	Id                                 int    `orm:"column(id);auto"`
	UserId                             int    `orm:"column(userId);null"`
	LoanId                             int    `orm:"column(loanId);null"`
	ProspectNo                         string `orm:"column(prospectNo);null"`
	ProspectCreationRequest            string `orm:"column(prospectCreationRequest);null"`
	ProspectCreationResponse           string `orm:"column(prospectCreationResponse);null"`
	ProspectCreationResponseStatusCode int    `orm:"column(prospectCreationResponseStatusCode);size(8);null"`
	LoanDecisionRequest                string `orm:"column(loanDecisionRequest);null"`
	LoanDecisionResponse               string `orm:"column(loanDecisionResponse);null"`
	LoanDecisionResponseStatusCode     int    `orm:"column(loanDecisionResponseStatusCode);size(8);null"`
	LoanDisbursalRequest               string `orm:"column(loanDisbursalRequest);null"`
	LoanDisbursalResponse              string `orm:"column(loanDisbursalResponse);null"`
	LoanDisbursalResponseStatusCode    int    `orm:"column(loanDisbursalResponseStatusCode);size(8);null"`
	UploadDocRequest                   string `orm:"column(uploadDocRequest);null"`
	UploadDocResponse                  string `orm:"column(uploadDocResponse);null"`
	UploadDocResponseStatusCode        int    `orm:"column(uploadDocResponseStatusCode);size(8);null"`
	ProspectStatusResponse             string `orm:"column(prospectStatusResponse);null"`
	ProspectStatusResponseStatusCode   int    `orm:"column(prospectStatusResponseStatusCode);null"`
	Status                             string `orm:"column(status);null"`
	UtrNumber                          string `orm:"column(utrNumber);null"`
	GetLoanStatusRequest               string `orm:"column(getLoanStatusRequest);null"`
	GetLoanStatusResponse              string `orm:"column(getLoanStatusResponse);null"`
	GetLoanStatusResponseStatusCode    int    `orm:"column(getLoanStatusResponseStatusCode);null"`
}

func (t *RestroMenu) TableName() string {
	return "yp_lp_iifl_loan"
}

func init() {
	orm.RegisterModel(new(RestroMenu))
}
