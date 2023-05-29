import React, { Fragment } from "react";
import {
  fireEvent,
  render,
  RenderResult,
  screen,
  waitFor,
} from "@testing-library/react";
import { ArrayInput } from "./ArrayInput";

describe("test", () => {
  const renderResult: () => RenderResult = () => {
    return render(<ArrayInput />);
  };

  it("render chart successfully", () => {
    const { container, getByText, getByLabelText } = renderResult();
    const input = getByLabelText("input");
    fireEvent.change(input, { target: { value: "3,2,1" } });

    expect(
      `<div style="width: 30px; background-color: red; height: 33.33333333333333px;">2`
    ).toBeInTheDocument;
  });

  it("click sort button and show different charts", () => {
    // TODO
  });
});
