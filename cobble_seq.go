package cobble

import (
	"fmt"
	"reflect"
	"sync"
)

// Seq is the holder of the operations wich may be chained if functional style
type Seq struct {
	val reflect.Value
	tp  reflect.Type
	err error
}

// NewSeq is  a factory to create new Sequence which provides set of operations
func NewSeq(source interface{}) (*Seq, error) {
	srcV := reflect.ValueOf(source)
	kind := srcV.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrSourceNotArray
	}

	elemT := reflect.TypeOf(source).Elem()

	var wrapped = reflect.MakeSlice(reflect.SliceOf(elemT), srcV.Len(), srcV.Cap())
	errs := make(chan error)
	wg := &sync.WaitGroup{}
	wg.Add(srcV.Len())

	for i := 0; i < srcV.Len(); i++ {
		go func(idx int, entry reflect.Value) {
			//Store the transformation result into array of result
			resultEntry := wrapped.Index(idx)
			if entry.Type() != elemT {
				// wg.Done()
				errs <- fmt.Errorf("Type mismatch on element %d; expected %s, found %s",
					idx, elemT, entry.Type())
			} else {
				resultEntry.Set(entry)
			}
			//this go routine is done
			wg.Done()
		}(i, srcV.Index(i))
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	// return the first error
	for err := range errs {
		if err != nil {
			return nil, err
		}
	}
	return &Seq{wrapped, elemT, nil}, nil
}

// IsError verifies, is the Sequence has the errors in previous operations
//
// If the error was occurred all consequent operations will be skipped
// (original sequence will be return to the chain)
func (s *Seq) IsError() bool {
	return s.err != nil
}

// ResetError drops the error and returns the sequence for the chain
func (s *Seq) ResetError() *Seq {
	s.err = nil
	return s
}

// Error returns the instance of error (may be nil if no errors occurred)
func (s *Seq) Error() error {
	return s.err
}

// *** Collection operations

// Append returns a new collection containing the elements from
// the called collection followed by the element given as operand.
//
// In case of incompatible type of operand the Error() flag for the
// source collection will be set and source collection will be returned
func (s *Seq) Append(e interface{}) *Seq {
	if s.err != nil {
		return s
	}
	et := reflect.TypeOf(e)
	if et != s.tp {
		s.err = fmt.Errorf(fmtTypeMismatch, s.tp, et)
		return s
	}
	nval := reflect.Append(s.val, reflect.ValueOf(e))
	return &Seq{nval, s.tp, s.err}
}

// Extend returns a new collection containing the elements from the called collection
// followed by the elements from the operand.
//
// In case of incompatible type of operand the Error() flag for the
// source collection will be set and source collection will be returned
func (s *Seq) Extend(e *Seq) *Seq {
	if s.err != nil {
		return s
	}
	if e.val.Len() < 1 {
		return s
	}

	var elemE = reflect.TypeOf(e.val.Interface()).Elem()
	if elemE.Kind() == reflect.Interface {
		elemE = reflect.TypeOf(e.val.Index(0).Interface())
	}

	if s.tp != elemE {
		s.err = fmt.Errorf(fmtTypeMismatch, s.tp, elemE)
		return s
	}

	nsz := s.val.Len() + e.val.Len()
	ncp := s.val.Cap() + e.val.Cap()
	var wrapped = reflect.MakeSlice(reflect.SliceOf(s.tp), nsz, ncp)

	errs := make(chan error)
	wg := &sync.WaitGroup{}
	wg.Add(nsz)

	for i := 0; i < s.val.Len(); i++ {
		go func(idx int, entry reflect.Value) {
			//Store the transformation result into array of result
			resultEntry := wrapped.Index(idx)
			resultEntry.Set(entry)
			//this go routine is done
			wg.Done()
		}(i, s.val.Index(i))
	}
	for i := 0; i < e.val.Len(); i++ {
		go func(idx int, entry reflect.Value) {
			//Store the transformation result into array of result
			resultEntry := wrapped.Index(idx)
			var et = entry.Type()
			var e = entry
			if et.Kind() == reflect.Interface {
				et = reflect.TypeOf(entry.Interface())
				e = reflect.ValueOf(entry.Interface())
			}
			if et != s.tp {
				// wg.Done()
				errs <- fmt.Errorf("Type mismatch on element %d; expected %s, found %s",
					idx, s.tp, entry.Type())
			} else {
				resultEntry.Set(e)
			}
			//this go routine is done
			wg.Done()
		}(i, e.val.Index(i))
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	// return the first error
	var res_err = s.err
	for err := range errs {
		if err != nil {
			res_err = err
			break
		}
	}

	return &Seq{wrapped, s.tp, res_err}
}

// ExtendN returns a new collection containing the elements from the called collection
// followed by the elements from the operand.
//
// Argument is the native collection (`Array` or `Slice`). N in the method's name is
// fo **Native**
//
// In case of incompatible type of operand the Error() flag for the
// source collection will be set and source collection will be returned
func (s *Seq) ExtendN(e interface{}) *Seq {
	col, err := NewSeq(e)
	if err != nil {
		s.err = err
		return s
	}
	return s.Extend(col)
}

// Get returns an element on sppecific indez, or error if index less tan zero or
// greater than collection size
func (s *Seq) Get(index int) (interface{}, error) {
	if index < 0 || index >= s.val.Len() {
		return nil, ErrOutOfRange
	}
	return s.val.Index(index).Interface(), nil
}

// Head selects the first element of this iterable collection.
//
// Returns the `ErrEmptySequence` error if no elements in collection
func (s *Seq) Head() (interface{}, error) {
	if s.val.Len() == 0 {
		return nil, ErrEmptySequence
	}
	return s.val.Index(0).Interface(), nil
}

// Size returns the size of this collection
func (s *Seq) Size() int {
	return s.val.Len()
}

// ToArray converts this iterable collection to a native array.
//
// As a generic, this function returns the `interface`, so result should be
// typecasted to  the desired array typee. Refer the examples how to do this.
func (s *Seq) ToArray() interface{} {
	out := reflect.MakeSlice(reflect.SliceOf(s.tp), s.val.Len(), s.val.Cap())

	wg := &sync.WaitGroup{}
	wg.Add(s.val.Len())

	for i := 0; i < s.val.Len(); i++ {
		go func(idx int, entry reflect.Value) {
			//Store the transformation result into array of result
			resultEntry := out.Index(idx)
			resultEntry.Set(entry)
			//this go routine is done
			wg.Done()
		}(i, s.val.Index(i))
	}

	wg.Wait()

	return out.Interface()
}
