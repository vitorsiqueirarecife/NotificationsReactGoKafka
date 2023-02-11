import styled from 'styled-components';
import { border } from 'styled-system';
import { OptionProps } from './types';

export const Option = styled.option<OptionProps>`
  ${border}
  display: block;
  border-radius: 6px;
  margin: 0;
  background: #fff;
  width: 100%;
  padding-top: 4px;
  padding-bottom: 4px;
  color: black;
  font-size: 18px;
  font-weight: 500px;
`;