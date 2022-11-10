// Code generated by github.com/oscerai/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"strconv"

	"github.com/oscerai/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type FieldsOrderInputResolver interface {
	OverrideFirstField(ctx context.Context, obj *FieldsOrderInput, data *string) error
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _FieldsOrderPayload_firstFieldValue(ctx context.Context, field graphql.CollectedField, obj *FieldsOrderPayload) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_FieldsOrderPayload_firstFieldValue(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.FirstFieldValue, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_FieldsOrderPayload_firstFieldValue(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "FieldsOrderPayload",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputFieldsOrderInput(ctx context.Context, obj interface{}) (FieldsOrderInput, error) {
	var it FieldsOrderInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"firstField", "overrideFirstField"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "firstField":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("firstField"))
			it.FirstField, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "overrideFirstField":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("overrideFirstField"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			if err = ec.resolvers.FieldsOrderInput().OverrideFirstField(ctx, &it, data); err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var fieldsOrderPayloadImplementors = []string{"FieldsOrderPayload"}

func (ec *executionContext) _FieldsOrderPayload(ctx context.Context, sel ast.SelectionSet, obj *FieldsOrderPayload) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, fieldsOrderPayloadImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("FieldsOrderPayload")
		case "firstFieldValue":

			out.Values[i] = ec._FieldsOrderPayload_firstFieldValue(ctx, field, obj)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNFieldsOrderInput2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐFieldsOrderInput(ctx context.Context, v interface{}) (FieldsOrderInput, error) {
	res, err := ec.unmarshalInputFieldsOrderInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNFieldsOrderPayload2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐFieldsOrderPayload(ctx context.Context, sel ast.SelectionSet, v FieldsOrderPayload) graphql.Marshaler {
	return ec._FieldsOrderPayload(ctx, sel, &v)
}

func (ec *executionContext) marshalNFieldsOrderPayload2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐFieldsOrderPayload(ctx context.Context, sel ast.SelectionSet, v *FieldsOrderPayload) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._FieldsOrderPayload(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
