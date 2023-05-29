export const bubbleSort = (
  array: Array<number>,
  sortType: string
): Array<number> => {
  if (sortType === "ascending") {
    for (let i = 0; i < array.length - 1; i++) {
      // from left->right , small->big
      if (array[i] > array[i + 1]) {
        //   return i,i+1
        [array[i], array[i + 1]] = [array[i + 1], array[i]];

      }
    }
  } else {
    // from left->right , big->small
    for (let i = array.length - 1; i >= 1; i--) {
      if (array[i] > array[i - 1]) {
        [array[i], array[i - 1]] = [array[i - 1], array[i]];
      }
    }
  }

  return array;
};

export const quickSort = (
  array: Array<number>,
  sortType: string
): Array<number> => {
  if (sortType === "ascending") {
    _quickSort(array, 0, array.length - 1);
    return array;
  } else {
    // from left->right , big->small
    for (let i = array.length - 1; i >= 1; i--) {
      if (array[i] > array[i - 1]) {
        [array[i], array[i - 1]] = [array[i - 1], array[i]];
      }
    }
  }

  return array;
};

const _quickSort = (array: Array<number>, left: number, right: number) => {
  if (left >= right) {
    return;
  }

  const index = _findPivot(array, left, right);
//   _quickSort(array, left, index - 1);
//   _quickSort(array, index + 1, right);
};

const _findPivot = (
  array: Array<number>,
  left: number,
  right: number
): number => {
  const pivot = array[left];

  while (left < right) {
    while (left < right && array[right] >= pivot) {
      right -= 1;
    }
    array[left] = array[right];

    while (left < right && array[left] <= pivot) {
      left += 1;
    }
    array[right] = array[left];
    [array[left], array[right]] = [array[right], array[left]];
  }

  array[left] = pivot;
  return left;
};

// A better way is to use lodash to compare array==sorted(array) directly
// which requires O(LogN) time complexity
// current it requires O(N) which is not so good
export const isArraySorted = (
  array: Array<number>,
  sortType: string
): boolean => {

  if (sortType === "ascending") {
    for (let i = 0; i < array.length - 1; i++) {
      // from left->right , small->big
      if (array[i] > array[i + 1]) {
        return false;
      }
    }
  } else {
    // from left->right , big->small
    for (let i = array.length - 1; i >= 1; i--) {
      if (array[i] > array[i - 1]) {
        return false;
      }
    }
  }

  return true;
};
