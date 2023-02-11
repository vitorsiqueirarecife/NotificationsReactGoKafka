import {
    FontSizeProps,
    LineHeightProps,
    TextColorProps,
    FontWeightProps,
    TypographyProps as TypographySSProps,
  } from 'styled-system';
  
  export type TypographyProps = TextColorProps &
    TypographySSProps &
    LineHeightProps &
    FontSizeProps &
    FontWeightProps;
  
  export interface Props extends TypographyProps {
    children: string;
  }