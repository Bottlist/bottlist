import { ButtonProps, Button as MuiButton } from '@mui/material';

export const Button = (props: ButtonProps & { children: string }) => (
  <MuiButton
    {...props}
    sx={{
      width: 100,
      height: 30,
      borderRadius: 30,
      fontSize: 16,
      backgroundColor: '#EAE8E1',
      boxShadow: '0px 4px 4px 0px rgba(0, 0, 0, 0.25);',
    }}
  />
);
