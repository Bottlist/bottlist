import { Card, Typography, Chip, Stack } from '@mui/material';
import dayjs, { Dayjs } from 'dayjs';

export const RejectedBottleCard = (props: {
  shopName: string;
  name: string;
  expires_at: Dayjs;
}) => (
  <Card
    sx={{
      width: 150,
      height: 100,
      borderRadius: '30px',
      boxShadow: '0px 4px 4px 0px rgba(0, 0, 0, 0.25)',
    }}
  >
    <Stack direction="row" justifyContent="space-evenly" paddingTop="0.5rem">
      <Typography>{props.shopName}</Typography>
      <Chip
        label="差戻し"
        color={dayjs().isBefore(props.expires_at) ? 'secondary' : 'primary'}
      />
    </Stack>
    <Typography variant="h5" textAlign="center">
      {props.name}
    </Typography>
    <Typography variant="body2" textAlign="center">
      {props.expires_at.format('YYYY年M月D日')}まで
    </Typography>
  </Card>
);
