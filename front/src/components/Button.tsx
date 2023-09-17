import { ButtonProps, Button as MuiButton } from '@mui/material';

export const Button = (props: ButtonProps & { children: string }) => (
  <MuiButton {...props} sx={{ width: 100, height: 30, fontSize: 100 }} />
);
