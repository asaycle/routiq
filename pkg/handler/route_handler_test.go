package handler

import (
	"context"
	"testing"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	"github.com/asaycle/routiq.git/pkg/domain/entity"
	"github.com/asaycle/routiq.git/pkg/domain/usecase"
	usecasemock "github.com/asaycle/routiq.git/pkg/domain/usecase/mocks"
	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/proto"
)

func TestRouteHandler_CreateRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		UnimplementedRouteServiceServer pb.UnimplementedRouteServiceServer
		useCase                         usecase.RouteUsecase
	}
	type args struct {
		ctx          context.Context
		req          *pb.CreateRouteRequest
		mockBehavior func(mock *usecasemock.MockRouteUsecase)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Route
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				useCase: usecasemock.NewMockRouteUsecase(ctrl),
			},
			args: args{
				ctx: context.Background(),
				req: &pb.CreateRouteRequest{
					Route: &pb.Route{
						DisplayName: "name-1",
						Description: "desc-1",
					},
				},
				mockBehavior: func(mock *usecasemock.MockRouteUsecase) {
					mock.EXPECT().CreateRoute(
						gomock.Any(),
						"name-1",
						"desc-1",
						gomock.Any(),
					).Return(&entity.Route{
						ID:          "id-1",
						Name:        "name-1",
						Description: "desc-1",
					}, nil)
				},
			},
			want: &pb.Route{
				Id:          "id-1",
				DisplayName: "name-1",
				Description: "desc-1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &RouteHandler{
				UnimplementedRouteServiceServer: pb.UnimplementedRouteServiceServer{},
				useCase:                         tt.fields.useCase,
			}
			tt.args.mockBehavior(tt.fields.useCase.(*usecasemock.MockRouteUsecase))
			got, err := h.CreateRoute(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RouteHandler.CreateRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("RouteHandler.CreateRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
