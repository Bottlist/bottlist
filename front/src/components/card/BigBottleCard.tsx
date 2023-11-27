import { Card, Typography, Stack, Box } from '@mui/material';
import { Dayjs } from 'dayjs';
import { components } from '../../schema';
import { Sake } from '../sake/Sake';

type BottleType = components['schemas']['bottle']['type'];

export const BigBottleCard = (props: {
  type: BottleType;
  shopName: string;
  name: string;
  expires_at: Dayjs;
  amount: number;
}) => (
  <Card
    sx={{
      width: 309,
      height: 110,
      borderRadius: '30px',
      boxShadow: '0px 4px 4px 0px rgba(0, 0, 0, 0.25)',
    }}
  >
    <Stack
      direction="row"
      height="100%"
      alignItems="center"
      justifyContent="space-evenly"
    >
      <Box
        width={100}
        height={100}
        borderRadius={30}
        alignItems="center"
        justifyContent="center"
        sx={{
          display: 'flex',
          background: `linear-gradient(180deg, #EDE9DA ${
            props.amount * 100
          }%, #D27C2C 82.08%)`,
        }}
      >
        <Sake type={props.type} />
      </Box>
      <Box>
        <Typography>{props.shopName}</Typography>
        <Typography fontSize={24}>{props.name}</Typography>
        <Typography>{props.expires_at.format('YYYY年M月D日')}まで</Typography>
      </Box>
    </Stack>
  </Card>
);
