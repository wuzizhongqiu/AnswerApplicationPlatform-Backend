package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/encoding"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"

	pb "wuzigoweb/api/http/question"
)

type QuestionService struct {
	uc *biz.QuestionUsecase
}

func NewQuestionService(uc *biz.QuestionUsecase) *QuestionService {
	return &QuestionService{uc: uc}
}

func (s *QuestionService) CreatQuestion(ctx context.Context, req *pb.CreatQuestionRequest) (*pb.CreatQuestionReply, error) {
	jsonCodec := encoding.GetCodec("json")
	bytes, err := jsonCodec.Marshal(req.Contents)
	if err != nil {
		fmt.Println("json 序列化出错，err:", err)
		return nil, err
	}
	content := string(bytes)
	err = s.uc.CreatQuestion(ctx, &model.Question{
		AppID:           req.AppId,
		QuestionContent: &content,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreatQuestionReply{}, nil
}

func (s *QuestionService) ModifyQuestion(ctx context.Context, req *pb.ModifyQuestionRequest) (*pb.ModifyQuestionReply, error) {
	jsonCodec := encoding.GetCodec("json")
	bytes, err := jsonCodec.Marshal(req.Contents)
	if err != nil {
		fmt.Println("json 序列化出错，err:", err)
		return nil, err
	}
	content := string(bytes)
	err = s.uc.ModifyQuestion(ctx, &model.Question{
		AppID:           req.AppId,
		QuestionContent: &content,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ModifyQuestionReply{}, nil
}

func (s *QuestionService) ListQuestionByPage(ctx context.Context, req *pb.ListQuestionByPageRequest) (*pb.ListQuestionByPageReply, error) {
	questions, err := s.uc.ListQuestionByPage(ctx, req.AppId)
	if err != nil {
		return nil, err
	}
	questionContent := make([]*pb.QuestionContent, 0)
	jsonCodec := encoding.GetCodec("json")
	err = jsonCodec.Unmarshal([]byte(*questions.QuestionContent), &questionContent)
	if err != nil {
		fmt.Println("json 反序列化出错，err:", err)
		return nil, err
	}
	return &pb.ListQuestionByPageReply{
		AppId:    questions.AppID,
		Contents: questionContent,
	}, nil
}

func (s *QuestionService) AddUserAnswer(ctx context.Context, req *pb.AddUserAnswerRequest) (*pb.AddUserAnswerReply, error) {
	resultId, err := s.uc.AddUserAnswer(ctx, req.AppId, req.Choices)
	if err != nil {
		return nil, err
	}
	return &pb.AddUserAnswerReply{
		Id: resultId,
	}, nil
}

func (s *QuestionService) GetResultById(ctx context.Context, req *pb.GetResultByIdRequest) (*pb.GetResultByIdReply, error) {
	result, err := s.uc.GetResultById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetResultByIdReply{
		Id:         result.ID,
		ResultDesc: *result.ResultDesc,
		ResultName: result.ResultName,
	}, nil
}
