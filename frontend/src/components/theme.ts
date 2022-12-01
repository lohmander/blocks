import { createStitches } from "@stitches/react";

export const { styled } = createStitches({
  theme: {
    colors: {
      primary: "blueviolet",
      border: "#ccc",
    },
    space: {
      1: "5px",
      2: "10px",
      3: "15px",
      4: "20px",
    },
  },
  utils: {
    mx: (v: number | string) => ({ marginLeft: v, marginRight: v }),
    spacingX: (v: number | string) => ({
      marginRight: v,
      "&:last-child": {
        marginRight: 0,
      },
    }),
    spacingY: (v: number | string) => ({
      marginBottom: v,
      "&:last-child": {
        marginBottom: 0,
      },
    }),
  },
});
