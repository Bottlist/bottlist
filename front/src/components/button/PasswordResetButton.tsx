import { Box, Typography } from '@mui/material';
import { Button } from '@mui/base/Button';

export const PasswordResetButton = (props: { onClick: () => void }) => (
  <Box textAlign="end">
    <Typography display="inline">パスワードをお忘れの方は</Typography>
    <Button
      style={{
        appearance: 'none',
        backgroundColor: 'transparent',
        border: 'none',
      }}
      onClick={props.onClick}
    >
      <Typography sx={{ textDecorationLine: 'underline' }}>こちら</Typography>
    </Button>
  </Box>
);
