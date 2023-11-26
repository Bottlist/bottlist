import { useQuery } from '@tanstack/react-query';
import { Logo } from '../../../components/Logo';
import {
  Avatar,
  Box,
  Button,
  Container,
  Grid,
  Paper,
  Stack,
  Typography,
} from '@mui/material';
import { request } from '../../../utils/axiosUtils';
import { useParams } from 'react-router';

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
    <Container>
      <Stack>
        <Logo />
        <Grid container spacing={4}>
          <Grid item xs={4}>
            <Avatar src={data?.img} />
          </Grid>
          <Grid item xs={8}>
            <Stack textAlign="center" spacing={2}>
              <Paper>
                <Typography variant="h4" margin={2}>
                  {data?.name} 様
                </Typography>
                <Typography>生年月日</Typography>
                <Typography>{data?.birthdate}</Typography>
              </Paper>
              <Box>
                <Button color="secondary">メッセージ</Button>
              </Box>
            </Stack>
          </Grid>
        </Grid>
        <Typography>キープボトル一覧</Typography>
        <Paper>
          <Container>
            {data?.bottles.map((bottle) => (
              // <BottleCard
              //   key={bottle.name}
              //   upperText={bottle.shop_name + ' ' + bottle.name}
              //   lowerText={bottle.expires_at + 'まで'}
              // >
              //   <Button>編集</Button>
              // </BottleCard>
              <p key={bottle.name}>a</p>
            ))}
          </Container>
        </Paper>
      </Stack>
    </Container>
  );
};
