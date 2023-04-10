import { Button, Container, Grid, Typography } from '@mui/material';
import { ReactNode } from 'react';
import { Link as RouterLink } from 'react-router-dom';

export function ChoicesLayout(props: {
  items: {
    text: string;
    link?: string;
  }[];
}) {
  return (
    <Container>
      <Grid
        container
        alignItems={'center'}
        justifyContent={'space-evenly'}
        minHeight={'100vh'}
        direction={'column'}
      >
        {props.items.map((item, i) => (
          <Grid item width="100%" key={i}>
            <PickButton link={item.link}>{item.text}</PickButton>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
}

const PickButton = (props: { children: ReactNode; link?: string }) => {
  const { children, link } = props;
  return (
    <Button
      color="secondary"
      fullWidth
      sx={{ padding: 2 }}
      {...(link && { component: RouterLink, to: link })}
    >
      <Typography variant="h3" letterSpacing={20}>
        {children}
      </Typography>
    </Button>
  );
};
