import { Box, Stack } from '@mui/material';
import { CenterLayout } from '../../components/CenterLayout';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { FormProvider, useForm } from 'react-hook-form';
import { request } from '../../utils/axiosUtils';
import { useNavigate } from 'react-router';
import { Button } from '../../components/Button';
import { GrayTextField } from '../../components/GrayTextField';
import { Button as RawButton } from '@mui/base/Button';

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
  );
};

const Footer = (
  <Box marginTop={5}>
    <RawButton
      style={{
        appearance: 'none',
        backgroundColor: 'transparent',
        border: 'none',
      }}
    >
      パスワードをお忘れの方は こちら
    </RawButton>
  </Box>
);
