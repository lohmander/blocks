import React from "react";
import { styled } from "./theme";

export const Button = styled("button", {
  backgroundColor: "$primary",
  border: "none",
  padding: "0 $2",
  height: 30,
  borderRadius: 9999,
  color: "white",
  fontSize: "x-small",
  textTransform: "uppercase",
  fontWeight: "600",
  letterSpacing: 0.5,
  cursor: "pointer",
  transition: ".2s transform ease, .2s color ease",

  "&:hover": {
    transform: "scale(1.05)",
  },

  "&:active": {
    transform: "scale(.95)",
    color: "rgba(255, 255, 255, 0.8)",
  },
});
