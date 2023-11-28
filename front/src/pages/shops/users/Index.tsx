import {
  Avatar,
  Box,
  Button,
  Card,
  CardMedia,
  Checkbox,
  Container,
  Paper,
  Stack,
  Typography,
} from '@mui/material';
import { Navigation } from '../components/Navigation';
import { request } from '../../../utils/axiosUtils';
import { useQuery } from '@tanstack/react-query';
import { useState } from 'react';
import { UpperLeftLogo } from '../../../components/logo/UpperLeftLogo';
import ForwardToInboxIcon from '@mui/icons-material/ForwardToInbox';
import { useNavigate } from 'react-router';

const ORDER_KEYS = {
  name: '名前順',
  birthday: '誕生日順',
} as const;
type Order = keyof typeof ORDER_KEYS;
export const Index = () => {
  const { data } = useQuery({
    queryKey: ['users'],
    queryFn: () =>
      request({
        url: '/users',
        method: 'get',
      }).then((r) => r.data.users),
  });
  const [order, setOrder] = useState<Order>('birthday');
  const navigate = useNavigate();
  const users = data?.sort((u1, u2) => (u1[order] < u2[order] ? 1 : -1)) ?? [];
  return (
    <>
      <UpperLeftLogo />
      <Container sx={{ paddingTop: '100px' }}>
        <Typography variant="h5" textAlign="center" marginBottom="15px">
          顧客様一覧
        </Typography>
        <Paper>
          <Stack paddingBottom="30px">
            <Stack
              direction="row"
              marginY="30px"
              justifyContent="space-between"
              paddingX="20px"
            >
              <Box>
                <Checkbox />
                <Button
                  onClick={() =>
                    setOrder(order === 'birthday' ? 'name' : 'birthday')
                  }
                >
                  {ORDER_KEYS[order]}
                </Button>
              </Box>
              <ForwardToInboxIcon />
            </Stack>
            <Stack spacing={3}>
              {users.map((user) => (
                <Stack key={user.id} direction="row">
                  <Checkbox />
                  <Card
                    sx={{
                      width: '275px',
                      borderRadius: '25px',
                      backgroundColor: '#EDE9DA',
                    }}
                    onClick={() => navigate('/shops/users/' + user.id)}
                  >
                    <Stack direction="row" alignItems="center">
                      <CardMedia
                        component={Avatar}
                        image={user.img}
                        sx={{ marginX: '30px' }}
                      />
                      <Box>
                        <Typography>{user.name}</Typography>
                        <Typography fontSize={20}>{user.name} 様</Typography>
                      </Box>
                    </Stack>
                  </Card>
                </Stack>
              ))}
            </Stack>
          </Stack>
        </Paper>
      </Container>
      <Navigation />
    </>
  );
};
