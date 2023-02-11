import styled from 'styled-components';
import { border } from 'styled-system';
import { InputBaseProps } from './types';

export const Textarea = styled.textarea<InputBaseProps>`
  ${border}
  display: block;
  border-radius: 10px;
  margin: 0;
  background: none;
  width: 100%;
  padding: 22px;
  color: green;
  font-size: 18px;
  font-weight: 500px;
`;