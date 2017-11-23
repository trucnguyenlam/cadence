// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package cadence

import (
	"errors"
	"fmt"
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type WorkflowService_DescribeWorkflowExecution_Args struct {
	DescribeRequest *shared.DescribeWorkflowExecutionRequest `json:"describeRequest,omitempty"`
}

func (v *WorkflowService_DescribeWorkflowExecution_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.DescribeRequest != nil {
		w, err = v.DescribeRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DescribeWorkflowExecutionRequest_Read(w wire.Value) (*shared.DescribeWorkflowExecutionRequest, error) {
	var v shared.DescribeWorkflowExecutionRequest
	err := v.FromWire(w)
	return &v, err
}

func (v *WorkflowService_DescribeWorkflowExecution_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DescribeRequest, err = _DescribeWorkflowExecutionRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *WorkflowService_DescribeWorkflowExecution_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.DescribeRequest != nil {
		fields[i] = fmt.Sprintf("DescribeRequest: %v", v.DescribeRequest)
		i++
	}
	return fmt.Sprintf("WorkflowService_DescribeWorkflowExecution_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *WorkflowService_DescribeWorkflowExecution_Args) Equals(rhs *WorkflowService_DescribeWorkflowExecution_Args) bool {
	if !((v.DescribeRequest == nil && rhs.DescribeRequest == nil) || (v.DescribeRequest != nil && rhs.DescribeRequest != nil && v.DescribeRequest.Equals(rhs.DescribeRequest))) {
		return false
	}
	return true
}

func (v *WorkflowService_DescribeWorkflowExecution_Args) MethodName() string {
	return "DescribeWorkflowExecution"
}

func (v *WorkflowService_DescribeWorkflowExecution_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var WorkflowService_DescribeWorkflowExecution_Helper = struct {
	Args           func(describeRequest *shared.DescribeWorkflowExecutionRequest) *WorkflowService_DescribeWorkflowExecution_Args
	IsException    func(error) bool
	WrapResponse   func(*shared.DescribeWorkflowExecutionResponse, error) (*WorkflowService_DescribeWorkflowExecution_Result, error)
	UnwrapResponse func(*WorkflowService_DescribeWorkflowExecution_Result) (*shared.DescribeWorkflowExecutionResponse, error)
}{}

func init() {
	WorkflowService_DescribeWorkflowExecution_Helper.Args = func(describeRequest *shared.DescribeWorkflowExecutionRequest) *WorkflowService_DescribeWorkflowExecution_Args {
		return &WorkflowService_DescribeWorkflowExecution_Args{DescribeRequest: describeRequest}
	}
	WorkflowService_DescribeWorkflowExecution_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.InternalServiceError:
			return true
		case *shared.EntityNotExistsError:
			return true
		default:
			return false
		}
	}
	WorkflowService_DescribeWorkflowExecution_Helper.WrapResponse = func(success *shared.DescribeWorkflowExecutionResponse, err error) (*WorkflowService_DescribeWorkflowExecution_Result, error) {
		if err == nil {
			return &WorkflowService_DescribeWorkflowExecution_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DescribeWorkflowExecution_Result.BadRequestError")
			}
			return &WorkflowService_DescribeWorkflowExecution_Result{BadRequestError: e}, nil
		case *shared.InternalServiceError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DescribeWorkflowExecution_Result.InternalServiceError")
			}
			return &WorkflowService_DescribeWorkflowExecution_Result{InternalServiceError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_DescribeWorkflowExecution_Result.EntityNotExistError")
			}
			return &WorkflowService_DescribeWorkflowExecution_Result{EntityNotExistError: e}, nil
		}
		return nil, err
	}
	WorkflowService_DescribeWorkflowExecution_Helper.UnwrapResponse = func(result *WorkflowService_DescribeWorkflowExecution_Result) (success *shared.DescribeWorkflowExecutionResponse, err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.InternalServiceError != nil {
			err = result.InternalServiceError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type WorkflowService_DescribeWorkflowExecution_Result struct {
	Success              *shared.DescribeWorkflowExecutionResponse `json:"success,omitempty"`
	BadRequestError      *shared.BadRequestError                   `json:"badRequestError,omitempty"`
	InternalServiceError *shared.InternalServiceError              `json:"internalServiceError,omitempty"`
	EntityNotExistError  *shared.EntityNotExistsError              `json:"entityNotExistError,omitempty"`
}

func (v *WorkflowService_DescribeWorkflowExecution_Result) ToWire() (wire.Value, error) {
	var (
		fields [4]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.InternalServiceError != nil {
		w, err = v.InternalServiceError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_DescribeWorkflowExecution_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _DescribeWorkflowExecutionResponse_Read(w wire.Value) (*shared.DescribeWorkflowExecutionResponse, error) {
	var v shared.DescribeWorkflowExecutionResponse
	err := v.FromWire(w)
	return &v, err
}

func (v *WorkflowService_DescribeWorkflowExecution_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _DescribeWorkflowExecutionResponse_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalServiceError, err = _InternalServiceError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.BadRequestError != nil {
		count++
	}
	if v.InternalServiceError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("WorkflowService_DescribeWorkflowExecution_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *WorkflowService_DescribeWorkflowExecution_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [4]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.InternalServiceError != nil {
		fields[i] = fmt.Sprintf("InternalServiceError: %v", v.InternalServiceError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}
	return fmt.Sprintf("WorkflowService_DescribeWorkflowExecution_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *WorkflowService_DescribeWorkflowExecution_Result) Equals(rhs *WorkflowService_DescribeWorkflowExecution_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.InternalServiceError == nil && rhs.InternalServiceError == nil) || (v.InternalServiceError != nil && rhs.InternalServiceError != nil && v.InternalServiceError.Equals(rhs.InternalServiceError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}
	return true
}

func (v *WorkflowService_DescribeWorkflowExecution_Result) MethodName() string {
	return "DescribeWorkflowExecution"
}

func (v *WorkflowService_DescribeWorkflowExecution_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}