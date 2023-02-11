import styled from 'styled-components';
import { border } from 'styled-system';
import { SelectProps } from './types';

export const Select = styled.select<SelectProps>`
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