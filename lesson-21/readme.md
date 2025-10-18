---
title: Types and Methods
description: 
date: 2024-02-05
tags: ['go', 'types', 'methods]
---


# #Types 
* a user defined types can be created with the type keyword 
* derive a type from existing primitive or composite types 
    * type todoList []string
* in the past lesson we have created struct types.
    * type student struct {...}

## Abstract types & Concrete Types 
* abstract type defines what the type should do. 
* example of abstract type is function type that we have created in previous lessons. 
* concrete defines what and how the type should do it. 
* example of concrete type is a struct with a method 

# Methods 
* methods are functions that are associated and belong to a type 
* you can define methods for both custom or derived type 
* syntax is slightly different from regualr function 
* their are no overloaded methods in go 
* methods can be accessed just like field i.e, dot notation. 


# Pointer and Value Receiver 

## Which receiver type to use ? 
* general rule of when to use pointers apply to method receivers as well. 
* method modifying the receiver will need to use pointer receiver. 
* such method should also handle the case for nil pointer receivers. 
* method accessing a value in the receiver can just use value receiver.

## Pointer Receiver
* in the last session we saw that we can call the pointer receiver method on a value as well. 
* this is because go automatically converts the passed in value to a pointer for a pointer receiver method. 
* pointer receiver only works when used with correct pointer. 

# Method Sets 
* with pointer you can call pointer receiver methods as well as value receiver methods . 
* if value is addressible you can call pointer receiver methods  as well as value receiver methods . 
* if value is not addressible, go cant get a pointer to it , so you can only call value receiver methods. 

# Methods and Functions
* methods are attachde to types while funcitons are standalone
* methods can be used for functions if the signature of method is same as function type. 
* similar to function method can be saved in variables and passed around

## Method expression 
* their is another way of using methods defined on types called method expressions. 
* this allows you to save a method in a variable 
* the first argument to this function will be the receiver 
* Examples : 
    * f1 := Person.getDetails; f1(p)
    * f2 := Rectangle.SetLength; f2(r,1)