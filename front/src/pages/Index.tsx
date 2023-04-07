import { Grid, Typography } from '@mui/material';

export function Index() {
  return (
    <Grid
      container
      alignItems={'center'}
      justifyContent={'center'}
      minHeight={'100vh'}
    >
      <Grid item>
        <Typography variant="h1">Bottlist</Typography>
      </Grid>
    </Grid>
  );
}
