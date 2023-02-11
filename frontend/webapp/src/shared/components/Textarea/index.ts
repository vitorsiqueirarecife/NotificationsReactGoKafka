import styled from 'styled-components';
import { border } from 'styled-system';
import { InputBaseProps } from './types';

export const Textarea = styled.textarea<InputBaseProps>`
  ${border}
  display: block;
  border-radius: 6px;
  margin: 0;
  background: #fff;
  padding-top: 4px;
  padding-bottom: 4px;
  width: 100%;
  color: black;
  font-size: 18px;
  font-weight: 500px;
`;