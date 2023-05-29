import React, { Fragment, useEffect, useState } from "react";

interface IProps {
  array: Array<number>;
}

export const ArrayChart = ({ array }: IProps) => {

  const [arraySum, setArraySum] = useState<number>(0);

  useEffect(() => {
    const sum_ = array.reduce((val1, val2) => val1 + val2, 0);
    setArraySum(sum_);
  }, [array]);

  const createShape = (num: number, index: number) => {
    // use ratios*height to handle too small/too big numbers
    const height = num.valueOf();
    return (
      <div
        key={index}
        style={{
          width: 30,
          height: 100 * (height / arraySum),
          backgroundColor: "red",
        }}
      >
        {num}
      </div>
    );
  };

  return (
    <Fragment>
      <h4>Below is chart </h4>
      <div style={{ display: "flex", justifyContent: "space-around" }}>
        {array.map((number, index) => {
          return createShape(number, index);
        })}
      </div>
    </Fragment>
  );
};
