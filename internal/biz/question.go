package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wuzigoweb/internal/data/model"
)

type QuestionRepo interface {
	Save(context.Context, *model.Question) error
	Update(context.Context, *model.Question) error
	ListQuestionPage(context.Context, int64) (*model.Question, error)

	SaveAnswer(context.Context, int64, []string) (int64, error)
	GetResultById(context.Context, int64) (*model.ScoringResult, error)
	//Login(context.Context, *model.User) (*model.User, error)
	//FindByID(context.Context, int64) (*Greeter, errcode)
	//ListByHello(context.Context, string) ([]*Greeter, errcode)
	//ListAll(context.Context) ([]*Greeter, errcode)
}

type QuestionUsecase struct {
	repo QuestionRepo
	log  *log.Helper
}

func NewQuestionUsecase(repo QuestionRepo, logger log.Logger) *QuestionUsecase {
	return &QuestionUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *QuestionUsecase) CreatQuestion(ctx context.Context, question *model.Question) error {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 CreatQuestion 创建题目操作")
	return uc.repo.Save(ctx, question)
}

func (uc *QuestionUsecase) ModifyQuestion(ctx context.Context, question *model.Question) error {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 CreatQuestion 创建题目操作")
	return uc.repo.Update(ctx, question)
}

func (uc *QuestionUsecase) ListQuestionByPage(ctx context.Context, appId int64) (*model.Question, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 ListQuestionByPage 获取题目信息分页")
	return uc.repo.ListQuestionPage(ctx, appId)
}

func (uc *QuestionUsecase) AddUserAnswer(ctx context.Context, appId int64, choices []string) (int64, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 AddUserAnswer 添加用户回答结果")
	return uc.repo.SaveAnswer(ctx, appId, choices)
}

func (uc *QuestionUsecase) GetResultById(ctx context.Context, resultId int64) (*model.ScoringResult, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 GetResultById 根据 id 获取答题结果")
	return uc.repo.GetResultById(ctx, resultId)
}
