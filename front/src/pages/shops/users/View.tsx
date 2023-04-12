import { BottleCard } from '../../../components/BottleCard';
import { Logo } from '../../../components/Logo';
import {
  Box,
  Button,
  Container,
  Grid,
  Paper,
  Stack,
  Typography,
} from '@mui/material';

export const View = () => {
  return (
    <Container>
      <Stack>
        <Logo />
        <Grid container spacing={4}>
          <Grid item xs={4}>
            <img
              src="https://www.yomeishu-online.jp/yomeishu-online_wp/wp-content/uploads/2020/08/herb_set700x300_p01.jpg"
              width="100%"
            />
          </Grid>
          <Grid item xs={8}>
            <Stack textAlign="center" spacing={2}>
              <Paper>
                <Typography variant="h4" margin={2}>
                  飯山 なぎ 様
                </Typography>
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
            <BottleCard upperText="白洲 300年" lowerText="開封日: 23/03/02">
              <Button>編集</Button>
            </BottleCard>
          </Container>
        </Paper>
      </Stack>
    </Container>
  );
};
