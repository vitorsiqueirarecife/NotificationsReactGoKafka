import styled from 'styled-components';

import {
  display,
  border,
  flexbox,
  space,
  position,
  width,
  color,

} from 'styled-system';

import { ButtonProps } from './types';

const Button = styled.button<ButtonProps>`
  ${color}
  ${display}
  ${flexbox}
  ${space}
  ${border}
  ${position}
  ${width}
  cursor: pointer;
`;
export default Button;