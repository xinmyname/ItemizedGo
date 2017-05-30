# itemizedgo

This is the Go version of Itemized. The program adds items to inventory and then prints out how many items you have. You can specify how many items to add by passing a number as an argument. If you don't specify any items, you will receive one item. 

Yes, it's very simple. That's the idea.

## Example

```
$> go run itemize.go
You have:
  An item
$> go run itemize.go 4
You have:
  Four items
$> go run itemize.go 42
You have: 
  42 items
```

Or, alternatively...

```
$> go build
$> ./itemizedgo 42
You have:
  42 items
```

## Prerequisites
- go 1.8.1

## Notes

Go made me think about pointers again. 

## The Good Parts
- Garbage collected
- Outputs a single platform-specific executable
- Very easy builds - no project file or makefile

## Airing of Grievances
- Forced folder workspace structure - separates go code from other code, even if related
- Can't override methods/constructors 
- Can't use "self" or "this" when declaring methods for a type
- Have to respecify type for every method associated with a type
- strconv.Atoi() - are you kidding? "Atoi"?
- No string interpolation - use Sprintf(). SPRINTF.
- Can't resolve package cycles
- No tuples
- Arrays not implicitly converted to slices

