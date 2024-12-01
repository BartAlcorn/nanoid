# nanoid

A tiny and fast Go unique string generator

Safe. It uses cryptographically strong random APIs and tests distribution of symbols.

Compact. It uses a larger alphabet than UUID (A-Z a-z 0-9 _ -). So ID size was reduced from 36 to 21 symbols.

## Not original work

Initially copied from: [github.com/aidarkhanov/nanoid](https://github.com/aidarkhanov/nanoid/blob/master/nanoid.go)

License: [original MIT license](https://github.com/aidarkhanov/nanoid/blob/master/LICENSE)

Any errors due to changes or modifications to the original code is on me.

## Usage

``` go

id := nanoid.WebSafeID()
// a 4 x 4 web safe id, first character is NOT a numeric, as html element IDs require
// e.g. aHd7-5gCW-8Hdt-Z9wh

```
