import { Box, Button, Stack } from '@mui/material';
import { CenterLayout } from '../../components/CenterLayout';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { FormProvider, useForm } from 'react-hook-form';
import { request } from '../../utils/axiosUtils';
import { TextField } from '../../components/TextField';
import { useNavigate } from 'react-router';

const schema = yup
  .object({
    email: yup.string().email().required(),
    password: yup.string().required(),
  })
  .required();
type FormType = yup.InferType<typeof schema>;

const TypedTextField = TextField<FormType>;

export const Login = () => {
  const navigate = useNavigate();
  const methods = useForm<FormType>({
    resolver: yupResolver(schema),
  });

  const onSubmit = methods.handleSubmit(({ email, password }) =>
    request({
      url: '/auth/login',
      method: 'post',
      data: { email, password },
    }).then(() => navigate('/top'))
  );

  return (
    <CenterLayout footer={Footer}>
      <FormProvider {...methods}>
        <form onSubmit={onSubmit}>
          <Stack spacing={5} paddingY={5} alignItems="center">
            <TypedTextField _key="email" placeholder="E-mail Address" />
            <TypedTextField _key="password" placeholder="password" />
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
    <Button color="secondary">Lost your password?</Button>
  </Box>
);
