import { Stack, Typography, Paper, Container } from '@mui/material';
import { VerticalCenter } from './VerticalCenter';
import { ReactNode } from 'react';

export const CenterLayout = (props: {
  children: ReactNode;
  spacing?: number;
  footer?: ReactNode;
}) => {
  return (
    <VerticalCenter>
      <Stack spacing={props.spacing ?? 5} marginTop={3}>
        <Typography variant="h4">BOTTLIST</Typography>
        <Container>
          <Paper>
            <Container>{props.children}</Container>
          </Paper>
          {props.footer}
        </Container>
      </Stack>
    </VerticalCenter>
  );
};
