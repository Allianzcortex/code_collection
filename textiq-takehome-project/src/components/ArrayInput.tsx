import React, { ChangeEvent, Fragment, useEffect, useState } from "react";
import { ArrayChart } from "./ArrayChart";
import { bubbleSort, isArraySorted, quickSort } from "./Sort";

export const ArrayInput = () => {
  const [array, setArray] = useState<Array<number>>([]);
  // TODO : should set sortType can only be one of two strings
  const [sortType, setSortType] = useState<string>("ascending");

  const handleInputArray = (event: ChangeEvent<HTMLInputElement>) => {
    let { value } = event.target;
    if (value[value.length - 1] === ",") {
      // 2, => [2,0]
      value = value.slice(0, -1);
    }
    setArray(value.split(",").map(Number));
  };

  const handleSortTypeChange = (event: ChangeEvent<HTMLInputElement>) => {
    setSortType(event.target.value);
  };

  const handleBubbleSort = () => {
    const res = bubbleSort(array, sortType);
    setArray([...res]);
  };

  useEffect(() => {
    console.log(array);
    console.log(sortType);
  }, [array, sortType]);

  return (
    <Fragment>
      <div>test input</div>
      <input
        type="text"
        aria-label="input"
        onChange={handleInputArray}
        required
      ></input>
      <div onChange={handleSortTypeChange}>
        <input type="radio" value="ascending" defaultChecked name="gender" />{" "}
        Ascending
        <input type="radio" value="descending" name="gender" /> Descending
      </div>
      <br />
      <br />
      <div id="result">
        Here is current array after sorting : <br />
        {array.toString()}
      </div>
      <br />
      <br />
      <button
        name="bubble"
        disabled={isArraySorted(array, sortType)}
        onClick={handleBubbleSort}
      >
        BubbleSort
      </button>
      <br />
      <ArrayChart array={array} />
    </Fragment>
  );
};
