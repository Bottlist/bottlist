import { useQuery } from '@tanstack/react-query';
import { Avatar, Container, Grid, Stack, Typography } from '@mui/material';
import { request } from '../../../utils/axiosUtils';
import { useParams } from 'react-router';
import ArrowBackIosNewIcon from '@mui/icons-material/ArrowBackIosNew';
import ForwardToInboxIcon from '@mui/icons-material/ForwardToInbox';
import { UpperLeftLogo } from '../../../components/logo/UpperLeftLogo';
import { BigBottleCard } from '../../../components/card/BigBottleCard';
import dayjs from 'dayjs';

export const View = () => {
  const { id } = useParams();
  const { data } = useQuery({
    queryKey: ['users'],
    queryFn: () =>
      request({
        url: `/users/${id}` as '/users/{id}',
        method: 'get',
      }).then((r) => r.data),
  });
  return (
    <>
      <UpperLeftLogo />
      <Container>
        <Stack marginTop="90px" spacing={3} alignItems="center">
          <Stack direction="row" justifyContent="space-between" width="100%">
            <ArrowBackIosNewIcon />
            <ForwardToInboxIcon />
          </Stack>
          <Grid container>
            <Grid item xs={5}>
              <Avatar
                src={data?.img}
                sx={{ width: '130px', height: '130px' }}
              />
            </Grid>
            <Grid item xs={7}>
              <Typography fontSize={10}>ニックネーム</Typography>
              <Typography>{data?.name}</Typography>
              <Typography fontSize={10}>名前</Typography>
              <Typography fontSize={10}>{data?.name}</Typography>
              <Typography>{data?.name}</Typography>
              <Typography fontSize={10}>生年月日</Typography>
              <Typography>{data?.birthdate}</Typography>
            </Grid>
          </Grid>
          <Typography textAlign="center" variant="h6">
            キープボトル一覧
          </Typography>
          {data?.bottles?.map((b) => (
            <BigBottleCard
              key={b.name}
              type={b.type}
              amount={b.amount}
              name={b.name}
              shopName={b.shop_name}
              expires_at={dayjs(b.expires_at)}
            />
          ))}
        </Stack>
      </Container>
    </>
  );
};
