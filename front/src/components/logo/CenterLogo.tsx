import { Box } from '@mui/material';
import { Logo } from './Logo';

export const CenterLogo = () => (
  <Box position="absolute" top={165} width="100%" textAlign="center">
    <Logo />
  </Box>
);
