import { Logo } from '../../components/Logo';
import {
  Box,
  Button,
  Container,
  Divider,
  Grid,
  MenuItem,
  Paper,
  Stack,
  TextField,
  Typography,
} from '@mui/material';
import { request } from '../../utils/axiosUtils';
import { useQuery } from '@tanstack/react-query';
import { Navigation } from './components/Navigation';
import { BigBottleCard } from '../../components/BigBottleCard';
import dayjs from 'dayjs';

export const Top = () => {
  const { data } = useQuery({
    queryKey: ['bottles'],
    queryFn: () =>
      request({
        url: '/bottles',
        method: 'get',
      }).then((r) => r.data.bottles),
  });
  return (
    <Stack>
      <Grid
        container
        direction="column"
        justifyContent="space-between"
        alignItems="center"
        paddingY={3}
      >
        <Grid item flexGrow={1}>
          <Container>
            <Stack direction="row" justifyContent="space-between">
              <Logo />
              <Box>
                <TextField select label="表示順▼" defaultValue="default">
                  <MenuItem value="default">表示順▼</MenuItem>
                </TextField>
              </Box>
            </Stack>
            {data
              ?.filter((b) => b.status.status === 'approved')
              .map((bottle) => (
                <BigBottleCard
                  key={bottle.id}
                  type={bottle.type}
                  name={bottle.name}
                  shopName={bottle.shop.name}
                  expires_at={dayjs(bottle.expires_at)}
                  amount={bottle.amount}
                />
              ))}
          </Container>
        </Grid>
        <Grid item flexGrow={0}>
          <Box marginBottom={3}>
            <Divider />
          </Box>
          <Button color="secondary">
            <Typography variant="h4" margin={2}>
              新規ボトル申請
            </Typography>
          </Button>
        </Grid>
      </Grid>
      <Divider />
      <Typography variant="h6">申請中・差戻しボトル一覧</Typography>
      {data
        ?.filter((b) => b.status.status !== 'approved')
        .map((bottle) => (
          <Paper key={bottle.id}>
            <Typography>{bottle.shop.name}</Typography>
            <Typography>{bottle.status.status}</Typography>
            <Typography>{bottle.name}</Typography>
            <Typography>{bottle.expires_at}</Typography>
          </Paper>
        ))}
      <Navigation />
    </Stack>
  );
};
