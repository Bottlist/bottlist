import { yupResolver } from '@hookform/resolvers/yup';
import { useForm, FormProvider } from 'react-hook-form';
import * as yup from 'yup';
import { Navigation } from '../../../components/Navigation';
import { Outlet } from 'react-router-dom';
import { Logo } from '../../../components/Logo';
import { Stack } from '@mui/material';

const schema = yup
  .object({
    date: yup.date().required(),
    shop: yup.string().required(),
    type: yup.string().oneOf(['shop'] as const),
  })
  .required();
export type FormType = yup.InferType<typeof schema>;

export const Register = () => {
  const methods = useForm<FormType>({
    resolver: yupResolver(schema),
  });

  const submit = () => {
    console.log('TBD');
  };

  return (
    <>
      <Stack spacing={2}>
        <Logo />
        <FormProvider {...methods}>
          <form onSubmit={methods.handleSubmit(submit)}>
            <Outlet />
          </form>
        </FormProvider>
      </Stack>
      <Navigation />
    </>
  );
};
