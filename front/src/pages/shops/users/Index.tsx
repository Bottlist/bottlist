import { Logo } from '../../../components/Logo';
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
  Checkbox,
  Container,
  Grid,
  MenuItem,
  Paper,
  Stack,
  TextField,
  Typography,
} from '@mui/material';
import { Navigation } from '../components/Navigation';
import { request } from '../../../utils/axiosUtils';
import { useQuery } from '@tanstack/react-query';
import { useState } from 'react';

type Order = 'default' | 'name' | 'birthday';
export const Index = () => {
  const { data } = useQuery({
    queryKey: ['users'],
    queryFn: () =>
      request({
        url: '/users',
        method: 'get',
      }).then((r) => r.data.users),
  });
  const [order, setOrder] = useState<Order>('default');
  const users =
    order === 'default'
      ? data
      : data?.sort((u1, u2) => (u1[order] < u2[order] ? 1 : -1));
  return (
    <>
      <Container>
        <Stack spacing={3}>
          <Logo />
          <Typography variant="h5" textAlign="center">
            顧客様一覧
          </Typography>
          <Paper>
            <Stack>
              <Grid container alignItems="center" padding={1}>
                <Grid item xs={4}>
                  <Typography variant="h6">申請待ち</Typography>
                </Grid>
                <Grid item xs={4}>
                  <Button>一斉送信</Button>
                </Grid>
                <Grid item xs={4}>
                  <TextField
                    select
                    label="表示順"
                    defaultValue="default"
                    onChange={(e) => setOrder(e.target.value as Order)}
                  >
                    <MenuItem value="default">初期値</MenuItem>
                    <MenuItem value="name">名前順</MenuItem>
                    <MenuItem value="birthdate">誕生日順</MenuItem>
                  </TextField>
                </Grid>
              </Grid>
              {users?.map((user) => (
                <Card key={user.id}>
                  <Grid container>
                    <Grid item xs={1}>
                      <Checkbox />
                    </Grid>
                    <Grid item xs={2}>
                      <CardMedia component="img" image={user.img} />
                    </Grid>
                    <Grid item xs={6}>
                      <CardContent>
                        <Typography>{user.name} 様</Typography>
                      </CardContent>
                    </Grid>
                    <Grid item xs={3} flexGrow={0} justifyContent="center">
                      <CardActions>
                        <Button size="small">
                          <Typography variant="caption">メッセージ</Typography>
                        </Button>
                      </CardActions>
                    </Grid>
                  </Grid>
                </Card>
              ))}
            </Stack>
          </Paper>
        </Stack>
      </Container>
      <Navigation />
    </>
  );
};
