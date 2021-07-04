package main

import "testing"

func TestLooping(test *testing.T) {
  result := looping()
  shouldBe := 249996287100.7386

  if result != shouldBe {
    test.Errorf("Função looping deveria %v, mas retornou %v", shouldBe, result);
  }
}