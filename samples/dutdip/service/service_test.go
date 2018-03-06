package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rdadbhawala/dutdip/samples/dutdip/model"
	"github.com/rdadbhawala/dutdip/samples/dutdip/model/mock"
)

func TestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tsaf := testAllFInit(ctrl)
	tsaf.MockDependency.EXPECT().Operation().Times(1)
	tsaf.MockDepFactory.EXPECT().NewDependency().Return(tsaf.MockDependency)
	model.F = tsaf

	srv := model.F.NewService()
	srv.Feature()
}

type testAllF struct {
	FactoryImpl
	*mock.MockDepFactory
	*mock.MockDependency
}

func testAllFInit(ctrl *gomock.Controller) *testAllF {
	return &testAllF{FactoryImpl{}, mock.NewMockDepFactory(ctrl), mock.NewMockDependency(ctrl)}
}
