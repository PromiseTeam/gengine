package base

import (
	"gengine/context"
	"gengine/core/errors"
	"reflect"
)

type MethodCall struct {
	MethodName       string
	MethodArgs       *Args
	knowledgeContext *KnowledgeContext
	dataCtx          *context.DataContext
}

func (mc *MethodCall) Initialize(ctx *KnowledgeContext, dataCtx *context.DataContext) {
	mc.knowledgeContext = ctx
	mc.dataCtx = dataCtx

	if mc.MethodArgs != nil {
		mc.MethodArgs.Initialize(ctx, dataCtx)
	}
}

func (mc *MethodCall) AcceptArgs(funcArg *Args) error {
	if mc.MethodArgs == nil{
		mc.MethodArgs = funcArg
		return nil
	}
	return errors.Errorf("methodArgs set twice")
}

func (mc *MethodCall) Evaluate(Vars map[string]interface{}) (interface{}, error) {
	var argumentValues []interface{}
	if mc.MethodArgs == nil {
		argumentValues = make([]interface{}, 0)
	} else {
		av, err := mc.MethodArgs.Evaluate(Vars)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		argumentValues = av
	}

	return mc.dataCtx.ExecMethod(mc.MethodName, argumentValues)
}