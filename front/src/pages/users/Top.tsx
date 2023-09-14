import { Logo } from '../../components/Logo';
import {
  Box,
  Button,
  Card,
  CardContent,
  CardMedia,
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
import { components } from '../../schema';
import { Navigation } from './components/Navigation';

const IMAGES: Record<components['schemas']['bottle']['type'], string> = {
  brandy: 'brandy.png',
  shochu: 'shochu.png',
  whisky: 'Whiske.png',
};

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
                <Card key={bottle.id}>
                  <Grid container>
                    <Grid item xs={5}>
                      <CardMedia component="img" image={IMAGES[bottle.type]} />
                    </Grid>
                    <Grid item xs={7}>
                      <CardContent>
                        <Typography>{bottle.shop.name}</Typography>
                        <Typography variant="h6">{bottle.name}</Typography>
                        <Divider />
                        <Typography>{bottle.expires_at}</Typography>
                      </CardContent>
                    </Grid>
                  </Grid>
                </Card>
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
