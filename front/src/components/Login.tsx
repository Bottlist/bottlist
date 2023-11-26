import { Box, Container, Paper, Stack, Typography } from '@mui/material';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { FormProvider, useForm } from 'react-hook-form';
import { useNavigate } from 'react-router';
import { Button as RawButton } from '@mui/base/Button';
import { useState } from 'react';
import { ResetModal } from './ResetModal';
import { request } from '../utils/axiosUtils';
import { GrayTextField } from './GrayTextField';
import { Button } from './button/Button';
import { CenterLogo } from './logo/CenterLogo';

const schema = yup
  .object({
    email: yup.string().email().required(),
    password: yup.string().required(),
  })
  .required();
type FormType = yup.InferType<typeof schema>;

const TextField = GrayTextField<FormType>;

export const Login = () => {
  const navigate = useNavigate();
  const [open, setOpen] = useState(false);
  const methods = useForm<FormType>({
    resolver: yupResolver(schema),
  });

  const onSubmit = methods.handleSubmit(({ email, password }) =>
    request({
      url: '/auth/login/user',
      method: 'post',
      data: { email, password },
    }).then(() => navigate('/top'))
  );

  return (
    <>
      <CenterLogo />
      <Box display="grid" sx={{ placeItems: 'center' }} height="100vh">
        <Container>
          <Stack spacing={4}>
            <Paper elevation={4}>
              <FormProvider {...methods}>
                <form onSubmit={onSubmit}>
                  <Stack spacing={5} paddingY={5} alignItems="center">
                    <TextField _key="email" label="メールアドレス" />
                    <TextField _key="password" label="パスワード" />
                    <Box>
                      <Button type="submit" color="primary">
                        ログイン
                      </Button>
                    </Box>
                  </Stack>
                </form>
              </FormProvider>
            </Paper>
            <Box textAlign="end">
              <Typography display="inline">パスワードをお忘れの方は</Typography>
              <RawButton
                style={{
                  appearance: 'none',
                  backgroundColor: 'transparent',
                  border: 'none',
                }}
                onClick={() => setOpen(true)}
              >
                <Typography sx={{ textDecorationLine: 'underline' }}>
                  こちら
                </Typography>
              </RawButton>
            </Box>
          </Stack>
        </Container>
      </Box>
      {open && <ResetModal onClose={() => setOpen(false)} />}
    </>
  );
};
