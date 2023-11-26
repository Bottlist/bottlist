import { Box } from '@mui/material';
import { Logo } from './Logo';

export const UpperLeftLogo = () => (
  <Box position="absolute" top={30} left={30} width="fit-content">
    <Logo fontSize={15} />
  </Box>
);
