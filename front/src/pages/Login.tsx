import { Box, Button, Stack, TextField } from '@mui/material';
import { CenterLayout } from '../components/CenterLayout';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { FormProvider, useForm } from 'react-hook-form';

const schema = yup
  .object({
    email: yup.string().required(),
    password: yup.string().required(),
  })
  .required();
type FormType = yup.InferType<typeof schema>;

export const Login = () => {
  const methods = useForm<FormType>({
    resolver: yupResolver(schema),
  });

  const submit = () => {
    console.log('TBD');
  };
  return (
    <CenterLayout footer={Footer}>
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(submit)}>
          <Stack spacing={5} paddingY={5} alignItems="center">
            <TextField placeholder="username or E-mail Address" />
            <TextField placeholder="Password" />
            <Box>
              <Button>ログイン</Button>
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
