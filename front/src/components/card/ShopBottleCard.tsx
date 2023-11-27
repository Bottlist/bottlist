import { Card, Typography, Stack, Box, SxProps } from '@mui/material';
import dayjs, { Dayjs } from 'dayjs';
import { components } from '../../schema';
import { Sake } from '../sake/Sake';

type BottleType = components['schemas']['bottle']['type'];

export const ShopBottleCard = (props: {
  type: BottleType;
  owner: string;
  name: string;
  expires_at: Dayjs;
  onClick: () => void;
}) => {
  const style: SxProps = {
    width: 120,
    height: 150,
    borderRadius: '30px',
    boxShadow: '0px 4px 4px 0px rgba(0, 0, 0, 0.25)',
  };
  if (dayjs().isAfter(props.expires_at)) {
    style.borderTop = '4px solid #D27C2C';
    style.borderBottom = '4px solid #D27C2C';
  }
  return (
    <Card sx={style} onClick={props.onClick}>
      <Box width="fit-content" height={0}>
        <Box
          zIndex={1}
          width="120px"
          height="150px"
          position="relative"
          sx={{
            background:
              'linear-gradient(90deg, rgba(255, 255, 255, 0.00) 30%, #FFF 55%)',
          }}
        />
        <Box position="relative" top="-150px">
          <Sake type={props.type} width="auto" height="150px" />
        </Box>
      </Box>
      <Stack
        alignItems="end"
        zIndex={2}
        position="relative"
        height="100%"
        justifyContent="space-around"
        paddingRight="8px"
        paddingY="15px"
      >
        <Typography fontSize={12}>{props.owner}</Typography>
        <Typography fontSize={12}>{props.name}</Typography>
        <Box>
          <Typography fontSize={7} textAlign="center">
            キープ期限
          </Typography>
          <Typography fontSize={12}>
            {props.expires_at.format('YY/MM/DD')}
          </Typography>
        </Box>
      </Stack>
    </Card>
  );
};
