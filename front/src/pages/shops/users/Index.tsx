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

export const Index = () => {
  return (
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
                <TextField select label="表示順▼" defaultValue="default">
                  <MenuItem value="default">表示順▼</MenuItem>
                </TextField>
              </Grid>
            </Grid>
            <Card>
              <Grid container>
                <Grid item xs={1}>
                  <Checkbox />
                </Grid>
                <Grid item xs={2}>
                  <CardMedia
                    component="img"
                    image="https://www.yomeishu-online.jp/yomeishu-online_wp/wp-content/uploads/2020/08/herb_set700x300_p01.jpg"
                  />
                </Grid>
                <Grid item xs={6}>
                  <CardContent>
                    <Typography>飯山かずや 様</Typography>
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
          </Stack>
        </Paper>
      </Stack>
    </Container>
  );
};
