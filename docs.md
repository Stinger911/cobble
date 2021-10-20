# cobble

Package cobble provides Scala-style collection operations

## Types

### type [Seq](/cobble_seq.go#L10)

`type Seq struct { ... }`

Seq is the holder of the operations wich may be chained if functional style

#### func [NewSeq](/cobble_seq.go#L17)

`func NewSeq(source interface{ ... }) (*Seq, error)`

NewSeq is  a factory to create new Sequence which provides set of operations

```golang

_, e1 := NewSeq([]int{}) // empty sequence of ints
if e1 == nil {
}   // everything is ok
c2, e2 := NewSeq([]int{1, 2, 3}) // sequence of 3 ints
if e2 == nil && c2.Size() == 3 {
}   // everything is ok
c3, e3 := NewSeq([]interface{}{1, "2", 3.1415}) // sequence of different types. yes, it's possible, but on your own risk
if e3 == nil && c3.Size() == 3 {
}   // everything is ok

```

#### func (*Seq) [Append](/cobble_seq.go#L87)

`func (s *Seq) Append(e interface{ ... }) *Seq`

Append returns a new collection containing the elements from
the called collection followed by the element given as operand.

In case of incompatible type of operand the Error() flag for the
source collection will be set and source collection will be returned

#### func (*Seq) [Error](/cobble_seq.go#L76)

`func (s *Seq) Error() error`

Error returns the instance of error (may be nil if no errors occurred)

#### func (*Seq) [Extend](/cobble_seq.go#L105)

`func (s *Seq) Extend(e *Seq) *Seq`

Extend returns a new collection containing the elements from the called collection
followed by the elements from the operand.

In case of incompatible type of operand the Error() flag for the
source collection will be set and source collection will be returned

#### func (*Seq) [ExtendN](/cobble_seq.go#L187)

`func (s *Seq) ExtendN(e interface{ ... }) *Seq`

ExtendN returns a new collection containing the elements from the called collection
followed by the elements from the operand.

Argument is the native collection (`Array` or `Slice`). N in the method's name is
fo **Native**

In case of incompatible type of operand the Error() flag for the
source collection will be set and source collection will be returned

#### func (*Seq) [Get](/cobble_seq.go#L198)

`func (s *Seq) Get(index int) (interface{ ... }, error)`

Get returns an element on sppecific indez, or error if index less tan zero or
greater than collection size

#### func (*Seq) [Head](/cobble_seq.go#L208)

`func (s *Seq) Head() (interface{ ... }, error)`

Head selects the first element of this iterable collection.

Returns the `ErrEmptySequence` error if no elements in collection

#### func (*Seq) [IsError](/cobble_seq.go#L65)

`func (s *Seq) IsError() bool`

IsError verifies, is the Sequence has the errors in previous operations

If the error was occurred all consequent operations will be skipped
(original sequence will be return to the chain)

#### func (*Seq) [ResetError](/cobble_seq.go#L70)

`func (s *Seq) ResetError() *Seq`

ResetError drops the error and returns the sequence for the chain

#### func (*Seq) [Size](/cobble_seq.go#L216)

`func (s *Seq) Size() int`

Size returns the size of this collection

#### func (*Seq) [ToArray](/cobble_seq.go#L224)

`func (s *Seq) ToArray() interface{ ... }`

ToArray converts this iterable collection to a native array.

As a generic, this function returns the `interface`, so result should be
typecasted to  the desired array typee. Refer the examples how to do this.

### Choose

```golang
c1, _ := NewSeq([]int{1, 2, 3})
array := c1.ToArray()
switch array.(type) {
case []int:
    res := array.([]int)
    fmt.Println("Result is ", res)
default:
    fmt.Println("Unknown type: ", reflect.TypeOf(array))
}
```

 Output:

```
Result is  [1 2 3]
```

### Sure

```golang
c1, _ := NewSeq([]int{1, 2, 3})
array := c1.ToArray().([]int)
fmt.Println("Result is ", array)
```

 Output:

```
Result is  [1 2 3]
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
