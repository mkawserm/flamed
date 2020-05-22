package graphql

import (
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/iface"
	"net/http"
)

type View struct {
	mFlamed                       *flamed.Flamed
	mTransactionProcessorList     []iface.ITransactionProcessor
	mPasswordHashAlgorithmFactory iface.IPasswordHashAlgorithmFactory
}

func (v *View) GetHTTPHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello world"))
	}
}

func NewView(flamed *flamed.Flamed,
	tpList []iface.ITransactionProcessor,
	passwordHashAlgorithmFactory iface.IPasswordHashAlgorithmFactory) *View {

	return &View{
		mFlamed:                       flamed,
		mTransactionProcessorList:     tpList,
		mPasswordHashAlgorithmFactory: passwordHashAlgorithmFactory,
	}
}
