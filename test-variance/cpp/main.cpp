#include <iostream>
using namespace std;

class Animal {
public:
  void breathe() {
    cout << "Breathe in... breathe out..." << endl;
  }
};

class Cat: public Animal {
public:
};

class AnimalSequence {
public:
  Animal next() {
    return Animal();
  }

  void insert(Cat cat) {

  }
};

class CatSequence : public AnimalSequence {

public:
  Cat next() {
    return Cat();
  }

  void insert(Animal animal) {

  }
};

int main() {
  Animal a1;

  Cat c1;

  a1.breathe();

  c1.breathe();

  return 0;
}

