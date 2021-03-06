package base

import (
	"gengine/context"
	"gengine/core/errors"
)

type RuleEntity struct {
	RuleName         string
	Salience         int64
	RuleDescription  string
	RuleContent      *RuleContent
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
	Vars             map[string]interface{}  //belongs to current rule,rule execute finish, it will be clear
}

func (r *RuleEntity) AcceptString(s string) error {
	if r.RuleName == "" {
		r.RuleName = s
		return nil
	}

	if r.RuleDescription == "" {
		r.RuleDescription = s
		return nil
	}
	return errors.Errorf("value = %s set twice!",s)
}

func (r *RuleEntity) AcceptInteger(val int64) error {
	r.Salience = val
	return nil
}

func (r *RuleEntity) Initialize(kc *KnowledgeContext,  dc *context.DataContext) {
	r.knowledgeContext = kc
	r.dataCtx = dc

	if r.RuleContent != nil {
		r.RuleContent.Initialize(kc, dc)
	}
}

func (r *RuleEntity) Execute() error {
	r.Vars = make(map[string]interface{})
	defer r.clearMap()
	return r.RuleContent.Execute(r.Vars)
}

func (r *RuleEntity)clearMap()  {
	r.Vars = make(map[string]interface{})
}