//
// This is only a SKELETON file for the 'Raindrops' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

function divisibleBy3(number) {
  if (number % 3 == 0) {
    return true;
  }
  return false;
}

function divisibleBy5(number) {
  if (number % 5 == 0) {
    return true;
  }
  return false;
}

function divisibleBy7(number) {
  if (number % 7 == 0) {
    return true;
  }
  return false;
}

export const convert = (number) => {
  let d3 = divisibleBy3(number);
  let d5 = divisibleBy5(number);
  let d7 = divisibleBy7(number);
  let res = ""

  if (d3) {
    res += 'Pling';
  }
  if (d5) {
    res += 'Plang';
  }
  if (d7) {
    res += 'Plong'
  }
  if(!(d3 || d5 || d7)) {
    return String(number);
  }
  return res;
};
