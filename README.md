# itemizedgo

This is the C# version of Itemized. The program adds items to inventory and then prints out how many items you have. You can specify how many items to add by passing a number as an argument. If you don't specify any items, you will receive one item. 

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