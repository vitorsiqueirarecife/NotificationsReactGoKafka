import styled from 'styled-components';

import { color, fontFamily, fontStyle, textAlign, fontWeight, fontSize } from 'styled-system';
import { TypographyProps } from './types';

const Typography = styled.label<TypographyProps>`
  line-height: 123%;
  margin-top: 2px;
  margin-bottom: 2px;
  ${fontSize}
  ${fontWeight}
  ${fontFamily}
  ${fontStyle}
  ${color}
  ${textAlign}
`;

export default Typography;