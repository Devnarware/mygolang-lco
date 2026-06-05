## for each loop

```
for idx, val := range  arr{
    fmt.print(idx, val)
}
```
##

## maps 

- to declare it
```
m := make(map[key_type] val_type)
```

- we can only use for each loop on map

## Struct

strut is just like the blueprint, it doesn't hold any data but help to create it, you can create your own data type using it 

```
type Box struct{
    height int
    width int
    length int
}
```

## method 

if we combine metod and struct it will beacome the class in java or c++, where struct help us definfing the variables, method helps us to perform some opertion using that variables

- methods are just the fucntion with the struct involved means you can perform various operation on the variable of the struct

```
 func (b Box) area() int{
    return 2 * (length * width + length * height + height + width)
 }
```