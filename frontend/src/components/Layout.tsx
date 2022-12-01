import React from "react";
import { styled } from "./theme";

export const Box = styled("div", {
  display: "block",
  variants: {
    variant: {
      bordered: {
        padding: "$4",
        border: "1px solid",
        borderColor: "$border",
        background: "white",
      },
    },
  },
});

export const Row = styled(Box, {
  display: "flex",
  flexDirection: "row",
  alignItems: "center",

  variants: {
    justify: {
      spaceBetween: {
        justifyContent: "space-between",
      },
    },
    align: {
      stretch: {
        alignItems: "stretch",
      },
      start: {
        alignItems: "flex-start",
      },
      end: {
        alignItems: "flex-end",
      },
    },
  },
});

export const Col = styled(Row, {
  flexDirection: "column",
});

export const PageWidth = styled(Box, {
  maxWidth: "1200px",
  mx: "auto",
});
