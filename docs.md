# cobble

Scala-style collection operations

## Types

### type [Seq](/cobble_seq.go#L9)

`type Seq struct { ... }`

#### func [NewSeq](/cobble_seq.go#L16)

`func NewSeq(source interface{ ... }) (*Seq, error)`

Creates new Sequence which provides set of operations

#### func (*Seq) [Append](/cobble_seq.go#L86)

`func (s *Seq) Append(e interface{ ... }) *Seq`

Returns a new collection containing the elements from
the called collection followed by the element given as operand.

In case of incompatible type of operand the Error() flag for the
source collection will be set and source collection will be returned

#### func (*Seq) [Error](/cobble_seq.go#L75)

`func (s *Seq) Error() error`

Gets the error (may be nil if no errors occured)

#### func (*Seq) [Extend](/cobble_seq.go#L104)

`func (s *Seq) Extend(e *Seq) *Seq`

Returns a new collection containing the elements from the called collection
followed by the elements from the operand.

In case of incompatible type of operand the Error() flag for the
source collection will be set and source collection will be returned

#### func (*Seq) [ExtendN](/cobble_seq.go#L185)

`func (s *Seq) ExtendN(e interface{ ... }) *Seq`

Returns a new collection containing the elements from the called collection
followed by the elements from the operand.

Argument is the native collection (`Array` or `Slice`)

In case of incompatible type of operand the Error() flag for the
source collection will be set and source collection will be returned

#### func (*Seq) [Get](/cobble_seq.go#L196)

`func (s *Seq) Get(index int) (interface{ ... }, error)`

Returns an element on sppecific indez, or error if index less tan zero or
greater than collection size

#### func (*Seq) [Head](/cobble_seq.go#L206)

`func (s *Seq) Head() (interface{ ... }, error)`

Selects the first element of this iterable collection.

Returns the `ErrEmptySequence` error if no elements in collection

#### func (*Seq) [IsError](/cobble_seq.go#L64)

`func (s *Seq) IsError() bool`

Verify, is the Sequence has the errors in previous operations

If the error was occured all consequent operations will be skipped
(original sequence will be return to the chain)

#### func (*Seq) [ResetError](/cobble_seq.go#L69)

`func (s *Seq) ResetError() *Seq`

Resets the error and returns the sequence for the chain

#### func (*Seq) [Size](/cobble_seq.go#L214)

`func (s *Seq) Size() int`

The size of this collection

#### func (*Seq) [ToArray](/cobble_seq.go#L222)

`func (s *Seq) ToArray() interface{ ... }`

Converts this iterable collection to an array.

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
