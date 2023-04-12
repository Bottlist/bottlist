import { BottleCard } from '../../components/BottleCard';
import { Logo } from '../../components/Logo';
import { Button, Container, Paper, Stack, Typography } from '@mui/material';

export const Top = () => {
  return (
    <Container>
      <Stack spacing={3}>
        <Logo />
        <Paper>顧客様一覧</Paper>
        <Paper>
          <Typography variant="h6">申請待ち</Typography>
          <BottleCard upperText="飯山かずや 様 23/3/1" lowerText="白洲 300年">
            <Stack>
              <Button>承認</Button>
              <Button>差戻し</Button>
            </Stack>
          </BottleCard>
        </Paper>
        <Paper>
          <Typography variant="h6">キープ期限</Typography>{' '}
          <BottleCard upperText="飯山かずや 様 23/3/1" lowerText="白洲 300年">
            <Button size="small">
              <Typography variant="caption">メッセージ</Typography>
            </Button>
          </BottleCard>
        </Paper>
      </Stack>
    </Container>
  );
};
