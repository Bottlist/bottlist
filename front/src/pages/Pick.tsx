import { Button, Container, Grid, Typography } from '@mui/material';
import { ReactNode } from 'react';

export function Pick() {
  return (
    <Container>
      <Grid
        container
        alignItems={'center'}
        justifyContent={'space-evenly'}
        minHeight={'100vh'}
        direction={'column'}
      >
        <Grid item width="100%">
          <PickButton>お客様</PickButton>
        </Grid>
        <Grid item width="100%">
          <PickButton>お店</PickButton>
        </Grid>
      </Grid>
    </Container>
  );
}

const PickButton = (props: { children: ReactNode }) => (
  <Button color="secondary" fullWidth sx={{ padding: 2 }}>
    <Typography variant="h3" letterSpacing={20}>
      {props.children}
    </Typography>
  </Button>
);
