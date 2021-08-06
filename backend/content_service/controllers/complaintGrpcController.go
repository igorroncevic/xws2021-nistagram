package controllers

import (
	"context"
	"github.com/igorroncevic/xws2021-nistagram/common"
	protopb "github.com/igorroncevic/xws2021-nistagram/common/proto"
	"github.com/igorroncevic/xws2021-nistagram/common/tracer"
	"github.com/igorroncevic/xws2021-nistagram/content_service/model/domain"
	"github.com/igorroncevic/xws2021-nistagram/content_service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ComplaintGrpcController struct {
	service    *services.ComplaintService
	jwtManager *common.JWTManager
}

func NewComplaintController(db *gorm.DB, jwtManager *common.JWTManager) (*ComplaintGrpcController, error) {
	service, err := services.NewComplaintService(db)
	if err != nil {
		return nil, err
	}

	return &ComplaintGrpcController{
		service,
		jwtManager,
	}, nil
}

func (c *ComplaintGrpcController) CreateContentComplaint(ctx context.Context, in *protopb.ContentComplaint) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateContentComplaint")
	defer span.Finish()
	claims, err := c.jwtManager.ExtractClaimsFromMetadata(ctx)
	ctx = tracer.ContextWithSpan(context.Background(), span)

	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	} else if claims.UserId == "" {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, err.Error())
	}

	var contentComplaint *domain.ContentComplaint
	contentComplaint = contentComplaint.ConvertFromGrpc(in)

	err = c.service.CreateContentComplaint(ctx, *contentComplaint)
	if err != nil {
		return &protopb.EmptyResponseContent{}, status.Errorf(codes.Unknown, "could not create report")
	}

	return &protopb.EmptyResponseContent{}, nil
}

func (c *ComplaintGrpcController) GetAllContentComplaints(ctx context.Context, in *protopb.EmptyRequestContent) (*protopb.ContentComplaintArray, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllContentComplaints")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	complaints, err := c.service.GetAllContentComplaints(ctx)
	if err != nil {
		return &protopb.ContentComplaintArray{}, status.Errorf(codes.Unknown, "could not get complaints")
	}

	responseComplaints := []*protopb.ContentComplaint{}
	for _, complaint := range complaints {
		responseComplaints = append(responseComplaints, complaint.ConvertToGrpc())
	}

	return &protopb.ContentComplaintArray{ContentComplaints: responseComplaints}, nil
}

func (c *ComplaintGrpcController) RejectById(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RejectById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := c.service.RejectById(ctx, in.Id)

	return &protopb.EmptyResponseContent{}, err
}

func (c *ComplaintGrpcController) DeleteComplaintByUserId(ctx context.Context, in *protopb.RequestId) (*protopb.EmptyResponseContent, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteComplaintByUserId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	err := c.service.DeleteComplaintByUserId(ctx, in.Id)
	return &protopb.EmptyResponseContent{}, err
}
