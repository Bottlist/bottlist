import { Box, Divider, Stack, Typography } from '@mui/material';
import dayjs from 'dayjs';
import { request } from '../utils/axiosUtils';
import { useQuery } from '@tanstack/react-query';
import { RejectedBottleCard } from './card/SmallBottleCard';

export const RejectedBottles = () => {
  const { data } = useQuery({
    queryKey: ['bottles'],
    queryFn: () =>
      request({
        url: '/bottles',
        method: 'get',
      }).then((r) => r.data.bottles),
  });
  const hasRejectedBottles = data && data.length > 0;
  return (
    <Box>
      <Divider>申請中・差戻しボトル一覧</Divider>
      <Stack direction="row" spacing={4} paddingLeft="2rem" marginTop="2rem">
        {hasRejectedBottles ? (
          data
            ?.filter((b) => b.status.status !== 'approved')
            .map((bottle) => (
              <RejectedBottleCard
                key={bottle.id}
                name={bottle.name}
                expires_at={dayjs(bottle.expires_at)}
                shopName={bottle.shop.name}
              />
            ))
        ) : (
          <Typography>差戻しリストはありません</Typography>
        )}
      </Stack>
    </Box>
  );
};
