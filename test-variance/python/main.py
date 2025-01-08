from typing import Callable

class Animal:
  def breathe(self):
    print("In... Out...")

class Dog(Animal):
  pass

animal = Animal()
animal.breathe()

dog = Dog()
dog.breathe()

# Functions Example

def initDog() -> Dog:
  return Dog()

def initAnimal() -> Animal:
  return Animal()

def checkDog(d: Dog):
  pass

def checkAnimal(a: Animal):
  pass

def test1(a: Callable[[], Animal]):
  pass
def test2(a: Callable[[], Dog]):
  pass
def test3(a: Callable[[Animal], None]):
  pass
def test4(a: Callable[[Dog], None]):
  pass

test1(initDog)
test2(initAnimal)
test3(checkDog)
test4(checkAnimal)
