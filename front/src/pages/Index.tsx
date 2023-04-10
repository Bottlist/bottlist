import { Grid, Typography } from '@mui/material';
import { useNavigate } from 'react-router';

export function Index() {
  const navigate = useNavigate();
  return (
    <Grid
      container
      alignItems={'center'}
      justifyContent={'center'}
      minHeight={'100vh'}
      onClick={() => navigate('pick')}
    >
      <Grid item>
        <Typography variant="h1">Bottlist</Typography>
      </Grid>
    </Grid>
  );
}
