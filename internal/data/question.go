package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	pb "wuzigoweb/api/http/question"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"
)

type questionRepo struct {
	data *Data
	log  *log.Helper
}

func NewQuestionRepo(data *Data, logger log.Logger) biz.QuestionRepo {
	return &questionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *questionRepo) Save(ctx context.Context, question *model.Question) error {
	u := r.data.query.Question
	return u.WithContext(ctx).Create(question)
}

func (r *questionRepo) Update(ctx context.Context, question *model.Question) error {
	u := r.data.query.Question
	_, err := u.WithContext(ctx).Where(u.ID.Eq(question.ID)).Update(u.QuestionContent, question)
	if err != nil {
		return err
	}
	return nil
}

func (r *questionRepo) ListQuestionPage(ctx context.Context, appId int64) (*model.Question, error) {
	u := r.data.query.Question
	question, err := u.WithContext(ctx).Where(u.AppID.Eq(appId)).First()
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (r *questionRepo) SaveAnswer(ctx context.Context, appId int64, choices []string) (int64, error) {
	u := r.data.query
	question, err := u.Question.WithContext(ctx).Where(u.Question.AppID.Eq(appId)).First()
	if err != nil {
		return 0, err
	}
	result, err := u.ScoringResult.WithContext(ctx).Where(u.ScoringResult.AppID.Eq(appId)).Find()
	if err != nil {
		return 0, err
	}

	// 题目内容 questionContent
	questionContent := make([]*pb.QuestionContent, 0)
	jsonCodec := encoding.GetCodec("json")
	err = jsonCodec.Unmarshal([]byte(*question.QuestionContent), &questionContent)
	if err != nil {
		fmt.Println("json 反序列化出错，err:", err)
		return 0, err
	}
	// 题目选项列表 resultList
	resultList := result

	// *************** 开始进行评分操作 *****************
	optionCount := make(map[string]int)

	// 遍历题目列表
	for _, questionContentDTO := range questionContent {
		// 遍历答案列表
		for _, answer := range choices {
			// 遍历题目中的选项
			for _, option := range questionContentDTO.Options {
				// 如果答案和选项的key匹配
				if option.Key == answer {
					// 获取选项的result属性
					result := option.Result

					// 如果result属性不在optionCount中，初始化为0
					if _, ok := optionCount[result]; !ok {
						optionCount[result] = 0
					}

					// 在optionCount中增加计数
					optionCount[result]++
				}
			}
		}
	}

	// 遍历每种评分结果，计算哪个结果的得分更高
	maxScore := 0
	var maxScoringResult *model.ScoringResult

	// 遍历评分结果列表
	for _, scoringResult := range resultList {
		var resultProp []string
		err = jsonCodec.Unmarshal([]byte(*scoringResult.ResultProp), &resultProp)
		if err != nil {
			fmt.Println("json 反序列化出错，err:", err)
			return 0, err
		}
		// 计算当前评分结果的分数
		score := 0
		for _, prop := range resultProp {
			score += optionCount[prop]
		}

		// 如果分数高于当前最高分数，更新最高分数和最高分数对应的评分结果
		if score > maxScore {
			maxScore = score
			maxScoringResult.ID = scoringResult.ID
		}
	}

	// 序列化 choices
	bytes, err := jsonCodec.Marshal(choices)
	if err != nil {
		fmt.Println("json 序列化出错，err:", err)
		return 0, err
	}
	choice := string(bytes)
	// 保存用户的选择
	err = u.UserAnswer.WithContext(ctx).Create(&model.UserAnswer{AppID: appId, Choices: &choice, ResultID: &maxScoringResult.ID})
	if err != nil {
		return 0, err
	}
	return maxScoringResult.ID, nil
}

func (r *questionRepo) GetResultById(ctx context.Context, resultId int64) (*model.ScoringResult, error) {
	u := r.data.query.ScoringResult
	scoring, err := u.WithContext(ctx).Where(u.ID.Eq(resultId)).First()
	if err != nil {
		return nil, err
	}
	return scoring, nil
}
