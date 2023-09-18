import { Box, Modal, Paper, Stack, Typography } from '@mui/material';
import { CenterLayout } from '../../components/CenterLayout';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { FormProvider, useForm } from 'react-hook-form';
import { request } from '../../utils/axiosUtils';
import { useNavigate } from 'react-router';
import { Button } from '../../components/Button';
import { GrayTextField } from '../../components/GrayTextField';
import { Button as RawButton } from '@mui/base/Button';
import { useState } from 'react';
import { ResetModal } from './ResetModal';

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

  const Footer = (
    <Box marginTop={5}>
      <RawButton
        style={{
          appearance: 'none',
          backgroundColor: 'transparent',
          border: 'none',
        }}
        onClick={() => setOpen(true)}
      >
        パスワードをお忘れの方は こちら
      </RawButton>
    </Box>
  );

  return (
    <>
      <CenterLayout footer={Footer}>
        <FormProvider {...methods}>
          <form onSubmit={onSubmit}>
            <Stack spacing={5} paddingY={5} alignItems="center">
              <TextField _key="email" label="メールアドレス" />
              <TextField _key="password" label="パスワード" />
              <Box>
                <Button type="submit">ログイン</Button>
              </Box>
            </Stack>
          </form>
        </FormProvider>
      </CenterLayout>
      {open && <ResetModal onClose={() => setOpen(false)} />}
    </>
  );
};
