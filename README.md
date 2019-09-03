# GO

# Introduction

* GO is created by small team @ Google.

* Creators:

    * Robert Griesemar
    * Rob Pike
    * Ken Thompson

 > Why New Language ?

    In google there're mainly 3 languages in Google

        1. Python
        2. Java
        3. C/C++

    1. Python :
        Easy to use, but slow ? interpreted.
    2.  Java :
        Increasingly complex type system
    3. C/C++ :
        Complex type system, also slow compile times.

    When all these 3 languages created was mainly focussed on SINGLE Thread.
 # GO

> Strong and statically typed

    1. So it inherits that Java & C/C++ features.
    2. Strong :
        1. Type of variable can not change over time.
             EX : var a int, its always int.
    3. Statically Typed :
        1. All  the variable had to  defined  at compile time.
> Excellent community

> Key features

     1. Simplicity
     2. Fast compile time
     3. Garbage collected
     4. Built in concurrency
     5. Compile to standalone binaries

> Run GO Program in official web-site

    main () is the entry point.
```
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

 # Setting up Dev env

 > Windows

 > Linux

 > Mac


 > Project structure

    home/users/workspace/code/
                             src/
                                  github.com/
                                             sanjayshr/
                                                       Main.go
                             pkg/
                             bin/

## RUN + Build + Install

> Run

    $go run src/github.com/sanjayshr/firstapp/Main.go

> Build

    $go build github.com/sanjayshr/firstapp
    Binary will be found in /code

> Install

    $go install github.com/sanjayshr/firstapp
    Binary will be found in /bin

# Variable

> variable declaration

    Ex: Ref: /variable

> Re-declaration

    Ex: Ref: /varReDeclare

> Visibility

    Ex:  Lower case variable means this variable scoped to this package.
    Ex: Unused variables, you have to use declared variables, clean-code
    No "private" cope in GO, however block-scope can be possible.


> Naming conventions

    Ex:  Lower case variable means this variable scoped to this package.
     Ref: varNaming.


>  Type conversation

    Ex: Ref: /typeconversion
    Into to String use "strconv" package.




# Primitive Types

> Boolean types

    Ref: /boolean
    Sample code:
    var n bool
    fmt.Printf("%v, %T\n", n, n)
    Output ?


> Numeric types

    1. Integer :

        1. Signed integers:

            Type	Size	Range
            int8	8 bits	-128 to 127
            int16	16 bits	-215 to 215 -1
            int32	32 bits	-231 to 231 -1
            int64	64 bits	-263 to 263 -1
            int 	Platform dependent	Platform dependent

        2. Unsigned integers:

            Type	Size	Range
            uint8	8 bits	0 to 255
            uint16	16 bits	0 to 216 -1
            uint32	32 bits	0 to 232 -1
            uint64	64 bits	0 to 264 -1
            uint	Platform dependent	Platform dependent

        Logic operators:
        Ref: /logicoperators

        Arithmetic operators:
        Ref: /arithmeticoperators



    2. Floating point
        Ex:

    3. Complex numbers
        Ex:

> String

    var s string = "This is string"
    s[3] = "u" // Strings are immutable

# Constants

> Naming convention

    Ref: /constants

> Typed constants

    const myConst int = 10

> Untyped constants

    const myConst = 10
    var n int16 = 20
    fmt.Println("%v, %T\n", myConst + n) // works

> Enumerated constants

    Ref: /enumeratedconstants
    iota is scoped to its constant block

> Enumeration expression

# Collection types:

## Arrays and Slices :

> Arrays

    Arrays should be the same type.

    1. Creation
        Ref: /arrays
    2. Built-in functions
    3. Working with arrays

> Slices

    1. Creation
        ref: /slices
    2. Built-in functions
    3. Working with Slices
        * Remove first value from the slice
        * Remove last value from the slice
        * Remove middle value from the slice

# Maps and Structs
# MPS

> Agenda:

    1. What they are ?
        Maps are a data type in GO, which maps keys to the values.
        Ref: /maps
    2. Creating
    3. Manipulation

# Structs

> Agenda:

	* Collection of disparate data types that describe a single concept.
	* Keyed by name field.
	* Normally created as "types", but anonymous structs are allowed.
	* Structs are value type.
	* No inheritance, but use composition via embedding.

    1. What they are ?
    	Typed collections of field. They are usefull for grouping data togrthere to form records.
	It gathers information about concept, like example the Village.
	Ref: /struct
    2. Creating
    	Different ways of initiating struct
    3. Naming conventions
    4. Embedding
    5. Tags

    	Meta data about struct field

# Function

	1. 

# Pointers

	1. Go lang provides pointers
	2. Pointers are the place to hold the address of value.
	3. A pionter is defined by * .

# Switch
	
	Ref: /switch




# Pass By Reference  & Pass By Value

	Ref: /passbyvalue & /passbyref
