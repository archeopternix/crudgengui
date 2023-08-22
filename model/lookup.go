package model

// Lookup is a string list
type Lookup struct {
  List []string     // list that contains the text values
}

// NewLookup creates a pointer to a new Lookup
func NewLookup() *Lookup {
  return new(Lookup)
}

// Add adds an text entry
func (f *Lookup)Add(text string){
  f.List=append(f.List,text)
}

// Delete deletes and text entry and preserves the order
func (f *Lookup)Delete(i int){
  copy(f.List[i:], f.List[i+1:]) // Shift a[i+1:] left one index.
  f.List[len(f.List)-1] = ""     // Erase last element (write zero value).
  f.List = f.List[:len(f.List)-1]     // Truncate slice.
}
