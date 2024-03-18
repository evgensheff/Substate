package substate

import (
	"errors"
	"fmt"
	"strings"
)

func NewSubstate(preState WorldState, postState WorldState, env *Env, message *Message, result *Result, block uint64, transaction int) *Substate {
	return &Substate{
		PreState:    preState,
		PostState:   postState,
		Env:         env,
		Message:     message,
		Result:      result,
		Block:       block,
		Transaction: transaction,
	}
}

type Substate struct {
	PreState    WorldState
	PostState   WorldState
	Env         *Env
	Message     *Message
	Result      *Result
	Block       uint64
	Transaction int
}

// Equal returns true if s is y or if values of s are equal to values of y.
// Otherwise, s and y are not equal hence false is returned.
func (s *Substate) Equal(y *Substate) (err error) {
	if s == y {
		return nil
	}

	if (s == nil || y == nil) && s != y {
		return errors.New("one of the substates is nil")
	}

	preState := s.PreState.Equal(y.PreState)
	postState := s.PostState.Equal(y.PostState)
	env := s.Env.Equal(y.Env)
	msg := s.Message.Equal(y.Message)
	res := s.Result.Equal(y.Result)

	if !preState {
		err = errors.Join(err, fmt.Errorf("preState is different\nwant: %v\n got: %v", s.PreState.String(), y.PreState.String()))
	}

	if !postState {
		err = errors.Join(err, fmt.Errorf("postState is different\nwant: %v\n got: %v", s.PostState.String(), y.PostState.String()))
	}

	if !env {
		err = errors.Join(err, fmt.Errorf("env is different\nwant: %v\n got: %v", s.Env.String(), y.Env.String()))
	}

	if !msg {
		err = errors.Join(err, fmt.Errorf("message is different\nwant: %v\n got: %v", s.Message.String(), y.Message.String()))
	}

	if !res {
		err = errors.Join(err, fmt.Errorf("result is different\nwant: %v\n got: %v", s.Result.String(), y.Result.String()))
	}

	return err
}

func (s *Substate) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("PreState: %v", s.PreState.String()))
	builder.WriteString(fmt.Sprintf("PostState: %v", s.PostState.String()))
	builder.WriteString(fmt.Sprintf("Env World State: %v", s.Env.String()))
	builder.WriteString(fmt.Sprintf("Message World State: %v", s.Message.String()))
	builder.WriteString(fmt.Sprintf("Result World State: %v", s.Result.String()))

	return builder.String()
}
