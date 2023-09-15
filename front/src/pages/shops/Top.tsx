import { BottleCard } from '../../components/BottleCard';
import { Logo } from '../../components/Logo';
import {
  Button,
  Container,
  Modal,
  Paper,
  Stack,
  TextField,
  Typography,
} from '@mui/material';
import { request } from '../../utils/axiosUtils';
import { useQuery } from '@tanstack/react-query';
import { Navigation } from './components/Navigation';
import { useState } from 'react';

export const Top = () => {
  const { data } = useQuery({
    queryKey: ['shops/bottles'],
    queryFn: () =>
      request({
        url: '/shops/bottles',
        method: 'get',
      }).then((r) => r.data.bottles),
  });
  const [selected, setSelected] = useState<string>();
  const [reason, setReason] = useState('');
  const selectedBottle = data?.find((b) => b.id === selected);
  const pendingBottles = data?.filter((b) => b.status === 'pending') ?? [];
  return (
    <>
      <Container>
        <Stack spacing={3}>
          <Logo />
          <Typography variant="h6">承認待ち</Typography>
          {pendingBottles.length > 0 ? (
            pendingBottles.map((bottle) => (
              <BottleCard
                key={bottle.id}
                upperText={bottle.username}
                lowerText={bottle.name + '\n' + bottle.expires_at}
              >
                <Stack>
                  <Button onClick={() => setSelected(bottle.id)}>承認</Button>
                  <Button>差戻し</Button>
                </Stack>
              </BottleCard>
            ))
          ) : (
            <Typography>現在申請されているボトルはございません</Typography>
          )}
          <Typography variant="h6">キープ期限</Typography>
          {data
            ?.filter((b) => b.status === 'approved')
            .map((bottle) => (
              <BottleCard
                key={bottle.id}
                upperText={bottle.username}
                lowerText={bottle.name + '\n' + bottle.expires_at}
              >
                <Stack>
                  <Button>承認</Button>
                  <Button>差戻し</Button>
                </Stack>
              </BottleCard>
            ))}
        </Stack>
      </Container>
      <Navigation />
      <Modal open={!!selected} onClose={() => setSelected(undefined)}>
        <Paper>
          <Typography>名前</Typography>
          <Typography>{selectedBottle?.username}</Typography>
          <Typography>キープボトル</Typography>
          <Typography>{selectedBottle?.name}</Typography>
          <Typography>申請日</Typography>
          <Typography>{selectedBottle?.expires_at}</Typography>
          <Button
            onClick={() =>
              request({
                url: '/bottles/{id}',
                params: { id: selectedBottle?.id },
                method: 'put',
                data: { status: 'approved' },
              }).finally(() => setSelected(undefined))
            }
          >
            承認
          </Button>
          <TextField
            multiline
            label="差戻し理由入力（必須）"
            onChange={(e) => setReason(e.target.value)}
          />
          <Button
            onClick={() =>
              request({
                url: '/bottles/{id}',
                params: { id: selectedBottle?.id },
                method: 'put',
                data: { status: 'rejected', reason },
              }).finally(() => setSelected(undefined))
            }
          >
            差戻し
          </Button>
        </Paper>
      </Modal>
    </>
  );
};
