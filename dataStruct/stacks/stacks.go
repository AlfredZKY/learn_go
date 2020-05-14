package stacks

import (
	"errors"
)

// Stack store any element
type Stack []interface{}

// Len stack length
func (stack Stack) Len() int {
	return len(stack)
}

// Cap stack Capacity
func (stack Stack) Cap() int {
	return cap(stack)
}

// Push push a element in stack
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

// Top stack top element
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index,len is 0")
	}
	return stack[len(stack)-1], nil
}

// Pop pop a element from stack
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index,len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

// IsEmpty Determines if a stack is empty
func (stack Stack)IsEmpty()bool{
	return len(stack) == 0
}