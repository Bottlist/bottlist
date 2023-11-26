import { VerticalCenter } from './VerticalCenter';
import { Stack, Paper, Container } from '@mui/material';
import { ReactNode } from 'react';
import { Logo } from './logo/Logo';

export const CenterLayout = (props: {
  children: ReactNode;
  spacing?: number;
  footer?: ReactNode;
}) => {
  return (
    <VerticalCenter>
      <Stack spacing={props.spacing ?? 5} marginTop={3}>
        <Logo />
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
