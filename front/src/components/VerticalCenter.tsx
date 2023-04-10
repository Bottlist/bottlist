import { Grid } from '@mui/material';
import { ComponentProps } from 'react';

export const VerticalCenter = ({
  children,
  ...props
}: ComponentProps<typeof Grid>) => (
  <Grid
    container
    alignItems={'center'}
    justifyContent={'center'}
    minHeight={'100vh'}
    {...props}
  >
    <Grid item width="100%">
      {children}
    </Grid>
  </Grid>
);
