import styled from "styled-components";

import {
  color,
  space,
  layout,
  flexbox,
  grid,
  border,
  position,
  textAlign,
} from "styled-system";

import { BoxProps } from "./types";

const Box = styled.section<BoxProps>`
  background-image: url(${(p) => p.backgroundImage});

  ${color}
  ${space}
  ${layout}
  ${flexbox}
  ${grid}
  ${border}
  ${position}
  ${textAlign}
`;

export default Box;