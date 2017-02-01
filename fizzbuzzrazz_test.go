package main

import "testing"

func TestFizz(t *testing.T) {
  if fizzbuzz(3) != "Fizz" {
    t.Error("expected Fizz")
  }
}

func TestBuzz(t *testing.T) {
  if fizzbuzz(5) != "Buzz" {
    t.Error("expected Buzz")
  }
}

func TestRazz(t *testing.T) {
  if fizzbuzz(7) != "Razz" {
    t.Error("expected Razz")
  }
}

func TestFizzBuzzRazz(t *testing.T) {
  if fizzbuzz(105) != "FizzBuzzRazz" {
    t.Error("expected FizzBuzzRazz")
  }
}

func TestBuzzRazz(t *testing.T) {
  if fizzbuzz(35) != "BuzzRazz" {
    t.Error("expected BuzzRazz")
  }
}

func TestNumber(t *testing.T) {
  if fizzbuzz(1) != "1" {
    t.Error("expected 1")
  }
}
