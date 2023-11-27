import { Box, Container, Stack, TextField, Typography } from '@mui/material';
import { request } from '../../utils/axiosUtils';
import { useQuery } from '@tanstack/react-query';
import { Navigation } from './components/Navigation';
import { useState } from 'react';
import { UpperLeftLogo } from '../../components/logo/UpperLeftLogo';
import { ShopBottleCard } from '../../components/card/ShopBottleCard';
import dayjs from 'dayjs';
import { Modal } from '../../components/Modal';
import ArrowBackIosNewIcon from '@mui/icons-material/ArrowBackIosNew';
import { Button } from '../../components/button/Button';
import { Sake } from '../../components/sake/Sake';
import { components, operations } from '../../schema';

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
  const selectedBottle = data?.find((b) => b.id === selected);
  const pendingBottles = data?.filter((b) => b.status === 'pending') ?? [];

  return (
    <>
      <UpperLeftLogo />
      <Container sx={{ marginTop: '120px' }}>
        <Stack spacing={3}>
          <Typography variant="h6">承認待ち</Typography>
          <Stack direction="row" spacing={2}>
            {pendingBottles.length > 0 ? (
              pendingBottles.map((bottle) => (
                <ShopBottleCard
                  key={bottle.id}
                  type={bottle.type}
                  owner={bottle.username}
                  name={bottle.name}
                  expires_at={dayjs(bottle.expires_at)}
                  onClick={() => setSelected(bottle.id)}
                />
              ))
            ) : (
              <Typography>現在申請されているボトルはございません</Typography>
            )}
          </Stack>
          <Typography variant="h6">キープ期限</Typography>
          {data
            ?.filter((b) => b.status === 'approved')
            .map((bottle) => (
              <ShopBottleCard
                key={bottle.id}
                type={bottle.type}
                owner={bottle.username}
                name={bottle.name}
                expires_at={dayjs(bottle.expires_at)}
                onClick={() => setSelected(bottle.id)}
              />
            ))}
        </Stack>
      </Container>
      <Navigation />
      {selectedBottle && (
        <SelectedBottleModal
          open={!!selected}
          onClose={() => setSelected(undefined)}
          bottle={selectedBottle}
        />
      )}
    </>
  );
};

const SelectedBottleModal = (props: {
  bottle: operations['get-shops-bottles']['responses'][200]['content']['application/json']['bottles'][number];
  open: boolean;
  onClose: () => void;
}) => {
  const { bottle } = props;
  const [reason, setReason] = useState('');
  return (
    <Modal open={props.open} onClose={props.onClose} width="320px">
      <ArrowBackIosNewIcon sx={{ position: 'absolute' }} />
      <Stack alignItems="center" padding="20px" spacing={2}>
        <Stack direction="row">
          <Sake type={bottle.type} height="150px" width="auto" />
          <Box>
            <Typography fontSize={10}>ニックネーム</Typography>
            <Typography>{bottle.username}</Typography>
            <Typography fontSize={10}>名前</Typography>
            <Typography>{bottle.username}</Typography>
            <Typography fontSize={10}>キープボトル</Typography>
            <Typography>{bottle.name}</Typography>
            <Typography fontSize={10}>申請日</Typography>
            <Typography>{bottle.expires_at}</Typography>
          </Box>
        </Stack>
        <Button
          onClick={() =>
            request({
              url: '/bottles/{id}',
              params: { id: bottle?.id },
              method: 'put',
              data: { status: 'approved' },
            }).finally(props.onClose)
          }
        >
          承認
        </Button>
        <Box width="255px">
          <TextField
            multiline
            label="差戻し理由入力（必須）"
            onChange={(e) => setReason(e.target.value)}
          />
        </Box>
        <Box textAlign="end" width="100%">
          <Button
            color="secondary"
            width="fit-content"
            onClick={() =>
              request({
                url: '/bottles/{id}',
                params: { id: bottle?.id },
                method: 'put',
                data: { status: 'rejected', reason },
              }).finally(props.onClose)
            }
          >
            差戻し
          </Button>
        </Box>
      </Stack>
    </Modal>
  );
};
