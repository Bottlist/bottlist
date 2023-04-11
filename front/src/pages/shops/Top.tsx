import { Logo } from '../../components/Logo';
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
  Container,
  Divider,
  Grid,
  Paper,
  Stack,
  Typography,
} from '@mui/material';
import { ReactNode } from 'react';

export const Top = () => {
  return (
    <Container>
      <Stack spacing={3}>
        <Logo />
        <Paper>顧客様一覧</Paper>
        <Paper>
          <Typography variant="h6">申請待ち</Typography>
          <BottleCard>
            <Stack>
              <Button>承認</Button>
              <Button>差戻し</Button>
            </Stack>
          </BottleCard>
        </Paper>
        <Paper>
          <Typography variant="h6">キープ期限</Typography>
          <BottleCard>
            <Button size="small">
              <Typography variant="caption">メッセージ</Typography>
            </Button>
          </BottleCard>
        </Paper>
      </Stack>
    </Container>
  );
};

const BottleCard = (props: { children: ReactNode }) => (
  <Card>
    <Grid container alignItems="center">
      <Grid item xs={2}>
        <CardMedia
          component="img"
          image="https://www.yomeishu-online.jp/yomeishu-online_wp/wp-content/uploads/2020/08/herb_set700x300_p01.jpg"
        />
      </Grid>
      <Grid item xs={7}>
        <CardContent>
          <Stack direction="row" spacing={2}>
            <Typography>飯山かずや 様</Typography>
            <Typography>23/3/1</Typography>
          </Stack>
          <Divider />
          <Stack direction="row" spacing={2}>
            <Typography>白洲</Typography>
            <Typography>300年</Typography>
          </Stack>
        </CardContent>
      </Grid>
      <Grid item xs={3} flexGrow={0}>
        <CardActions>{props.children}</CardActions>
      </Grid>
    </Grid>
  </Card>
);
