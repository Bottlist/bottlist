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
  Stack,
  TextField,
  Typography,
} from '@mui/material';

export const Top = () => (
  <Grid
    container
    direction="column"
    justifyContent="space-between"
    alignItems="center"
    minHeight={'100vh'}
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
        <Card>
          <Grid container>
            <Grid item xs={5}>
              <CardMedia
                component="img"
                image="https://www.yomeishu-online.jp/yomeishu-online_wp/wp-content/uploads/2020/08/herb_set700x300_p01.jpg"
              />
            </Grid>
            <Grid item xs={7}>
              <CardContent>
                <Typography variant="h6">ボトル名</Typography>
                <Divider />
                <Typography>店名</Typography>
                <Typography>最終更新日/更新者</Typography>
              </CardContent>
            </Grid>
          </Grid>
        </Card>
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
);
