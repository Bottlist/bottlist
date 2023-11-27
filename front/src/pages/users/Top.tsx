import {
  Box,
  Container,
  Divider,
  Grid,
  Stack,
  Tab,
  Tabs,
  Typography,
} from '@mui/material';
import { request } from '../../utils/axiosUtils';
import { useQuery } from '@tanstack/react-query';
import { Navigation } from './components/Navigation';
import { BigBottleCard } from '../../components/card/BigBottleCard';
import dayjs from 'dayjs';
import { UpperLeftLogo } from '../../components/logo/UpperLeftLogo';
import { useState } from 'react';
import { ResponseData } from '../../utils/schemaHelper';
import { RejectedBottles } from '../../components/RejectedBottles';

const KEYS = [
  { key: 'deadline', label: '期限' },
  { key: 'amount', label: '残量' },
  { key: 'shop', label: 'お店' },
] as const;
type Tabs = (typeof KEYS)[number]['key'];

const renderCards = (
  bottles: ResponseData<'/bottles', 'get'>['bottles'],
  sortKey: Tabs
) =>
  bottles
    .filter((b) => b.status.status === 'approved')
    .sort((a, b) => {
      switch (sortKey) {
        case 'deadline':
          return dayjs(a.expires_at).diff(dayjs(b.expires_at), 'D');
        case 'amount':
          return a.amount - b.amount;
        case 'shop':
          return a.shop.name.localeCompare(b.shop.name);
        default:
          return 0;
      }
    })
    .map((bottle) => (
      <BigBottleCard
        key={bottle.id}
        type={bottle.type}
        name={bottle.name}
        shopName={bottle.shop.name}
        expires_at={dayjs(bottle.expires_at)}
        amount={bottle.amount}
      />
    ));

export const Top = () => {
  const { data } = useQuery({
    queryKey: ['bottles'],
    queryFn: () =>
      request({
        url: '/bottles',
        method: 'get',
      }).then((r) => r.data.bottles),
  });
  const [tab, setTab] = useState<Tabs>('deadline');
  return (
    <>
      <UpperLeftLogo />
      <Stack marginTop="3rem">
        <Typography textAlign="center" variant="h6">
          キープボトル一覧
        </Typography>
        <Tabs value={tab} onChange={(_, v) => setTab(v)} variant="fullWidth">
          {KEYS.map((k) => (
            <Tab label={k.label} value={k.key} key={k.key} />
          ))}
        </Tabs>
        <Grid
          container
          direction="column"
          justifyContent="space-between"
          alignItems="center"
          paddingY={3}
        >
          <Grid item flexGrow={1}>
            <Container>{renderCards(data ?? [], tab)}</Container>
          </Grid>
          <Grid item flexGrow={0}>
            <Box marginBottom={3}>
              <Divider />
            </Box>
          </Grid>
        </Grid>
        <RejectedBottles />
        <Navigation />
      </Stack>
    </>
  );
};
