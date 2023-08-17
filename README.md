# Set
An implementation of `Set` data structure in **GoLang**.

## How To Install
Get the latest version (maybe untagged):
```bash
go get github.com/MhmdRyhn/set
```
Get a specific version:
```bash
go get github.com/MhmdRyhn/set@<tag>
```

## How To Use
```go
package yourpackage

import "github.com/MhmdRyhn/set"

func main() {
    // Initialize
    si := set.New(1, 5, 6)

    // Get set items
    fmt.Println(si.Members())

    // Get the items count in the set
    fmt.Println(si.Len())

    // Check if the set is empty or null
    fmt.Println(si.IsEmpty())
    fmt.Println(si.IsNull())

    // Add a item(s) in the set
    si.Add(4, 5)

    // Check if the set contains an item
    fmt.Println(si.Contains(4))

    // Remove an item from the set
    si.Remove(1)

    // -- Set Operations --
    ss1 := set.new("a", "b", "c")
    ss2 := set.New("b", "d")

    // Union of two sets
    resultUnion := ss1.Union(*ss2)

    // Intersection of two sets
    resultIntersection := ss1.Intersection(*ss2)

    // Difference between two sets
    resultDifference := ss1.Difference(*ss2)

    // -- Set Relationshipd --
    // Check if two sets are equal
    fmt.Println(ss1.IsEqual(*ss2))

    // Check if two sets are disjoint
    fmt.Println(ss1.IsEqual(*ss2))

    // Check if a set is a subset of another
    fmt.Println(ss2.IsSubsetOf(*ss1, false))

    // Check if a set is a strictly proper subset of another
    fmt.Println(ss2.IsSubsetOf(*ss1, true))
}
```
